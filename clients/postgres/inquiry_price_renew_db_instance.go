package postgres

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询实例续费价格
// https://cloud.tencent.com/document/api/409/18103

type InquiryPriceRenewDbInstanceRequest struct {
	// 实例ID
	DBInstanceId string `name:"DBInstanceId"`
	// 续费周期，按月计算，最大不超过48
	Period int64 `name:"Period"`
	// 区域
	Region string `name:"Region"`
}

func (req *InquiryPriceRenewDbInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*InquiryPriceRenewDbInstanceResponse, error) {
	resp := &InquiryPriceRenewDbInstanceResponse{}
	err := client.Request("postgres", "InquiryPriceRenewDBInstance", "2017-03-12").Do(req, resp)
	return resp, err
}

type InquiryPriceRenewDbInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 总费用，打折前的。比如24650表示246.5元
	OriginalPrice int64 `json:"OriginalPrice"`
	// 实际需要付款金额。比如24650表示246.5元
	Price int64 `json:"Price"`
}
