package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询可回档时间
// https://cloud.tencent.com/document/api/236/18726

type DescribeRollbackRangeTimeRequest struct {
	// 实例ID列表，单个实例Id的格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例ID相同。
	InstanceIds []*string `name:"InstanceIds"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeRollbackRangeTimeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeRollbackRangeTimeResponse, error) {
	resp := &DescribeRollbackRangeTimeResponse{}
	err := client.Request("cdb", "DescribeRollbackRangeTime", "2017-03-20").Do(req, resp)
	return resp, err
}

type DescribeRollbackRangeTimeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 返回的参数信息。
	Items []*InstanceRollbackRangeTime `json:"Items"`
	// 符合查询条件的实例总数。
	TotalCount int64 `json:"TotalCount"`
}
