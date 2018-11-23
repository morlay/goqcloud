package cis

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建容器实例
// https://cloud.tencent.com/document/api/858/17775

type CreateContainerInstanceRequest struct {
	// 容器列表
	Containers []*Container `name:"Containers"`
	// 容器实例名称，由小写字母、数字和 - 组成，由小写字母开头，小写字母或数字结尾，且长度不超过 40个字符
	InstanceName string `name:"InstanceName"`
	// 区域
	Region string `name:"Region"`
	// 重启策略（Always,OnFailure,Never）
	RestartPolicy string `name:"RestartPolicy"`
	// subnetId
	SubnetId string `name:"SubnetId"`
	// vpcId
	VpcId string `name:"VpcId"`
	// 可用区
	Zone string `name:"Zone"`
}

func (req *CreateContainerInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateContainerInstanceResponse, error) {
	resp := &CreateContainerInstanceResponse{}
	err := client.Request("cis", "CreateContainerInstance", "2018-04-08").Do(req, resp)
	return resp, err
}

type CreateContainerInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 容器实例ID
	InstanceId string `json:"InstanceId"`
}
