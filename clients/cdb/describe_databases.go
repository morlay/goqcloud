package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询数据库
// https://cloud.tencent.com/document/api/236/17493

type DescribeDatabasesRequest struct {
	// 匹配数据库库名的正则表达式，规则同MySQL官网
	DatabaseRegexp *string `name:"DatabaseRegexp,omitempty"`
	// 实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId string `name:"InstanceId"`
	// 单次请求数量，默认值为20，最小值为1，最大值为100。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，最小值为0。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeDatabasesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDatabasesResponse, error) {
	resp := &DescribeDatabasesResponse{}
	err := client.Request("cdb", "DescribeDatabases", "2017-03-20").Do(req, resp)
	return resp, err
}

type DescribeDatabasesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 返回的实例信息。
	Items []*string `json:"Items"`
	// 符合查询条件的实例总数。
	TotalCount int64 `json:"TotalCount"`
}
