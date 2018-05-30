package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询实例带宽配置
// https://cloud.tencent.com/document/api/213/15734
type DescribeInstanceInternetBandwidthConfigsRequest struct {
	// 待操作的实例ID。可通过DescribeInstances接口返回值中的InstanceId获取。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeInstanceInternetBandwidthConfigsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeInstanceInternetBandwidthConfigsResponse, error) {
	resp := &DescribeInstanceInternetBandwidthConfigsResponse{}
	err := client.Request("cvm", "DescribeInstanceInternetBandwidthConfigs", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeInstanceInternetBandwidthConfigsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 带宽配置信息列表。
	InternetBandwidthConfigSet []*InternetBandwidthConfig `json:"InternetBandwidthConfigSet"`
}
