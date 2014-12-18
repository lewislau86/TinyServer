///////////////////////////////////////
// TinyServer
// Golang Server Framework
//  Lewis	( lewislau86@gmail.com )
// A session manager for server
///////////////////////////////////////
package TinyServer

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"time"
)

///////////////////////////////////////
type Session struct {
	// About network
	id   uint64
	conn net.Conn
}

///////////////////////////////////////
func NewSession(id uint64, conn net.Conn) *Session {
	session := &Session{
		id:   id,
		conn: conn,
	}
	return session
}

///////////////////////////////////////
///////////////////////////////////////
// addr  ”127.0.0.1：9090“
// 有异常就关闭连接
// 分别创建接收/发送的线程通过chan来同步
func (session *Session) recvServerRoutine(tcp *TcpServer) {
	var buffer = new(bytes.Buffer)

	//	tcp.conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	defer tcp.closeSession(session)
	defer func() {
		if r := recover(); r != nil {
			log.Printf("error")
		}
	}()
	fmt.Println("wait to read")
	msg := make([]byte, MAX_PACKET_SIZE)
	ph := ProtocolHeader{}
	for {
		len, err := session.conn.Read(msg)
		if err != nil {
			break
		}
		if len > 0 {
			// format buffer
			err := binary.Write(buffer, binary.BigEndian, msg)
			CheckErr(err)
			err = binary.Read(buffer, binary.BigEndian, &ph)
			CheckErr(err)
			status := session.parseProtocal(ph)
			if false == status {
				fmt.Println("1")
				break
			}
			//	fmt.Println(string(ph.Name[0:16]))
			buffer.Reset()
		} else {
			fmt.Println(len)
		}
		time.Sleep(time.Millisecond * 100)
	}
}

///////////////////////////////////////
func (session *Session) parseProtocal(proto ProtocolHeader) bool {
	if proto.ProtocolFlag == 0x55ff {
		switch proto.ControlCode {
		case CMD_LOGIN:
			err := RespLogin()
			if err != nil {
				return false
			}
		case CMD_LOGOUT:
			err := RespLogout()
			if err != nil {
				return false
			}
		default:
			break
		}
	}
	return true
}

// EOF
///////////////////////////////////////
