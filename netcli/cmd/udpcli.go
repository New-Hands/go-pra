package cmd

import "github.com/spf13/cobra"

type Udp struct {
}

func UdpCmd() *cobra.Command {
	return &cobra.Command{
		Use:  "udp",
		Long: "create tcp communication",
		Run:  CommonProcess,
	}
}

func (netT *Udp) Start() error {

	return nil
}

func (netT *Udp) Read() []byte {

	return nil
}

func (netT *Udp) Write(d []byte) {

}

func (netT *Udp) Stop() error {
	return nil
}
