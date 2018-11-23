package tbaas

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询交易
// https://cloud.tencent.com/document/api/663/19463

type QueryRequest struct {
	// 函数参数列表
	Args []*string `name:"Args,omitempty"`
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

func (req *QueryRequest) Invoke(client github_com_morlay_goqcloud.Client) (*QueryResponse, error) {
	resp := &QueryResponse{}
	err := client.Request("tbaas", "Query", "2018-04-16").Do(req, resp)
	return resp, err
}

type QueryResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 查询结果数据
	Data []*string `json:"Data"`
}
