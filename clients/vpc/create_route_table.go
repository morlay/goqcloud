package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建路由表
// https://cloud.tencent.com/document/api/215/15768

type CreateRouteTableRequest struct {
	// 区域
	Region string `name:"Region"`
	// 路由表名称，最大长度不能超过60个字节。
	RouteTableName string `name:"RouteTableName"`
	// 待操作的VPC实例ID。可通过DescribeVpcs接口返回值中的VpcId获取。
	VpcId string `name:"VpcId"`
}

func (req *CreateRouteTableRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateRouteTableResponse, error) {
	resp := &CreateRouteTableResponse{}
	err := client.Request("vpc", "CreateRouteTable", "2017-03-12").Do(req, resp)
	return resp, err
}

type CreateRouteTableResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 路由表对象。
	RouteTable RouteTable `json:"RouteTable"`
}
