package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/nippati/moff/pkg/playbook"
	"github.com/nippati/moff/pkg/scan"
	"github.com/nippati/moff/pkg/ui"
)

// hoge
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

	// Convert []string to []int
	var vulns []int
	for _, v := range results.Vulnerabilities {
		id, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("error converting vulnerability ID to int: %v", err)
		}
		vulns = append(vulns, id)
	}

	// Convert []int to []ui.Vulnerability
	var uiVulns []ui.Vulnerability
	for _, v := range vulns {
		uiVulns = append(uiVulns, ui.Vulnerability{ID: v})
	}

	// Display the list of vulnerabilities and allow the user to select which ones to deploy
	selected, err := ui.SelectVulnerabilities(uiVulns)
	if err != nil {
		fmt.Println("Error selecting vulnerabilities:", err)
		return
	}

	// Generate an Ansible playbook based on the selected vulnerabilities
	playbook := playbook.Generate(selected)

	// Print the playbook to the console
	fmt.Println(playbook)
}
