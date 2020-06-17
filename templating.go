package main

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func walk(in string, out string, m JsonMap) error {
	wg := &sync.WaitGroup{}
	filepath.Walk(in, func(path string, info os.FileInfo, err error) error {
		file := relativeTo(in, path)
		//log.Println("check file", path, file)
		if s, err := os.Stat(file); err == nil {
			if !s.IsDir() && s.Name() != "scuff.json" {
				log.Println("found template:", file)
				execute(in, out, file, m, wg)
			}
		}
		return nil
	})

	wg.Wait()
	return nil
}

func execute(in string, outputDir string, file string, m JsonMap, group *sync.WaitGroup) {
	group.Add(1)
	defer group.Done()
	out, err := filepath.Rel(in, file)
	if err != nil {
		out = file
	}
	log.Println("processing:", out)
	out = relativeTo(outputDir, out)
	os.MkdirAll(filepath.Dir(out), os.ModePerm)
	outfile, err := os.Create(out)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()

	t := makeTemplate(filepath.Base(file), m)
	t, err = t.ParseFiles(file)
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(outfile, m)
	if err != nil {
		log.Println(err)
	}
}

func makeTemplate(name string, m JsonMap) *template.Template {
	delims := m.UnpackMap("scuff.delim")
	left := delims.AsString("left", "")
	right := delims.AsString("right", "")
	return template.New(name).Delims(left, right)
}
