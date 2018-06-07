package mariadb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询账号列表
// https://cloud.tencent.com/document/api/237/16167

type DescribeAccountsRequest struct {
	// 实例ID，形如：tdsql-ow728lmc，可以通过 DescribeDBInstances 查询实例详情获得。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeAccountsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeAccountsResponse, error) {
	resp := &DescribeAccountsResponse{}
	err := client.Request("mariadb", "DescribeAccounts", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeAccountsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 实例ID，透传入参。
	InstanceId string `json:"InstanceId"`
	// 实例用户列表。
	Users []*DBAccount `json:"Users"`
}
