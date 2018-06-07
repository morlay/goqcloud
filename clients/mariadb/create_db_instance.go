package mariadb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建实例（包年包月）
// https://cloud.tencent.com/document/api/237/16180

type CreateDbInstanceRequest struct {
	// 是否自动使用代金券进行支付，默认不使用。
	AutoVoucher *bool `name:"AutoVoucher,omitempty"`
	// 欲购买的数量，默认查询购买1个实例的价格。
	Count *int64 `name:"Count,omitempty"`
	// 数据库引擎版本，当前可选：10.0.10，10.1.9，5.7.17
	DbVersionId *string `name:"DbVersionId,omitempty"`
	// 内存大小，单位：GB，可以通过 DescribeDBInstanceSpecs 查询实例规格获得。
	Memory int64 `name:"Memory"`
	// 节点个数大小，可以通过 DescribeDBInstanceSpecs 查询实例规格获得。
	NodeCount int64 `name:"NodeCount"`
	// 欲购买的时长，单位：月。
	Period *int64 `name:"Period,omitempty"`
	// 项目 ID，可以通过查看项目列表获取，不传则关联到默认项目
	ProjectId *int64 `name:"ProjectId,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 存储空间大小，单位：GB，可以通过 DescribeDBInstanceSpecs 查询实例规格获得不同内存大小对应的磁盘规格下限和上限。
	Storage int64 `name:"Storage"`
	// 虚拟私有网络子网 ID，VpcId 不为0时必填
	SubnetId *string `name:"SubnetId,omitempty"`
	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `name:"VoucherIds,omitempty"`
	// 虚拟私有网络 ID，不传或传 0 表示创建为基础网络
	VpcId *string `name:"VpcId,omitempty"`
	// 实例节点可用区分布，最多可填两个可用区。当分片规格为一主两从时，其中两个节点在第一个可用区。
	Zones []*string `name:"Zones"`
}

func (req *CreateDbInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateDbInstanceResponse, error) {
	resp := &CreateDbInstanceResponse{}
	err := client.Request("mariadb", "CreateDBInstance", "2017-03-12").Do(req, resp)
	return resp, err
}

type CreateDbInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 长订单号。可以据此调用 DescribeOrders 查询订单详细信息，或在支付失败时调用用户账号相关接口进行支付。
	DealName string `json:"DealName"`
	// 订单对应的实例 ID 列表，如果此处没有返回实例 ID，可以通过订单查询接口获取。还可通过实例查询接口查询实例是否创建完成。
	InstanceIds []*string `json:"InstanceIds"`
}
