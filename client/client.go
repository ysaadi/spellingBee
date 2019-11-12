package main

import (
	"bufio"
	"fmt"
	"net"
)

func openConnection() {
	conn, err := net.Dial("tcp", "localhost:8080")
	defer conn.Close()
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Fprintf(conn, "word")
	status, err := bufio.NewReader(conn).ReadByte()

	fmt.Printf(string(status))

}
func validation(status []byte) []byte {
	if status[0] == byte(1) {
		fmt.Printf("TRUE!")
	} else if status[0] == byte(0) {
		fmt.Printf("FALSE")
	}
	return []byte{}
}

func main() {
	openConnection()
}
