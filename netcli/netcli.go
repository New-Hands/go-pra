package main

import (
	"github.com/spf13/cobra"
	cli "lstfight.cn/go-pra/netcli/cmd"
	"os"
)

// NetCmd receive NetCmd to get cobarCommand
type NetCmd interface {
	// Cmd get cobra Command
	Cmd() *cobra.Command

	// Execute 指令处理
	Execute()
}

func main() {
	root := cobra.Command{
		Use:     "netcli",
		Long:    "like netasist for a network tool",
		Short:   "ddd",
		Version: "1.0.0",
	}

	// child command
	tcp := cli.Tcp{}.Cmd()
	root.AddCommand(tcp)
	tcpServer := cli.TcpServer{}.Cmd()
	root.AddCommand(tcpServer)
	root.AddCommand(cli.Udp{}.Cmd())

	// exec root command
	err := root.Execute()
	if err != nil {
		os.Exit(1)
	}
}
