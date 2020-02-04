package main

import (
	"fmt"
	"log"
	"os"

	"github.com/RajaSrinivasan/srctrace/adagen"
	"github.com/RajaSrinivasan/srctrace/cgen"
	"github.com/RajaSrinivasan/srctrace/gogen"
	"github.com/RajaSrinivasan/srctrace/inigen"
	"github.com/RajaSrinivasan/srctrace/versions"
	"github.com/akamensky/argparse"
)

var VERBOSE bool

var VERSION versions.Version
var outputFile string
var generator versions.Generator

func ProcessCommandLine() {

	parser := argparse.NewParser("srctrace", "generate source trace info")

	v := parser.Flag("v", "verbose", &argparse.Options{Help: "Verbose", Default: false})
	ver := parser.Flag("r", "report-version", &argparse.Options{Help: "report version", Default: false})

	m := parser.Int("m", "major", &argparse.Options{Help: "Major version", Default: 0})
	minor := parser.Int("n", "minor", &argparse.Options{Help: "Minor version", Default: 0})
	build := parser.Int("b", "build", &argparse.Options{Help: "Build Number", Default: 999})

	lang := parser.Selector("L", "language", []string{"go", "C", "Ada", "ini"}, &argparse.Options{Help: "Language to output"})
	out := parser.String("o", "output", &argparse.Options{Help: "Output file base name", Default: "revisions"})

	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		log.Print(parser.Usage(err))
	}

	if *ver {
		versions.Report()
		os.Exit(0)
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
		generator = gogen.GoGen(1)
	case "ini":
		generator = inigen.IniGen(1)
	default:
		fmt.Printf("Language is not recognized\n")
		os.Exit(1)
	}

}

func main() {
	log.Printf("srctrace\n")
	ProcessCommandLine()
	rem := versions.GetRemoteURL(".")
	VERSION.Repo = rem
	br := versions.GetBranchWithHead(".")
	VERSION.Branch = br
	cid, lcid := versions.GetCommitId(".", "qqq")
	VERSION.ShortCommitId = cid
	VERSION.LongCommitId = lcid
	generator.Generate(VERSION, outputFile)

}
