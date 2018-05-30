package partners

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询待审核客户列表
// https://cloud.tencent.com/document/api/563/16046
type DescribeAgentClientsRequest struct {
	// 客户类型，a/b，类型定义参考代理商相关政策文档
	ClientFlag *string `name:"ClientFlag,omitempty"`
	// 客户名称。由于涉及隐私，名称打码显示，故名称仅支持打码后的模糊搜索
	ClientName *string `name:"ClientName,omitempty"`
	// 客户账号ID
	ClientUin *string `name:"ClientUin,omitempty"`
	// 限制数目
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量
	Offset *int64 `name:"Offset,omitempty"`
	// ASC/DESC， 不区分大小写，按申请时间排序
	OrderDirection *string `name:"OrderDirection,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeAgentClientsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeAgentClientsResponse, error) {
	resp := &DescribeAgentClientsResponse{}
	err := client.Request("partners", "DescribeAgentClients", "2018-03-21").Do(req, resp)
	return resp, err
}

type DescribeAgentClientsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 代客列表
	AgentClientSet []*AgentClientElem `json:"AgentClientSet"`
	// 符合条件的代客总数
	TotalCount int64 `json:"TotalCount"`
}
