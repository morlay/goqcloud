package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 启动实例
// https://cloud.tencent.com/document/api/213/15735

type StartInstancesRequest struct {
	// 一个或多个待操作的实例ID。可通过DescribeInstances接口返回值中的InstanceId获取。每次请求批量实例的上限为100。
	InstanceIds []*string `name:"InstanceIds"`
	// 区域
	Region string `name:"Region"`
}

func (req *StartInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*StartInstancesResponse, error) {
	resp := &StartInstancesResponse{}
	err := client.Request("cvm", "StartInstances", "2017-03-12").Do(req, resp)
	return resp, err
}

type StartInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
