package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 重启实例
// https://cloud.tencent.com/document/api/236/17488

type RestartDbInstancesRequest struct {
	// 实例ID数组，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceIds []*string `name:"InstanceIds"`
	// 区域
	Region string `name:"Region"`
}

func (req *RestartDbInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*RestartDbInstancesResponse, error) {
	resp := &RestartDbInstancesResponse{}
	err := client.Request("cdb", "RestartDBInstances", "2017-03-20").Do(req, resp)
	return resp, err
}

type RestartDbInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步任务的请求ID，可使用此ID查询异步任务的执行结果。
	AsyncRequestId string `json:"AsyncRequestId"`
}
