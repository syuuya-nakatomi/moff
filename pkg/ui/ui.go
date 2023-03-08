package ui

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Vulnerability struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func SelectVulnerabilities(vulns []Vulnerability) ([]int, error) {
	var selected []int

	// Create a map of vulnerabilities for easy lookups
	vulnMap := make(map[int]Vulnerability)
	for _, vuln := range vulns {
		vulnMap[vuln.ID] = vuln
	}

	// Prompt the user to select vulnerabilities to deploy
	for {
		// Print the list of available vulnerabilities
		for _, vuln := range vulns {
			fmt.Printf("[%d] %s\n", vuln.ID, vuln.Name)
		}

		// Prompt the user to select a vulnerability
		fmt.Print("Select a vulnerability to deploy (0 to finish): ")
		var input string
		fmt.Scanln(&input)

		// Convert the input to an integer
		id, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		// Check if the user wants to finish selecting vulnerabilities
		if id == 0 {
			break
		}

		// Check if the selected vulnerability exists
		if _, ok := vulnMap[id]; !ok {
			fmt.Println("Invalid selection")
			continue
		}

		// Add the selected vulnerability to the list
		selected = append(selected, id)
	}

	return selected, nil
}

func LoadVulnerabilities(filePath string) ([]Vulnerability, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var vulns []Vulnerability
	err = json.Unmarshal(data, &vulns)
	if err != nil {
		return nil, err
	}

	return vulns, nil
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	vulns, err := LoadVulnerabilities("vuls-results.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, vulns)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func SelectedHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	selectedIDs := r.Form["vulns"]
	var selected []Vulnerability
	for _, idStr := range selectedIDs {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		selected = append(selected, Vulnerability{ID: id})
	}

	playbook := Generate(selected)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=ansible-playbook.yml")
	w.Write([]byte(playbook))
}

func Generate(selected []Vulnerability) string {
	// Generate the Ansible playbook
	var playbook strings.Builder
	playbook.WriteString("hosts: all\n")
	playbook.WriteString("tasks:\n")
	for _, vuln := range selected {
		playbook.WriteString(fmt.Sprintf("- name: Deploy %s\n", vuln.Name))
		playbook.WriteString("  # TODO: Add deployment steps here\n")
		playbook.WriteString(fmt.Sprintf("  shell: echo 'Deploying %s'\n", vuln.Name))
	}
	return playbook.String()
}
