package dcdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询分布式数据库可售卖地域和可用区信息
// https://cloud.tencent.com/document/api/557/16141
type DescribeDcdbSaleInfoRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeDcdbSaleInfoRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDcdbSaleInfoResponse, error) {
	resp := &DescribeDcdbSaleInfoResponse{}
	err := client.Request("dcdb", "DescribeDCDBSaleInfo", "2018-04-11").Do(req, resp)
	return resp, err
}

type DescribeDcdbSaleInfoResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 可售卖地域信息列表
	RegionList []*RegionInfo `json:"RegionList"`
}
