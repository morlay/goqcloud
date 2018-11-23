package tbaas

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 新增交易
// https://cloud.tencent.com/document/api/663/19464

type InvokeRequest struct {
	// 函数参数列表
	Args []*string `name:"Args,omitempty"`
	// 同步调用标识
	AsyncFlag *int64 `name:"AsyncFlag,omitempty"`
	// 合约名称
	ChaincodeName string `name:"ChaincodeName"`
	// 通道名称
	ChannelName string `name:"ChannelName"`
	// cluster标识
	ClusterId string `name:"ClusterId"`
	// 函数名
	FuncName string `name:"FuncName"`
	// 模块名
	Module string `name:"Module"`
	// 操作名
	Operation string `name:"Operation"`
	// 使用的节点名称及对应组织名称
	Peers []*PeerSet `name:"Peers"`
	// 区域
	Region string `name:"Region"`
}

func (req *InvokeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*InvokeResponse, error) {
	resp := &InvokeResponse{}
	err := client.Request("tbaas", "Invoke", "2018-04-16").Do(req, resp)
	return resp, err
}

type InvokeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 返回内容
	Events string `json:"Events"`
	// 交易编号
	Txid string `json:"Txid"`
}
