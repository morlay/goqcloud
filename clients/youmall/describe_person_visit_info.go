package youmall

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 获取顾客到访信息列表
// https://cloud.tencent.com/document/api/860/18463

type DescribePersonVisitInfoRequest struct {
	// 公司ID
	CompanyId string `name:"CompanyId"`
	// 结束日期，格式yyyy-MM-dd，已废弃，请使用EndDateTime
	EndDate *time.Time `name:"EndDate,omitempty"`
	// 结束时间，格式yyyy-MM-dd HH:mm:ss
	EndDateTime *time.Time `name:"EndDateTime,omitempty"`
	// Limit:每页的数据项，最大100，超过100会被强制指定为100
	Limit int64 `name:"Limit"`
	// 偏移量：分页控制参数，第一页传0，第n页Offset=(n-1)*Limit
	Offset int64 `name:"Offset"`
	// 图片url过期时间：在当前时间+PictureExpires秒后，图片url无法继续正常访问；单位s；默认值12460*60（1天）
	PictureExpires *int64 `name:"PictureExpires,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 门店ID
	ShopId int64 `name:"ShopId"`
	// 开始日期，格式yyyy-MM-dd，已废弃，请使用StartDateTime
	StartDate *time.Time `name:"StartDate,omitempty"`
	// 开始时间，格式yyyy-MM-dd HH:mm:ss
	StartDateTime *time.Time `name:"StartDateTime,omitempty"`
}

func (req *DescribePersonVisitInfoRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribePersonVisitInfoResponse, error) {
	resp := &DescribePersonVisitInfoResponse{}
	err := client.Request("youmall", "DescribePersonVisitInfo", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribePersonVisitInfoResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 公司ID
	CompanyId string `json:"CompanyId"`
	// 用户到访明细
	PersonVisitInfoSet []*PersonVisitInfo `json:"PersonVisitInfoSet"`
	// 门店ID
	ShopId int64 `json:"ShopId"`
	// 总数
	TotalCount int64 `json:"TotalCount"`
}
