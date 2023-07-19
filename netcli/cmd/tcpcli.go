package cmd

import (
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

func (netT *Tcp) Read() (*model.MsgForm, error) {
	rt := netT.NetParam.ReceiveTimeOut
	if rt > 0 {
		_ = netT.conn.SetReadDeadline(time.Now().Add(time.Duration(rt) * time.Second))
	}

	bytes := make([]byte, 10240, 10240)
	rNum, err := netT.conn.Read(bytes)
	if err != nil {
		// 关闭连接
		if err == io.EOF {
			_ = netT.conn.Close()
			return nil, err
		}
		return nil, err
	}

	return &model.MsgForm{
		Data: bytes[:rNum],
		Ip:   netT.NetParam.Ip,
		Port: netT.NetParam.Port,
	}, nil
}

func (netT *Tcp) Write(to *model.MsgTo) error {
	_, err := netT.conn.Write(to.Data)
	if nil != err {
		return err
	}
	return nil
}

func (netT *Tcp) Stop() error {
	err := netT.conn.Close()
	if err != nil {
		return err
	}
	return nil
}
