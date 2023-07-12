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

// flag option arg
// 执行shell脚本时是没有这些区分 可都视作参数
// flag 或 option

func main() {
	root := cobra.Command{
		Use:     "netcli",
		Long:    "like netAssist for a network tool",
		Short:   "ddd",
		Version: "1.0.0",
	}

	// 配置全局flg
	root.PersistentFlags().StringVarP(&cli.FlagContext.Encode, "encode", "e", "hex",
		"data encode hex x or ascii a")
	root.PersistentFlags().Uint16VarP(&cli.FlagContext.Port, "port", "p", 6111,
		"port to connect and bind if not listen port")
	root.PersistentFlags().IntVarP(&cli.FlagContext.ListenPort, "listen-port", "l", 6111,
		"port to connect and bind if not listen port")
	root.PersistentFlags().IntVarP(&cli.FlagContext.ConnectTimeOut, "con-time", "c", 3,
		"connect timeout")

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
