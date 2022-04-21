package cmd

import (
	"fmt"

	"github.com/1gkx/goget/internal/arguments"
	"github.com/1gkx/goget/internal/goget"
)

func Start(opts *arguments.Options, link string) error {

	fmt.Println("Download Started")

	if opts.Version {
		fmt.Printf("Version: v0.0.1")
		return nil
	}
	if opts.Output != "" && link != "" {
		if err := goget.DownloadFile(link, opts.Output); err != nil {
			return err
		}
	}

	fmt.Println(link)

	fmt.Println("Download Finished")

	return nil

}
