package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询数据库价格
// https://cloud.tencent.com/document/api/236/18566

type DescribeDbPriceRequest struct {
	// 实例数量，默认值为1, 最小值1，最大值为100
	GoodsNum int64 `name:"GoodsNum"`
	// 实例类型，默认为 master，支持值包括：master-表示主实例，ro-表示只读实例，dr-表示灾备实例
	InstanceRole *string `name:"InstanceRole,omitempty"`
	// 实例内存大小，单位：MB
	Memory int64 `name:"Memory"`
	// 付费类型，支持值包括：PRE_PAID - 包年包月，HOUR_PAID - 按量计费
	PayType string `name:"PayType"`
	// 实例时长，单位：月，最小值1，最大值为36；查询按量计费价格时，该字段无效
	Period int64 `name:"Period"`
	// 数据复制方式，默认为0，支持值包括：0-表示异步复制，1-表示半同步复制，2-表示强同步复制
	ProtectMode *int64 `name:"ProtectMode,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 实例硬盘大小，单位：GB
	Volume int64 `name:"Volume"`
	// 可用区信息，格式如"ap-guangzhou-1"
	Zone string `name:"Zone"`
}

func (req *DescribeDbPriceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDbPriceResponse, error) {
	resp := &DescribeDbPriceResponse{}
	err := client.Request("cdb", "DescribeDBPrice", "2017-03-20").Do(req, resp)
	return resp, err
}

type DescribeDbPriceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 实例原价，单位：分（人民币）
	OriginalPrice int64 `json:"OriginalPrice"`
	// 实例价格，单位：分（人民币）
	Price int64 `json:"Price"`
}
