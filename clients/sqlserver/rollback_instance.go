package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 回档实例
// https://cloud.tencent.com/document/api/238/19949

type RollbackInstanceRequest struct {
	// 需要回档的数据库
	DBs []*string `name:"DBs"`
	// 实例ID
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
	// 回档目标时间点
	Time time.Time `name:"Time"`
	// 回档类型，0-回档的数据库覆盖原库；1-回档的数据库以重命名的形式生成，不覆盖原库
	Type int64 `name:"Type"`
}

func (req *RollbackInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*RollbackInstanceResponse, error) {
	resp := &RollbackInstanceResponse{}
	err := client.Request("sqlserver", "RollbackInstance", "2018-03-28").Do(req, resp)
	return resp, err
}

type RollbackInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步任务ID
	FlowId int64 `json:"FlowId"`
}
