// Package storage contains storage functions
package storage

import (
	"encoding/json"
	"os"

	"purple.learn/3-struct/bins"

	"purple.learn/3-struct/file"
)

func SaveToJSONFile(bin *bins.BinList) error {
	data, err := json.Marshal(bin)
	if err != nil {
		return err
	}

	err = file.SaveFile("bin.json", data)
	if err != nil {
		return err
	}

	return nil
}

func ReadJSONFile(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var result []byte
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
