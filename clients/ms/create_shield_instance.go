package ms

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 提交加固基础数据
// https://cloud.tencent.com/document/api/283/17753

type CreateShieldInstanceRequest struct {
	// 待加固的应用信息
	AppInfo AppInfo `name:"AppInfo"`
	// 区域
	Region string `name:"Region"`
	// 加固服务信息
	ServiceInfo ServiceInfo `name:"ServiceInfo"`
}

func (req *CreateShieldInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateShieldInstanceResponse, error) {
	resp := &CreateShieldInstanceResponse{}
	err := client.Request("ms", "CreateShieldInstance", "2018-04-08").Do(req, resp)
	return resp, err
}

type CreateShieldInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 任务唯一标识
	ItemId string `json:"ItemId"`
	// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	Progress int64 `json:"Progress"`
}
