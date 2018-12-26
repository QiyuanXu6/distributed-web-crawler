package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"web-crawler/rpc"
)

func main() {
	rpc.Register(rpcdemo.DemoService{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error %v\n", err)
		}
		go jsonrpc.ServeConn(conn)
	}
}
