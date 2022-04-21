package arguments

import (
	"flag"
)

var link string

type Options struct {
	Version bool
	Output  string
}

var (
	version = flag.Bool("v", false, "version")
	output  = flag.String("o", "", "usage")
)

func Parse() (*Options, string) {
	flag.Parse()
	args := flag.CommandLine.Args()
	if len(args) > 0 {
		link = args[0]
	}
	return &Options{
		Version: *version,
		Output:  *output,
	}, link
}
