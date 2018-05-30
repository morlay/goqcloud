package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改密钥对属性
// https://cloud.tencent.com/document/api/213/15701
type ModifyKeyPairAttributeRequest struct {
	// 修改后的密钥对描述信息。可任意命名，但不得超过60个字符。
	Description *string `name:"Description,omitempty"`
	// 密钥对ID，密钥对ID形如：skey-xxxxxxxx。可以通过以下方式获取可用的密钥 ID：通过登录控制台查询密钥 ID。通过调用接口 DescribeKeyPairs ，取返回信息中的 KeyId 获取密钥对 ID。
	KeyId string `name:"KeyId"`
	// 修改后的密钥对名称，可由数字，字母和下划线组成，长度不超过25个字符。
	KeyName *string `name:"KeyName,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyKeyPairAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyKeyPairAttributeResponse, error) {
	resp := &ModifyKeyPairAttributeResponse{}
	err := client.Request("cvm", "ModifyKeyPairAttribute", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyKeyPairAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
