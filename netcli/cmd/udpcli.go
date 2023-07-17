package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"lstfight.cn/go-pra/netcli/model"
	"net"
)

type Udp struct {
	netParam model.NetParam
	conn     *net.UDPConn
	DefAddr  *net.UDPAddr
}

func UdpCmd() *cobra.Command {
	return &cobra.Command{
		Use:  "udp",
		Long: "create tcp communication",
		Run:  CommonProcess,
	}
}

func (netU *Udp) Start() error {
	port := netU.netParam.Port
	lPort := netU.netParam.ListenPort
	if lPort < 1 {
		lPort = port
	}
	// 后续使用可修改默认发送地址
	netU.DefAddr = &net.UDPAddr{
		IP:   net.ParseIP(netU.netParam.Ip),
		Port: port,
	}
	lAddr := &net.UDPAddr{
		Port: lPort,
	}

	udp, err := net.ListenUDP("udp", lAddr)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	netU.conn = udp
	return nil
}

func (netU *Udp) Read() []byte {
	// get data
	data := make([]byte, 128)
	udp, u, err := netU.conn.ReadFromUDP(data)
	fmt.Println(u)
	if err != nil {
		fmt.Println(err)
	}
	if udp > 0 {
		return data[:udp]
	}

	return nil
}

func (netU *Udp) Write(d []byte) {
	// 指定remote udp
	write, err := netU.conn.Write(d)
	if err != nil {
		return
	}
	fmt.Printf("write to %d bytes\n", write)
}

func (netU *Udp) Stop() error {
	err := netU.conn.Close()
	if err != nil {
		return err
	}
	return nil
}
