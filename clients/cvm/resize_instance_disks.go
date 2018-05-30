package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 扩容实例磁盘
// https://cloud.tencent.com/document/api/213/15731
type ResizeInstanceDisksRequest struct {
	// 待扩容的数据盘配置信息。只支持扩容随实例购买的数据盘，且数据盘类型为：CLOUD_BASIC、CLOUD_PREMIUM、CLOUD_SSD。数据盘容量单位：GB。最小扩容步长：10G。关于数据盘类型的选择请参考硬盘产品简介。可选数据盘类型受到实例类型InstanceType限制。另外允许扩容的最大容量也因数据盘类型的不同而有所差异。
	DataDisks []*DataDisk `name:"DataDisks"`
	// 是否对运行中的实例选择强制关机。建议对运行中的实例先手动关机，然后再重置用户密码。取值范围：TRUE：表示在正常关机失败后进行强制关机FALSE：表示在正常关机失败后不进行强制关机默认取值：FALSE。强制关机的效果等同于关闭物理计算机的电源开关。强制关机可能会导致数据丢失或文件系统损坏，请仅在服务器不能正常关机时使用。
	ForceStop *bool `name:"ForceStop,omitempty"`
	// 待操作的实例ID。可通过DescribeInstances接口返回值中的InstanceId获取。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *ResizeInstanceDisksRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ResizeInstanceDisksResponse, error) {
	resp := &ResizeInstanceDisksResponse{}
	err := client.Request("cvm", "ResizeInstanceDisks", "2017-03-12").Do(req, resp)
	return resp, err
}

type ResizeInstanceDisksResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
