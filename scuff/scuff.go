package scuff

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)


func Of(jsonFileLocation string) *scuff {

	data, err := ioutil.ReadFile(jsonFileLocation)
	if handledError("can't read file", err, jsonFileLocation) {
		return nil
	}
	return OfData(jsonFileLocation, data)
}

func OfData(location string, data []byte) *scuff {
	if !isDir(location) {
		location = filepath.Dir(location)
	}
	scuff := &scuff{
		Config: scuffConfig{
			Location: location,
			Delim: delim{
				Left:  "{{",
				Right: "}}",
			},
			In:  "./templates/in",
			Out: "./out",
			Rewrite: []string{
				"generated_*",
			},
		},
		AsMap: map[string]interface{}{},
	}
	err := json.Unmarshal(data, scuff)
	justLog("invalid Of Of section", err)

	err = json.Unmarshal(data, &scuff.AsMap)

	justLog("parsing json", err)
	return scuff
}

func (f *scuff) Fold() error {
	inputDir, outputDir := f.resolvePaths()
	return filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if s, err := os.Stat(path); err == nil {
			if !s.IsDir() {
				fmt.Println("found template:", path)
				f.execute(inputDir, outputDir, path)
			}
		}
		return err
	})
}

func (f *scuff) resolvePaths() (inputDir string, outputDir string) {
	scuffDir := f.Config.Location

	var err error
	inputDir, err = filepath.Rel(scuffDir, f.Config.In)
	if err != nil {
		inputDir = f.Config.In
	}
	outputDir, err = filepath.Rel(scuffDir, f.Config.Out)
	if err != nil {
		outputDir = f.Config.Out
	}
	return inputDir, outputDir
}
