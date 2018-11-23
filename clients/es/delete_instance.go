package es

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 销毁ES集群实例
// https://cloud.tencent.com/document/api/845/30632

type DeleteInstanceRequest struct {
	// 要销毁的实例ID
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteInstanceResponse, error) {
	resp := &DeleteInstanceResponse{}
	err := client.Request("es", "DeleteInstance", "2018-04-16").Do(req, resp)
	return resp, err
}

type DeleteInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
