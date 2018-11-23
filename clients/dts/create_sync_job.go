package dts

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建灾备同步任务
// https://cloud.tencent.com/document/api/571/18573

type CreateSyncJobRequest struct {
	// 需要同步的源数据库表信息，用json格式的字符串描述。对于database-table两级结构的数据库：[{Database:db1,Table:[table1,table2]},{Database:db2}]
	DatabaseInfo *string `name:"DatabaseInfo,omitempty"`
	// 目标实例接入类型，目前仅包括：cdb(云上cdb实例)
	DstAccessType string `name:"DstAccessType"`
	// 目标实例数据库类型，目前仅包括：mysql
	DstDatabaseType string `name:"DstDatabaseType"`
	// 目标实例信息
	DstInfo SyncInstanceInfo `name:"DstInfo"`
	// 灾备同步任务名
	JobName string `name:"JobName"`
	// 区域
	Region string `name:"Region"`
	// 源实例接入类型，目前仅包括：cdb(云上cdb实例)
	SrcAccessType string `name:"SrcAccessType"`
	// 源实例数据库类型，目前仅包括：mysql
	SrcDatabaseType string `name:"SrcDatabaseType"`
	// 源实例信息
	SrcInfo SyncInstanceInfo `name:"SrcInfo"`
	// 灾备同步任务配置选项
	SyncOption SyncOption `name:"SyncOption"`
}

func (req *CreateSyncJobRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateSyncJobResponse, error) {
	resp := &CreateSyncJobResponse{}
	err := client.Request("dts", "CreateSyncJob", "2018-03-30").Do(req, resp)
	return resp, err
}

type CreateSyncJobResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 灾备同步任务ID
	JobId string `json:"JobId"`
}
