package partners

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询客户余额
// https://cloud.tencent.com/document/api/563/19924

type DescribeClientBalanceRequest struct {
	// 客户(代客)账号ID
	ClientUin string `name:"ClientUin"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeClientBalanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeClientBalanceResponse, error) {
	resp := &DescribeClientBalanceResponse{}
	err := client.Request("partners", "DescribeClientBalance", "2018-03-21").Do(req, resp)
	return resp, err
}

type DescribeClientBalanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 账户余额，单位分
	Balance int64 `json:"Balance"`
}
