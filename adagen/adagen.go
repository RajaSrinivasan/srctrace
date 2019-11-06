package adagen

import (
	"fmt"

	"../versions"
)

type AdaGen int

func (a AdaGen) Generate(v versions.Version, filename string) {
	fmt.Printf("package %s is\n", filename)
	fmt.Printf("   VERSION_MAJOR : constant := %d ;\n", v.Major)
	fmt.Printf("end %s ;\n", filename)
}
