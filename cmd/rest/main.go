package main

import (
	"flag"
	"fmt"

	"Sharykhin/buffstream-questionnaire/http"
)

var (
	addr string
)

func main() {
	flag.StringVar(&addr, "addr", ":8000", "web server addr")
	flag.Parse()

	fmt.Printf("Run a web server on %s\n", addr)
	http.ListenAndServe(addr)
}
