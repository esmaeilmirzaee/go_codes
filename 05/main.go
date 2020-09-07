package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, resp.Body)
}
