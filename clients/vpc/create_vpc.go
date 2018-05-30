package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建VPC
// https://cloud.tencent.com/document/api/215/15774
type CreateVpcRequest struct {
	// vpc的cidr，只能为10.0.0.0/16，172.16.0.0/12，192.168.0.0/16这三个内网网段内。
	CidrBlock string `name:"CidrBlock"`
	// DNS地址，最多支持4个
	DnsServers []*string `name:"DnsServers,omitempty"`
	// 域名
	DomainName *string `name:"DomainName,omitempty"`
	// 是否开启组播。true: 开启, false: 不开启。
	EnableMulticast *string `name:"EnableMulticast,omitempty"`
	// 区域
	Region string `name:"Region"`
	// vpc名称，最大长度不能超过60个字节。
	VpcName string `name:"VpcName"`
}

func (req *CreateVpcRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateVpcResponse, error) {
	resp := &CreateVpcResponse{}
	err := client.Request("vpc", "CreateVpc", "2017-03-12").Do(req, resp)
	return resp, err
}

type CreateVpcResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// Vpc对象。
	Vpc Vpc `json:"Vpc"`
}
