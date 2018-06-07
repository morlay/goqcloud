package cws

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 新增监测任务
// https://cloud.tencent.com/document/api/692/16743

type CreateMonitorsRequest struct {
	// 扫描周期，单位小时，每X小时执行一次
	Crontab int64 `name:"Crontab"`
	// 首次扫描开始时间
	FirstScanStartTime time.Time `name:"FirstScanStartTime"`
	// 任务名称
	Name string `name:"Name"`
	// 扫描速率限制，每秒发送X个HTTP请求
	RateLimit int64 `name:"RateLimit"`
	// 区域
	Region string `name:"Region"`
	// 扫描模式，normal-正常扫描；deep-深度扫描
	ScannerType string `name:"ScannerType"`
	// 站点的url列表
	Urls []*string `name:"Urls"`
}

func (req *CreateMonitorsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateMonitorsResponse, error) {
	resp := &CreateMonitorsResponse{}
	err := client.Request("cws", "CreateMonitors", "2018-03-12").Do(req, resp)
	return resp, err
}

type CreateMonitorsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
