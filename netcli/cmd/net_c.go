package cmd

import (
	"encoding/hex"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"io"
	"log"
	"lstfight.cn/go-pra/netcli/model"
	"lstfight.cn/go-pra/netcli/ui"
	"strconv"
	"strings"
)

type NetCh interface {
	Start() error

	Read() (*model.MsgForm, error)

	Write(to *model.MsgTo) error

	Stop() error
}

// create net component
func newNet(p model.NetParam) NetCh {
	switch p.Type {
	case model.TCP:
		return &Tcp{
			NetParam: p,
		}
	case model.TcpServer:
		return &TcpServer{
			NetParam: p,
		}
	case model.UDP:
		return &Udp{
			NetParam: p,
		}
	default:
		panic("not match net type")
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
	"tcp":       model.TCP,
	"tcpserver": model.TcpServer,
	"udp":       model.UDP,
}

// CommonProcess 指令模板方式处理
func CommonProcess(cmd *cobra.Command, args []string) {
	use := cmd.Use
	netType := cmdMap[use]

	param := model.NetParam{
		Type:           netType,
		Port:           FlagContext.Port,
		Ip:             FlagContext.Ip,
		ListenPort:     FlagContext.ListenPort,
		ConnectTimeOut: FlagContext.ConnectTimeOut,
		ReceiveTimeOut: FlagContext.ReceiveTimeOut,
	}

	// 创建网络组件
	net := newNet(param)
	err := net.Start()
	if err != nil {
		log.Fatal(err)
	}

	// 创建数据交互面板
	p := tea.NewProgram(ui.InitialModel(func(input string) (string, error) {
		var toBytes []byte
		switch FlagContext.Encode {
		case "ascii":
			toBytes = []byte(input)
		default:
			toBytes, err = hex.DecodeString(strings.ReplaceAll(input, " ", ""))
			if nil != err {
				return "", err
			}
		}

		// 发送消息
		err := net.Write(&model.MsgTo{
			Ip:   FlagContext.Ip,
			Port: FlagContext.Port,
			Data: toBytes,
		})
		if err != nil {
			// 连接断开
			if err == io.EOF {
				return "(" + FlagContext.Ip + ":" + strconv.Itoa(FlagContext.Port) + "):disconnect", nil
			}
			return "", err
		}

		sendTo := "(" + FlagContext.Ip + ":" + strconv.Itoa(FlagContext.Port) + "):" + input
		return sendTo, nil
	}))

	// 读取网络组件数据 发送到终端
	go func() {
		for true {
			read, err := net.Read()
			if nil != err {
				p.Send(ui.NetInMsg(err.Error()))
				if err == io.EOF {
					p.Quit()
				}
				continue
			}
			switch FlagContext.Encode {
			case "ascii":
				p.Send(ui.NetInMsg(read.ToString()))
			default:
				p.Send(ui.NetInMsg(read.ToHexString()))
			}
		}
	}()

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
