package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询云数据库实例的字符集
// https://cloud.tencent.com/document/api/236/15866

type DescribeDbInstanceCharsetRequest struct {
	// 实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同，可使用查询实例列表 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeDbInstanceCharsetRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDbInstanceCharsetResponse, error) {
	resp := &DescribeDbInstanceCharsetResponse{}
	err := client.Request("cdb", "DescribeDBInstanceCharset", "2017-03-20").Do(req, resp)
	return resp, err
}

type DescribeDbInstanceCharsetResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 实例的默认字符集，如"latin1", "utf8"等。
	Charset string `json:"Charset"`
}
