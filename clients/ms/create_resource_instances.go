package ms

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建资源
// https://cloud.tencent.com/document/api/283/19823

type CreateResourceInstancesRequest struct {
	// 资源类型id。13624：加固专业版。
	Pid int64 `name:"Pid"`
	// 区域
	Region string `name:"Region"`
	// 资源数量。
	ResourceNum int64 `name:"ResourceNum"`
	// 时间数量。
	TimeSpan int64 `name:"TimeSpan"`
	// 时间单位，取值为d，m，y，分别表示天，月，年。
	TimeUnit string `name:"TimeUnit"`
}

func (req *CreateResourceInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateResourceInstancesResponse, error) {
	resp := &CreateResourceInstancesResponse{}
	err := client.Request("ms", "CreateResourceInstances", "2018-04-08").Do(req, resp)
	return resp, err
}

type CreateResourceInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 新创建的资源列表。
	ResourceSet []*string `json:"ResourceSet"`
}
