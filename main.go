package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	directory := flag.String("d", ".", "directory to serve")
	port := flag.String("p", "8888", "port of server")
	flag.Parse()

	log.Printf("Serving %s on port %s", *directory, *port)

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
