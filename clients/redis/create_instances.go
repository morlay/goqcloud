package redis

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建Redis实例
// https://cloud.tencent.com/document/api/239/20026

type CreateInstancesRequest struct {
	// 自动续费表示。0 - 默认状态（手动续费）；1 - 自动续费；2 - 明确不自动续费
	AutoRenew *int64 `name:"AutoRenew,omitempty"`
	// 付费方式:0-按量计费，1-包年包月。
	BillingMode int64 `name:"BillingMode"`
	// 实例数量，单次购买实例数量以 查询售卖规格接口返回的规格为准
	GoodsNum int64 `name:"GoodsNum"`
	// 实例容量，单位MB， 取值大小以 查询售卖规格接口返回的规格为准
	MemSize int64 `name:"MemSize"`
	// 实例密码，密码规则：1.长度为8-16个字符；2:至少包含字母、数字和字符!@^*()中的两种
	Password string `name:"Password"`
	// 购买时长，单位：月，取值范围 [1,2,3,4,5,6,7,8,9,10,11,12,24,36]
	Period int64 `name:"Period"`
	// 项目id，取值以用户账户>用户账户相关接口查询>项目列表返回的projectId为准
	ProjectId *int64 `name:"ProjectId,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 安全组id数组
	SecurityGroupIdList []*string `name:"SecurityGroupIdList,omitempty"`
	// 基础网络下， subnetId无效； vpc子网下，取值以查询查询子网列表
	SubnetId *string `name:"SubnetId,omitempty"`
	// 实例类型：2 – 主从版，5-单机版
	TypeId int64 `name:"TypeId"`
	// 用户自定义的端口 不填则默认为6379
	VPort *int64 `name:"VPort,omitempty"`
	// 私有网络ID，如果不传则默认选择基础网络，请使用私有网络列表 查询
	VpcId *string `name:"VpcId,omitempty"`
	// 实例所属的可用区id
	ZoneId int64 `name:"ZoneId"`
}

func (req *CreateInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateInstancesResponse, error) {
	resp := &CreateInstancesResponse{}
	err := client.Request("redis", "CreateInstances", "2018-04-12").Do(req, resp)
	return resp, err
}

type CreateInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 交易的Id
	DealId string `json:"DealId"`
}
