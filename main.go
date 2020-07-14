package main

import (
	"flag"
	"fmt"
	"github.com/iamgoroot/scuff/scuff"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("can't determine working directory", "wd:", wd, err)
	}
	fmt.Println("scanning scuff files in ", wd)
	flag.Parse()
	if len(flag.Args()) == 0 {
		scanDir(err, wd)
	}

	for _, file := range flag.Args() {
		fmt.Println("Folding argument files")
		scuff.Fold(filepath.Join(wd, file))
	}
}

func scanDir(err error, wd string) {
	dir, err := ioutil.ReadDir(wd)
	if err != nil {
		fmt.Println("can't access working directory", "wd:", wd, err)
	}
	for _, info := range dir {
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
			fmt.Println("found scuff", info.Name())
			scuff.Fold(filepath.Join(wd, info.Name()))
		}
	}
}
