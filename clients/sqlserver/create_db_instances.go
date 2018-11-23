package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建实例
// https://cloud.tencent.com/document/api/238/19973

type CreateDbInstancesRequest struct {
	// 是否自动使用代金券；1 - 是，0 - 否，默认不使用
	AutoVoucher *int64 `name:"AutoVoucher,omitempty"`
	// 数据库版本号，目前取值有2012SP3，表示SQL Server 2012；2008R2，表示SQL Server 2008 R2；2016SP1，表示SQL Server 2016 SP1。每个地域支持售卖的版本可能不一样，可以通过DescribeZones接口来拉取每个地域可售卖的版本信息。不填的话，默认为版本2008R2
	DBVersion *string `name:"DBVersion,omitempty"`
	// 本次购买几个实例，默认值为1。取值不超过10
	GoodsNum *int64 `name:"GoodsNum,omitempty"`
	// 付费模式，目前只支持预付费，其值为PREPAID。可不填，默认值为PREPAID
	InstanceChargeType *string `name:"InstanceChargeType,omitempty"`
	// 实例内存大小，单位GB
	Memory int64 `name:"Memory"`
	// 购买实例周期，默认取值为1，表示一个月。取值不超过48
	Period *int64 `name:"Period,omitempty"`
	// 项目ID
	ProjectId *int64 `name:"ProjectId,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 实例磁盘大小，单位GB
	Storage int64 `name:"Storage"`
	// VPC子网ID，形如subnet-bdoe83fa；SubnetId和VpcId需同时设置或者同时不设置
	SubnetId *string `name:"SubnetId,omitempty"`
	// 代金券ID数组，目前单个订单只能使用一张
	VoucherIds []*string `name:"VoucherIds,omitempty"`
	// VPC网络ID，形如vpc-dsp338hz；SubnetId和VpcId需同时设置或者同时不设置
	VpcId *string `name:"VpcId,omitempty"`
	// 实例可用区，类似ap-guangzhou-1（广州一区）；实例可售卖区域可以通过接口DescribeZones获取
	Zone string `name:"Zone"`
}

func (req *CreateDbInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateDbInstancesResponse, error) {
	resp := &CreateDbInstancesResponse{}
	err := client.Request("sqlserver", "CreateDBInstances", "2018-03-28").Do(req, resp)
	return resp, err
}

type CreateDbInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 订单名称
	DealName string `json:"DealName"`
}
