package main

import (
	"fmt"
	"net/http"
	"os"

	flag "github.com/ogier/pflag"
)

func main() {
	dir, _ := os.Getwd()
	port := flag.IntP("port", "p", 8080, "help message for flagname")

	flag.Parse()

	if flag.NArg() > 1 {
		flag.Usage()
		os.Exit(1)
	} else if flag.NArg() == 1 {
		dir = flag.Args()[0]
	}

	bindTo := fmt.Sprintf(":%d", *port)

	fmt.Printf("Serving '%s' on %s...\n", dir, bindTo)

	panic(http.ListenAndServe(bindTo, http.FileServer(http.Dir(dir))))
}
