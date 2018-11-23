package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 拉取实例账户列表
// https://cloud.tencent.com/document/api/238/19970

type DescribeAccountsRequest struct {
	// 实例ID
	InstanceId string `name:"InstanceId"`
	// 分页返回，每页返回的数目，取值为1-100，默认值为20
	Limit *int64 `name:"Limit,omitempty"`
	// 分页返回，从第几页开始返回。从第0页开始，默认第0页
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeAccountsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeAccountsResponse, error) {
	resp := &DescribeAccountsResponse{}
	err := client.Request("sqlserver", "DescribeAccounts", "2018-03-28").Do(req, resp)
	return resp, err
}

type DescribeAccountsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 账户信息列表
	Accounts []*AccountDetail `json:"Accounts"`
	// 实例ID
	InstanceId string `json:"InstanceId"`
}
