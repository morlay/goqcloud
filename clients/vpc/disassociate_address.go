package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 解绑定弹性公网IP
// https://cloud.tencent.com/document/api/215/16703
type DisassociateAddressRequest struct {
	// 标识 EIP 的唯一 ID。EIP 唯一 ID 形如：eip-11112222。
	AddressId string `name:"AddressId"`
	// 表示解绑 EIP 之后是否分配普通公网 IP。取值范围：TRUE：表示解绑 EIP 之后分配普通公网 IP。FALSE：表示解绑 EIP 之后不分配普通公网 IP。默认取值：FALSE。只有满足以下条件时才能指定该参数： 只有在解绑主网卡的主内网 IP 上的 EIP 时才能指定该参数。解绑 EIP 后重新分配普通公网 IP 操作一个账号每天最多操作 10 次；详情可通过 DescribeAddressQuota 接口获取。
	ReallocateNormalPublicIp *bool `name:"ReallocateNormalPublicIp,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DisassociateAddressRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DisassociateAddressResponse, error) {
	resp := &DisassociateAddressResponse{}
	err := client.Request("vpc", "DisassociateAddress", "2017-03-12").Do(req, resp)
	return resp, err
}

type DisassociateAddressResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
