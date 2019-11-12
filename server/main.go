package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
)

func main() {
	ch := make(chan []byte, 1024)
	defer close(ch)
	go serve(ch)
	for {
		byteArray := <-ch
		fmt.Print(string(byteArray))
	}
}

func handleConnection(conn net.Conn, ch chan []byte) {
	defer conn.Close()
	// result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			if err == io.EOF {
				break
			}
			ch <- []byte(err.Error())
		}
		word := strings.TrimSpace(string(buf[0:n]))
		validated := ValidateWord(word)
		if validated {
			conn.Write([]byte{1})
		} else {
			conn.Write([]byte{0})
		}
		ch <- []byte(strconv.FormatBool(validated))
	}
}
func serve(ch chan []byte) {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		ch <- []byte(err.Error())
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			ch <- []byte(err.Error())
		}

		go handleConnection(conn, ch)
	}
}
