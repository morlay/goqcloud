package dcdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询分布式数据库可售卖分片规格
// https://cloud.tencent.com/document/api/557/16134

type DescribeShardSpecRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeShardSpecRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeShardSpecResponse, error) {
	resp := &DescribeShardSpecResponse{}
	err := client.Request("dcdb", "DescribeShardSpec", "2018-04-11").Do(req, resp)
	return resp, err
}

type DescribeShardSpecResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 按机型分类的可售卖规格列表
	SpecConfig []*SpecConfig `json:"SpecConfig"`
}
