package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 重启实例
// https://cloud.tencent.com/document/api/213/15742
type RebootInstancesRequest struct {
	// 是否在正常重启失败后选择强制重启实例。取值范围：TRUE：表示在正常重启失败后进行强制重启FALSE：表示在正常重启失败后不进行强制重启默认取值：FALSE。
	ForceReboot *bool `name:"ForceReboot,omitempty"`
	// 一个或多个待操作的实例ID。可通过DescribeInstances接口返回值中的InstanceId获取。每次请求批量实例的上限为100。
	InstanceIds []*string `name:"InstanceIds"`
	// 区域
	Region string `name:"Region"`
}

func (req *RebootInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*RebootInstancesResponse, error) {
	resp := &RebootInstancesResponse{}
	err := client.Request("cvm", "RebootInstances", "2017-03-12").Do(req, resp)
	return resp, err
}

type RebootInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
