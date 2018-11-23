package dc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除专用通道
// https://cloud.tencent.com/document/api/216/19820

type DeleteDirectConnectTunnelRequest struct {
	// 专用通道ID
	DirectConnectTunnelId string `name:"DirectConnectTunnelId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteDirectConnectTunnelRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteDirectConnectTunnelResponse, error) {
	resp := &DeleteDirectConnectTunnelResponse{}
	err := client.Request("dc", "DeleteDirectConnectTunnel", "2018-04-10").Do(req, resp)
	return resp, err
}

type DeleteDirectConnectTunnelResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
