package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 调整实例带宽上限询价
// https://cloud.tencent.com/document/api/213/15732

type InquiryPriceResetInstancesInternetMaxBandwidthRequest struct {
	// 带宽生效的终止时间。格式：YYYY-MM-DD，例如：2016-10-30。新设置的带宽的有效期包含终止时间此日期。终止时间不能晚于包年包月实例的到期时间。实例的到期时间可通过DescribeInstances接口返回值中的ExpiredTime获取。该参数只对包年包月带宽有效，其他模式带宽不支持该参数，否则接口会以相应错误码返回。
	EndTime *string `name:"EndTime,omitempty"`
	// 一个或多个待操作的实例ID。可通过DescribeInstances接口返回值中的InstanceId获取。每次请求批量实例的上限为100。当调整 BANDWIDTH_PREPAID 和 BANDWIDTH_POSTPAID_BY_HOUR 计费方式的带宽时，只支持一个实例。
	InstanceIds []*string `name:"InstanceIds"`
	// 公网出带宽配置。不同机型带宽上限范围不一致，具体限制详见带宽限制对账表。暂时只支持InternetMaxBandwidthOut参数。
	InternetAccessible InternetAccessible `name:"InternetAccessible"`
	// 区域
	Region string `name:"Region"`
	// 带宽生效的起始时间。格式：YYYY-MM-DD，例如：2016-10-30。起始时间不能早于当前时间。如果起始时间是今天则新设置的带宽立即生效。该参数只对包年包月带宽有效，其他模式带宽不支持该参数，否则接口会以相应错误码返回。
	StartTime *string `name:"StartTime,omitempty"`
}

func (req *InquiryPriceResetInstancesInternetMaxBandwidthRequest) Invoke(client github_com_morlay_goqcloud.Client) (*InquiryPriceResetInstancesInternetMaxBandwidthResponse, error) {
	resp := &InquiryPriceResetInstancesInternetMaxBandwidthResponse{}
	err := client.Request("cvm", "InquiryPriceResetInstancesInternetMaxBandwidth", "2017-03-12").Do(req, resp)
	return resp, err
}

type InquiryPriceResetInstancesInternetMaxBandwidthResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 该参数表示带宽调整为对应大小之后的价格。
	Price Price `json:"Price"`
}
