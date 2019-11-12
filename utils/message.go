package utils

import (
	"bufio"
	"net"
)

func parseMessage(conn net.Conn) (string, []byte) {
	buff := bufio.NewReader(conn)
	buff.Readbytes()
}
