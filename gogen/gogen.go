package gogen

import (
	"fmt"
	"log"
	"os"
	"time"

	"../versions"
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
	fmt.Fprintf(gofile, "const BUILD_TIME = \"%s\"\n", time.Now().Format("Mon Jan 2 2006 15:04:05"))
	fmt.Fprintf(gofile, "const VERSION_MAJOR = %d\n", v.Major)
	fmt.Fprintf(gofile, "const VERSION_MINOR = %d\n", v.Minor)
	fmt.Fprintf(gofile, "const VERSION_BUILD = %d\n", v.Build)
	fmt.Fprintf(gofile, "const REPO_URL = \"%s\"\n", v.Repo)
	fmt.Fprintf(gofile, "const BRANCH_NAME = \"%s\"\n", v.Branch)
	fmt.Fprintf(gofile, "const SHORT_COMMIT_ID = \"%s\"\n", v.ShortCommitId)
	fmt.Fprintf(gofile, "const LONG_COMMIT_ID = \"%s\"\n", v.LongCommitId)
}
