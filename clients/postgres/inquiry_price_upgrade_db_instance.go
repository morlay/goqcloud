package postgres

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询升级实例价格
// https://cloud.tencent.com/document/api/409/18102

type InquiryPriceUpgradeDbInstanceRequest struct {
	// 实例ID，形如postgres-hez4fh0v
	DBInstanceId string `name:"DBInstanceId"`
	// 实例计费类型，预付费或者后付费。PREPAID-预付费。目前只支持预付费。
	InstanceChargeType *string `name:"InstanceChargeType,omitempty"`
	// 实例的内存大小，单位GB
	Memory int64 `name:"Memory"`
	// 区域
	Region string `name:"Region"`
	// 实例的磁盘大小，单位GB
	Storage int64 `name:"Storage"`
}

func (req *InquiryPriceUpgradeDbInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*InquiryPriceUpgradeDbInstanceResponse, error) {
	resp := &InquiryPriceUpgradeDbInstanceResponse{}
	err := client.Request("postgres", "InquiryPriceUpgradeDBInstance", "2017-03-12").Do(req, resp)
	return resp, err
}

type InquiryPriceUpgradeDbInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 总费用，打折前的
	OriginalPrice int64 `json:"OriginalPrice"`
	// 实际需要付款金额
	Price int64 `json:"Price"`
}
