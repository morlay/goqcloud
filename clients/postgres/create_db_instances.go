package postgres

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建DB实例
// https://cloud.tencent.com/document/api/409/16771

type CreateDbInstancesRequest struct {
	// 是否自动使用代金券。1（是），0（否），默认不使用。
	AutoVoucher *int64 `name:"AutoVoucher,omitempty"`
	// PostgreSQL内核版本，目前只支持：9.3.5、9.5.4两种版本。
	DBVersion string `name:"DBVersion"`
	// 实例计费类型。目前只支持：PREPAID（预付费，即包年包月）。
	InstanceChargeType *string `name:"InstanceChargeType,omitempty"`
	// 一次性购买的实例数量。取值1-100
	InstanceCount int64 `name:"InstanceCount"`
	// 购买时长，单位：月。目前只支持1,2,3,4,5,6,7,8,9,10,11,12,24,36这些值。
	Period int64 `name:"Period"`
	// 项目ID。
	ProjectId *int64 `name:"ProjectId,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 售卖规格ID。该参数可以通过调用DescribeProductConfig的返回值中的SpecCode字段来获取。
	SpecCode string `name:"SpecCode"`
	// 实例容量大小，单位：GB。
	Storage int64 `name:"Storage"`
	// 私有网络子网ID。
	SubnetId *string `name:"SubnetId,omitempty"`
	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `name:"VoucherIds,omitempty"`
	// 私有网络ID。
	VpcId *string `name:"VpcId,omitempty"`
	// 可用区ID。该参数可以通过调用 DescribeZones 接口的返回值中的Zone字段来获取。
	Zone string `name:"Zone"`
}

func (req *CreateDbInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateDbInstancesResponse, error) {
	resp := &CreateDbInstancesResponse{}
	err := client.Request("postgres", "CreateDBInstances", "2017-03-12").Do(req, resp)
	return resp, err
}

type CreateDbInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 订单号列表。每个实例对应一个订单号。
	DealNames []*string `json:"DealNames"`
}
