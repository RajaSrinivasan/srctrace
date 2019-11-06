package main

import (
	"log"
	"os"

	"./adagen"
	"./cgen"
	"./versions"
	"github.com/akamensky/argparse"
)

var VERBOSE bool

var VERSION versions.Version
var outputFile string

func ProcessCommandLine() {
	VERBOSE = true
	parser := argparse.NewParser("srctrace", "generate source trace info")
	m := parser.Int("m", "major", &argparse.Options{Help: "Major version", Default: 0})
	out := parser.String("o", "output", &argparse.Options{Help: "Output file base name", Default: "versions"})

	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		log.Print(parser.Usage(err))
	}

	VERSION.Major = *m
	outputFile = *out
	if VERBOSE {
		VERSION.Show()
	}
}

func main() {
	log.Printf("srctrace\n")
	ProcessCommandLine()

	cg := cgen.CGen(1)
	cg.Generate(VERSION, outputFile)
	ag := adagen.AdaGen(1)
	ag.Generate(VERSION, outputFile)
}
