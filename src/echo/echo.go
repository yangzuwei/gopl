// echo project echo.go
package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

//func handleConn(c net.Conn) {
//	io.Copy(c, c)
//	c.Close()
//}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		handleConn(conn)
	}
}
