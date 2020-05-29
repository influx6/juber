package main

import (
	"fmt"
	"net/http"
	"os"
)

var (
	PORT = os.Getenv("PORT")
	MESSAGE = os.Getenv("MESSAGE")
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello %s\n",MESSAGE)
}

func hello(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
	fmt.Fprintf(w, "hello %s\n",MESSAGE)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil)
}