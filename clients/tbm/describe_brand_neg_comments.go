package tbm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 获取品牌用户差评列表
// https://cloud.tencent.com/document/api/853/18396

type DescribeBrandNegCommentsRequest struct {
	// 品牌ID
	BrandId string `name:"BrandId"`
	// 查询结束时间
	EndDate time.Time `name:"EndDate"`
	// 查询条数上限，默认20
	Limit *int64 `name:"Limit,omitempty"`
	// 查询偏移，默认从0开始
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 查询开始时间
	StartDate time.Time `name:"StartDate"`
}

func (req *DescribeBrandNegCommentsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeBrandNegCommentsResponse, error) {
	resp := &DescribeBrandNegCommentsResponse{}
	err := client.Request("tbm", "DescribeBrandNegComments", "2018-01-29").Do(req, resp)
	return resp, err
}

type DescribeBrandNegCommentsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 评论列表
	BrandCommentSet []*CommentInfo `json:"BrandCommentSet"`
	// 总的差评个数
	TotalComments int64 `json:"TotalComments"`
}
