package playbook

import (
	"fmt"
	"strings"
)

func Generate(selected []string) string {
	var playbook strings.Builder

	// Add header
	playbook.WriteString("---\n")

	// Add tasks for each selected vulnerability
	for _, vuln := range selected {
		playbook.WriteString(fmt.Sprintf("- name: deploy %s\n", vuln))
		playbook.WriteString("  apt:\n")
		playbook.WriteString("    name: " + vuln + "\n")
		playbook.WriteString("    state: latest\n\n")
	}

	return playbook.String()
}
