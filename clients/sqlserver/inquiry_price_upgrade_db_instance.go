package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询升级实例价格
// https://cloud.tencent.com/document/api/238/19960

type InquiryPriceUpgradeDbInstanceRequest struct {
	// 实例ID，形如mssql-njj2mtpl
	InstanceId string `name:"InstanceId"`
	// 实例升级后的内存大小，单位GB，其值不能比当前实例内存小
	Memory int64 `name:"Memory"`
	// 区域
	Region string `name:"Region"`
	// 实例升级后的磁盘大小，单位GB，其值不能比当前实例磁盘小
	Storage int64 `name:"Storage"`
}

func (req *InquiryPriceUpgradeDbInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*InquiryPriceUpgradeDbInstanceResponse, error) {
	resp := &InquiryPriceUpgradeDbInstanceResponse{}
	err := client.Request("sqlserver", "InquiryPriceUpgradeDBInstance", "2018-03-28").Do(req, resp)
	return resp, err
}

type InquiryPriceUpgradeDbInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 未打折的原价，其值除以100表示最终的价格。比如10094表示100.94元
	OriginalPrice int64 `json:"OriginalPrice"`
	// 实际需要支付价格，其值除以100表示最终的价格。比如10094表示100.94元
	Price int64 `json:"Price"`
}
