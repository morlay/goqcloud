package youmall

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询设备历史网络状态
// https://cloud.tencent.com/document/api/860/19897

type DescribeHistoryNetworkInfoRequest struct {
	// 优mall集团id，通过"指定身份标识获取客户门店列表"接口获取
	CompanyId string `name:"CompanyId"`
	// 拉取结束日期，格式L:2018-09-05，超过StartDay 90天，按StartDay+90天算
	EndDay string `name:"EndDay"`
	// 拉取条数，默认10
	Limit *int64 `name:"Limit,omitempty"`
	// 拉取偏移，返回offset之后的数据
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 优mall店铺id，通过"指定身份标识获取客户门店列表"接口获取，为0则拉取集团全部店铺当前
	ShopId int64 `name:"ShopId"`
	// 拉取开始日期，格式：2018-09-05
	StartDay string `name:"StartDay"`
	// 请求时间戳
	Time int64 `name:"Time"`
}

func (req *DescribeHistoryNetworkInfoRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeHistoryNetworkInfoResponse, error) {
	resp := &DescribeHistoryNetworkInfoResponse{}
	err := client.Request("youmall", "DescribeHistoryNetworkInfo", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeHistoryNetworkInfoResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 网络状态数据
	InstanceSet NetworkHistoryInfo `json:"InstanceSet"`
}
