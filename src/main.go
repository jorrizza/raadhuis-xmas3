package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

var remote string
var q chan []byte

func parseColor(c string) []byte {
	if len(c) > 0 {
		i, err := strconv.ParseInt(c, 16, 32)

		if err == nil {
			return []byte{
				byte((i >> 16) & 0xFF),
				byte((i >> 8) & 0xFF),
				byte(i & 0xFF),
			}
		}
	}

	return []byte{0xFF, 0x00, 0x00}
}

func setColor(color []byte) {
	conn, err := net.Dial("tcp", remote)

	if err != nil {
		log.Print(err.Error())
		return
	}

	conn.Write([]byte{0x73})
	conn.Write(color)
	conn.Write([]byte{0x00})
	conn.Close()
}

func queueColor(w http.ResponseWriter, r *http.Request) {
	q <- parseColor(r.FormValue("value"))
}

func queue() {
	for {
		newColor := <-q

		setColor(newColor)

		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	q = make(chan []byte)

	port := flag.String("p", "localhost:4321", "socket to listen on")
	public := flag.String("d", "public", "directory of public files")
	flag.StringVar(&remote, "c", "localhost:9000", "host to connect to")

	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*public)))
	http.HandleFunc("/set-color", queueColor)

	go queue()

	log.Fatal(http.ListenAndServe(*port, nil))
}
