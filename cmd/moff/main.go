package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"

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
		var id string
		if r.ID != "" {
			id = r.ID
		} else if r.VulnID != "" {
			id = r.VulnID
		} else {
			continue
		}

		v := ui.Vulnerability{
			ID: id,
		}
		vulns = append(vulns, v)
	}

	// Render the HTML template with the vulns data
	tmpl := template.Must(template.ParseFiles("vulns.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, vulns)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Open the browser to the app URL
	url := "http://localhost:8080"
	cmd := exec.Command("open", url)
	err = cmd.Run()
	if err != nil {
		log.Fatalf("error opening URL in browser: %v", err)
	}

	// Start the app server
	fmt.Printf("Starting server at %s\n", url)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
