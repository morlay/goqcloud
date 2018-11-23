package cws

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 生成漏洞报告
// https://cloud.tencent.com/document/api/692/18089

type CreateVulsReportRequest struct {
	// 监控任务ID
	MonitorId *int64 `name:"MonitorId,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 站点ID
	SiteId *int64 `name:"SiteId,omitempty"`
}

func (req *CreateVulsReportRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateVulsReportResponse, error) {
	resp := &CreateVulsReportResponse{}
	err := client.Request("cws", "CreateVulsReport", "2018-03-12").Do(req, resp)
	return resp, err
}

type CreateVulsReportResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 报告下载地址
	ReportFileUrl string `json:"ReportFileUrl"`
}
