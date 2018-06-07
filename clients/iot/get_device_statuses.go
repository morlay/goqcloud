package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 批量获取设备状态
// https://cloud.tencent.com/document/api/568/16524

type GetDeviceStatusesRequest struct {
	// 设备名称列表（单次限制1000个设备）
	DeviceNames []*string `name:"DeviceNames"`
	// 产品ID
	ProductId string `name:"ProductId"`
	// 区域
	Region string `name:"Region"`
}

func (req *GetDeviceStatusesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*GetDeviceStatusesResponse, error) {
	resp := &GetDeviceStatusesResponse{}
	err := client.Request("iot", "GetDeviceStatuses", "2018-01-23").Do(req, resp)
	return resp, err
}

type GetDeviceStatusesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 设备状态列表
	DeviceStatuses []*DeviceStatus `json:"DeviceStatuses"`
}
