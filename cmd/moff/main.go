package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nippati/moff/pkg/ansible"
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

	// Convert []string to []ui.Vulnerability
	var vulns []ui.Vulnerability
	for _, v := range results.Vulnerabilities {
		vulns = append(vulns, ui.Vulnerability{ID: v})
	}

	// Create a HTTP server and listen for requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Display the list of vulnerabilities and allow the user to select which ones to deploy
		selected, err := ui.SelectVulnerabilities(vulns)
		if err != nil {
			fmt.Println("Error selecting vulnerabilities:", err)
			return
		}

		// Generate an Ansible playbook based on the selected vulnerabilities
		playbook := ansible.GeneratePlaybook(selected)

		// Write the playbook to a file
		file, err := os.Create("playbook.yml")
		if err != nil {
			fmt.Println("Error creating playbook file:", err)
			return
		}
		defer file.Close()

		_, err = file.WriteString(playbook)
		if err != nil {
			fmt.Println("Error writing playbook file:", err)
			return
		}

		// Return a success message to the user
		w.Write([]byte("Ansible playbook successfully generated!"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
