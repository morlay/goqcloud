package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询云数据库的所有账号信息
// https://cloud.tencent.com/document/api/236/17499

type DescribeAccountsRequest struct {
	// 实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId string `name:"InstanceId"`
	// 单次请求返回的数量，默认值为20，最小值为1，最大值为100。
	Limit *int64 `name:"Limit,omitempty"`
	// 记录偏移量，默认值为0。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeAccountsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeAccountsResponse, error) {
	resp := &DescribeAccountsResponse{}
	err := client.Request("cdb", "DescribeAccounts", "2017-03-20").Do(req, resp)
	return resp, err
}

type DescribeAccountsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 符合查询条件的账号详细信息。
	Items []*AccountInfo `json:"Items"`
	// 符合查询条件的账号数量。
	TotalCount int64 `json:"TotalCount"`
}
