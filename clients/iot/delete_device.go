package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除设备
// https://cloud.tencent.com/document/api/568/16466

type DeleteDeviceRequest struct {
	// 设备名称
	DeviceName string `name:"DeviceName"`
	// 产品Id
	ProductId string `name:"ProductId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteDeviceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteDeviceResponse, error) {
	resp := &DeleteDeviceResponse{}
	err := client.Request("iot", "DeleteDevice", "2018-01-23").Do(req, resp)
	return resp, err
}

type DeleteDeviceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
