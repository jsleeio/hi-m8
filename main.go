package main

import (
	"io"
	"net/http"
  "flag"
  "log"
)

func main() {
  path    := flag.String("path","/","specify a path to respond to")
  message := flag.String("message","hi m8","specify a message to be returned to clients")
  listen  := flag.String("listen",":3000","[address]:port to bind to")
  flag.Parse()
	http.HandleFunc(*path, func(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, *message)
  })
	log.Fatal(http.ListenAndServe(*listen, nil))
}


