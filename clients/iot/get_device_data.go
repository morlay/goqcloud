package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取设备数据
// https://cloud.tencent.com/document/api/568/16522
type GetDeviceDataRequest struct {
	// 设备名称
	DeviceName string `name:"DeviceName"`
	// 产品Id
	ProductId string `name:"ProductId"`
	// 区域
	Region string `name:"Region"`
}

func (req *GetDeviceDataRequest) Invoke(client github_com_morlay_goqcloud.Client) (*GetDeviceDataResponse, error) {
	resp := &GetDeviceDataResponse{}
	err := client.Request("iot", "GetDeviceData", "2018-01-23").Do(req, resp)
	return resp, err
}

type GetDeviceDataResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 设备数据
	DeviceData Object `json:"DeviceData"`
}
