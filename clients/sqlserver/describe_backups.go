package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 查询备份列表
// https://cloud.tencent.com/document/api/238/19943

type DescribeBackupsRequest struct {
	// 结束时间(yyyy-MM-dd HH:mm:ss)
	EndTime time.Time `name:"EndTime"`
	// 实例ID，形如mssql-njj2mtpl
	InstanceId string `name:"InstanceId"`
	// 分页返回，每页返回数量，默认为20，最大值为 100
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为 0
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 开始时间(yyyy-MM-dd HH:mm:ss)
	StartTime time.Time `name:"StartTime"`
}

func (req *DescribeBackupsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeBackupsResponse, error) {
	resp := &DescribeBackupsResponse{}
	err := client.Request("sqlserver", "DescribeBackups", "2018-03-28").Do(req, resp)
	return resp, err
}

type DescribeBackupsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 备份列表详情
	Backups []*Backup `json:"Backups"`
	// 备份总数量
	TotalCount int64 `json:"TotalCount"`
}
