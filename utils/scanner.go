package utils

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func FileInputScanner(filename string) *bufio.Scanner {
	path := ""
	if wd, err := os.Getwd(); err == nil { 
		path = wd + "/" + filename
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	return bufio.NewScanner(strings.NewReader(string(data)))
}