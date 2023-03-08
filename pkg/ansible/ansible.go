package ansible

import (
	"fmt"
	"strings"
)

// GeneratePlaybook generates an Ansible playbook for the selected vulnerabilities
func GeneratePlaybook(selected []string) string {
	var sb strings.Builder

	sb.WriteString("---\n")
	sb.WriteString("- name: Update packages\n")
	sb.WriteString("  hosts: all\n")
	sb.WriteString("  become: true\n\n")

	for _, vuln := range selected {
		sb.WriteString(fmt.Sprintf("  - name: Update packages for %s\n", vuln))
		sb.WriteString("    apt:\n")
		sb.WriteString(fmt.Sprintf("      name: %s\n", vuln))
		sb.WriteString("      state: latest\n\n")
	}

	return sb.String()
}
