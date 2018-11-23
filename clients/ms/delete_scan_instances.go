package ms

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 批量删除提交过的app扫描信息
// https://cloud.tencent.com/document/api/283/17757

type DeleteScanInstancesRequest struct {
	// 删除一个或多个扫描的app，最大支持20个
	AppSids []*string `name:"AppSids"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteScanInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteScanInstancesResponse, error) {
	resp := &DeleteScanInstancesResponse{}
	err := client.Request("ms", "DeleteScanInstances", "2018-04-08").Do(req, resp)
	return resp, err
}

type DeleteScanInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	Progress int64 `json:"Progress"`
}
