package dts

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改灾备同步任务
// https://cloud.tencent.com/document/api/571/18569

type ModifySyncJobRequest struct {
	// 当选择'指定库表'灾备同步的时候, 需要设置待同步的源数据库表信息,用符合json数组格式的字符串描述, 如下所例。对于database-table两级结构的数据库：[{"Database":"db1","Table":["table1","table2"]},{"Database":"db2"}]
	DatabaseInfo *string `name:"DatabaseInfo,omitempty"`
	// 待修改的灾备同步任务ID
	JobId string `name:"JobId"`
	// 灾备同步任务名称
	JobName *string `name:"JobName,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 灾备同步任务配置选项
	SyncOption *SyncOption `name:"SyncOption,omitempty"`
}

func (req *ModifySyncJobRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifySyncJobResponse, error) {
	resp := &ModifySyncJobResponse{}
	err := client.Request("dts", "ModifySyncJob", "2018-03-30").Do(req, resp)
	return resp, err
}

type ModifySyncJobResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
