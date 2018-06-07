package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询密钥对列表
// https://cloud.tencent.com/document/api/213/15699

type DescribeKeyPairsRequest struct {
	// 过滤条件。 project-id - Integer - 是否必填：否 -（过滤条件）按照项目ID过滤。可以通过项目列表查询项目ID，或者调用接口 DescribeProject，取返回信息中的projectId获取项目ID。 key-name - String - 是否必填：否 -（过滤条件）按照密钥对名称过滤。参数不支持同时指定 KeyIds 和 Filters。
	Filters []*Filter `name:"Filters,omitempty"`
	// 密钥对ID，密钥对ID形如：skey-11112222（此接口支持同时传入多个ID进行过滤。此参数的具体格式可参考 API 简介的 id.N 一节）。参数不支持同时指定 KeyIds 和 Filters。密钥对ID可以通过登录控制台查询。
	KeyIds []*string `name:"KeyIds,omitempty"`
	// 返回数量，默认为20，最大值为100。关于 Limit 的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。关于 Offset 的更进一步介绍请参考 API 简介中的相关小节。返回数量，默认为20，最大值为100。关于 Limit 的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeKeyPairsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeKeyPairsResponse, error) {
	resp := &DescribeKeyPairsResponse{}
	err := client.Request("cvm", "DescribeKeyPairs", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeKeyPairsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 密钥对详细信息列表。
	KeyPairSet []*KeyPair `json:"KeyPairSet"`
	// 符合条件的密钥对数量。
	TotalCount int64 `json:"TotalCount"`
}
