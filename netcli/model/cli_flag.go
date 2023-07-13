package model

type CliFlags struct {
	// 连接超时时间 单位秒
	ConnectTimeOut int
	// 端口
	Port int
	// 监听端口
	ListenPort int
	// ip
	Ip string
	// 数据接收超时时间 单位秒
	ReceiveTimeOut int
	// 数据格式
	Encode string
}

// Valid 参数验证
func (f CliFlags) Valid(profile string) {

}
