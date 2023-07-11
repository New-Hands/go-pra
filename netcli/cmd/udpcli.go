package cmd

import "github.com/spf13/cobra"

type Udp struct {
}

func (netT Udp) Cmd() *cobra.Command {
	return &cobra.Command{
		Use: "udp",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
}
