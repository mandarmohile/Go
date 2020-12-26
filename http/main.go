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
		fmt.Println("Error occured while fetching data ", err)
		os.Exit(1)
	}

	//io.Copy(os.Stdout, resp.Body)
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(b []byte) (int, error) {
	fmt.Println(string(b))
	fmt.Println("Bytes Written : ", len(b))
	return len(b), nil
}
