package main

import (
	"flag"
	"fmt"
	lab2 "github.com/whole-lotta-go/lab_2"
	"io"
	"os"
	"strings"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "File to read from")
	outputFile      = flag.String("o", "", "File to write to")
)

func main() {
	flag.Parse()
	if (*inputExpression == "" && *inputFile == "") || (*inputExpression != "" && *inputFile != "") {
		flag.Usage()
		os.Exit(1)
	}

	var input io.Reader
	if *inputExpression != "" {
		input = strings.NewReader(*inputExpression)
	} else if *inputFile != "" {
		file, err := os.Open(*inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	}

	var output io.Writer
	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		output = file
	} else {
		output = os.Stdout
	}

	handler := &lab2.ComputeHandler{Reader: input, Writer: output}

	if err := handler.Compute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error computing expression: %v\n", err)
		os.Exit(1)
	}
}
