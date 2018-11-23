package dcdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询账号列表
// https://cloud.tencent.com/document/api/557/19993

type DescribeAccountsRequest struct {
	// 实例ID，形如：dcdbt-ow728lmc。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeAccountsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeAccountsResponse, error) {
	resp := &DescribeAccountsResponse{}
	err := client.Request("dcdb", "DescribeAccounts", "2018-04-11").Do(req, resp)
	return resp, err
}

type DescribeAccountsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 实例ID，透传入参。
	InstanceId string `json:"InstanceId"`
	// 实例用户列表。
	Users []*DBAccount `json:"Users"`
}
