package main

import (
	"fmt"
	"log"
	"net"
)

// run "02-read"
// run "06-dial-write"

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	fmt.Fprintln(conn, "I dialed you.")
}
