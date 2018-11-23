package iotcloud

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取创建多设备结果
// https://cloud.tencent.com/document/api/634/19491

type DescribeMultiDevicesRequest struct {
	// 分页大小，每页返回的设备个数
	Limit int64 `name:"Limit"`
	// 分页偏移
	Offset int64 `name:"Offset"`
	// 产品 ID，创建产品时腾讯云为用户分配全局唯一的 ID
	ProductId string `name:"ProductId"`
	// 区域
	Region string `name:"Region"`
	// 任务 ID，由批量创建设备接口返回
	TaskId string `name:"TaskId"`
}

func (req *DescribeMultiDevicesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeMultiDevicesResponse, error) {
	resp := &DescribeMultiDevicesResponse{}
	err := client.Request("iotcloud", "DescribeMultiDevices", "2018-06-14").Do(req, resp)
	return resp, err
}

type DescribeMultiDevicesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 设备详细信息列表
	DevicesInfo []*MultiDevicesInfo `json:"DevicesInfo"`
	// 任务 ID，由批量创建设备接口返回
	TaskId string `json:"TaskId"`
	// 该任务创建设备的总数
	TotalDevNum int64 `json:"TotalDevNum"`
}
