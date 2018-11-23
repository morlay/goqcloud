package youmall

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取指定区域每日客流量
// https://cloud.tencent.com/document/api/860/20416

type DescribeZoneFlowDailyByZoneIDRequest struct {
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
	// 区域ID
	ZoneId int64 `name:"ZoneId"`
}

func (req *DescribeZoneFlowDailyByZoneIDRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeZoneFlowDailyByZoneIDResponse, error) {
	resp := &DescribeZoneFlowDailyByZoneIDResponse{}
	err := client.Request("youmall", "DescribeZoneFlowDailyByZoneId", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeZoneFlowDailyByZoneIDResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 集团id
	CompanyId string `json:"CompanyId"`
	// 每日人流量
	Data []*ZoneDayFlow `json:"Data"`
	// 店铺id
	ShopId int64 `json:"ShopId"`
	// 区域ID
	ZoneId int64 `json:"ZoneId"`
	// 区域名称
	ZoneName string `json:"ZoneName"`
}
