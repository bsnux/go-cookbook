package main

import (
	"io/ioutil"
	"log"
	"net"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "www.google.com:80")
	if err != nil {
		log.Println(err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Println(err)
	}

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	if err != nil {
		log.Println(err)
	}

	result, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(result))
}
