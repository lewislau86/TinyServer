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
	MAX_PACKET_SIZE = 1024
)

const (
	CMD_LOGIN     = 0xeef0
	CMD_LOGOUT    = 0xeef1
	CMG_HEARTBEAT = 0xeef2
	CMG_REPLY     = 0xeef3
)

type ProtoHeader struct {
	ProtoSize uint8
	ProtoFlag uint16
	CtrlCode  uint16
}

type ProtoLogin struct {
	Header ProtoHeader
	Name   [64]rune
	Passwd [64]byte
}

type ReplyMsg struct {
	Header ProtoHeader
	Result uint8
}

///////////////////////////////////////
// helper function\
func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

///////////////////////////////////////
func GetRuneStrLen(str []rune) uint32 {
	var i uint32 = 0
	for {
		if str[i] == 0 {
			return i
		}
		i++
	}
}

///////////////////////////////////////

func GetByteStrLen(str []byte) uint32 {
	var i uint32 = 0
	for {
		if str[i] == 0 {
			return i
		}
		i++
	}
}
