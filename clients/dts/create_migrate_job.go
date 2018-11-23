package dts

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建数据迁移任务
// https://cloud.tencent.com/document/api/571/18141

type CreateMigrateJobRequest struct {
	// 需要迁移的源数据库表信息，用json格式的字符串描述。对于database-table两级结构的数据库：[{Database:db1,Table:[table1,table2]},{Database:db2}]对于database-schema-table三级结构：[{Database:db1,Schema:s1Table:[table1,table2]},{Database:db1,Schema:s2Table:[table1,table2]},{Database:db2,Schema:s1Table:[table1,table2]},{Database:db3},{Database:db4Schema:s1}]
	DatabaseInfo *string `name:"DatabaseInfo,omitempty"`
	// 目标实例接入类型，值包括：extranet(外网),cvm(cvm自建实例),dcg(专线接入的实例),vpncloud(云vpn接入的实例),vpnselfbuild(自建vpn接入的实例)，cdb(云上cdb实例). 目前只支持cdb.
	DstAccessType string `name:"DstAccessType"`
	// 目标实例数据库类型,mysql,redis,mongodb
	DstDatabaseType string `name:"DstDatabaseType"`
	// 目标实例信息
	DstInfo DstInfo `name:"DstInfo"`
	// 数据迁移任务名称
	JobName string `name:"JobName"`
	// 迁移任务配置选项
	MigrateOption MigrateOption `name:"MigrateOption"`
	// 区域
	Region string `name:"Region"`
	// 源实例接入类型，值包括：extranet(外网),cvm(cvm自建实例),dcg(专线接入的实例),vpncloud(云vpn接入的实例),vpnselfbuild(自建vpn接入的实例)，cdb(云上cdb实例)
	SrcAccessType string `name:"SrcAccessType"`
	// 源实例数据库类型:mysql,redis,mongodb
	SrcDatabaseType string `name:"SrcDatabaseType"`
	// 源实例信息，具体内容跟迁移任务类型相关
	SrcInfo SrcInfo `name:"SrcInfo"`
}

func (req *CreateMigrateJobRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateMigrateJobResponse, error) {
	resp := &CreateMigrateJobResponse{}
	err := client.Request("dts", "CreateMigrateJob", "2018-03-30").Do(req, resp)
	return resp, err
}

type CreateMigrateJobResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 数据迁移任务ID
	JobId string `json:"JobId"`
}
