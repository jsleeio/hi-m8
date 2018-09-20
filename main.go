package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	path := flag.String("path", "/", "specify a path to respond to")
	message := flag.String("message", "hi m8", "specify a message to be returned to clients")
	listen := flag.String("listen", ":3000", "[address]:port to bind to")
	delay := flag.Duration("delay", 0, "sleep this duration before responding")
	flag.Parse()
	http.HandleFunc(*path, func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(*delay)
		io.WriteString(w, *message)
	})
	log.Fatal(http.ListenAndServe(*listen, nil))
}
