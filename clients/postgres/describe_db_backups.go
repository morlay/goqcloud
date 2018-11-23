package postgres

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 查询实例备份列表
// https://cloud.tencent.com/document/api/409/18094

type DescribeDbBackupsRequest struct {
	// 实例ID，形如postgres-4wdeb0zv。
	DBInstanceId string `name:"DBInstanceId"`
	// 查询结束时间，形如2018-06-10 17:06:38
	EndTime time.Time `name:"EndTime"`
	// 备份列表分页返回，每页返回数量，默认为 20，最小为1，最大值为 100。
	Limit *int64 `name:"Limit,omitempty"`
	// 返回结果中的第几页，从第0页开始。默认为0。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 查询开始时间，形如2018-06-10 17:06:38，起始时间不得小于7天以前
	StartTime time.Time `name:"StartTime"`
	// 备份方式（1-全量）。目前只支持全量，取值为1。
	Type int64 `name:"Type"`
}

func (req *DescribeDbBackupsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDbBackupsResponse, error) {
	resp := &DescribeDbBackupsResponse{}
	err := client.Request("postgres", "DescribeDBBackups", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeDbBackupsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 备份列表
	BackupList []*DBBackup `json:"BackupList"`
	// 返回备份列表中备份文件的个数
	TotalCount int64 `json:"TotalCount"`
}
