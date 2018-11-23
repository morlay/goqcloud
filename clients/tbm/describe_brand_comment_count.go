package tbm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 获取品牌好差评数
// https://cloud.tencent.com/document/api/853/18399

type DescribeBrandCommentCountRequest struct {
	// 品牌ID
	BrandId string `name:"BrandId"`
	// 查询结束日期
	EndDate time.Time `name:"EndDate"`
	// 区域
	Region string `name:"Region"`
	// 查询开始日期
	StartDate time.Time `name:"StartDate"`
}

func (req *DescribeBrandCommentCountRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeBrandCommentCountResponse, error) {
	resp := &DescribeBrandCommentCountResponse{}
	err := client.Request("tbm", "DescribeBrandCommentCount", "2018-01-29").Do(req, resp)
	return resp, err
}

type DescribeBrandCommentCountResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 按天统计好评/差评数
	CommentSet []*Comment `json:"CommentSet"`
}
