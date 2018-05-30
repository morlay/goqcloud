package mariadb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询云数据库可售卖规格
// https://cloud.tencent.com/document/api/237/16188
type DescribeDbInstanceSpecsRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeDbInstanceSpecsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDbInstanceSpecsResponse, error) {
	resp := &DescribeDbInstanceSpecsResponse{}
	err := client.Request("mariadb", "DescribeDBInstanceSpecs", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeDbInstanceSpecsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 按机型分类的可售卖规格列表
	Specs []*InstanceSpec `json:"Specs"`
}
