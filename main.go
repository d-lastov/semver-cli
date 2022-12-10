package main

import (
	"flag"
	"fmt"
	"github.com/d-lastov/semver-go"
	"log"
	"os"
)

func main() {
	bmaj := flag.Bool("bump-major", false, "bump major version (e.g. 1.0.0 -> 2.0.0)")
	bmin := flag.Bool("bump-minor", false, "bump minor version (e.g. 1.0.0 -> 1.1.0)")
	bpat := flag.Bool("bump-patch", true, "bump patch version (e.g. 1.0.0 -> 1.0.1)")
	h := flag.Bool("h", false, "print help")
	flag.Parse()

	if h != nil && *h {
		fmt.Fprintf(os.Stderr, "Usage: %s <version>\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}

	if flag.NArg() != 1 {
		fmt.Fprint(os.Stderr, "Expected exactly 1 argument\n")
		fmt.Fprintf(os.Stderr, "Usage: %s <version>\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}

	argVersion := flag.Arg(0)
	sv, err := semver.Parse(argVersion)
	if err != nil {
		log.Fatal(err.Error())
	}

	if bmaj != nil && *bmaj {
		fmt.Print(sv.BumpMajor().String())
		os.Exit(0)
	}

	if bmin != nil && *bmin {
		fmt.Print(sv.BumpMinor().String())
		os.Exit(0)
	}

	if bpat != nil && *bpat {
		fmt.Print(sv.BumpPatch().String())
		os.Exit(0)
	}

	flag.PrintDefaults()
	os.Exit(1)
}
