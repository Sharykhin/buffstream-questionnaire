package main

import (
	"Sharykhin/buffstream-questionnaire/database/postgres"
	"Sharykhin/buffstream-questionnaire/http"
	"flag"
	"fmt"
)

var (
	addr string
)

func main() {
	flag.StringVar(&addr, "addr", ":8000", "web server addr")
	flag.Parse()

	postgres.NewConnection("postgres", "root", "postgres", "buffstreams", "5432")

	fmt.Printf("Run a web server on %s\n", addr)
	http.ListenAndServe(addr)
}
