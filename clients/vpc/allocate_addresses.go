package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建弹性公网IP
// https://cloud.tencent.com/document/api/215/16699

type AllocateAddressesRequest struct {
	// 申请 EIP 数量，默认值为1。
	AddressCount *int64 `name:"AddressCount,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *AllocateAddressesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*AllocateAddressesResponse, error) {
	resp := &AllocateAddressesResponse{}
	err := client.Request("vpc", "AllocateAddresses", "2017-03-12").Do(req, resp)
	return resp, err
}

type AllocateAddressesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 申请到的 EIP 的唯一 ID 列表。
	AddressSet []*string `json:"AddressSet"`
}
