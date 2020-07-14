package scuff

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func Fold(jsonFileLocation string) {

	data, err := ioutil.ReadFile(jsonFileLocation)
	if handledError("can't read file", err, jsonFileLocation) {
		return
	}

	//get scuff scuff
	scuff := Scuff(jsonFileLocation, data)

	fmt.Println("folding", scuff.Config.In, "into", scuff.Config.Out)
	scuff.walk()
}

func (scuff *scuff) resolvePaths() (inputDir string, outputDir string) {
	scuffDir := scuff.Config.Location

	var err error
	inputDir, err = filepath.Rel(scuffDir, scuff.Config.In)
	if err != nil {
		inputDir = scuff.Config.In
	}
	outputDir, err = filepath.Rel(scuffDir, scuff.Config.Out)
	if err != nil {
		outputDir = scuff.Config.Out
	}
	return inputDir, outputDir
}

func Scuff(location string, data []byte) *scuff {
	if !isDir(location) {
		location = filepath.Dir(location)
	}
	scuff := &scuff{
		Config: scuffConfig{
			Location: location,
			Delim:    delim{
				Left:  "{{",
				Right: "}}",
			},
			In:       "./templates/in",
			Out:      "./out",
			Rewrite: []string{
				"generated_*",
			},
		},
		AsMap: map[string]interface{}{},
	}
	err := json.Unmarshal(data, scuff)
	justLog("invalid scuff scuff section", err)


	err = json.Unmarshal(data, &scuff.AsMap)

	justLog("parsing json", err)
	return scuff
}
