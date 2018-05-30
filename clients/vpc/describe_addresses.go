package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询弹性公网IP列表
// https://cloud.tencent.com/document/api/215/16702
type DescribeAddressesRequest struct {
	// 标识 EIP 的唯一 ID 列表。EIP 唯一 ID 形如：eip-11112222。参数不支持同时指定AddressIds和Filters。
	AddressIds []*string `name:"AddressIds,omitempty"`
	// 每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定AddressIds和Filters。详细的过滤条件如下： address-id - String - 是否必填：否 - （过滤条件）按照 EIP 的唯一 ID 过滤。EIP 唯一 ID 形如：eip-11112222。 address-name - String - 是否必填：否 - （过滤条件）按照 EIP 名称过滤。不支持模糊过滤。 address-ip - String - 是否必填：否 - （过滤条件）按照 EIP 的 IP 地址过滤。 address-status - String - 是否必填：否 - （过滤条件）按照 EIP 的状态过滤。取值范围：详见EIP状态列表。 instance-id - String - 是否必填：否 - （过滤条件）按照 EIP 绑定的实例 ID 过滤。实例 ID 形如：ins-11112222。 private-ip-address - String - 是否必填：否 - （过滤条件）按照 EIP 绑定的内网 IP 过滤。 network-interface-id - String - 是否必填：否 - （过滤条件）按照 EIP 绑定的弹性网卡 ID 过滤。弹性网卡 ID 形如：eni-11112222。 is-arrears - String - 是否必填：否 - （过滤条件）按照 EIP 是否欠费进行过滤。（TRUE：EIP 处于欠费状态|FALSE：EIP 费用状态正常）
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeAddressesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeAddressesResponse, error) {
	resp := &DescribeAddressesResponse{}
	err := client.Request("vpc", "DescribeAddresses", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeAddressesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// EIP 详细信息列表。
	AddressSet []*Address `json:"AddressSet"`
	// 符合条件的 EIP 数量。
	TotalCount int64 `json:"TotalCount"`
}
