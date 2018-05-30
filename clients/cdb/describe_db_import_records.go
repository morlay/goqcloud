package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询数据库导入任务记录
// https://cloud.tencent.com/document/api/236/15856
type DescribeDbImportRecordsRequest struct {
	// 结束时间，时间格式如：2016-01-01 23:59:59。
	EndTime *string `name:"EndTime,omitempty"`
	// 实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId string `name:"InstanceId"`
	// 分页参数 , 单次请求返回的数量 , 默认值为20。
	Limit *int64 `name:"Limit,omitempty"`
	// 分页参数 , 偏移量 , 默认值为0。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 开始时间，时间格式如：2016-01-01 00:00:01。
	StartTime *string `name:"StartTime,omitempty"`
}

func (req *DescribeDbImportRecordsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDbImportRecordsResponse, error) {
	resp := &DescribeDbImportRecordsResponse{}
	err := client.Request("cdb", "DescribeDBImportRecords", "2017-03-20").Do(req, resp)
	return resp, err
}

type DescribeDbImportRecordsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 返回的导入操作记录列表。
	Items []*ImportRecord `json:"Items"`
	// 符合查询条件的导入任务操作日志总数。
	TotalCount int64 `json:"TotalCount"`
}
