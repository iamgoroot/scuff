package scuff

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func (scuff *scuff) walk() {
	inputDir, outputDir := scuff.resolvePaths()
	filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if s, err := os.Stat(path); err == nil {
			if !s.IsDir() {
				fmt.Println("found template:", path)
				scuff.execute(inputDir, outputDir, path)
			}
		}
		return nil
	})
}

func (scuff *scuff) execute(inputDir string, outputDir string, file string) {
	outPath, err := filepath.Rel(inputDir, file)
	if err != nil {
		outPath = file
	}
	outPath = filepath.Join(outputDir, scuff.processText(outPath))
	if !filepath.IsAbs(outPath) {
		outPath = filepath.Join(scuff.Config.Location, outPath)
	}
	os.MkdirAll(filepath.Dir(outPath), os.ModePerm)
	outfile, err := os.Create(outPath)
	defer outfile.Close()
	if handledError("creating file", err) {
		return
	}
	t, err := scuff.makeTemplate(filepath.Base(file)).ParseFiles(file)
	if handledError("parsing file", err) {
		return
	}
	err = t.Execute(outfile, &scuff.AsMap)
	handledError("error generating file", err)
}

func (scuff *scuff) makeTemplate(name string) *template.Template {
	return template.New(name).Delims(scuff.Config.Delim.Left, scuff.Config.Delim.Right)
}

func (scuff *scuff) processText(text string) string {
	t, err := scuff.makeTemplate("text").Parse(text)
	if handledError("parsing error", err, text) {
		return text
	}
	b := &bytes.Buffer{}
	err = t.Execute(b, &scuff.AsMap)
	if handledError("failed template processing", err, text) {
		return text
	}
	return b.String()
}
