package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	//"io/ioutil"
	"net"
	"time"
)

const (
	CMD_LOGIN  = 0xeef0
	CMD_LOGOUT = 0xeef1
	CMG_EWWW   = 0xeef2
)

type ProtocolHeader struct {
	ProtocolSize uint8
	ProtocolFlag uint16
	ControlCode  uint16
	Name         [64]rune
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	var ph ProtocolHeader
	var buffer bytes.Buffer

	ph.ProtocolSize = 4
	ph.ProtocolFlag = 0x55ff
	ph.ControlCode = CMD_LOGIN
	copy(ph.Name[:], []rune("abcde"))
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:18081")
	CheckErr(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	CheckErr(err)
	err = binary.Write(&buffer, binary.BigEndian, ph)
	CheckErr(err)
	fmt.Println(buffer.Len())
	fmt.Println("1")
	_, err = conn.Write(buffer.Bytes())
	CheckErr(err)
	//result, err := ioutil.ReadAll(conn)
	//fmt.Println(string(result))

	///////////////////////////////////////
	time.Sleep(time.Second * 1)
	fmt.Println("2")
	fmt.Println(buffer.Len())
	buffer.Reset()
	ph.ControlCode = CMD_LOGOUT
	err = binary.Write(&buffer, binary.BigEndian, ph)
	_, err = conn.Write(buffer.Bytes())
	CheckErr(err)
	conn.Close()
}
