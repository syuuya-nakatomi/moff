package scan

import (
	"encoding/json"
	"os"
)

type Results struct {
	Vulnerabilities []string `json:"Vulnerabilities"`
}

func GetVulnerabilities(filepath string) ([]string, error) {
	// Open the JSON file of vuls scan results
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode the JSON into a Results struct
	var results Results
	err = json.NewDecoder(file).Decode(&results)
	if err != nil {
		return nil, err
	}

	return results.Vulnerabilities, nil
}
