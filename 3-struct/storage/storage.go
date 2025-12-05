// Package storage contains storage functions
package storage

import (
	"encoding/json"

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

func ReadJSONFile(fileName string) (bins.BinList, error) {
	data, err := file.ReadFile(fileName)
	if err != nil {
		return bins.BinList{}, err
	}

	var result bins.BinList
	err = json.Unmarshal(data, &result)
	if err != nil {
		return bins.BinList{}, err
	}
	return result, nil
}
