package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建密钥对
// https://cloud.tencent.com/document/api/213/15702
type CreateKeyPairRequest struct {
	// 密钥对名称，可由数字，字母和下划线组成，长度不超过25个字符。
	KeyName string `name:"KeyName"`
	// 密钥对创建后所属的项目ID。可以通过以下方式获取项目ID：通过项目列表查询项目ID。通过调用接口DescribeProject，取返回信息中的projectId获取项目ID。
	ProjectId int64 `name:"ProjectId"`
	// 区域
	Region string `name:"Region"`
}

func (req *CreateKeyPairRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateKeyPairResponse, error) {
	resp := &CreateKeyPairResponse{}
	err := client.Request("cvm", "CreateKeyPair", "2017-03-12").Do(req, resp)
	return resp, err
}

type CreateKeyPairResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 密钥对信息。
	KeyPair KeyPair `json:"KeyPair"`
}
