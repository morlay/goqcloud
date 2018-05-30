package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建CDH实例
// https://cloud.tencent.com/document/api/213/16473
type AllocateHostsRequest struct {
	// 用于保证请求幂等性的字符串。
	ClientToken *string `name:"ClientToken,omitempty"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	HostChargePrepaid *ChargePrepaid `name:"HostChargePrepaid,omitempty"`
	// 实例计费类型。目前仅支持：PREPAID（预付费，即包年包月模式）。
	HostChargeType *string `name:"HostChargeType,omitempty"`
	// 购买CDH实例数量。
	HostCount *int64 `name:"HostCount,omitempty"`
	// CDH实例机型，默认为：'HS1'。
	HostType *string `name:"HostType,omitempty"`
	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目等属性。
	Placement Placement `name:"Placement"`
	// 区域
	Region string `name:"Region"`
}

func (req *AllocateHostsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*AllocateHostsResponse, error) {
	resp := &AllocateHostsResponse{}
	err := client.Request("cvm", "AllocateHosts", "2017-03-12").Do(req, resp)
	return resp, err
}

type AllocateHostsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 新创建云子机的实例id列表。
	HostIdSet []*string `json:"HostIdSet"`
}
