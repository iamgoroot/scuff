package scuff

import (
	"log"
	"os"
)

func handledError(text string, err error, etc ...interface{}) bool {
	if err != nil {
		justLog(text, err, etc)
	}
	return err != nil
}

func justLog(text string, err error, etc ...interface{}) {
	if err == nil {
		return
	}
	log.Println(text)
	log.Println(err)
	log.Println(etc...)
}

func isDir(path string) bool {
	if s, err := os.Stat(path); err == nil {
		return s.IsDir()
	}
	return false
}
