package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body) //1st arg uses writer interface while resp.Body implements reader interface
}

//custom implementation of the writer interface used above
// the writer interface has a func called Write, we gonna modify that

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println("Just processed: ", len(bs))
	return len(bs), nil
}
