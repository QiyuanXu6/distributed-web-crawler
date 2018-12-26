package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
	"web-crawler/rpc"
)

func main()  {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	var result float64
	client := jsonrpc.NewClient(conn)
	err = client.Call("DemoService.Div", rpcdemo.Args{1, 2}, &result)
	fmt.Println(result, err)
}
