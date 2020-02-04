package adagen

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/RajaSrinivasan/srctrace/versions"
)

type AdaGen int

func (a AdaGen) Generate(v versions.Version, filename string) {

	specfilename := filename + ".ads"
	specfile, err := os.Create(specfilename)
	if err != nil {
		log.Printf("%v creating %s\n", err, specfilename)
		return
	}
	defer specfile.Close()

	fmt.Fprintf(specfile, "package %s is\n", filename)
	fmt.Fprintln(specfile, "-- Ada spec generator")
	fmt.Fprintf(specfile, "-- File: %s.ads\n", filename)
	fmt.Fprintf(specfile, "    BUILD_TIME : constant String := \"%s\" ;\n", time.Now().Format("Mon Jan 2 2006 15:04:05"))
	fmt.Fprintf(specfile, "    VERSION_MAJOR : constant := %d ;\n", v.Major)
	fmt.Fprintf(specfile, "    VERSION_MINOR : constant := %d ;\n", v.Minor)
	fmt.Fprintf(specfile, "    VERSION_BUILD : constant := %d ;\n", v.Build)
	fmt.Fprintf(specfile, "    REPO_URL : constant String := \"%s\" ;\n", v.Repo)
	fmt.Fprintf(specfile, "    BRANCH_NAME : constant String := \"%s\" ;\n", v.Branch)
	fmt.Fprintf(specfile, "    SHORT_COMMIT_ID : constant String := \"%s\" ;\n", v.ShortCommitId)
	fmt.Fprintf(specfile, "    LONG_COMMIT_ID : constant String := \"%s\" ;\n", v.LongCommitId)
	fmt.Fprintf(specfile, "end %s ;\n", filename)
}
