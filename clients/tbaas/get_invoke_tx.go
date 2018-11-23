package tbaas

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// Invoke异步调用结果查询
// https://cloud.tencent.com/document/api/663/19465

type GetInvokeTxRequest struct {
	// 通道名称
	ChannelName string `name:"ChannelName"`
	// cluster标识
	ClusterId string `name:"ClusterId"`
	// 模块名
	Module string `name:"Module"`
	// 操作名
	Operation string `name:"Operation"`
	// 节点所属组织名称
	PeerGroup string `name:"PeerGroup"`
	// 节点名称
	PeerName string `name:"PeerName"`
	// 区域
	Region string `name:"Region"`
	// 事务ID
	TxId string `name:"TxId"`
}

func (req *GetInvokeTxRequest) Invoke(client github_com_morlay_goqcloud.Client) (*GetInvokeTxResponse, error) {
	resp := &GetInvokeTxResponse{}
	err := client.Request("tbaas", "GetInvokeTx", "2018-04-16").Do(req, resp)
	return resp, err
}

type GetInvokeTxResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 状态码
	TxValidationCode int64 `json:"TxValidationCode"`
	// 消息
	TxValidationMsg string `json:"TxValidationMsg"`
}
