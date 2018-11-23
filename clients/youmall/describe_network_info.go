package youmall

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询设备最新网络状态
// https://cloud.tencent.com/document/api/860/19896

type DescribeNetworkInfoRequest struct {
	// 优mall集团id，通过"指定身份标识获取客户门店列表"接口获取
	CompanyId string `name:"CompanyId"`
	// 区域
	Region string `name:"Region"`
	// 优mall店铺id，通过"指定身份标识获取客户门店列表"接口获取，不填则拉取集团全部店铺当前
	ShopId *int64 `name:"ShopId,omitempty"`
	// 请求时间戳
	Time int64 `name:"Time"`
}

func (req *DescribeNetworkInfoRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeNetworkInfoResponse, error) {
	resp := &DescribeNetworkInfoResponse{}
	err := client.Request("youmall", "DescribeNetworkInfo", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeNetworkInfoResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 网络状态详情
	InstanceSet NetworkLastInfo `json:"InstanceSet"`
}
