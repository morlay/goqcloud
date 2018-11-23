package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询实例管理终端地址
// https://cloud.tencent.com/document/api/213/18150

type DescribeInstanceVncURLRequest struct {
	// 一个操作的实例ID。可通过DescribeInstances API返回值中的InstanceId获取。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeInstanceVncURLRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeInstanceVncURLResponse, error) {
	resp := &DescribeInstanceVncURLResponse{}
	err := client.Request("cvm", "DescribeInstanceVncUrl", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeInstanceVncURLResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 实例的管理终端地址。
	InstanceVncUrl string `json:"InstanceVncUrl"`
}
