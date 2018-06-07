package mariadb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 扩容实例询价
// https://cloud.tencent.com/document/api/237/16183

type DescribeUpgradePriceRequest struct {
	// 待升级的实例ID。形如：tdsql-ow728lmc，可以通过 DescribeDBInstances 查询实例详情获得。
	InstanceId string `name:"InstanceId"`
	// 内存大小，单位：GB，可以通过 DescribeDBInstanceSpecs 查询实例规格获得。
	Memory int64 `name:"Memory"`
	// 区域
	Region string `name:"Region"`
	// 存储空间大小，单位：GB，可以通过 DescribeDBInstanceSpecs 查询实例规格获得不同内存大小对应的磁盘规格下限和上限。
	Storage int64 `name:"Storage"`
}

func (req *DescribeUpgradePriceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeUpgradePriceResponse, error) {
	resp := &DescribeUpgradePriceResponse{}
	err := client.Request("mariadb", "DescribeUpgradePrice", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeUpgradePriceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 原价，单位：分
	OriginalPrice int64 `json:"OriginalPrice"`
	// 实际价格，单位：分。受折扣等影响，可能和原价不同。
	Price int64 `json:"Price"`
}
