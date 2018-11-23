package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取云镜破解事件列表
// https://cloud.tencent.com/document/api/296/19838

type DescribeBruteAttacksRequest struct {
	// 过滤条件。Keywords - String - 是否必填：否 -  查询关键字Status - String - 是否必填：否 -  查询状态（FAILED：破解失败 |SUCCESS：破解成功）
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为10，最大值为100。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 客户端唯一Uuid。
	Uuid *string `name:"Uuid,omitempty"`
}

func (req *DescribeBruteAttacksRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeBruteAttacksResponse, error) {
	resp := &DescribeBruteAttacksResponse{}
	err := client.Request("yunjing", "DescribeBruteAttacks", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeBruteAttacksResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 暴力破解事件列表
	BruteAttacks []*BruteAttack `json:"BruteAttacks"`
	// 事件数量
	TotalCount int64 `json:"TotalCount"`
}
