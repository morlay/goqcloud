package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 重置实例密码
// https://cloud.tencent.com/document/api/213/15736

type ResetInstancesPasswordRequest struct {
	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再重置用户密码。取值范围：TRUE：表示在正常关机失败后进行强制关机FALSE：表示在正常关机失败后不进行强制关机默认取值：FALSE。强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。
	ForceStop *bool `name:"ForceStop,omitempty"`
	// 一个或多个待操作的实例ID。可通过DescribeInstances API返回值中的InstanceId获取。每次请求允许操作的实例数量上限是100。
	InstanceIds []*string `name:"InstanceIds"`
	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：Linux实例密码必须8到16位，至少包括两项[a-z，A-Z]、[0-9]和[( ) ~ ~ ! @ # $ % ^ & * - + = _ &#124; { } [ ] : ; ' < > , . ? /]中的符号。密码不允许以/符号开头。Windows实例密码必须12到16位，至少包括三项[a-z]，[A-Z]，[0-9]和[( ) ~ ~ ! @ # $ % ^ & * - + = _ &#124; { } [ ] : ; ' < > , . ? /]中的符号。密码不允许以/符号开头。如果实例即包含Linux实例又包含Windows实例，则密码复杂度限制按照Windows实例的限制。
	Password string `name:"Password"`
	// 区域
	Region string `name:"Region"`
	// 待重置密码的实例操作系统用户名。不得超过64个字符。
	UserName *string `name:"UserName,omitempty"`
}

func (req *ResetInstancesPasswordRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ResetInstancesPasswordResponse, error) {
	resp := &ResetInstancesPasswordResponse{}
	err := client.Request("cvm", "ResetInstancesPassword", "2017-03-12").Do(req, resp)
	return resp, err
}

type ResetInstancesPasswordResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
