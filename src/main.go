package main

import (
	"fmt"
	"net/http"
	"html"
	"log"
	"flag"
)

func main() {
	var port = flag.Int("p", 4321, "port to listen on")

	flag.Parse()

	http.HandleFunc("/set-color", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Color: #%s", html.EscapeString(r.FormValue("value")))
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
