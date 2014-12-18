///////////////////////////////////////
// TinyServer
// Golang Server Framework
//  Lewis	( lewislau86@gmail.com )
///////////////////////////////////////
package TinyServer

import (
	"fmt"
	//"io/ioutil"
	//"bufio"
	"bytes"
	"encoding/binary"
	"log"
	"net"
	"sync/atomic"
	"time"
)

type TcpServer struct {
	name       string
	listener   *net.TCPListener
	addr       *net.TCPAddr
	createFlag bool
	// About sessions
	maxSessionId uint64
	sessionStack map[uint64]*Session
	sessionMutex sync.Mutex
	//	session  map[uint64]*Session
}

//+++++++++++++++++++++++++++++++++++++++++++++++
// public
///////////////////////////////////////
//
func (tcp *TcpServer) CreateTcpServer(name, addr string) {
	var err error
	tcp.name = name
	tcp.addr, err = net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		fmt.Println(err)
	}
	tcp.listener, err = net.ListenTCP("tcp", tcp.addr)
	if err != nil {
		fmt.Println("ListenTCP error")
		fmt.Println(err)
	} else {
		tcp.createFlag = true
		fmt.Println("Create a server successfully.")
	}
}

///////////////////////////////////////
// addr  ”127.0.0.1：9090“
func (tcp *TcpServer) Start() {
	if tcp.createFlag != true {
		fmt.Println("Server don't created")
		return
	}
	for {
		// 要在这里做 Session管理
		session, err := tcp.accept()
		if err == nil {
			go tcp.recvServerRoutine(session)
			fmt.Println("Server Success")
		} else {
			fmt.Println("Server Failed")
			break
		}

	}
}

///////////////////////////////////////
// addr  ”127.0.0.1：9090“
func (tcp *TcpServer) Stop() {
	if tcp.createFlag != true {
		return
	}
	tcp.listener.Close()
}

//-----------------------------------------------
// private
///////////////////////////////////////
func (tcp TcpServer) parseProtocal(proto ProtocolHeader) {
	if proto.ProtocolFlag == 0x55ff {
		switch proto.ControlCode {
		case CMD_LOGIN:
			RespLogin()
		case CMD_LOGOUT:
			RespLogout()
		default:
			break
		}
	}
}

///////////////////////////////////////
// addr  ”127.0.0.1：9090“
// 有异常就关闭连接
// 分别创建接收/发送的线程通过chan来同步
func (tcp *TcpServer) recvServerRoutine(session *Session) {
	var buffer = new(bytes.Buffer)

	//	tcp.conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	defer session.conn.Close()
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
			fmt.Println("read out")
			fmt.Println(err)
			break
		}
		if len > 0 {
			// format buffer
			err := binary.Write(buffer, binary.BigEndian, msg)
			CheckErr(err)
			err = binary.Read(buffer, binary.BigEndian, &ph)
			CheckErr(err)

			tcp.parseProtocal(ph)
			//	fmt.Println(string(ph.Name[0:16]))
			buffer.Reset()
		} else {
			fmt.Println(len)
		}
		time.Sleep(time.Millisecond * 100)
	}
}

///////////////////////////////////////

func (tcp *TcpServer) accept() (*Session, error) {
	conn, err := tcp.listener.Accept()
	if err != nil {
		return nil, err
	}
	session, err := tcp.createSession(conn)
	fmt.Print("*****conn****\t")
	return session, nil
}

///////////////////////////////////////

func (tcp *TcpServer) createSession(conn net.Conn) (*Session, error) {
	session := NewSession(conn)
	err := tcp.pushSession(session)
	return session, err
}

///////////////////////////////////////

func (tcp *TcpServer) pushSession(session Session) error {
	tcp.sessionMutex.Lock()
	defer tcp.sessionMutex.Unlock()
	tcp.sessionStack[session.id] = session
	return err
}

///////////////////////////////////////

func (tcp *TcpServer) popSession(session Session) error {
	tcp.sessionMutex.Lock()
	defer tcp.sessionMutex.Unlock()
	tcp.sessionStack[session.id] = session
	return err
}

//EOF
///////////////////////////////////////
