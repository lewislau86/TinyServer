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

	"net"
	"sync"
	"sync/atomic"
)

type TcpServer struct {
	name       string
	listener   *net.TCPListener
	addr       *net.TCPAddr
	createFlag bool
	// About sessions
	maxSessionId  uint64
	onlineSession uint32
	sessionStack  map[uint64]*Session
	sessionMutex  sync.Mutex
	mysql         TinyDatabase
}

//+++++++++++++++++++++++++++++++++++++++++++++++
// public
///////////////////////////////////////
//
func (tcp *TcpServer) CreateTcpServer(name, addr string, dataSource string) {
	var err error

	tcp.name = name
	tcp.onlineSession = 0
	tcp.maxSessionId = 0
	tcp.sessionStack = make(map[uint64]*Session)
	tcp.mysql = TinyDatabase{}
	tcp.addr, err = net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		fmt.Println(err)
	}

	tcp.listener, err = net.ListenTCP("tcp", tcp.addr)
	if err != nil {
		fmt.Println("ListenTCP error")
		fmt.Println(err)
	} else {
		tcp.mysql.Open(dataSource)
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
			go session.recvServerRoutine(tcp)
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

func (tcp *TcpServer) accept() (*Session, error) {
	conn, err := tcp.listener.Accept()
	if err != nil {
		return nil, err
	}
	fmt.Println("New connection :\t", conn.RemoteAddr().String())
	session := tcp.createSession(conn)
	return session, nil
}

///////////////////////////////////////

func (tcp *TcpServer) createSession(conn net.Conn) *Session {
	session := NewSession(atomic.AddUint64(&tcp.maxSessionId, 1), conn)
	if session != nil {
		tcp.pushSession(session)
		fmt.Println("Create session here. Online :", tcp.onlineSession)
	}
	return session
}

///////////////////////////////////////
func (tcp *TcpServer) closeSession(session *Session) error {
	tcp.popSession(session)
	fmt.Println("Close session here. Online :", tcp.onlineSession)
	return nil
}

///////////////////////////////////////

func (tcp *TcpServer) pushSession(session *Session) {
	tcp.sessionMutex.Lock()
	defer tcp.sessionMutex.Unlock()
	tcp.sessionStack[session.id] = session
	tcp.onlineSession = tcp.onlineSession + 1
}

///////////////////////////////////////

func (tcp *TcpServer) popSession(session *Session) {
	tcp.sessionMutex.Lock()
	defer tcp.sessionMutex.Unlock()
	delete(tcp.sessionStack, session.id)
	tcp.onlineSession = tcp.onlineSession - 1
}

//EOF
///////////////////////////////////////
