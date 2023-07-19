package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"lstfight.cn/go-pra/netcli/model"
	"net"
)

type Udp struct {
	NetParam model.NetParam
	conn     *net.UDPConn
	DefAddr  *net.UDPAddr
	udpCache []byte
}

func UdpCmd() *cobra.Command {
	return &cobra.Command{
		Use:  "udp",
		Long: "create udp communication",
		Run:  CommonProcess,
	}
}

func (netU *Udp) Start() error {
	port := netU.NetParam.Port
	lPort := netU.NetParam.ListenPort
	if lPort < 1 {
		lPort = port
	}
	// 后续使用可修改默认发送地址
	netU.DefAddr = &net.UDPAddr{
		IP:   net.ParseIP(netU.NetParam.Ip),
		Port: port,
	}
	lAddr := &net.UDPAddr{
		Port: lPort,
	}

	// udp listen
	udp, err := net.ListenUDP("udp", lAddr)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	netU.conn = udp
	netU.udpCache = make([]byte, 65535, 65535)
	return nil
}

func (netU *Udp) Read() (*model.MsgForm, error) {
	// get data
	// 并发访问有问题
	num, u, err := netU.conn.ReadFromUDP(netU.udpCache)
	if err != nil {
		fmt.Println(err)
	}

	return &model.MsgForm{
		Data: netU.udpCache[:num],
		Ip:   u.IP.String(),
		Port: u.Port,
	}, nil
}

func (netU *Udp) Write(to *model.MsgTo) error {
	// 指定remote udp
	_, err := netU.conn.WriteToUDP(to.Data, &net.UDPAddr{
		IP:   net.ParseIP(to.Ip),
		Port: to.Port,
	})
	if err != nil {
		return err
	}

	return nil
}

func (netU *Udp) Stop() error {
	err := netU.conn.Close()
	if err != nil {
		return err
	}
	return nil
}
