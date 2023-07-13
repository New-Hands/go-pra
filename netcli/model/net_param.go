package model

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
