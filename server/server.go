package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	check(err)
	conn, err := ln.Accept()
	check(err)
	message := make(chan string)

	for {
		go asyncWrite(conn, message)
		go asyncRead(conn, message)
		fmt.Print(<-message)
	}
}
func asyncRead(conn net.Conn, message chan string) {
	reader, err := bufio.NewReader(conn).ReadString('\n')
	check(err)
	message <- ("Message from client: " + reader)
}
func asyncWrite(conn net.Conn, message chan string) {
	reader, err := bufio.NewReader(os.Stdin).ReadString('\n')
	check(err)
	conn.Write([]byte(reader))
	message <- ""
}
