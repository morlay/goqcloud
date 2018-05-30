package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除密钥对
// https://cloud.tencent.com/document/api/213/15700
type DeleteKeyPairsRequest struct {
	// 一个或多个待操作的密钥对ID。每次请求批量密钥对的上限为100。可以通过以下方式获取可用的密钥ID：通过登录控制台查询密钥ID。通过调用接口 DescribeKeyPairs ，取返回信息中的 KeyId 获取密钥对ID。
	KeyIds []*string `name:"KeyIds"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteKeyPairsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteKeyPairsResponse, error) {
	resp := &DeleteKeyPairsResponse{}
	err := client.Request("cvm", "DeleteKeyPairs", "2017-03-12").Do(req, resp)
	return resp, err
}

type DeleteKeyPairsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
