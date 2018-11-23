package youmall

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取指定区域人流各年龄占比
// https://cloud.tencent.com/document/api/860/20418

type DescribeZoneFlowAgeInfoByZoneIDRequest struct {
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

func (req *DescribeZoneFlowAgeInfoByZoneIDRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeZoneFlowAgeInfoByZoneIDResponse, error) {
	resp := &DescribeZoneFlowAgeInfoByZoneIDResponse{}
	err := client.Request("youmall", "DescribeZoneFlowAgeInfoByZoneId", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeZoneFlowAgeInfoByZoneIDResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 集团ID
	CompanyId string `json:"CompanyId"`
	// 当前年龄段占比
	Data []*float64 `json:"Data"`
	// 店铺ID
	ShopId int64 `json:"ShopId"`
	// 区域ID
	ZoneId int64 `json:"ZoneId"`
	// 区域名称
	ZoneName string `json:"ZoneName"`
}
