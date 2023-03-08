package ui

import (
	"fmt"
	"html/template"
	"net/http"
)

type Vulnerability struct {
	ID          int
	Title       string
	Description string
}

func SelectVulnerabilities(vulns []Vulnerability) ([]string, error) {
	selected := make(map[string]bool)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
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
			return
		}

		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for k := range r.PostForm {
			selected[k] = true
		}

		tmpl, err := template.ParseFiles("templates/result.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, vulns)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Server started at http://localhost:8080")
	fmt.Println("Press CTRL-C to stop")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return nil, err
	}

	var selectedSlice []string
	for v := range selected {
		selectedSlice = append(selectedSlice, v)
	}
	return selectedSlice, nil
}
