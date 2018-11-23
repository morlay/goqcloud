package tbm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 获取品牌用户好评列表
// https://cloud.tencent.com/document/api/853/18395

type DescribeBrandPosCommentsRequest struct {
	// 品牌ID
	BrandId string `name:"BrandId"`
	// 查询结束时间
	EndDate time.Time `name:"EndDate"`
	// 查询条数上限，默认20
	Limit *int64 `name:"Limit,omitempty"`
	// 查询偏移，从0开始
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 查询开始时间
	StartDate time.Time `name:"StartDate"`
}

func (req *DescribeBrandPosCommentsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeBrandPosCommentsResponse, error) {
	resp := &DescribeBrandPosCommentsResponse{}
	err := client.Request("tbm", "DescribeBrandPosComments", "2018-01-29").Do(req, resp)
	return resp, err
}

type DescribeBrandPosCommentsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 评论列表
	BrandCommentSet []*CommentInfo `json:"BrandCommentSet"`
	// 总的好评个数
	TotalComments int64 `json:"TotalComments"`
}
