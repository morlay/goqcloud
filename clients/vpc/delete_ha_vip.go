package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除HAVIP
// https://cloud.tencent.com/document/api/215/30651

type DeleteHaVipRequest struct {
	// HAVIP唯一ID，形如：havip-9o233uri。
	HaVipId string `name:"HaVipId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteHaVipRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteHaVipResponse, error) {
	resp := &DeleteHaVipResponse{}
	err := client.Request("vpc", "DeleteHaVip", "2017-03-12").Do(req, resp)
	return resp, err
}

type DeleteHaVipResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
