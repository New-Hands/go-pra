package cmd

import (
	"github.com/spf13/cobra"
	"net"
)

type TcpServer struct {
}

func (netT TcpServer) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:  "tcpserver",
		Long: "create tcpserver communication",
		Run:  CommonProcess,
	}
}

func (netT TcpServer) Start() error {
	net.Listen("tcp", "127.0.0.1:6001")
	return nil
}

func (netT TcpServer) Read() []byte {

	return nil
}

func (netT TcpServer) Write(d []byte) {

}
