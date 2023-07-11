package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type TcpServer struct {
}

func (netT TcpServer) Cmd() *cobra.Command {
	return &cobra.Command{
		Use: "tcpserver",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(len(args))
			fmt.Println(args)
		},
	}
}
