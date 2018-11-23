package youmall

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取客户所属门店列表
// https://cloud.tencent.com/document/api/860/18457

type DescribeShopInfoRequest struct {
	// Limit:每页的数据项，最大100，超过100会被强制指定为100
	Limit int64 `name:"Limit"`
	// 偏移量：分页控制参数，第一页传0，第n页Offset=(n-1)*Limit
	Offset int64 `name:"Offset"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeShopInfoRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeShopInfoResponse, error) {
	resp := &DescribeShopInfoResponse{}
	err := client.Request("youmall", "DescribeShopInfo", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeShopInfoResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 门店列表信息
	ShopInfoSet []*ShopInfo `json:"ShopInfoSet"`
	// 门店总数
	TotalCount int64 `json:"TotalCount"`
}
