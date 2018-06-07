package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 下发设备控制指令
// https://cloud.tencent.com/document/api/568/16525

type IssueDeviceControlRequest struct {
	// 控制数据（json）
	ControlData string `name:"ControlData"`
	// 设备名称
	DeviceName string `name:"DeviceName"`
	// 产品Id
	ProductId string `name:"ProductId"`
	// 区域
	Region string `name:"Region"`
}

func (req *IssueDeviceControlRequest) Invoke(client github_com_morlay_goqcloud.Client) (*IssueDeviceControlResponse, error) {
	resp := &IssueDeviceControlResponse{}
	err := client.Request("iot", "IssueDeviceControl", "2018-01-23").Do(req, resp)
	return resp, err
}

type IssueDeviceControlResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
