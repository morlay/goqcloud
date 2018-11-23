package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 根据备份文件恢复实例
// https://cloud.tencent.com/document/api/238/19950

type RestoreInstanceRequest struct {
	// 备份文件ID，该ID可以通过DescribeBackups接口返回数据中的Id字段获得
	BackupId int64 `name:"BackupId"`
	// 实例ID，形如mssql-j8kv137v
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *RestoreInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*RestoreInstanceResponse, error) {
	resp := &RestoreInstanceResponse{}
	err := client.Request("sqlserver", "RestoreInstance", "2018-03-28").Do(req, resp)
	return resp, err
}

type RestoreInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步流程任务ID，使用FlowId调用DescribeFlowStatus接口获取任务执行状态
	FlowId int64 `json:"FlowId"`
}
