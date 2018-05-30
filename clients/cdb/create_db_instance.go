package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建云数据库实例（包年包月）
// https://cloud.tencent.com/document/api/236/15871
type CreateDbInstanceRequest struct {
	// 自动续费标记，可选值为：0-不自动续费；1-自动续费
	AutoRenewFlag *int64 `name:"AutoRenewFlag,omitempty"`
	// 备库2的可用区ID，默认为0，购买主实例时可指定该参数，购买只读实例或者灾备实例时指定该参数无意义
	BackupZone *string `name:"BackupZone,omitempty"`
	// 多可用区域，默认为0，支持值包括：0-表示单可用区，1-表示多可用区
	DeployMode *int64 `name:"DeployMode,omitempty"`
	// MySQL版本，值包括：5.5、5.6和5.7，请使用获取云数据库可售卖规格接口获取可创建的实例版本
	EngineVersion *string `name:"EngineVersion,omitempty"`
	// 实例数量，默认值为1, 最小值1，最大值为100
	GoodsNum int64 `name:"GoodsNum"`
	// 实例名称
	InstanceName *string `name:"InstanceName,omitempty"`
	// 实例类型，默认为 master，支持值包括：master-表示主实例，dr-表示灾备实例，ro-表示只读实例
	InstanceRole *string `name:"InstanceRole,omitempty"`
	// 实例ID，购买只读实例时必填，该字段表示只读实例的主实例ID，请使用查询实例列表接口查询云数据库实例ID
	MasterInstanceId *string `name:"MasterInstanceId,omitempty"`
	// 主实例地域信息，购买灾备实例时，该字段必填
	MasterRegion *string `name:"MasterRegion,omitempty"`
	// 实例内存大小，单位：MB，请使用获取云数据库可售卖规格接口获取可创建的内存规格
	Memory int64 `name:"Memory"`
	// 参数列表，参数格式如ParamList.0.Name=auto_increment&ParamList.0.Value=1。可通过查询参数列表查询支持设置的参数
	ParamList []*ParamInfo `name:"ParamList,omitempty"`
	// 设置root帐号密码，密码规则：8-64个字符，至少包含字母、数字、字符（支持的字符：_+-&=!@#$%^*()）中的两种，购买主实例时可指定该参数，购买只读实例或者灾备实例时指定该参数无意义
	Password *string `name:"Password,omitempty"`
	// 实例时长，单位：月，可选值包括[1,2,3,4,5,6,7,8,9,10,11,12,24,36]
	Period int64 `name:"Period"`
	// 自定义端口，端口支持范围：[ 1024-65535 ]
	Port *int64 `name:"Port,omitempty"`
	// 项目ID，不填为默认项目。请使用查询项目列表接口获取项目ID
	ProjectId *int64 `name:"ProjectId,omitempty"`
	// 数据复制方式，默认为0，支持值包括：0-表示异步复制，1-表示半同步复制，2-表示强同步复制
	ProtectMode *int64 `name:"ProtectMode,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 只读实例参数
	RoGroup *RoGroup `name:"RoGroup,omitempty"`
	// 安全组参数
	SecurityGroup []*string `name:"SecurityGroup,omitempty"`
	// 备库1的可用区信息，默认为zone的值
	SlaveZone *string `name:"SlaveZone,omitempty"`
	// 私有网络下的子网ID，如果设置了 UniqVpcId，则 UniqSubnetId 必填，请使用查询子网列表
	UniqSubnetId *string `name:"UniqSubnetId,omitempty"`
	// 私有网络ID，如果不传则默认选择基础网络，请使用查询私有网络列表
	UniqVpcId *string `name:"UniqVpcId,omitempty"`
	// 实例硬盘大小，单位：GB，请使用获取云数据库可售卖规格接口获取可创建的硬盘范围
	Volume int64 `name:"Volume"`
	// 可用区信息，该参数缺省时，系统会自动选择一个可用区，请使用获取云数据库可售卖规格接口获取可创建的可用区
	Zone *string `name:"Zone,omitempty"`
}

func (req *CreateDbInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateDbInstanceResponse, error) {
	resp := &CreateDbInstanceResponse{}
	err := client.Request("cdb", "CreateDBInstance", "2017-03-20").Do(req, resp)
	return resp, err
}

type CreateDbInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 短订单ID，用于调用云API相关接口，如获取订单信息
	DealIds string `json:"DealIds"`
	// 实例ID列表
	InstanceIds []*string `json:"InstanceIds"`
}
