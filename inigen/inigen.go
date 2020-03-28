package inigen

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/RajaSrinivasan/srctrace/versions"
)

type IniGen int

// Generate (v versionss.Version, filename string)
//     generates a C header file with the given revision information
func (ig IniGen) Generate(v versions.Version, filename string) {

	inifile, err := os.Create(filename)
	if err != nil {
		log.Printf("%v creating %s\n", err, filename)
		return
	}
	defer inifile.Close()

	fmt.Fprintf(inifile, "[versions]\n")
	fmt.Fprintf(inifile, "buildTime = \"%s\"\n", time.Now().Format("Mon Jan 2 2006 15:04:05"))
	fmt.Fprintf(inifile, "versionMajor = %d\n", v.Major)
	fmt.Fprintf(inifile, "versionMinor = %d\n", v.Minor)
	fmt.Fprintf(inifile, "versionBuild = %d\n", v.Build)
	fmt.Fprintf(inifile, "repoURL = \"%s\"\n", v.Repo)
	fmt.Fprintf(inifile, "branchName = \"%s\"\n", v.Branch)
	fmt.Fprintf(inifile, "shortCommitId = \"%s\"\n", v.ShortCommitId)
	fmt.Fprintf(inifile, "longCommitId = \"%s\"\n", v.LongCommitId)
	fmt.Fprintf(inifile, "assogmedTags = \"%s\"\n", v.Tags)
}
