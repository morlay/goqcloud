package dc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 拒绝专用通道申请
// https://cloud.tencent.com/document/api/216/19817

type RejectDirectConnectTunnelRequest struct {
	// 无
	DirectConnectTunnelId string `name:"DirectConnectTunnelId"`
	// 区域
	Region string `name:"Region"`
}

func (req *RejectDirectConnectTunnelRequest) Invoke(client github_com_morlay_goqcloud.Client) (*RejectDirectConnectTunnelResponse, error) {
	resp := &RejectDirectConnectTunnelResponse{}
	err := client.Request("dc", "RejectDirectConnectTunnel", "2018-04-10").Do(req, resp)
	return resp, err
}

type RejectDirectConnectTunnelResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
