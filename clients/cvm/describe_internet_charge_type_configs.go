package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询网络计费类型
// https://cloud.tencent.com/document/api/213/15729

type DescribeInternetChargeTypeConfigsRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeInternetChargeTypeConfigsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeInternetChargeTypeConfigsResponse, error) {
	resp := &DescribeInternetChargeTypeConfigsResponse{}
	err := client.Request("cvm", "DescribeInternetChargeTypeConfigs", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeInternetChargeTypeConfigsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 网络计费类型配置。
	InternetChargeTypeConfigSet []*InternetChargeTypeConfig `json:"InternetChargeTypeConfigSet"`
}
