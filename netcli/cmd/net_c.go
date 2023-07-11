package cmd

import "lstfight.cn/go-pra/netcli/model"

type NetCh interface {
	Read() []byte

	Write(d []byte)
}

// create net component
func newNet(p model.NetParam) *NetCh {
	return nil
}
