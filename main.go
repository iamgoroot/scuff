package main

import (
	"flag"
	"fmt"
	"github.com/iamgoroot/scuff/scuff"
	"os"
	"path/filepath"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("can't determine working directory", "wd:", wd, err)
	}
	fmt.Println("scanning scuff files in ", wd)
	flag.Parse()
	if len(flag.Args()) == 0 {
		scuff.ScanDir(wd)
	}

	for _, file := range flag.Args() {
		fmt.Println("Folding argument files")
		scuff.Of(filepath.Join(wd, file))
	}
}
