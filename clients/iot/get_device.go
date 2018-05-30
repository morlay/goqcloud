package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取设备信息
// https://cloud.tencent.com/document/api/568/16467
type GetDeviceRequest struct {
	// 设备名称
	DeviceName string `name:"DeviceName"`
	// 产品Id
	ProductId string `name:"ProductId"`
	// 区域
	Region string `name:"Region"`
}

func (req *GetDeviceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*GetDeviceResponse, error) {
	resp := &GetDeviceResponse{}
	err := client.Request("iot", "GetDevice", "2018-01-23").Do(req, resp)
	return resp, err
}

type GetDeviceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 设备信息
	Device Device `json:"Device"`
}
