package main

import (

	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"sharing/rpc/lib"
)


func main() {

	rpc.Register(new(lib.Arith)) // 注册rpc服务
	lis, err := net.Listen("tcp", "127.0.0.1:8095")
	if err != nil {
		log.Fatalln("fatal error: ", err)
	}

	fmt.Fprintf(os.Stdout, "%s", "start connection")
	//rpc.Accept(lis)
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go rpc.ServeConn(conn)
	}
}
