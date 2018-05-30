package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询实例机型列表
// https://cloud.tencent.com/document/api/213/15749
type DescribeInstanceTypeConfigsRequest struct {
	// 过滤条件。 zone - String - 是否必填：否 -（过滤条件）按照可用区过滤。 instance-family - String - 是否必填：否 -（过滤条件）按照实例机型系列过滤。实例机型系列形如：S1、I1、M1等。每次请求的Filters的上限为10，Filter.Values的上限为1。
	Filters []*Filter `name:"Filters,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeInstanceTypeConfigsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeInstanceTypeConfigsResponse, error) {
	resp := &DescribeInstanceTypeConfigsResponse{}
	err := client.Request("cvm", "DescribeInstanceTypeConfigs", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeInstanceTypeConfigsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 实例机型配置列表。
	InstanceTypeConfigSet []*InstanceTypeConfig `json:"InstanceTypeConfigSet"`
}
