package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Println("no dir, run it within the directory with scuff.json inside", "wd:"+wd, err)
	}
	flag.Parse()
	location := flag.Arg(0)
	if location == "" {
		location = "scuff.json"
	}
	location = relativeTo(wd, location)
	log.Println("location", location)
	Scuff(location)
}

func Scuff(fileJson string) {
	st, err := os.Stat(fileJson)
	if err != nil {
		log.Fatal(`
								Usage:
										
								scuff myproject/scuff.json
										
									
								scuff.json example:
										
								{
								  "scuff": {
									"delim": {
									  "left": "[[",
									  "right": "]]"
									},
									"out": "./out",
									"in": "./in"
								  },
								  "project": {
									  "shortName": "project1"
								  }
								}`, err)
	}
	if st.IsDir() {
		fileJson = filepath.Join(fileJson, "scuff.json")
	}
	data, err := ioutil.ReadFile(fileJson)
	if err != nil {
		log.Println("run it within the directory with scuff.json inside", fileJson, err)
	}
	m := JsonMap{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		log.Fatalf("cannot parse json: %s", err)
	}
	scuff := m.UnpackMap("scuff")

	outputDir := scuff.AsString("out", "./out")
	inputDir := scuff.AsString("in", "./in")
	scuffDir := filepath.Dir(fileJson)
	outputDir = relativeTo(scuffDir, outputDir)
	inputDir = relativeTo(scuffDir, inputDir)

	log.Println("input dir:", inputDir)
	log.Println("output dir:", outputDir)

	err = walk(inputDir, outputDir, m)
	if err != nil {
		log.Fatal("cannot parse directory ", err)
	}
}

func relativeTo(dir string, path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(dir, path)
}
