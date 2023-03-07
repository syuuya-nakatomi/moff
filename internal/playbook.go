package internal

import (
	"os"
	"text/template"
)

func GeneratePlaybook(selectedItems []string) error {
	// Define the playbook template
	playbookTemplate := `
    ---
    - hosts: all
      tasks:
        {{- range $item := . }}
        - name: Deploy {{ $item }}
          include: {{ $item }}.yml
        {{- end }}
    `

	// Parse the playbook template
	tmpl, err := template.New("playbook").Parse(playbookTemplate)
	if err != nil {
		return err
	}

	// Execute the template with the selected items
	err = tmpl.Execute(os.Stdout, selectedItems)
	if err != nil {
		return err
	}

	return nil
}
