package main

import (
	"encoding/json"
	"io/ioutil"
)

func readConfigFromFile(filepath string) (*Configuration, error) {
	cfg := new(Configuration)
	fileBytes, errRead := ioutil.ReadFile(filepath)
	if errRead != nil {
		return cfg, errRead
	}
	errJSON := json.Unmarshal(fileBytes, cfg)
	if errJSON != nil {
		return cfg, errJSON
	}
	return cfg, nil
}

func saveConfigToFile(filepath string, cfg *Configuration) error {
	fileBytes, errJSON := json.Marshal(cfg)
	if errJSON != nil {
		return errJSON
	}
	errWrite := ioutil.WriteFile(filepath, fileBytes, 0644)
	if errWrite != nil {
		return errWrite
	}
	return nil
}
