package dc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 接受专用通道申请
// https://cloud.tencent.com/document/api/216/19822

type AcceptDirectConnectTunnelRequest struct {
	// 物理专线拥有者接受共享专用通道申请
	DirectConnectTunnelId string `name:"DirectConnectTunnelId"`
	// 区域
	Region string `name:"Region"`
}

func (req *AcceptDirectConnectTunnelRequest) Invoke(client github_com_morlay_goqcloud.Client) (*AcceptDirectConnectTunnelResponse, error) {
	resp := &AcceptDirectConnectTunnelResponse{}
	err := client.Request("dc", "AcceptDirectConnectTunnel", "2018-04-10").Do(req, resp)
	return resp, err
}

type AcceptDirectConnectTunnelResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
