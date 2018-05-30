package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除子网
// https://cloud.tencent.com/document/api/215/15783
type DeleteSubnetRequest struct {
	// 区域
	Region string `name:"Region"`
	// 子网实例ID。可通过DescribeSubnets接口返回值中的SubnetId获取。
	SubnetId string `name:"SubnetId"`
}

func (req *DeleteSubnetRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteSubnetResponse, error) {
	resp := &DeleteSubnetResponse{}
	err := client.Request("vpc", "DeleteSubnet", "2017-03-12").Do(req, resp)
	return resp, err
}

type DeleteSubnetResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
