package iotcloud

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取设备影子
// https://cloud.tencent.com/document/api/634/19489

type DescribeDeviceShadowRequest struct {
	// 设备名称。命名规则：[a-zA-Z0-9:_-]{1,48}
	DeviceName string `name:"DeviceName"`
	// 产品 ID
	ProductId string `name:"ProductId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeDeviceShadowRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDeviceShadowResponse, error) {
	resp := &DescribeDeviceShadowResponse{}
	err := client.Request("iotcloud", "DescribeDeviceShadow", "2018-06-14").Do(req, resp)
	return resp, err
}

type DescribeDeviceShadowResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 设备影子数据
	Data string `json:"Data"`
}
