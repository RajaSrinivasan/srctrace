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
}
