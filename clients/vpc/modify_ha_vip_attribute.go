package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改HAVIP属性
// https://cloud.tencent.com/document/api/215/30647

type ModifyHaVipAttributeRequest struct {
	// HAVIP唯一ID，形如：havip-9o233uri。
	HaVipId string `name:"HaVipId"`
	// HAVIP名称，可任意命名，但不得超过60个字符。
	HaVipName string `name:"HaVipName"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyHaVipAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyHaVipAttributeResponse, error) {
	resp := &ModifyHaVipAttributeResponse{}
	err := client.Request("vpc", "ModifyHaVipAttribute", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyHaVipAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
