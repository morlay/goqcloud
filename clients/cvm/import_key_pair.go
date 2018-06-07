package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 导入密钥对
// https://cloud.tencent.com/document/api/213/15703

type ImportKeyPairRequest struct {
	// 密钥对名称，可由数字，字母和下划线组成，长度不超过25个字符。
	KeyName string `name:"KeyName"`
	// 密钥对创建后所属的项目ID。可以通过以下方式获取项目ID：通过项目列表查询项目ID。通过调用接口 DescribeProject，取返回信息中的 projectId 获取项目ID。如果是默认项目，直接填0就可以。
	ProjectId int64 `name:"ProjectId"`
	// 密钥对的公钥内容，OpenSSH RSA 格式。
	PublicKey string `name:"PublicKey"`
	// 区域
	Region string `name:"Region"`
}

func (req *ImportKeyPairRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ImportKeyPairResponse, error) {
	resp := &ImportKeyPairResponse{}
	err := client.Request("cvm", "ImportKeyPair", "2017-03-12").Do(req, resp)
	return resp, err
}

type ImportKeyPairResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 密钥对ID。
	KeyId string `json:"KeyId"`
}
