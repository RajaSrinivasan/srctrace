package cgen

import (
	"fmt"

	"../versions"
)

type CGen int

func (cg CGen) Generate(v versions.Version, filename string) {
	fmt.Println("// C header generator")
	fmt.Printf("// File: %s.h\n", filename)
	fmt.Printf("#define VERSION_MAJOR (%d)\n", v.Major)
	fmt.Printf("#define REPO_URL \"%s\"\n", v.Repo)
	fmt.Printf("#define SHORT_COMMIT_ID \"%s\"\n", v.ShortCommitId)
	fmt.Printf("#define LONG_COMMIT_ID \"%s\"\n", v.LongCommitId)
}
