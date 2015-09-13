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
	conn, err := net.Dial("tcp", "192.168.1.184:8080")
	check(err)
	message := make(chan string)

	for {
		go asyncWrite(conn, message)
		go asyncRead(conn, message)
		fmt.Print(<-message)
	}
}
func asyncRead(conn net.Conn, message chan string) {
	reader, _ := bufio.NewReader(conn).ReadString('\n')
	message <- ("Message from server: " + reader + "\n")
}
func asyncWrite(conn net.Conn, message chan string) {
	reader, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	conn.Write([]byte(reader))
	message <- ""
}
