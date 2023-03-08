
import (
	"errors"
	"strings"

	"github.com/gomarkdown/markdown/parser"
	"github.com/russross/blackfriday/v2"
)

// ToMarkdown converts the given HTML content to Markdown.
func ToMarkdown(html string) (string, error) {
	// Create a parser with the necessary options
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)
	renderer := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{Flags: blackfriday.UseXHTML})
	md := blackfriday.New(blackfriday.WithExtensions(extensions), blackfriday.WithRenderer(renderer), blackfriday.WithParser(parser))

	// Convert the HTML to Markdown using Blackfriday
	mdBytes := md.Run([]byte(html))

	// Remove any leading/trailing white space and convert to a string
	mdStr := strings.TrimSpace(string(mdBytes))

	// Check if the resulting Markdown is empty
	if len(mdStr) == 0 {
		return "", errors.New("empty Markdown")
	}

	return mdStr, nil
}
