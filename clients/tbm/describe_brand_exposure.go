package tbm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 获取品牌总热度
// https://cloud.tencent.com/document/api/853/18398

type DescribeBrandExposureRequest struct {
	// 品牌ID
	BrandId string `name:"BrandId"`
	// 查询结束时间
	EndDate time.Time `name:"EndDate"`
	// 区域
	Region string `name:"Region"`
	// 查询开始时间
	StartDate time.Time `name:"StartDate"`
}

func (req *DescribeBrandExposureRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeBrandExposureResponse, error) {
	resp := &DescribeBrandExposureResponse{}
	err := client.Request("tbm", "DescribeBrandExposure", "2018-01-29").Do(req, resp)
	return resp, err
}

type DescribeBrandExposureResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 按天计算的统计数据
	DateCountSet []*DateCount `json:"DateCountSet"`
	// 累计曝光量
	TotalCount int64 `json:"TotalCount"`
}
