package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Failed to bind to port 8080: ", err)
	}
	defer listener.Close()

	log.Println("Server listening on 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept connection: ", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Accepted connection from: ", conn.RemoteAddr())

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Read error: ", err)
			return
		}

		_, err = conn.Write(buffer[:n])
		if err != nil {
			fmt.Println("Write error: ", err)
			return
		}
	}
}
