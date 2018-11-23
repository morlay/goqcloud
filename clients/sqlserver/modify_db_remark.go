package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改数据库备注
// https://cloud.tencent.com/document/api/238/19954

type ModifyDbRemarkRequest struct {
	// 数据库名称及备注数组，每个元素包含数据库名和对应的备注
	DBRemarks []*DBRemark `name:"DBRemarks"`
	// 实例ID，形如mssql-rljoi3bf
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyDbRemarkRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyDbRemarkResponse, error) {
	resp := &ModifyDbRemarkResponse{}
	err := client.Request("sqlserver", "ModifyDBRemark", "2018-03-28").Do(req, resp)
	return resp, err
}

type ModifyDbRemarkResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
