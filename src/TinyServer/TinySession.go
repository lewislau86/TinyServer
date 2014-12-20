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
	proto := ProtoHeader{}
	for {
		len, err := session.conn.Read(msg)
		if err != nil {
			break
		}
		if len > 0 {
			// format buffer
			err := binary.Write(buffer, binary.BigEndian, msg)
			CheckErr(err)
			err = binary.Read(buffer, binary.BigEndian, &proto)
			CheckErr(err)

			// 解析协议
			if proto.ProtoFlag == 0x55ff {
				switch proto.CtrlCode {
				case CMD_LOGIN:
					err := RespLogin(session, tcp, msg)
					if err != nil {
						break
					}
				case CMD_LOGOUT:
					err := RespLogout(session, tcp, msg)
					if err != nil {
						break
					}
				case CMG_HEARTBEAT:
					err := RespHeartbeat(session, tcp, msg)
					if err != nil {
						break
					}
				default:
					break
				}
				buffer.Reset()
			}
		} else {
			fmt.Println("Read len error", len)
		}
		time.Sleep(time.Millisecond * 100)
	}
}

// EOF
///////////////////////////////////////
