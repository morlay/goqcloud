package billing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 支付订单
// https://cloud.tencent.com/document/api/555/19178

type PayDealsRequest struct {
	// 是否自动使用代金券,1:是,0否,默认0
	AutoVoucher *int64 `name:"AutoVoucher,omitempty"`
	// 需要支付的一个或者多个订单号
	OrderIds []*string `name:"OrderIds"`
	// 区域
	Region string `name:"Region"`
	// 代金券ID列表,目前仅支持指定一张代金券
	VoucherIds []*string `name:"VoucherIds,omitempty"`
}

func (req *PayDealsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*PayDealsResponse, error) {
	resp := &PayDealsResponse{}
	err := client.Request("billing", "PayDeals", "2018-07-09").Do(req, resp)
	return resp, err
}

type PayDealsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 此次操作支付成功的订单号数组
	OrderIds []*string `json:"OrderIds"`
	// 此次操作支付成功的资源Id数组
	ResourceIds []*string `json:"ResourceIds"`
}
