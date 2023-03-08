package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"github.com/nippati/moff/pkg/playbook"
	"github.com/nippati/moff/pkg/scan"
	"github.com/nippati/moff/pkg/ui"
)

func main() {
	// Read the JSON file of vuls scan results
	file, err := os.Open("vuls-results.json")
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	var results scan.Results
	err = json.NewDecoder(file).Decode(&results)
	if err != nil {
		log.Fatalf("error decoding JSON: %v", err)
	}

	// Extract vulnerability IDs
	re := regexp.MustCompile(`CVE-(\d{4})-(\d+)`)
	var vulns []int
	for _, v := range results.Vulnerabilities {
		matches := re.FindStringSubmatch(v)
		if len(matches) == 3 {
			id, err := strconv.Atoi(matches[2])
			if err != nil {
				log.Fatalf("error converting vulnerability ID to int: %v", err)
			}
			vulns = append(vulns, id)
		}
	}

	// Convert []int to []ui.Vulnerability
	var uiVulns []ui.Vulnerability
	for _, v := range vulns {
		uiVulns = append(uiVulns, ui.Vulnerability{ID: v})
	}

	// Create a HTTP server and listen for requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Display the list of vulnerabilities and allow the user to select which ones to deploy
		selected, err := ui.SelectVulnerabilities(uiVulns)
		if err != nil {
			fmt.Println("Error selecting vulnerabilities:", err)
			return
		}
		// Convert []int to []string
		var selectedStr []string
		for _, v := range selected {
			selectedStr = append(selectedStr, strconv.Itoa(v))
		}

		// Generate an Ansible playbook based on the selected vulnerabilities
		playbook := playbook.Generate(selectedStr)

		// Generate an Ansible playbook based on the selected vulnerabilities

		// Print the playbook to the console
		fmt.Println(playbook)

		// Return the playbook to the user
		w.Write([]byte(playbook))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
