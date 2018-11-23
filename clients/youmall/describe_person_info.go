package youmall

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取顾客详情列表
// https://cloud.tencent.com/document/api/860/18464

type DescribePersonInfoRequest struct {
	// 公司ID
	CompanyId string `name:"CompanyId"`
	// Limit:每页的数据项，最大100，超过100会被强制指定为100
	Limit int64 `name:"Limit"`
	// 偏移量：分页控制参数，第一页传0，第n页Offset=(n-1)*Limit
	Offset int64 `name:"Offset"`
	// 身份类型(0表示普通顾客，1 白名单，2 表示黑名单）
	PersonType *int64 `name:"PersonType,omitempty"`
	// 图片url过期时间：在当前时间+PictureExpires秒后，图片url无法继续正常访问；单位s；默认值12460*60（1天）
	PictureExpires *int64 `name:"PictureExpires,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 门店ID
	ShopId int64 `name:"ShopId"`
	// 起始ID，第一次拉取时StartPersonId传0，后续送入的值为上一页最后一条数据项的PersonId
	StartPersonId int64 `name:"StartPersonId"`
}

func (req *DescribePersonInfoRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribePersonInfoResponse, error) {
	resp := &DescribePersonInfoResponse{}
	err := client.Request("youmall", "DescribePersonInfo", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribePersonInfoResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 公司ID
	CompanyId string `json:"CompanyId"`
	// 用户信息
	PersonInfoSet []*PersonInfo `json:"PersonInfoSet"`
	// 门店ID
	ShopId int64 `json:"ShopId"`
	// 总数
	TotalCount int64 `json:"TotalCount"`
}
