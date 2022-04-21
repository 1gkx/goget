package cmd

import (
	"fmt"

	"github.com/1gkx/goget/internal/arguments"
)

func Start(opts *arguments.Options, link string) error {

	fmt.Println("Download Started")

	if opts.Version {
		fmt.Printf("Version: v0.0.1")
		return nil
	}

	fmt.Println(link)

	// if err := goget.DownloadFile(url, filename); err != nil {
	// 	return err
	// }

	fmt.Println("Download Finished")

	return nil

}
