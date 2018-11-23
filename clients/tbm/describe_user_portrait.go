package tbm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取品牌画像结果
// https://cloud.tencent.com/document/api/853/18392

type DescribeUserPortraitRequest struct {
	// 品牌ID
	BrandId string `name:"BrandId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeUserPortraitRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeUserPortraitResponse, error) {
	resp := &DescribeUserPortraitResponse{}
	err := client.Request("tbm", "DescribeUserPortrait", "2018-01-29").Do(req, resp)
	return resp, err
}

type DescribeUserPortraitResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 年龄画像
	Age AgePortraitInfo `json:"Age"`
	// 性别画像
	Gender GenderPortraitInfo `json:"Gender"`
	// 电影喜好画像
	Movie MoviePortraitInfo `json:"Movie"`
	// 省份画像
	Province ProvincePortraitInfo `json:"Province"`
	// 明星喜好画像
	Star StarPortraitInfo `json:"Star"`
}
