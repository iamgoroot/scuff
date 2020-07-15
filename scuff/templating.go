package scuff

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"text/template"
)



func (f *scuff) execute(inputDir string, outputDir string, file string) error {
	outPath, err := filepath.Rel(inputDir, file)
	if err != nil {
		outPath = file
	}
	outPath = filepath.Join(outputDir, f.processText(outPath))
	if !filepath.IsAbs(outPath) {
		outPath = filepath.Join(f.Config.Location, outPath)
	}
	os.MkdirAll(filepath.Dir(outPath), os.ModePerm)
	log.Println("making", outPath)
	outfile, err := os.Create(outPath)
	defer outfile.Close()
	if handledError("creating file", err) {
		return err
	}
	t, err := f.makeTemplate(filepath.Base(file)).ParseFiles(file)
	if handledError("parsing file", err) {
		return err
	}
	return t.Execute(outfile, &f.AsMap)
}

func (f *scuff) makeTemplate(name string) *template.Template {
	return template.New(name).Delims(f.Config.Delim.Left, f.Config.Delim.Right)
}

func (f *scuff) processText(text string) string {
	t, err := f.makeTemplate("text").Parse(text)
	if handledError("parsing error", err, text) {
		return text
	}
	b := &bytes.Buffer{}
	err = t.Execute(b, &f.AsMap)
	if handledError("failed template processing", err, text) {
		return text
	}
	return b.String()
}
