package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/nippati/moff/pkg/scan"
	"github.com/nippati/moff/pkg/ui"
	"github.comi/nippati/moff/pkg/playbook"
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

	// Display the list of vulnerabilities and allow the user to select which ones to deploy
	selected, err := ui.SelectVulnerabilities(results.Vulnerabilities)
	if err != nil {
		log.Fatalf("error selecting vulnerabilities: %v", err)
	}

	// Generate an Ansible playbook based on the selected vulnerabilities
	playbook := playbook.Generate(selected)

	// Print the playbook to the console
	fmt.Println(playbook)
}
