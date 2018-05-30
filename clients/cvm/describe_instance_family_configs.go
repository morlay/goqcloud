package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询所支持的实例机型族信息
// https://cloud.tencent.com/document/api/213/15748
type DescribeInstanceFamilyConfigsRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeInstanceFamilyConfigsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeInstanceFamilyConfigsResponse, error) {
	resp := &DescribeInstanceFamilyConfigsResponse{}
	err := client.Request("cvm", "DescribeInstanceFamilyConfigs", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeInstanceFamilyConfigsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 实例机型组配置的列表信息
	InstanceFamilyConfigSet []*InstanceFamilyConfig `json:"InstanceFamilyConfigSet"`
}
