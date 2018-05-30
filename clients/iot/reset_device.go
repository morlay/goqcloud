package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 重置设备
// https://cloud.tencent.com/document/api/568/16526
type ResetDeviceRequest struct {
	// 设备名称
	DeviceName string `name:"DeviceName"`
	// 产品Id
	ProductId string `name:"ProductId"`
	// 区域
	Region string `name:"Region"`
}

func (req *ResetDeviceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ResetDeviceResponse, error) {
	resp := &ResetDeviceResponse{}
	err := client.Request("iot", "ResetDevice", "2018-01-23").Do(req, resp)
	return resp, err
}

type ResetDeviceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
