///////////////////////////////////////
// TinyServer
// Golang Server Framework
//  Lewis	( lewislau86@gmail.com )
///////////////////////////////////////
package TinyServer

import (
	"fmt"
	"net"
)

type Session struct {
	// About network
	id   uint32
	conn net.Conn
}

func NewSession(id uint64, conn net.Conn) *Session {
	session := &Session{
		conn: conn,
	}
	return session
}
