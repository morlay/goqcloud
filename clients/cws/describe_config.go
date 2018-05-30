package cws

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 查看用户配置列表
// https://cloud.tencent.com/document/api/692/16757
type DescribeConfigRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeConfigRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeConfigResponse, error) {
	resp := &DescribeConfigResponse{}
	err := client.Request("cws", "DescribeConfig", "2018-03-12").Do(req, resp)
	return resp, err
}

type DescribeConfigResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 记录创建时间。
	CreatedAt time.Time `json:"CreatedAt"`
	// 配置ID。
	Id int64 `json:"Id"`
	// 漏洞告警通知等级，4位分别代表：高危、中危、低危、提示。
	NoticeLevel string `json:"NoticeLevel"`
	// 记录更新新建。
	UpdatedAt time.Time `json:"UpdatedAt"`
}
