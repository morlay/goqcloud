package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 重装实例
// https://cloud.tencent.com/document/api/213/15724
type ResetInstanceRequest struct {
	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `name:"EnhancedService,omitempty"`
	// 指定有效的镜像ID，格式形如img-xxx。镜像类型分为四种：公共镜像自定义镜像共享镜像服务市场镜像可通过以下方式获取可用的镜像ID：公共镜像、自定义镜像、共享镜像的镜像ID可通过登录控制台查询；服务镜像市场的镜像ID可通过云市场查询。通过调用接口 DescribeImages ，取返回信息中的ImageId字段。
	ImageId *string `name:"ImageId,omitempty"`
	// 实例ID。可通过 DescribeInstances API返回值中的InstanceId获取。
	InstanceId string `name:"InstanceId"`
	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。
	LoginSettings *LoginSettings `name:"LoginSettings,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 实例系统盘配置信息。系统盘为云盘的实例可以通过该参数指定重装后的系统盘大小来实现对系统盘的扩容操作，若不指定则默认系统盘大小保持不变。系统盘大小只支持扩容不支持缩容；重装只支持修改系统盘的大小，不能修改系统盘的类型。
	SystemDisk *SystemDisk `name:"SystemDisk,omitempty"`
}

func (req *ResetInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ResetInstanceResponse, error) {
	resp := &ResetInstanceResponse{}
	err := client.Request("cvm", "ResetInstance", "2017-03-12").Do(req, resp)
	return resp, err
}

type ResetInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
