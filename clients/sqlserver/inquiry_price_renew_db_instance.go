package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询实例续费价格
// https://cloud.tencent.com/document/api/238/19961

type InquiryPriceRenewDbInstanceRequest struct {
	// 实例ID
	InstanceId string `name:"InstanceId"`
	// 续费周期。按月续费最多48个月。默认查询续费一个月的价格
	Period *int64 `name:"Period,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 续费周期单位。month表示按月续费，当前只支持按月付费查询价格
	TimeUnit *string `name:"TimeUnit,omitempty"`
}

func (req *InquiryPriceRenewDbInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*InquiryPriceRenewDbInstanceResponse, error) {
	resp := &InquiryPriceRenewDbInstanceResponse{}
	err := client.Request("sqlserver", "InquiryPriceRenewDBInstance", "2018-03-28").Do(req, resp)
	return resp, err
}

type InquiryPriceRenewDbInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 未打折的原价，其值除以100表示最终的价格。比如10094表示100.94元
	OriginalPrice int64 `json:"OriginalPrice"`
	// 实际需要支付价格，其值除以100表示最终的价格。比如10094表示100.94元
	Price int64 `json:"Price"`
}
