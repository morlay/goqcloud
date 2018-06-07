package cws

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 修改监测任务的属性
// https://cloud.tencent.com/document/api/692/16746

type ModifyMonitorAttributeRequest struct {
	// 扫描周期，单位小时，每X小时执行一次
	Crontab int64 `name:"Crontab"`
	// 首次扫描开始时间
	FirstScanStartTime time.Time `name:"FirstScanStartTime"`
	// 监测任务ID
	MonitorId int64 `name:"MonitorId"`
	// 监测状态：1-监测中；2-暂停监测
	MonitorStatus int64 `name:"MonitorStatus"`
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

func (req *ModifyMonitorAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyMonitorAttributeResponse, error) {
	resp := &ModifyMonitorAttributeResponse{}
	err := client.Request("cws", "ModifyMonitorAttribute", "2018-03-12").Do(req, resp)
	return resp, err
}

type ModifyMonitorAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
