package iotcloud

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建多个设备
// https://cloud.tencent.com/document/api/634/19495

type CreateMultiDeviceRequest struct {
	// 批量创建的设备名数组，单次最多创建 100 个设备。命名规则：[a-zA-Z0-9:_-]{1,48}
	DeviceNames []*string `name:"DeviceNames"`
	// 产品 ID。创建产品时腾讯云为用户分配全局唯一的 ID
	ProductId string `name:"ProductId"`
	// 区域
	Region string `name:"Region"`
}

func (req *CreateMultiDeviceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateMultiDeviceResponse, error) {
	resp := &CreateMultiDeviceResponse{}
	err := client.Request("iotcloud", "CreateMultiDevice", "2018-06-14").Do(req, resp)
	return resp, err
}

type CreateMultiDeviceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 任务ID，腾讯云生成全局唯一的任务 ID，有效期一个月，一个月之后任务失效。可以调用获取创建多设备任务状态接口获取该任务的执行状态，当状态为成功时，可以调用获取创建多设备任务结果接口获取该任务的结果
	TaskId string `json:"TaskId"`
}
