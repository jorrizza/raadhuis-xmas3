package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var remote string
var q chan []byte
var lastColor []byte
var colorMutex sync.Mutex

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

	return lastColor
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

	colorMutex.Lock()
	copy(lastColor, color)
	colorMutex.Unlock()
}

func getColor(w http.ResponseWriter, r *http.Request) {
	colorMutex.Lock()
	fmt.Fprintf(w, "%02X%02X%02X", lastColor[0], lastColor[1], lastColor[2])
	colorMutex.Unlock()
}

func queueColor(w http.ResponseWriter, r *http.Request) {
	q <- parseColor(r.FormValue("value"))
}

func initializeColor() {
	conn, err := net.Dial("tcp", remote)

	if err != nil {
		log.Print(err.Error())
		return
	}

	conn.Write([]byte{0xDE, 0xAD, 0xBA, 0xBE, 0x00})
	time.Sleep(100 * time.Millisecond)
	_, err = bufio.NewReader(conn).Read(lastColor)

	if err != nil {
		log.Print(err.Error())
		return
	}

	fmt.Printf("initialized with color #%02X%02X%02X\n",
		lastColor[0], lastColor[1], lastColor[2])
}

func queue() {
	for {
		newColor := <-q

		setColor(newColor)

		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	q = make(chan []byte, 120)
	lastColor = make([]byte, 3)

	port := flag.String("p", "localhost:4321", "socket to listen on")
	public := flag.String("d", "public", "directory of public files")
	flag.StringVar(&remote, "c", "localhost:9000", "host to connect to")

	go queue()

	flag.Parse()

	initializeColor()

	http.Handle("/", http.FileServer(http.Dir(*public)))
	http.HandleFunc("/set-color", queueColor)
	http.HandleFunc("/get-color", getColor)

	log.Fatal(http.ListenAndServe(*port, nil))
}
