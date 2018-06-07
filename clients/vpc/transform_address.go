package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 普通IP转弹性IP
// https://cloud.tencent.com/document/api/215/16706

type TransformAddressRequest struct {
	// 待操作有普通公网 IP 的实例 ID。实例 ID 形如：ins-11112222。可通过登录控制台查询，也可通过 DescribeInstances 接口返回值中的InstanceId获取。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *TransformAddressRequest) Invoke(client github_com_morlay_goqcloud.Client) (*TransformAddressResponse, error) {
	resp := &TransformAddressResponse{}
	err := client.Request("vpc", "TransformAddress", "2017-03-12").Do(req, resp)
	return resp, err
}

type TransformAddressResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
