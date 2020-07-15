package scuff

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func ScanDir(wd string) {
	dir, err := ioutil.ReadDir(wd)
	if err != nil {
		fmt.Println("can't access working directory", "wd:", wd, err)
	}
	for _, info := range dir {
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
			fmt.Println("found scuff", info.Name())
			Of(filepath.Join(wd, info.Name())).Fold()
		}
	}
}
