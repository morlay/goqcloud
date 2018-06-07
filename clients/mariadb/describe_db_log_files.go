package mariadb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取日志列表
// https://cloud.tencent.com/document/api/237/16162

type DescribeDbLogFilesRequest struct {
	// 实例 ID，形如：tdsql-ow728lmc。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
	// 请求日志类型，取值只能为1、2、3或者4。1-binlog，2-冷备，3-errlog，4-slowlog。
	Type int64 `name:"Type"`
}

func (req *DescribeDbLogFilesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDbLogFilesResponse, error) {
	resp := &DescribeDbLogFilesResponse{}
	err := client.Request("mariadb", "DescribeDBLogFiles", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeDbLogFilesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 包含uri、length、mtime（修改时间）等信息
	Files []*LogFileInfo `json:"Files"`
	// 实例 ID，形如：tdsql-ow728lmc。
	InstanceId string `json:"InstanceId"`
	// 如果是普通网络的实例，做用本前缀加上URI为下载地址
	Normalprefix string `json:"Normalprefix"`
	// 请求日志总数
	Total int64 `json:"Total"`
	// 请求日志类型，取值只能为1、2、3或者4。1-binlog，2-冷备，3-errlog，4-slowlog。
	Type int64 `json:"Type"`
	// 如果是VPC网络的实例，做用本前缀加上URI为下载地址
	Vpcprefix string `json:"Vpcprefix"`
}
