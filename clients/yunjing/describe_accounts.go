package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取帐号列表
// https://cloud.tencent.com/document/api/296/30339

type DescribeAccountsRequest struct {
	// 过滤条件。Username - String - 是否必填：否 - 帐号名Privilege - String - 是否必填：否 - 帐号类型（ORDINARY: 普通帐号 | SUPPER: 超级管理员帐号）MachineIp - String - 是否必填：否 - 主机内网IP
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为10，最大值为100。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 云镜客户端唯一Uuid。Username和Uuid必填其一，使用Username表示，查询该用户名下列表信息。
	Username *string `name:"Username,omitempty"`
	// 云镜客户端唯一Uuid。Username和Uuid必填其一，使用Uuid表示，查询该主机下列表信息。
	Uuid *string `name:"Uuid,omitempty"`
}

func (req *DescribeAccountsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeAccountsResponse, error) {
	resp := &DescribeAccountsResponse{}
	err := client.Request("yunjing", "DescribeAccounts", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeAccountsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 帐号数据列表。
	Accounts []*Account `json:"Accounts"`
	// 帐号列表记录总数。
	TotalCount int64 `json:"TotalCount"`
}
