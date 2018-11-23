package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建HAVIP
// https://cloud.tencent.com/document/api/215/30652

type CreateHaVipRequest struct {
	// HAVIP名称。
	HaVipName string `name:"HaVipName"`
	// 区域
	Region string `name:"Region"`
	// HAVIP所在子网ID。
	SubnetId string `name:"SubnetId"`
	// 指定虚拟IP地址，必须在VPC网段内且未被占用。不指定则自动分配。
	Vip *string `name:"Vip,omitempty"`
	// HAVIP所在私有网络ID。
	VpcId string `name:"VpcId"`
}

func (req *CreateHaVipRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateHaVipResponse, error) {
	resp := &CreateHaVipResponse{}
	err := client.Request("vpc", "CreateHaVip", "2017-03-12").Do(req, resp)
	return resp, err
}

type CreateHaVipResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// HAVIP对象。
	HaVip HaVip `json:"HaVip"`
}
