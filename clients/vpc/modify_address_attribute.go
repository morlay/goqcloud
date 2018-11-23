package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改弹性公网IP属性
// https://cloud.tencent.com/document/api/215/16704

type ModifyAddressAttributeRequest struct {
	// 标识 EIP 的唯一 ID。EIP 唯一 ID 形如：eip-11112222。
	AddressId string `name:"AddressId"`
	// 修改后的 EIP 名称。长度上限为20个字符。
	AddressName *string `name:"AddressName,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyAddressAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyAddressAttributeResponse, error) {
	resp := &ModifyAddressAttributeResponse{}
	err := client.Request("vpc", "ModifyAddressAttribute", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyAddressAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
