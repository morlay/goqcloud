package postgres

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询实例详情
// https://cloud.tencent.com/document/api/409/16772
type DescribeDbInstanceAttributeRequest struct {
	// 实例ID。
	DBInstanceId string `name:"DBInstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeDbInstanceAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDbInstanceAttributeResponse, error) {
	resp := &DescribeDbInstanceAttributeResponse{}
	err := client.Request("postgres", "DescribeDBInstanceAttribute", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeDbInstanceAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 实例详细信息。
	DBInstance DBInstance `json:"DBInstance"`
}
