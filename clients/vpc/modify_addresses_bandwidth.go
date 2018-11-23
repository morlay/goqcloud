package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 调整弹性公网ip带宽
// https://cloud.tencent.com/document/api/215/19214

type ModifyAddressesBandwidthRequest struct {
	// EIP唯一标识id，形如'eip-xxxx'
	AddressIds []*string `name:"AddressIds"`
	// 包月带宽结束时间
	EndTime *time.Time `name:"EndTime,omitempty"`
	// 调整带宽目标值
	InternetMaxBandwidthOut int64 `name:"InternetMaxBandwidthOut"`
	// 区域
	Region string `name:"Region"`
	// 包月带宽起始时间
	StartTime *time.Time `name:"StartTime,omitempty"`
}

func (req *ModifyAddressesBandwidthRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyAddressesBandwidthResponse, error) {
	resp := &ModifyAddressesBandwidthResponse{}
	err := client.Request("vpc", "ModifyAddressesBandwidth", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyAddressesBandwidthResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
