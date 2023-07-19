package model

import (
	"encoding/hex"
	"strconv"
)

type NetType uint8

const (
	TCP       NetType = 1
	TcpServer NetType = 2
	UDP       NetType = 3
)

type NetParam struct {
	Ip             string
	Port           int
	ListenPort     int
	ConnectTimeOut int
	ReceiveTimeOut int
	Type           NetType
}

type MsgForm struct {
	Ip   string
	Port int
	Data []byte
}

func (msg *MsgForm) ToHexString() string {
	toString := hex.EncodeToString(msg.Data)
	return "(" + msg.Ip + ":" + strconv.Itoa(msg.Port) + "):" + toString
}

func (msg *MsgForm) ToString() string {
	return "(" + msg.Ip + ":" + strconv.Itoa(msg.Port) + "):" + string(msg.Data)
}

type MsgTo struct {
	Ip   string
	Port int
	Data []byte
}
