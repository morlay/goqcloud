package cis

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取容器实例信息
// https://cloud.tencent.com/document/api/858/17773

type DescribeContainerInstanceRequest struct {
	// 容器实例名称
	InstanceName string `name:"InstanceName"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeContainerInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeContainerInstanceResponse, error) {
	resp := &DescribeContainerInstanceResponse{}
	err := client.Request("cis", "DescribeContainerInstance", "2018-04-08").Do(req, resp)
	return resp, err
}

type DescribeContainerInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 容器实例详细信息
	ContainerInstance ContainerInstance `json:"ContainerInstance"`
}
