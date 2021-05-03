package goget

import (
	"fmt"
	"io"
	"net/http"
)

type Responce struct {
	body io.ReadCloser
}

func Fetch(url string) *Responce {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil
	}
	return &Responce{body: resp.Body}
}

func (r *Responce) Close() {
	r.body.Close()
}
