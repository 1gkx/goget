package cmd

import (
	"fmt"
	"os"

	"github.com/1gkx/goget/internal/goget"
)

func Start() error {

	fmt.Println("Download Started")

	url := os.Args[1]
	filename := os.Args[2]
	if err := goget.DownloadFile(url, filename); err != nil {
		return err
	}

	fmt.Println("Download Finished")

	return nil

}
