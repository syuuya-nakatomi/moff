package main

import (
	"encoding/json"
	"html/template"
	"log"
	"os"

	"github.com/nippati/moff/pkg/scan"
	"github.com/nippati/moff/pkg/ui"
)

type Vulnerability struct {
	ID string `json:"id"`
}

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

	// Convert []scan.Vulnerability to []ui.Vulnerability
	var vulns []ui.Vulnerability
	for _, r := range results.Vulnerabilities {
		v := ui.Vulnerability{
			ID: r.VulnID,
		}
		vulns = append(vulns, v)
	}

	// Render the template with the vulnerability data
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatalf("error parsing template: %v", err)
	}
	data := struct {
		Vulns []ui.Vulnerability
	}{
		Vulns: vulns,
	}
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("error rendering template: %v", err)
	}
}
