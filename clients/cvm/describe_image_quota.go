package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询镜像配额上限
// https://cloud.tencent.com/document/api/213/15719

type DescribeImageQuotaRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeImageQuotaRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeImageQuotaResponse, error) {
	resp := &DescribeImageQuotaResponse{}
	err := client.Request("cvm", "DescribeImageQuota", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeImageQuotaResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 账户的镜像配额
	ImageNumQuota int64 `json:"ImageNumQuota"`
}
