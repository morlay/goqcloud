package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取区域主机列表
// https://cloud.tencent.com/document/api/296/19850

type DescribeMachinesRequest struct {
	// 过滤条件。Keywords - String - 是否必填：否 - 查询关键字 Status - String - 是否必填：否 - 客户端在线状态（OFFLINE: 离线 | ONLINE: 在线）Version - String  是否必填：否 - 当前防护版本（ PRO_VERSION：专业版 | BASIC_VERSION：基础版）每个过滤条件只支持一个值，暂不支持多个值“或”关系查询
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为10，最大值为100。
	Limit *int64 `name:"Limit,omitempty"`
	// 机器所属地域。如：ap-guangzhou，ap-shanghai
	MachineRegion string `name:"MachineRegion"`
	// 云主机类型。CVM：表示虚拟主机BM:  表示黑石物理机
	MachineType string `name:"MachineType"`
	// 偏移量，默认为0。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeMachinesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeMachinesResponse, error) {
	resp := &DescribeMachinesResponse{}
	err := client.Request("yunjing", "DescribeMachines", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeMachinesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 主机列表
	Machines []*Machine `json:"Machines"`
	// 主机数量
	TotalCount int64 `json:"TotalCount"`
}
