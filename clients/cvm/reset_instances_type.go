package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 调整实例配置
// https://cloud.tencent.com/document/api/213/15744
type ResetInstancesTypeRequest struct {
	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再重置用户密码。取值范围：TRUE：表示在正常关机失败后进行强制关机FALSE：表示在正常关机失败后不进行强制关机默认取值：FALSE。强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。
	ForceStop *bool `name:"ForceStop,omitempty"`
	// 一个或多个待操作的实例ID。可通过DescribeInstances接口返回值中的InstanceId获取。每次请求批量实例的上限为1。
	InstanceIds []*string `name:"InstanceIds"`
	// 实例机型。不同实例机型指定了不同的资源规格，具体取值可参见附表实例资源规格对照表，也可以调用查询实例资源规格列表接口获得最新的规格表。
	InstanceType string `name:"InstanceType"`
	// 区域
	Region string `name:"Region"`
}

func (req *ResetInstancesTypeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ResetInstancesTypeResponse, error) {
	resp := &ResetInstancesTypeResponse{}
	err := client.Request("cvm", "ResetInstancesType", "2017-03-12").Do(req, resp)
	return resp, err
}

type ResetInstancesTypeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
