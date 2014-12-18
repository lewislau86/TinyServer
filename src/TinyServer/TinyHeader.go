///////////////////////////////////////
// TinyServer
// Golang Server Framework
//  Lewis	( lewislau86@gmail.com )
// 数据包格式
// 目前只提供两种包
//		1. 登陆验证
//		2. 心跳
///////////////////////////////////////
package TinyServer

import (
	"fmt"
)

const (
	MAX_PACKET_SIZE = 512
)

const (
	CMD_LOGIN     = 0xeef0
	CMD_LOGOUT    = 0xeef1
	CMG_HEARTBEAT = 0xeef2
)

type ProtocolHeader struct {
	ProtocolSize uint8
	ProtocolFlag uint16
	ControlCode  uint16
	Name         [64]rune
}

///////////////////////////////////////
// helper function\
func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
