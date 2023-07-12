package cmd

import "github.com/spf13/cobra"

type Tcp struct {
}

func (netT Tcp) Cmd() *cobra.Command {
	return &cobra.Command{
		Use:  "tcp",
		Long: "create tcp communication",
		Run:  CommonProcess,
	}
}
