package iotcloud

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除设备
// https://cloud.tencent.com/document/api/634/19494

type DeleteDeviceRequest struct {
	// 需要删除的设备名称
	DeviceName string `name:"DeviceName"`
	// 设备所属的产品 ID
	ProductId string `name:"ProductId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteDeviceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteDeviceResponse, error) {
	resp := &DeleteDeviceResponse{}
	err := client.Request("iotcloud", "DeleteDevice", "2018-06-14").Do(req, resp)
	return resp, err
}

type DeleteDeviceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
