package dts

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改数据迁移任务
// https://cloud.tencent.com/document/api/571/18138

type ModifyMigrateJobRequest struct {
	// 当选择'指定库表'迁移的时候, 需要设置待迁移的源数据库表信息,用符合json数组格式的字符串描述, 如下所例。对于database-table两级结构的数据库：[{"Database":"db1","Table":["table1","table2"]},{"Database":"db2"}]对于database-schema-table三级结构：[{"Database":"db1","Schema":"s1","Table":["table1","table2"]},{"Database":"db1","Schema":"s2","Table":["table1","table2"]},{"Database":"db2","Schema":"s1","Table":["table1","table2"]},{"Database":"db3"},{"Database":"db4","Schema":"s1"}]如果是'整个实例'的迁移模式,不需设置该字段
	DatabaseInfo *string `name:"DatabaseInfo,omitempty"`
	// 目标实例接入类型，值包括：extranet(外网),cvm(cvm自建实例),dcg(专线接入的实例),vpncloud(云vpn接入的实例),vpnselfbuild(自建vpn接入的实例)，cdb(云上cdb实例). 目前只支持cdb.
	DstAccessType *string `name:"DstAccessType,omitempty"`
	// 目标实例信息, 其中目标实例地域不允许修改.
	DstInfo *DstInfo `name:"DstInfo,omitempty"`
	// 待修改的数据迁移任务ID
	JobId string `name:"JobId"`
	// 数据迁移任务名称
	JobName *string `name:"JobName,omitempty"`
	// 迁移任务配置选项
	MigrateOption *MigrateOption `name:"MigrateOption,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 源实例接入类型，值包括：extranet(外网),cvm(cvm自建实例),dcg(专线接入的实例),vpncloud(云vpn接入的实例),vpnselfbuild(自建vpn接入的实例)，cdb(云上cdb实例)
	SrcAccessType *string `name:"SrcAccessType,omitempty"`
	// 源实例信息，具体内容跟迁移任务类型相关
	SrcInfo *SrcInfo `name:"SrcInfo,omitempty"`
}

func (req *ModifyMigrateJobRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyMigrateJobResponse, error) {
	resp := &ModifyMigrateJobResponse{}
	err := client.Request("dts", "ModifyMigrateJob", "2018-03-30").Do(req, resp)
	return resp, err
}

type ModifyMigrateJobResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
