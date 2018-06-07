package mariadb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除账号
// https://cloud.tencent.com/document/api/237/16171

type DeleteAccountRequest struct {
	// 用户允许的访问 host
	Host string `name:"Host"`
	// 实例ID，形如：tdsql-ow728lmc，可以通过 DescribeDBInstances 查询实例详情获得。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
	// 用户名
	UserName string `name:"UserName"`
}

func (req *DeleteAccountRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteAccountResponse, error) {
	resp := &DeleteAccountResponse{}
	err := client.Request("mariadb", "DeleteAccount", "2017-03-12").Do(req, resp)
	return resp, err
}

type DeleteAccountResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
