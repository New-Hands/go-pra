package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"lstfight.cn/go-pra/netcli/model"
	"net"
	"strconv"
	"time"
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
	rt := netT.NetParam.ReceiveTimeOut
	if rt > 0 {
		_ = netT.conn.SetReadDeadline(time.Now().Add(time.Duration(rt) * time.Second))
	}
	all, err := io.ReadAll(netT.conn)
	if err != nil && nil != io.EOF {
		fmt.Println(err)
	}

	return all
}

func (netT *Tcp) Write(d []byte) {
	write, err := netT.conn.Write(d)
	if err != nil && err != io.EOF {
		fmt.Println(err)
	}
	fmt.Printf("write %d bytes\n", write)
}

func (netT *Tcp) Stop() error {
	err := netT.conn.Close()
	if err != nil {
		return err
	}
	return nil
}
