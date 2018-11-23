package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询账户属性
// https://cloud.tencent.com/document/api/215/17875

type DescribeAccountAttributesRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeAccountAttributesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeAccountAttributesResponse, error) {
	resp := &DescribeAccountAttributesResponse{}
	err := client.Request("vpc", "DescribeAccountAttributes", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeAccountAttributesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 用户账号属性对象
	AccountAttributeSet []*AccountAttribute `json:"AccountAttributeSet"`
}
