package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询分库分表下载位点
// https://cloud.tencent.com/document/api/236/15838
type DescribeBackupDownloadDbTableCodeRequest struct {
	// 待下载的数据库和数据表列表。
	DatabaseTableList []*DatabaseTableList `name:"DatabaseTableList"`
	// 实例ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例ID相同。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
	// 开始时间，格式为：2017-07-12 10:29:20。
	StartTime string `name:"StartTime"`
}

func (req *DescribeBackupDownloadDbTableCodeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeBackupDownloadDbTableCodeResponse, error) {
	resp := &DescribeBackupDownloadDbTableCodeResponse{}
	err := client.Request("cdb", "DescribeBackupDownloadDbTableCode", "2017-03-20").Do(req, resp)
	return resp, err
}

type DescribeBackupDownloadDbTableCodeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 下载位点
	DatabaseTableCode int64 `json:"DatabaseTableCode"`
}
