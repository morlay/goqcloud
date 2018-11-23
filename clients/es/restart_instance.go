package es

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 重启ES集群实例
// https://cloud.tencent.com/document/api/845/30630

type RestartInstanceRequest struct {
	// 要重启的实例ID
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *RestartInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*RestartInstanceResponse, error) {
	resp := &RestartInstanceResponse{}
	err := client.Request("es", "RestartInstance", "2018-04-16").Do(req, resp)
	return resp, err
}

type RestartInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
