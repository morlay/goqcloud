package mariadb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查看数据库参数
// https://cloud.tencent.com/document/api/237/16154

type DescribeDbParametersRequest struct {
	// 实例 ID，形如：tdsql-ow728lmc。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeDbParametersRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDbParametersResponse, error) {
	resp := &DescribeDbParametersResponse{}
	err := client.Request("mariadb", "DescribeDBParameters", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeDbParametersResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 实例 ID，形如：tdsql-ow728lmc。
	InstanceId string `json:"InstanceId"`
	// 请求DB的当前参数值
	Params []*ParamDesc `json:"Params"`
}
