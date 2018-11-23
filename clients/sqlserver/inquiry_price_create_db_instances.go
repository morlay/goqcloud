package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询申请实例价格
// https://cloud.tencent.com/document/api/238/19962

type InquiryPriceCreateDbInstancesRequest struct {
	// sqlserver版本，目前只支持：2008R2（SQL Server 2008 R2），2012SP3（SQL Server 2012），2016SP1（SQL Server 2016 SP1）两种版本。默认为2008R2版本
	DBVersion *string `name:"DBVersion,omitempty"`
	// 一次性购买的实例数量。取值1-100，默认取值为1
	GoodsNum *int64 `name:"GoodsNum,omitempty"`
	// 计费类型，当前只支持预付费，即包年包月，取值为PREPAID。默认值为PREPAID
	InstanceChargeType *string `name:"InstanceChargeType,omitempty"`
	// 内存大小，单位：GB
	Memory int64 `name:"Memory"`
	// 购买时长，单位：月。取值为1到48，默认为1
	Period *int64 `name:"Period,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 实例容量大小，单位：GB。
	Storage int64 `name:"Storage"`
	// 可用区ID。该参数可以通过调用 DescribeZones 接口的返回值中的Zone字段来获取。
	Zone string `name:"Zone"`
}

func (req *InquiryPriceCreateDbInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*InquiryPriceCreateDbInstancesResponse, error) {
	resp := &InquiryPriceCreateDbInstancesResponse{}
	err := client.Request("sqlserver", "InquiryPriceCreateDBInstances", "2018-03-28").Do(req, resp)
	return resp, err
}

type InquiryPriceCreateDbInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 未打折前价格，其值除以100表示多少钱。比如10010表示100.10元
	OriginalPrice int64 `json:"OriginalPrice"`
	// 实际需要支付的价格，其值除以100表示多少钱。比如10010表示100.10元
	Price int64 `json:"Price"`
}
