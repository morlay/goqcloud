package youmall

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取指定区域不同年龄段男女平均停留时间
// https://cloud.tencent.com/document/api/860/20415

type DescribeZoneFlowGenderAvrStayTimeByZoneIDRequest struct {
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

func (req *DescribeZoneFlowGenderAvrStayTimeByZoneIDRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeZoneFlowGenderAvrStayTimeByZoneIDResponse, error) {
	resp := &DescribeZoneFlowGenderAvrStayTimeByZoneIDResponse{}
	err := client.Request("youmall", "DescribeZoneFlowGenderAvrStayTimeByZoneId", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeZoneFlowGenderAvrStayTimeByZoneIDResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 集团ID
	CompanyId string `json:"CompanyId"`
	// 不同年龄段男女停留时间（返回格式为数组，从第 1 个到最后一个数据，年龄段分别为 0-17，18 - 23,  24 - 30, 31 - 40, 41 - 50, 51 - 60, 61 - 100）
	Data []*ZoneAgeGroupAvrStayTime `json:"Data"`
	// 店铺ID
	ShopId int64 `json:"ShopId"`
	// 区域ID
	ZoneId int64 `json:"ZoneId"`
	// 区域名称
	ZoneName string `json:"ZoneName"`
}
