package main

import (
	"fmt"
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
var generator versions.Generator

func ProcessCommandLine() {

	parser := argparse.NewParser("srctrace", "generate source trace info")

	v := parser.Flag("v", "verbose", &argparse.Options{Help: "Verbose", Default: true})

	m := parser.Int("m", "major", &argparse.Options{Help: "Major version", Default: 0})
	minor := parser.Int("n", "minor", &argparse.Options{Help: "Minor version", Default: 0})
	build := parser.Int("b", "build", &argparse.Options{Help: "Build Number", Default: 999})

	lang := parser.Selector("L", "language", []string{"C", "C++", "Ada", "python", "go"}, &argparse.Options{Help: "Language to output"})
	out := parser.String("o", "output", &argparse.Options{Help: "Output file base name", Default: "versions"})

	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		log.Print(parser.Usage(err))
	}

	VERBOSE = *v
	VERSION.Major = *m
	VERSION.Minor = *minor
	VERSION.Build = *build

	outputFile = *out

	switch *lang {
	case "C":
		generator = cgen.CGen(1)
	case "Ada":
		generator = adagen.AdaGen(1)
	case "go":
		fmt.Println("go not yet supported")
		os.Exit(1)
	case "python":
		fmt.Println("python not yet supported")
		os.Exit(1)
	default:
		fmt.Printf("Language is not recognized\n")
		os.Exit(1)
	}
	if VERBOSE {
		VERSION.Show()
	}
}

func main() {
	log.Printf("srctrace\n")
	ProcessCommandLine()
	generator.Generate(VERSION, outputFile)
}
