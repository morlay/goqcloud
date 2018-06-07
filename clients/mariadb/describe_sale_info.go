package mariadb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询云数据库可售卖地域和可用区信息
// https://cloud.tencent.com/document/api/237/16178

type DescribeSaleInfoRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeSaleInfoRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeSaleInfoResponse, error) {
	resp := &DescribeSaleInfoResponse{}
	err := client.Request("mariadb", "DescribeSaleInfo", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeSaleInfoResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 可售卖地域信息列表
	RegionList []*RegionInfo `json:"RegionList"`
}
