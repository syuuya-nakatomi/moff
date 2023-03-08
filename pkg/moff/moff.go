package moff

import (
	"bytes"
	"errors"
	"io/ioutil"
	"markdown"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gomarkdown/markdown"
)

// Fetch fetches the contents of a URL and returns it as a string.
func Fetch(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// ToMarkdown converts the given HTML content to Markdown.
func ToMarkdown(html string) (string, error) {
	// Convert the HTML to Markdown using the gomarkdown package
	md := markdown.ToMarkdown([]byte(html), nil)

	// Remove any leading/trailing white space and convert to a string
	mdStr := strings.TrimSpace(string(md))

	// Check if the resulting Markdown is empty
	if len(mdStr) == 0 {
		return "", errors.New("empty Markdown")
	}

	return mdStr, nil
}

// ToHTML converts the given Markdown content to HTML.
func ToHTML(md string) (string, error) {
	// Convert the Markdown to HTML using the gomarkdown package
	html := markdown.ToHTML([]byte(md), nil, nil)

	// Check if the resulting HTML is empty
	if len(html) == 0 {
		return "", errors.New("empty HTML")
	}

	return string(html), nil
}

// ToPDF converts the given Markdown content to a PDF file.
func ToPDF(md string) ([]byte, error) {
	// Convert the Markdown to HTML
	html, err := ToHTML(md)
	if err != nil {
		return nil, err
	}

	// Use the wkhtmltopdf command-line tool to convert the HTML to PDF
	cmd := exec.Command("wkhtmltopdf", "-", "-")
	cmd.Stdin = bytes.NewBufferString(html)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}
