package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除分散置放群组
// https://cloud.tencent.com/document/api/213/17812

type DeleteDisasterRecoverGroupsRequest struct {
	// 分散置放群组ID列表，可通过DescribeDisasterRecoverGroups接口获取。
	DisasterRecoverGroupIds []*string `name:"DisasterRecoverGroupIds"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteDisasterRecoverGroupsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteDisasterRecoverGroupsResponse, error) {
	resp := &DeleteDisasterRecoverGroupsResponse{}
	err := client.Request("cvm", "DeleteDisasterRecoverGroups", "2017-03-12").Do(req, resp)
	return resp, err
}

type DeleteDisasterRecoverGroupsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
