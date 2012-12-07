package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
)

var remote string

func parseColor(c string) string {
	if len(c) > 0 {
		i, err := strconv.ParseInt(c, 16, 32)

		if err == nil {
			return fmt.Sprintf("%c%c%c", (i >> 16) & 0xFF, (i >> 8) & 0xFF, i & 0xFF)
		}
	}

	return "\xff\x00\x00"
}

func setColor(w http.ResponseWriter, r *http.Request) {
	conn, err := net.Dial("tcp", remote)

	if err != nil {
		log.Print(err.Error())
		return
	}

	fmt.Fprintf(conn, "%s%s", "s", parseColor(r.FormValue("value")))
	conn.Close()
}

func main() {
	port := flag.Int("p", 4321, "port to listen on")
	public := flag.String("d", "public", "directory of public files")
	flag.StringVar(&remote, "c", "localhost:9000", "host to connect to")

	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*public)))
	http.HandleFunc("/set-color", setColor)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
