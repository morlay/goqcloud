package iotcloud

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取设备列表
// https://cloud.tencent.com/document/api/634/19493

type DescribeDevicesRequest struct {
	// 设备固件版本号，若不带此参数会返回所有固件版本的设备
	FirmwareVersion *string `name:"FirmwareVersion,omitempty"`
	// 分页的大小，数值范围 10-100
	Limit int64 `name:"Limit"`
	// 分页偏移
	Offset int64 `name:"Offset"`
	// 需要查看设备列表的产品 ID
	ProductId string `name:"ProductId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeDevicesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDevicesResponse, error) {
	resp := &DescribeDevicesResponse{}
	err := client.Request("iotcloud", "DescribeDevices", "2018-06-14").Do(req, resp)
	return resp, err
}

type DescribeDevicesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 设备详细信息列表
	Devices []*DeviceInfo `json:"Devices"`
	// 设备总数
	TotalCount int64 `json:"TotalCount"`
}
