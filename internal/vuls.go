package internal

import (
	"encoding/json"
	"os"
)

type VulsScanResult struct {
	// ...
}

func ParseVulsScanResults(filePath string) ([]VulsScanResult, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var results []VulsScanResult
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&results)
	if err != nil {
		return nil, err
	}

	return results, nil
}
