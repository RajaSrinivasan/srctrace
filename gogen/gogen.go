package gogen

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/RajaSrinivasan/srctrace/versions"
)

type GoGen int

// Generate (v versionss.Version, filename string)
//     generates a C header file with the given revision information
func (gg GoGen) Generate(v versions.Version, filename string) {

	gofilename := filename + ".go"
	gofile, err := os.Create(gofilename)
	if err != nil {
		log.Printf("%v creating %s\n", err, gofilename)
		return
	}
	defer gofile.Close()

	fmt.Fprintf(gofile, "package %s\n", filename)
	fmt.Fprintln(gofile, "// Go package generator")
	fmt.Fprintf(gofile, "// File: %s.h\n", filename)
	fmt.Fprintf(gofile, "const buildTime = \"%s\"\n", time.Now().Format("Mon Jan 2 2006 15:04:05"))
	fmt.Fprintf(gofile, "const versionMajor = %d\n", v.Major)
	fmt.Fprintf(gofile, "const versionMinor = %d\n", v.Minor)
	fmt.Fprintf(gofile, "const versionBuild = %d\n", v.Build)
	fmt.Fprintf(gofile, "const repoURL = \"%s\"\n", v.Repo)
	fmt.Fprintf(gofile, "const branchName = \"%s\"\n", v.Branch)
	fmt.Fprintf(gofile, "const shortCommitId = \"%s\"\n", v.ShortCommitId)
	fmt.Fprintf(gofile, "const longCommitId = \"%s\"\n", v.LongCommitId)
	fmt.Fprintf(gofile, "const assignedTags = \"%s\"\n", v.Tags)
}
