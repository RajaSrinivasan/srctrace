package cgen

import (
	"fmt"
	"log"
	"os"
	"time"

	"../versions"
)

type CGen int

// Generate (v versionss.Version, filename string)
//     generates a C header file with the given revision information
func (cg CGen) Generate(v versions.Version, filename string) {

	hfilename := filename + ".h"
	hfile, err := os.Create(hfilename)
	if err != nil {
		log.Printf("%v creating %s\n", err, hfilename)
		return
	}
	defer hfile.Close()

	fmt.Fprintln(hfile, "// C header generator")
	fmt.Fprintf(hfile, "// File: %s.h\n", filename)
	fmt.Fprintf(hfile, "#define BUILD_TIME \"%s\"\n", time.Now().Format("Mon Jan 2 2006 15:04:05"))
	fmt.Fprintf(hfile, "#define VERSION_MAJOR (%d)\n", v.Major)
	fmt.Fprintf(hfile, "#define VERSION_MINOR (%d)\n", v.Minor)
	fmt.Fprintf(hfile, "#define VERSION_BUILD (%d)\n", v.Build)
	fmt.Fprintf(hfile, "#define REPO_URL \"%s\"\n", v.Repo)
	fmt.Fprintf(hfile, "#define SHORT_COMMIT_ID \"%s\"\n", v.ShortCommitId)
	fmt.Fprintf(hfile, "#define LONG_COMMIT_ID \"%s\"\n", v.LongCommitId)
}
