package youmall

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 获取每小时客流人数
// https://cloud.tencent.com/document/api/860/18461

type DescribeShopHourTrafficInfoRequest struct {
	// 公司ID
	CompanyId string `name:"CompanyId"`
	// 结束日期，格式：yyyy-MM-dd
	EndDate time.Time `name:"EndDate"`
	// Limit:每页的数据项，最大100，超过100会被强制指定为100
	Limit int64 `name:"Limit"`
	// 偏移量：分页控制参数，第一页传0，第n页Offset=(n-1)*Limit
	Offset int64 `name:"Offset"`
	// 区域
	Region string `name:"Region"`
	// 门店ID
	ShopId int64 `name:"ShopId"`
	// 开始日期，格式：yyyy-MM-dd
	StartDate time.Time `name:"StartDate"`
}

func (req *DescribeShopHourTrafficInfoRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeShopHourTrafficInfoResponse, error) {
	resp := &DescribeShopHourTrafficInfoResponse{}
	err := client.Request("youmall", "DescribeShopHourTrafficInfo", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeShopHourTrafficInfoResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 公司ID
	CompanyId string `json:"CompanyId"`
	// 分时客流信息
	ShopHourTrafficInfoSet []*ShopHourTrafficInfo `json:"ShopHourTrafficInfoSet"`
	// 门店ID
	ShopId int64 `json:"ShopId"`
	// 查询结果总数
	TotalCount int64 `json:"TotalCount"`
}
