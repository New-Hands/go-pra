package cmd

import (
	"encoding/hex"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
	"io"
	"log"
	dlog "lstfight.cn/go-pra/netcli/log"
	"lstfight.cn/go-pra/netcli/model"
	"lstfight.cn/go-pra/netcli/ui"
	"os"
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

// 记录上一次数据来源
var lastMsg *model.MsgForm

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

	// 创建数据交互文件
	open, fErr := os.OpenFile("netcli.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if fErr != nil {
		fmt.Println(fErr)
	}

	// 创建数据交互面板
	forceAnsi()

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

		// 不考虑并 只需保持发送地址和显示一致即可
		to := &model.MsgTo{
			Ip:   FlagContext.Ip,
			Port: FlagContext.Port,
			Data: toBytes,
		}
		if lastMsg != nil {
			to.Ip = strings.Clone(lastMsg.Ip)
			to.Port = lastMsg.Port
		}
		err := net.Write(to)

		if err != nil {
			// 连接断开
			if err == io.EOF {
				return "(" + FlagContext.Ip + ":" + strconv.Itoa(FlagContext.Port) + "):disconnect", nil
			}
			return "", err
		}

		sendTo := "(" + to.Ip + ":" + strconv.Itoa(to.Port) + "):" + input
		dlog.Log(open, sendTo)

		return sendTo, nil
	}))

	// 读取网络组件数据 发送到终端
	readNetData(net, p, open)

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func forceAnsi() {
	profile := termenv.EnvColorProfile()
	if profile == termenv.Ascii {
		// 强制设置颜色
		envErr := os.Setenv("TERM", "linux")
		if envErr != nil {
			panic(envErr)
		}
	}
}

func readNetData(net NetCh, p *tea.Program, open *os.File) {
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

			// changeLastMsg
			lastMsg = read

			switch FlagContext.Encode {
			case "ascii":
				recv := read.ToString()
				p.Send(ui.NetInMsg(recv))
				dlog.Log(open, recv)
			default:
				recv := read.ToHexString()
				p.Send(ui.NetInMsg(recv))
				dlog.Log(open, recv)
			}
		}
	}()
}
