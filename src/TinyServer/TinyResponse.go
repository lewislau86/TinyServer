///////////////////////////////////////
// TinyServer
// Golang Server Framework
//  Lewis	( lewislau86@gmail.com )
///////////////////////////////////////

package TinyServer

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

///////////////////////////////////////

func RespLogin(session *Session, tcp *TcpServer, msg []byte) error {
	var loginBuf = new(bytes.Buffer)
	pLogin := ProtoLogin{}
	fmt.Println("RespLogin")

	err := binary.Write(loginBuf, binary.BigEndian, msg)
	CheckErr(err)
	err = binary.Read(loginBuf, binary.BigEndian, &pLogin)
	CheckErr(err)
	userLen := GetRuneStrLen(pLogin.Name[:])
	// modified by liuhaiping
	// 如果密码使用hash值，这里就是固定长度
	pwdLen := GetByteStrLen(pLogin.Passwd[:])
	ret := tcp.mysql.CheckLogin(string(pLogin.Name[:userLen]), string(pLogin.Passwd[:pwdLen]))

	// replay msg
	pRep := ReplyMsg{}
	var replyBuffer = new(bytes.Buffer)
	pRep.Header = pLogin.Header
	pRep.Result = 0
	if ret {
		pRep.Result = 1
	} else {
		pRep.Result = 0
	}
	err = binary.Write(replyBuffer, binary.BigEndian, pRep)
	session.conn.Write(replyBuffer.Bytes())
	return nil
}

///////////////////////////////////////

func RespLogout(session *Session, tcp *TcpServer, buffer []byte) error {
	fmt.Println("RespLogout")
	return nil
}

///////////////////////////////////////

func RespHeartbeat(session *Session, tcp *TcpServer, buffer []byte) error {
	fmt.Println("RespHeartbeat")
	return nil
}

// EOF
///////////////////////////////////////
