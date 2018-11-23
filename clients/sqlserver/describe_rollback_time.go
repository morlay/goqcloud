package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询实例可回档时间范围
// https://cloud.tencent.com/document/api/238/19964

type DescribeRollbackTimeRequest struct {
	// 需要查询的数据库列表
	DBs []*string `name:"DBs"`
	// 实例ID
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeRollbackTimeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeRollbackTimeResponse, error) {
	resp := &DescribeRollbackTimeResponse{}
	err := client.Request("sqlserver", "DescribeRollbackTime", "2018-03-28").Do(req, resp)
	return resp, err
}

type DescribeRollbackTimeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 数据库可回档实例信息
	Details []*DbRollbackTimeInfo `json:"Details"`
}
