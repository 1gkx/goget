package goget

import (
	"fmt"
)

func DownloadFile(url string, filepath string) error {

	out := CreateFile(filepath)
	defer out.Close()

	res := Fetch(url)
	defer res.Close()

	out.Save(res.body)
	fmt.Println()

	return nil
}
