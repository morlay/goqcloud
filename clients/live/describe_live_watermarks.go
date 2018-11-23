package live

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询水印列表
// https://cloud.tencent.com/document/api/267/30152

type DescribeLiveWatermarksRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeLiveWatermarksRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeLiveWatermarksResponse, error) {
	resp := &DescribeLiveWatermarksResponse{}
	err := client.Request("live", "DescribeLiveWatermarks", "2018-08-01").Do(req, resp)
	return resp, err
}

type DescribeLiveWatermarksResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 水印总个数。
	TotalNum int64 `json:"TotalNum"`
	// 水印信息列表。
	WatermarkList []*WatermarkInfo `json:"WatermarkList"`
}
