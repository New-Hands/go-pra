package cmd

import "github.com/spf13/cobra"

type Tcp struct {
}

func TcpCmd() *cobra.Command {
	return &cobra.Command{
		Use:  "tcp",
		Long: "create tcp communication",
		Run:  CommonProcess,
	}
}

func (netT *Tcp) Start() error {
	return nil
}

func (netT *Tcp) Read() []byte {

	return nil
}

func (netT *Tcp) Write(d []byte) {

}
