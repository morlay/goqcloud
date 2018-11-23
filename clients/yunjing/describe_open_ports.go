package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取端口列表
// https://cloud.tencent.com/document/api/296/30327

type DescribeOpenPortsRequest struct {
	// 过滤条件。Port - Uint64 - 是否必填：否 - 端口号ProcessName - String - 是否必填：否 - 进程名MachineIp - String - 是否必填：否 - 主机内网IP
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为10，最大值为100。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。
	Offset *int64 `name:"Offset,omitempty"`
	// 开放端口号。Port和Uuid必填其一，使用Port表示查询该端口的列表信息。
	Port *int64 `name:"Port,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 云镜客户端唯一Uuid。Port和Uuid必填其一，使用Uuid表示，查询该主机列表信息。
	Uuid *string `name:"Uuid,omitempty"`
}

func (req *DescribeOpenPortsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeOpenPortsResponse, error) {
	resp := &DescribeOpenPortsResponse{}
	err := client.Request("yunjing", "DescribeOpenPorts", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeOpenPortsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 端口列表。
	OpenPorts []*OpenPort `json:"OpenPorts"`
	// 端口列表记录总数。
	TotalCount int64 `json:"TotalCount"`
}
