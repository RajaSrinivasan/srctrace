package versions

import (
	"fmt"
)

type Version struct {
	Major int
	Minor int
	Build int
}

type Generator interface {
	Generate(v Version, filename string)
}

func (v Version) Show() {
	fmt.Printf("Major %d\n", v.Major)
}
