package moff

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Vulnerability struct {
	CveID    string `json:"cve_id"`
	Package  string `json:"package"`
	Severity string `json:"severity"`
}

func ParseVulsJSON(filePath string) ([]Vulnerability, error) {
	var vulnerabilities []Vulnerability

	jsonFile, err := os.Open(filePath)
	if err != nil {
		return vulnerabilities, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &vulnerabilities)
	if err != nil {
		return vulnerabilities, err
	}

	return vulnerabilities, nil
}
