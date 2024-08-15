package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/nikkbh/htmllinkparser"
)

func main() {
	filename := flag.String("path", "htmls/ex2.html", "HTML file path to parse")
	flag.Parse()
	htmlFile, err := ReadHTML(*filename)
	if err != nil {
		log.Fatalf("Error reading the HTML File you specified")
		return
	}
	r := strings.NewReader(htmlFile)
	links, err := htmllinkparser.Parse(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", links)
}

func ReadHTML(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	htmlFile, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}
	return string(htmlFile), nil
}
