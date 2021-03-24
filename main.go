package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"flag"

	"github.com/wellington/go-libsass"
)

// Usage prints how to use the command-line utility
func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

func generateCSS(inputSASS, outputCSS string) {
	r, err := os.Open(inputSASS)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create(outputCSS)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	w := bufio.NewWriter(file)

	comp, err := libsass.New(w, r)
	if err != nil {
		log.Fatal(err)
	}

	if err := comp.Run(); err != nil {
		log.Printf("Error while generating sass: %s\n", err)
	}
	w.Flush()
}

func main() {
	flag.Usage = Usage
	var help bool
	flag.BoolVar(&help, "help", false, "Print this help")
	var input string
	flag.StringVar(&input, "i", "", "Input SASS file")
	var output string
	flag.StringVar(&output, "o", "", "Output CSS file")

	flag.Parse()

	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) > 0 && strings.ToLower(argsWithoutProg[0]) == "help" {
		help = true
	}

	if help || !(input != "" && output != "") {
		Usage()
		os.Exit(2)
	}

	generateCSS(input, output)
}
