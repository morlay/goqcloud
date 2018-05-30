package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改VPC属性
// https://cloud.tencent.com/document/api/215/15773
type ModifyVpcAttributeRequest struct {
	// DNS地址，最多支持4个，第1个默认为主，其余为备
	DnsServers []*string `name:"DnsServers,omitempty"`
	// 域名
	DomainName *string `name:"DomainName,omitempty"`
	// 是否开启组播。true: 开启, false: 关闭。
	EnableMulticast *string `name:"EnableMulticast,omitempty"`
	// 区域
	Region string `name:"Region"`
	// VPC实例ID。形如：vpc-f49l6u0z。每次请求的实例的上限为100。参数不支持同时指定VpcIds和Filters。
	VpcId string `name:"VpcId"`
	// 私有网络名称，可任意命名，但不得超过60个字符。
	VpcName *string `name:"VpcName,omitempty"`
}

func (req *ModifyVpcAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyVpcAttributeResponse, error) {
	resp := &ModifyVpcAttributeResponse{}
	err := client.Request("vpc", "ModifyVpcAttribute", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyVpcAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
