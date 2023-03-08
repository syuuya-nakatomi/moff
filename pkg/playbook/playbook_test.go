package playbook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	selected := []string{"CVE-2021-1234", "CVE-2021-5678"}

	expected := "---\n" +
		"- name: deploy CVE-2021-1234\n" +
		"  apt:\n" +
		"    name: CVE-2021-1234\n" +
		"    state: latest\n\n" +
		"- name: deploy CVE-2021-5678\n" +
		"  apt:\n" +
		"    name: CVE-2021-5678\n" +
		"    state: latest\n\n"

	actual := Generate(selected)

	assert.Equal(t, expected, actual)
}
