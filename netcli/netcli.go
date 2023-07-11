package main

import (
	"fmt"
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
		Use:   "netcli",
		Long:  "like netasist for a network tool",
		Short: "ddd",
		Run: func(cmd *cobra.Command, args []string) {
			// 传入的command为
			fmt.Println(args)
		},
	}

	root.AddCommand(cli.Tcp{}.Cmd())
	root.AddCommand(cli.TcpServer{}.Cmd())
	root.AddCommand(cli.Udp{}.Cmd())

	err := root.Execute()
	if err != nil {
		os.Exit(1)
	}
}
