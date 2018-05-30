package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改镜像分享信息
// https://cloud.tencent.com/document/api/213/15710
type ModifyImageSharePermissionRequest struct {
	// 接收分享镜像的账号Id列表，array型参数的格式可以参考API简介。帐号ID不同于QQ号，查询用户帐号ID请查看帐号信息中的帐号ID栏。
	AccountIds []*string `name:"AccountIds"`
	// 镜像ID，形如img-gvbnzy6f。镜像Id可以通过如下方式获取：通过DescribeImages接口返回的ImageId获取。通过镜像控制台获取。 镜像ID必须指定为状态为NORMAL的镜像。镜像状态请参考镜像数据表。
	ImageId string `name:"ImageId"`
	// 操作，包括 SHARE，CANCEL。其中SHARE代表分享操作，CANCEL代表取消分享操作。
	Permission string `name:"Permission"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyImageSharePermissionRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyImageSharePermissionResponse, error) {
	resp := &ModifyImageSharePermissionResponse{}
	err := client.Request("cvm", "ModifyImageSharePermission", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyImageSharePermissionResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
