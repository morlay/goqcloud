package batch

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取批量计算可用的CVM机型配置信息
// https://cloud.tencent.com/document/api/599/15887

type DescribeAvailableCvmInstanceTypesRequest struct {
	// 过滤条件。 zone - String - 是否必填：否 -（过滤条件）按照可用区过滤。 instance-family String - 是否必填：否 -（过滤条件）按照机型系列过滤。实例机型系列形如：S1、I1、M1等。
	Filters []*Filter `name:"Filters,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeAvailableCvmInstanceTypesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeAvailableCvmInstanceTypesResponse, error) {
	resp := &DescribeAvailableCvmInstanceTypesResponse{}
	err := client.Request("batch", "DescribeAvailableCvmInstanceTypes", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeAvailableCvmInstanceTypesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 机型配置数组
	InstanceTypeConfigSet []*InstanceTypeConfig `json:"InstanceTypeConfigSet"`
}
