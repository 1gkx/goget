package arguments

import (
	"flag"
)

var link string

type Options struct {
	Version bool
}

var (
	version = flag.Bool("v", false, "version")
	Output  = flag.String("o", "", "usage")
)

func Parse() (*Options, string) {
	flag.Parse()
	args := flag.CommandLine.Args()
	if len(args) > 0 {
		link = args[0]
	}
	return &Options{
		Version: *version,
	}, link
}
