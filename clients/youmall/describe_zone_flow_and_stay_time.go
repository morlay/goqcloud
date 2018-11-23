package youmall

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取区域人流和停留时间
// https://cloud.tencent.com/document/api/860/20417

type DescribeZoneFlowAndStayTimeRequest struct {
	// 集团ID
	CompanyId string `name:"CompanyId"`
	// 结束日期，格式yyyy-MM-dd
	EndDate string `name:"EndDate"`
	// 区域
	Region string `name:"Region"`
	// 店铺ID
	ShopId int64 `name:"ShopId"`
	// 开始日期，格式yyyy-MM-dd
	StartDate string `name:"StartDate"`
}

func (req *DescribeZoneFlowAndStayTimeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeZoneFlowAndStayTimeResponse, error) {
	resp := &DescribeZoneFlowAndStayTimeResponse{}
	err := client.Request("youmall", "DescribeZoneFlowAndStayTime", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeZoneFlowAndStayTimeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 集团id
	CompanyId string `json:"CompanyId"`
	// 各区域人流数目和停留时长
	Data []*ZoneFlowAndAvrStayTime `json:"Data"`
	// 店铺id
	ShopId int64 `json:"ShopId"`
}
