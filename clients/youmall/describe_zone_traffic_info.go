package youmall

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 获取区域客流人次及停留时间
// https://cloud.tencent.com/document/api/860/18459

type DescribeZoneTrafficInfoRequest struct {
	// 公司ID
	CompanyId string `name:"CompanyId"`
	// 结束日期，格式yyyy-MM-dd
	EndDate time.Time `name:"EndDate"`
	// Limit:每页的数据项，最大100，超过100会被强制指定为100
	Limit int64 `name:"Limit"`
	// 偏移量：分页控制参数，第一页传0，第n页Offset=(n-1)*Limit
	Offset int64 `name:"Offset"`
	// 区域
	Region string `name:"Region"`
	// 店铺ID
	ShopId int64 `name:"ShopId"`
	// 开始日期，格式yyyy-MM-dd
	StartDate time.Time `name:"StartDate"`
}

func (req *DescribeZoneTrafficInfoRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeZoneTrafficInfoResponse, error) {
	resp := &DescribeZoneTrafficInfoResponse{}
	err := client.Request("youmall", "DescribeZoneTrafficInfo", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeZoneTrafficInfoResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 公司ID
	CompanyId string `json:"CompanyId"`
	// 门店ID
	ShopId int64 `json:"ShopId"`
	// 查询结果总数
	TotalCount int64 `json:"TotalCount"`
	// 区域客流信息列表
	ZoneTrafficInfoSet []*ZoneTrafficInfo `json:"ZoneTrafficInfoSet"`
}
