package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 获取慢查询日志文件信息
// https://cloud.tencent.com/document/api/238/19936

type DescribeSlowlogsRequest struct {
	// 查询结束时间
	EndTime time.Time `name:"EndTime"`
	// 实例ID，形如mssql-k8voqdlz
	InstanceId string `name:"InstanceId"`
	// 分页返回结果，分页大小，默认20，不超过100
	Limit *int64 `name:"Limit,omitempty"`
	// 从第几页开始返回，起始页，从0开始，默认为0
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 查询开始时间
	StartTime time.Time `name:"StartTime"`
}

func (req *DescribeSlowlogsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeSlowlogsResponse, error) {
	resp := &DescribeSlowlogsResponse{}
	err := client.Request("sqlserver", "DescribeSlowlogs", "2018-03-28").Do(req, resp)
	return resp, err
}

type DescribeSlowlogsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 慢查询日志信息列表
	Slowlogs []*SlowlogInfo `json:"Slowlogs"`
	// 查询总数
	TotalCount int64 `json:"TotalCount"`
}
