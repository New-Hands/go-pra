package cmd

import (
	"github.com/spf13/cobra"
	"lstfight.cn/go-pra/netcli/model"
)

type NetCh interface {
	Start() error

	Read() []byte

	Write(d []byte)

	Stop() error
}

// create net component
func newNet(p model.NetParam) NetCh {
	switch p.Type {
	case 1:
	case 2:
	case 3:
	default:
		panic("")
	}
	return nil
}

// FlagContext 默认执行参数
var FlagContext = model.CliFlags{
	ConnectTimeOut: 3,
	Port:           0,
	ReceiveTimeOut: 3,
	Encode:         "hex",
}
var cmdMap = map[string]model.NetType{
	"tcp":       1,
	"tcpserver": 2,
	"udp":       3,
}

// CommonProcess 指令模板方式处理
func CommonProcess(cmd *cobra.Command, args []string) {
	use := cmd.Use
	netType := cmdMap[use]

	param := model.NetParam{
		Type: netType,
		Port: FlagContext.Port,
		Ip:   FlagContext.Ip,
	}

	net := newNet(param)
	err := net.Start()
	if err != nil {
		return
	}

}
