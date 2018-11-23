package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询实例的可设置参数列表
// https://cloud.tencent.com/document/api/236/20411

type DescribeInstanceParamsRequest struct {
	// 实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同，可使用查询实例列表 接口获取，其值为输出参数中字段 InstanceId 的值
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeInstanceParamsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeInstanceParamsResponse, error) {
	resp := &DescribeInstanceParamsResponse{}
	err := client.Request("cdb", "DescribeInstanceParams", "2017-03-20").Do(req, resp)
	return resp, err
}

type DescribeInstanceParamsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 参数详情
	Items []*ParameterDetail `json:"Items"`
	// 实例的参数总数
	TotalCount int64 `json:"TotalCount"`
}
