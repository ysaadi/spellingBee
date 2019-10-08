package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	ch := make(chan []byte, 1024)
	for {
		select {
		case byteArray := <-ch:
			fmt.Print(string(byteArray))
		default:
			conn, err := ln.Accept()
			if err != nil {
				fmt.Print(err)
				// handle error
			}

			go handleConnection(conn, ch)
		}
	}
}

func handleConnection(conn net.Conn, ch chan []byte) {
	defer conn.Close()
	defer close(ch)
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			ch <- []byte("why u error me")
		}
	}
	ch <- result.Bytes()
}
