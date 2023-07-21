package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"io"
	"lstfight.cn/go-pra/netcli/model"
	"net"
	"net/netip"
	"strconv"
)

type TcpServer struct {
	NetParam model.NetParam
	connMap  map[string]net.Conn
	channel  chan *model.MsgForm
	server   *net.TCPListener
}

func TcpServerCmd() *cobra.Command {
	return &cobra.Command{
		Use:  "tcpserver",
		Long: "create tcpserver communication",
		Run:  CommonProcess,
	}
}

func (netT *TcpServer) Start() error {
	port := netT.NetParam.Port
	listenPort := netT.NetParam.ListenPort
	if listenPort < 1 {
		listenPort = port
	}
	// 地址参数
	addr := net.TCPAddr{
		Port: listenPort,
	}
	server, err := net.ListenTCP("tcp", &addr)

	if err != nil {
		return err
	}

	// make chan
	netT.channel = make(chan *model.MsgForm)
	netT.connMap = map[string]net.Conn{}

	// 接收client
	go func() {
		for {
			conn, _ := server.Accept()
			remoteAddr := conn.RemoteAddr()
			netT.connMap[remoteAddr.String()] = conn
			// 读取线程数据
			netT.readClient(conn)
		}
	}()

	return nil
}

func (netT *TcpServer) Read() (*model.MsgForm, error) {
	// 从channel 获取数据
	form := <-netT.channel

	// 没有获取到数据 说明断开连接
	if form.Data == nil {
		return nil, errors.New("(" + form.Ip + ":" + strconv.Itoa(form.Port) + "):disconnect")
	}

	return form, nil
}

func (netT *TcpServer) Write(to *model.MsgTo) error {
	// 获取到 client socket 进行发送
	key := to.Ip + ":" + strconv.Itoa(to.Port)
	conn := netT.connMap[key]
	if nil != conn {
		_, err := conn.Write(to.Data)
		if err != nil {
			return err
		}
	} else {
		return errors.New("not conn " + key)
	}

	return nil
}

func (netT *TcpServer) Stop() error {
	for _, conn := range netT.connMap {
		_ = conn.Close()
	}
	err := netT.server.Close()
	if err != nil {
		return err
	}
	return nil
}

func (netT *TcpServer) readClient(conn net.Conn) {
	// 获取客户端数据
	go func() {
		for true {
			addPort, _ := netip.ParseAddrPort(conn.RemoteAddr().String())

			bytes := make([]byte, 10240, 10240)
			rNum, err := conn.Read(bytes)
			if err != nil {
				// 关闭连接
				if err == io.EOF {
					_ = conn.Close()

					// 写入data nil msg 说明断开连接
					netT.channel <- &model.MsgForm{
						Ip:   addPort.Addr().String(),
						Port: int(addPort.Port()),
					}

					// 删除记录
					delete(netT.connMap, conn.RemoteAddr().String())
					// 结束轮询获取
					return
				}
			}
			// 将数据写入通道
			netT.channel <- &model.MsgForm{
				Data: bytes[:rNum],
				Ip:   addPort.Addr().String(),
				Port: int(addPort.Port()),
			}
		}
	}()
}
