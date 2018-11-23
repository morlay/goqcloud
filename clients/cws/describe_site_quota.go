package cws

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查看站点购买配额
// https://cloud.tencent.com/document/api/692/16751

type DescribeSiteQuotaRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeSiteQuotaRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeSiteQuotaResponse, error) {
	resp := &DescribeSiteQuotaResponse{}
	err := client.Request("cws", "DescribeSiteQuota", "2018-03-12").Do(req, resp)
	return resp, err
}

type DescribeSiteQuotaResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 剩余可用的扫描次数。
	Available int64 `json:"Available"`
	// 已购买的扫描次数。
	Total int64 `json:"Total"`
	// 已使用的扫描次数。
	Used int64 `json:"Used"`
}
