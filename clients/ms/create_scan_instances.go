package ms

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 批量提交安全扫描
// https://cloud.tencent.com/document/api/283/17758

type CreateScanInstancesRequest struct {
	// 待扫描的app信息列表，一次最多提交20个
	AppInfos []*AppInfo `name:"AppInfos"`
	// 区域
	Region string `name:"Region"`
	// 扫描信息
	ScanInfo ScanInfo `name:"ScanInfo"`
}

func (req *CreateScanInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateScanInstancesResponse, error) {
	resp := &CreateScanInstancesResponse{}
	err := client.Request("ms", "CreateScanInstances", "2018-04-08").Do(req, resp)
	return resp, err
}

type CreateScanInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 提交成功的app的md5集合
	AppMd5s []*string `json:"AppMd5s"`
	// 任务唯一标识
	ItemId string `json:"ItemId"`
	// 剩余可用次数
	LimitCount int64 `json:"LimitCount"`
	// 到期时间
	LimitTime int64 `json:"LimitTime"`
	// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	Progress int64 `json:"Progress"`
}
