package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nippati/moff/pkg/moff"
)

func main() {
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	// Parse the query parameters
	q := r.URL.Query()
	url := q.Get("url")

	// Check if the "url" parameter is present
	if url == "" {
		http.Error(w, "missing 'url' parameter", http.StatusBadRequest)
		return
	}

	// Fetch the URL contents
	content, err := moff.Fetch(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("error fetching URL: %s", err), http.StatusInternalServerError)
		return
	}

	// Convert the content to Markdown
	md, err := moff.ToMarkdown(content)
	if err != nil {
		http.Error(w, fmt.Sprintf("error converting content to Markdown: %s", err), http.StatusInternalServerError)
		return
	}

	// Write the Markdown response
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if _, err := w.Write([]byte(md)); err != nil {
		log.Printf("error writing response: %s", err)
	}
}
