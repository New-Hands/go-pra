package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"lstfight.cn/go-pra/netcli/model"
	"net"
	"strconv"
)

type Tcp struct {
	NetParam model.NetParam
	conn     net.Conn
}

func TcpCmd() *cobra.Command {
	return &cobra.Command{
		Use:  "tcp",
		Long: "create tcp communication",
		Run:  CommonProcess,
	}
}

func (netT *Tcp) Start() error {
	dial, err := net.Dial("tcp", netT.NetParam.Ip+":"+strconv.Itoa(netT.NetParam.Port))
	if err != nil {
		return err
	}
	netT.conn = dial
	return nil
}

func (netT *Tcp) Read() []byte {
	all, err := io.ReadAll(netT.conn)
	if err != nil && nil != io.EOF {
		panic(err)
	}

	return all
}

func (netT *Tcp) Write(d []byte) {
	write, err := netT.conn.Write(d)
	if err != nil && err != io.EOF {
		panic(err)
	}
	fmt.Printf("write %d bytes", write)
}

func (netT *Tcp) Stop() error {
	return nil
}
