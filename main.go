package main

import (
	"golang.org/x/tools/go/packages"
)

// Depend on tools, which has a transitive dependency on goldmark, but not
// from 'packages'.
var _ = packages.NeedName

func main() {}
