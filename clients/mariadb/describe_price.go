package mariadb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 新购实例询价
// https://cloud.tencent.com/document/api/237/16175
type DescribePriceRequest struct {
	// 欲购买的数量，默认查询购买1个实例的价格。
	Count *int64 `name:"Count,omitempty"`
	// 内存大小，单位：GB，可以通过 DescribeDBInstanceSpecs 查询实例规格获得。
	Memory int64 `name:"Memory"`
	// 实例节点个数，可以通过 DescribeDBInstanceSpecs 查询实例规格获得。
	NodeCount int64 `name:"NodeCount"`
	// 欲购买的时长，单位：月。
	Period *int64 `name:"Period,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 存储空间大小，单位：GB，可以通过 DescribeDBInstanceSpecs 查询实例规格获得不同内存大小对应的磁盘规格下限和上限。
	Storage int64 `name:"Storage"`
	// 欲新购实例的可用区ID。
	Zone string `name:"Zone"`
}

func (req *DescribePriceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribePriceResponse, error) {
	resp := &DescribePriceResponse{}
	err := client.Request("mariadb", "DescribePrice", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribePriceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 原价，单位：分
	OriginalPrice int64 `json:"OriginalPrice"`
	// 实际价格，单位：分。受折扣等影响，可能和原价不同。
	Price int64 `json:"Price"`
}
