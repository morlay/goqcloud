package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取设备列表
// https://cloud.tencent.com/document/api/568/16468

type GetDevicesRequest struct {
	// 长度
	Length *int64 `name:"Length,omitempty"`
	// 偏移
	Offset *int64 `name:"Offset,omitempty"`
	// 产品Id
	ProductId string `name:"ProductId"`
	// 区域
	Region string `name:"Region"`
}

func (req *GetDevicesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*GetDevicesResponse, error) {
	resp := &GetDevicesResponse{}
	err := client.Request("iot", "GetDevices", "2018-01-23").Do(req, resp)
	return resp, err
}

type GetDevicesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 设备列表
	Devices []*Device `json:"Devices"`
	// 设备总数
	Total int64 `json:"Total"`
}
