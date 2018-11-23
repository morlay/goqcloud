package billing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 查询订单数据
// https://cloud.tencent.com/document/api/555/19179

type DescribeDealsByCondRequest struct {
	// 结束时间
	EndTime time.Time `name:"EndTime"`
	// 一页多少条数据，默认是20条，最大不超过1000
	Limit int64 `name:"Limit"`
	// 第多少页，从0开始，默认是0
	Offset *int64 `name:"Offset,omitempty"`
	// 订单号
	OrderId *string `name:"OrderId,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 开始时间
	StartTime time.Time `name:"StartTime"`
	// 订单状态,默认为4（成功的订单）订单的状态1：未支付2：已支付3：发货中4：已发货5：发货失败6：已退款7：已关单8：订单过期9：订单已失效10：产品已失效11：代付拒绝12：支付中
	Status *int64 `name:"Status,omitempty"`
}

func (req *DescribeDealsByCondRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDealsByCondResponse, error) {
	resp := &DescribeDealsByCondResponse{}
	err := client.Request("billing", "DescribeDealsByCond", "2018-07-09").Do(req, resp)
	return resp, err
}

type DescribeDealsByCondResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 订单列表
	Deals []*Deal `json:"Deals"`
	// 订单总数
	TotalCount int64 `json:"TotalCount"`
}
