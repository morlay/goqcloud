package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 续费CDH实例
// https://cloud.tencent.com/document/api/213/16476
type RenewHostsRequest struct {
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	HostChargePrepaid ChargePrepaid `name:"HostChargePrepaid"`
	// 一个或多个待操作的CDH实例ID。
	HostIds []*string `name:"HostIds"`
	// 区域
	Region string `name:"Region"`
}

func (req *RenewHostsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*RenewHostsResponse, error) {
	resp := &RenewHostsResponse{}
	err := client.Request("cvm", "RenewHosts", "2017-03-12").Do(req, resp)
	return resp, err
}

type RenewHostsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
