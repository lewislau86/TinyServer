///////////////////////////////////////
// TinyServer
// Golang Server Framework
//  Lewis	( lewislau86@gmail.com )
///////////////////////////////////////
package main

import (
	"TinyServer"
	//"fmt"
	"time"
)

///////////////////////////////////////

func main() {
	var Server TinyServer.TcpServer
	TinyServer.StartHttpServer()
	Server.CreateTcpServer("Server", ":18081")
	Server.Start()
	time.Sleep(time.Minute * 10)
}

//EOF
///////////////////////////////////////
