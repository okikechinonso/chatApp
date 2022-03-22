package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Blog struct {
	member map[net.Addr]*User
}

type User struct {
	conn net.Conn
	name string
}

type Message struct {
	msg  string
	user *User
}

var MessageCmd = make(chan *Message)

func main() {
	listener, err := net.Listen("tcp", ":4000")
	if err != nil {
		panic(err)
	}
	go Mesage()
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go Handler(conn)
	}

}

var blog = &Blog{
	member: make(map[net.Addr]*User),
}

func Handler(conn net.Conn) {
	log.Println("New Client Connected")
	user := &User{
		conn: conn,
		name: "No name",
	}
	blog.member[user.conn.RemoteAddr()] = user
	for {
		var command = bufio.NewReader(conn)
		input, err := command.ReadString('\n')
		if err != nil {
			continue
		}
		MessageCmd <- &Message{msg: input, user: user}
	}
}

func Mesage() {
	for msg := range MessageCmd {
		for i, v := range blog.member {
			if i != msg.user.conn.RemoteAddr() {
				v.conn.Write([]byte(fmt.Sprintf("%s : %v", v.name, msg.msg)))
			}
		}

	}
}
