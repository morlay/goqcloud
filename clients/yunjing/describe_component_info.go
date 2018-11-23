package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取组件信息
// https://cloud.tencent.com/document/api/296/30338

type DescribeComponentInfoRequest struct {
	// 组件ID。
	ComponentId int64 `name:"ComponentId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeComponentInfoRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeComponentInfoResponse, error) {
	resp := &DescribeComponentInfoResponse{}
	err := client.Request("yunjing", "DescribeComponentInfo", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeComponentInfoResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 组件名称。
	ComponentName string `json:"ComponentName"`
	// 组件类型。WEB：web组件SYSTEM：系统组件
	ComponentType string `json:"ComponentType"`
	// 组件描述。
	Description string `json:"Description"`
	// 组件官网。
	Homepage string `json:"Homepage"`
	// 组件ID。
	Id int64 `json:"Id"`
}
