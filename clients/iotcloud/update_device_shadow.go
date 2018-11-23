package iotcloud

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 更新设备影子
// https://cloud.tencent.com/document/api/634/19488

type UpdateDeviceShadowRequest struct {
	// 设备名称
	DeviceName string `name:"DeviceName"`
	// 产品ID
	ProductId string `name:"ProductId"`
	// 区域
	Region string `name:"Region"`
	// 当前版本号，需要和后台的version保持一致，才能更新成功
	ShadowVersion int64 `name:"ShadowVersion"`
	// 虚拟设备的状态，JSON字符串格式，由desired结构组成
	State string `name:"State"`
}

func (req *UpdateDeviceShadowRequest) Invoke(client github_com_morlay_goqcloud.Client) (*UpdateDeviceShadowResponse, error) {
	resp := &UpdateDeviceShadowResponse{}
	err := client.Request("iotcloud", "UpdateDeviceShadow", "2018-06-14").Do(req, resp)
	return resp, err
}

type UpdateDeviceShadowResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 设备影子数据，JSON字符串格式
	Data string `json:"Data"`
}
