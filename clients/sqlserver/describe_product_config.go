package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询售卖规格配置
// https://cloud.tencent.com/document/api/238/19966

type DescribeProductConfigRequest struct {
	// 区域
	Region string `name:"Region"`
	// 可用区英文ID，形如ap-guangzhou-1
	Zone string `name:"Zone"`
}

func (req *DescribeProductConfigRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeProductConfigResponse, error) {
	resp := &DescribeProductConfigResponse{}
	err := client.Request("sqlserver", "DescribeProductConfig", "2018-03-28").Do(req, resp)
	return resp, err
}

type DescribeProductConfigResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 规格信息数组
	SpecInfoList []*SpecInfo `json:"SpecInfoList"`
	// 返回总共多少条数据
	TotalCount int64 `json:"TotalCount"`
}
