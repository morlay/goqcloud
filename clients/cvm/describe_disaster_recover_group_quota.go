package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询置放群组配额
// https://cloud.tencent.com/document/api/213/17811

type DescribeDisasterRecoverGroupQuotaRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeDisasterRecoverGroupQuotaRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDisasterRecoverGroupQuotaResponse, error) {
	resp := &DescribeDisasterRecoverGroupQuotaResponse{}
	err := client.Request("cvm", "DescribeDisasterRecoverGroupQuota", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeDisasterRecoverGroupQuotaResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 当前用户已经创建的置放群组数量。
	CurrentNum int64 `json:"CurrentNum"`
	// 物理机类型容灾组内实例的配额数。
	CvmInHostGroupQuota int64 `json:"CvmInHostGroupQuota"`
	// 机架类型容灾组内实例的配额数。
	CvmInRackGroupQuota int64 `json:"CvmInRackGroupQuota"`
	// 交换机类型容灾组内实例的配额数。
	CvmInSwGroupQuota int64 `json:"CvmInSwGroupQuota"`
	// 可创建置放群组数量的上限。
	GroupQuota int64 `json:"GroupQuota"`
}
