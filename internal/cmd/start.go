package cmd

import (
	"fmt"

	"github.com/1gkx/goget/internal/arguments"
	"github.com/1gkx/goget/internal/goget"
)

func Start(opts *arguments.Options, links []string) error {

	fmt.Println("Download Started")

	if opts.Help {
		fmt.Printf("Help: help string")
		return nil
	}

	if opts.Version {
		fmt.Printf("Version: v0.0.1")
		return nil
	}
	if opts.Output != "" && len(links) == 1 {
		if err := goget.DownloadFile(links[0], opts.Output); err != nil {
			return err
		}
	}

	fmt.Println(links)

	fmt.Println("Download Finished")

	return nil

}
