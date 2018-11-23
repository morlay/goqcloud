package youmall

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 获取每日客流人数
// https://cloud.tencent.com/document/api/860/18460

type DescribeShopTrafficInfoRequest struct {
	// 公司ID
	CompanyId string `name:"CompanyId"`
	// 介绍日期，格式yyyy-MM-dd
	EndDate time.Time `name:"EndDate"`
	// Limit:每页的数据项，最大100，超过100会被强制指定为100
	Limit int64 `name:"Limit"`
	// 偏移量：分页控制参数，第一页传0，第n页Offset=(n-1)*Limit
	Offset int64 `name:"Offset"`
	// 区域
	Region string `name:"Region"`
	// 门店ID
	ShopId int64 `name:"ShopId"`
	// 开始日期，格式yyyy-MM-dd
	StartDate time.Time `name:"StartDate"`
}

func (req *DescribeShopTrafficInfoRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeShopTrafficInfoResponse, error) {
	resp := &DescribeShopTrafficInfoResponse{}
	err := client.Request("youmall", "DescribeShopTrafficInfo", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeShopTrafficInfoResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 公司ID
	CompanyId string `json:"CompanyId"`
	// 客流信息列表
	ShopDayTrafficInfoSet []*ShopDayTrafficInfo `json:"ShopDayTrafficInfoSet"`
	// 门店ID
	ShopId int64 `json:"ShopId"`
	// 查询结果总数
	TotalCount int64 `json:"TotalCount"`
}
