package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 新增设备
// https://cloud.tencent.com/document/api/568/16465

type AddDeviceRequest struct {
	// 设备名称，唯一标识某产品下的一个设备
	DeviceName string `name:"DeviceName"`
	// 产品Id
	ProductId string `name:"ProductId"`
	// 区域
	Region string `name:"Region"`
}

func (req *AddDeviceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*AddDeviceResponse, error) {
	resp := &AddDeviceResponse{}
	err := client.Request("iot", "AddDevice", "2018-01-23").Do(req, resp)
	return resp, err
}

type AddDeviceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 设备信息
	Device Device `json:"Device"`
}
