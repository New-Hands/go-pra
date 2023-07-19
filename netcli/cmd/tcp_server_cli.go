package cmd

import (
	"github.com/spf13/cobra"
	"lstfight.cn/go-pra/netcli/model"
	"net"
)

type TcpServer struct {
	NetParam model.NetParam
	connMap  map[string]net.Conn
}

func TcpServerCmd() *cobra.Command {
	return &cobra.Command{
		Use:  "tcpserver",
		Long: "create tcpserver communication",
		Run:  CommonProcess,
	}
}

func (netT *TcpServer) Start() error {
	port := netT.NetParam.Port
	// 地址参数
	addr := net.TCPAddr{
		Port: port,
	}
	server, err := net.ListenTCP("tcp", &addr)

	if err != nil {
		return err
	}
	for {
		_, err := server.Accept()
		if err != nil {
			return err
		}
	}
}

func (netT *TcpServer) Read() (*model.MsgForm, error) {

	return nil, nil
}

func (netT *TcpServer) Write(to *model.MsgTo) error {
	// 获取到 client socket 进行发送

	return nil
}

func (netT *TcpServer) Stop() error {
	return nil
}
