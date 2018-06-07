package postgres

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询售卖规格配置
// https://cloud.tencent.com/document/api/409/16776

type DescribeProductConfigRequest struct {
	// 区域
	Region string `name:"Region"`
	// 可用区名称
	Zone *string `name:"Zone,omitempty"`
}

func (req *DescribeProductConfigRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeProductConfigResponse, error) {
	resp := &DescribeProductConfigResponse{}
	err := client.Request("postgres", "DescribeProductConfig", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeProductConfigResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 售卖规格列表。
	SpecInfoList []*SpecInfo `json:"SpecInfoList"`
}
