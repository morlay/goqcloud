package cis

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除容器实例
// https://cloud.tencent.com/document/api/858/17774

type DeleteContainerInstanceRequest struct {
	// 容器实例名称
	InstanceName string `name:"InstanceName"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteContainerInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteContainerInstanceResponse, error) {
	resp := &DeleteContainerInstanceResponse{}
	err := client.Request("cis", "DeleteContainerInstance", "2018-04-08").Do(req, resp)
	return resp, err
}

type DeleteContainerInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 操作信息
	Msg string `json:"Msg"`
}
