package dcdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建DCDB分布式实例
// https://cloud.tencent.com/document/api/557/16135
type CreateDcdbInstanceRequest struct {
	// 是否自动使用代金券进行支付，默认不使用。
	AutoVoucher *bool `name:"AutoVoucher,omitempty"`
	// 欲购买实例的数量，目前只支持购买1个实例
	Count *int64 `name:"Count,omitempty"`
	// 数据库引擎版本，当前可选：10.0.10，10.1.9，5.7.17
	DbVersionId *string `name:"DbVersionId,omitempty"`
	// 欲购买的时长，单位：月。
	Period int64 `name:"Period"`
	// 项目 ID，可以通过查看项目列表获取，不传则关联到默认项目
	ProjectId *int64 `name:"ProjectId,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 实例分片个数，可选范围2-8，可以通过升级实例进行新增分片到最多64个分片。
	ShardCount int64 `name:"ShardCount"`
	// 分片内存大小，单位：GB，可以通过 DescribeShardSpec 查询实例规格获得。
	ShardMemory int64 `name:"ShardMemory"`
	// 单个分片节点个数，可以通过 DescribeShardSpec 查询实例规格获得。
	ShardNodeCount int64 `name:"ShardNodeCount"`
	// 分片存储空间大小，单位：GB，可以通过 DescribeShardSpec 查询实例规格获得。
	ShardStorage int64 `name:"ShardStorage"`
	// 虚拟私有网络子网 ID，VpcId不为空时必填
	SubnetId *string `name:"SubnetId,omitempty"`
	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `name:"VoucherIds,omitempty"`
	// 虚拟私有网络 ID，不传或传空表示创建为基础网络
	VpcId *string `name:"VpcId,omitempty"`
	// 分片节点可用区分布，最多可填两个可用区。当分片规格为一主两从时，其中两个节点在第一个可用区。
	Zones []*string `name:"Zones"`
}

func (req *CreateDcdbInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateDcdbInstanceResponse, error) {
	resp := &CreateDcdbInstanceResponse{}
	err := client.Request("dcdb", "CreateDCDBInstance", "2018-04-11").Do(req, resp)
	return resp, err
}

type CreateDcdbInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 长订单号。可以据此调用 DescribeOrders 查询订单详细信息，或在支付失败时调用用户账号相关接口进行支付。
	DealName string `json:"DealName"`
	// 订单对应的实例 ID 列表，如果此处没有返回实例 ID，可以通过订单查询接口获取。还可通过实例查询接口查询实例是否创建完成。
	InstanceIds []*string `json:"InstanceIds"`
}
