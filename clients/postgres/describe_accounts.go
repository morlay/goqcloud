package postgres

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取实例用户列表
// https://cloud.tencent.com/document/api/409/18109

type DescribeAccountsRequest struct {
	// 实例ID，形如postgres-6fego161
	DBInstanceId string `name:"DBInstanceId"`
	// 分页返回，每页最大返回数目，默认20，取值范围为1-100
	Limit *int64 `name:"Limit,omitempty"`
	// 分页返回，返回第几页的用户数据。页码从0开始计数
	Offset *int64 `name:"Offset,omitempty"`
	// 返回数据按照创建时间或者用户名排序。取值只能为createTime或者name。createTime-按照创建时间排序；name-按照用户名排序
	OrderBy *string `name:"OrderBy,omitempty"`
	// 返回结果是升序还是降序。取值只能为desc或者asc。desc-降序；asc-升序
	OrderByType *string `name:"OrderByType,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeAccountsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeAccountsResponse, error) {
	resp := &DescribeAccountsResponse{}
	err := client.Request("postgres", "DescribeAccounts", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeAccountsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 帐号列表详细信息。
	Details []*AccountInfo `json:"Details"`
	// 本次调用接口共返回了多少条数据。
	TotalCount int64 `json:"TotalCount"`
}
