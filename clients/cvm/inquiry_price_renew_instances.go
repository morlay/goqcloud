package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 续费实例询价
// https://cloud.tencent.com/document/api/213/15725

type InquiryPriceRenewInstancesRequest struct {
	// 试运行。
	DryRun *bool `name:"DryRun,omitempty"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的续费时长、是否设置自动续费等属性。
	InstanceChargePrepaid InstanceChargePrepaid `name:"InstanceChargePrepaid"`
	// 一个或多个待操作的实例ID。可通过DescribeInstances接口返回值中的InstanceId获取。每次请求批量实例的上限为100。
	InstanceIds []*string `name:"InstanceIds"`
	// 区域
	Region string `name:"Region"`
	// 是否续费弹性数据盘。取值范围：TRUE：表示续费包年包月实例同时续费其挂载的弹性数据盘FALSE：表示续费包年包月实例同时不再续费其挂载的弹性数据盘默认取值：TRUE。
	RenewPortableDataDisk *bool `name:"RenewPortableDataDisk,omitempty"`
}

func (req *InquiryPriceRenewInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*InquiryPriceRenewInstancesResponse, error) {
	resp := &InquiryPriceRenewInstancesResponse{}
	err := client.Request("cvm", "InquiryPriceRenewInstances", "2017-03-12").Do(req, resp)
	return resp, err
}

type InquiryPriceRenewInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 该参数表示对应配置实例的价格。
	Price Price `json:"Price"`
}
