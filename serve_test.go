package main

import (
	"log"
	"net"
	"testing"
)

func TestChat(t *testing.T) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "localhost:4000")
	if err != nil {
		log.Println(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil{
		log.Println(err)
	}
	Handler(conn)
}
