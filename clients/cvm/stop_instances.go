package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 关闭实例
// https://cloud.tencent.com/document/api/213/15743
type StopInstancesRequest struct {
	// 是否在正常关闭失败后选择强制关闭实例。取值范围：TRUE：表示在正常关闭失败后进行强制关闭FALSE：表示在正常关闭失败后不进行强制关闭默认取值：FALSE。
	ForceStop *bool `name:"ForceStop,omitempty"`
	// 一个或多个待操作的实例ID。可通过DescribeInstances接口返回值中的InstanceId获取。每次请求批量实例的上限为100。
	InstanceIds []*string `name:"InstanceIds"`
	// 区域
	Region string `name:"Region"`
	// 实例的关闭模式。取值范围：SOFT_FIRST：表示在正常关闭失败后进行强制关闭HARD：直接强制关闭SOFT：仅软关机默认取值：SOFT。
	StopType *string `name:"StopType,omitempty"`
	// 关机收费模式KEEP_CHARGING：关机继续收费STOP_CHARGING：关机停止收费默认取值：KEEP_CHARGING。
	StoppedMode *string `name:"StoppedMode,omitempty"`
}

func (req *StopInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*StopInstancesResponse, error) {
	resp := &StopInstancesResponse{}
	err := client.Request("cvm", "StopInstances", "2017-03-12").Do(req, resp)
	return resp, err
}

type StopInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
