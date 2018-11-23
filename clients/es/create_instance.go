package es

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建ES集群实例
// https://cloud.tencent.com/document/api/845/30633

type CreateInstanceRequest struct {
	// 是否自动使用代金券，1是，0否，默认不使用
	AutoVoucher *int64 `name:"AutoVoucher,omitempty"`
	// 包年包月购买时长，单位由TimeUint决定，默认为月
	ChargePeriod *int64 `name:"ChargePeriod,omitempty"`
	// 计费类型: PREPAID：预付费，即包年包月 POSTPAID_BY_HOUR：按小时后付费，默认值
	ChargeType *string `name:"ChargeType,omitempty"`
	// 节点存储容量，单位GB
	DiskSize int64 `name:"DiskSize"`
	// 节点存储类型,取值范围:  LOCAL_BASIC: 本地硬盘  LOCAL_SSD: 本地SSD硬盘，默认值  CLOUD_BASIC: 普通云硬盘  CLOUD_PREMIUM: 高硬能云硬盘  CLOUD_SSD: SSD云硬盘
	DiskType *string `name:"DiskType,omitempty"`
	// 实例版本,当前只支持5.6.4
	EsVersion string `name:"EsVersion"`
	// 实例名称，1-50 个英文、汉字、数字、连接线-或下划线_
	InstanceName *string `name:"InstanceName,omitempty"`
	// 节点数量
	NodeNum int64 `name:"NodeNum"`
	// 节点规格： ES.S1.SMALL2: 1核2GES.S1.MEDIUM4: 2核4GES.S1.MEDIUM8: 2核8GES.S1.LARGE16: 4核16GES.S1.2XLARGE32: 8核32GES.S1.4XLARGE64: 16核64G
	NodeType string `name:"NodeType"`
	// 访问密码，密码需8到16位，至少包括两项（[a-z,A-Z],[0-9]和[()`~!@#$%^&*-+=_|{}:;' <>,.?/]的特殊符号
	Password string `name:"Password"`
	// 区域
	Region string `name:"Region"`
	// 自动续费标识，取值范围： RENEW_FLAG_AUTO：自动续费RENEW_FLAG_MANUAL：不自动续费，用户手动续费如不传递该参数，普通用于默认不自动续费，SVIP用户自动续费
	RenewFlag *string `name:"RenewFlag,omitempty"`
	// 子网ID
	SubnetId string `name:"SubnetId"`
	// 计费时长单位，当前只支持“m”，表示月
	TimeUnit *string `name:"TimeUnit,omitempty"`
	// 代金券ID列表，目前仅支持指定一张代金券
	VoucherIds []*string `name:"VoucherIds,omitempty"`
	// 私有网络ID
	VpcId string `name:"VpcId"`
	// 可用区
	Zone string `name:"Zone"`
}

func (req *CreateInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateInstanceResponse, error) {
	resp := &CreateInstanceResponse{}
	err := client.Request("es", "CreateInstance", "2018-04-16").Do(req, resp)
	return resp, err
}

type CreateInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 实例ID
	InstanceId string `json:"InstanceId"`
}
