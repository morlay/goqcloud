package dcdb

import (
	time "time"
)

// 分布式数据库实例信息
type DCDBInstanceInfo struct {
	// APPID
	AppId int64 `json:"AppId"`
	// 自动续费标志
	AutoRenewFlag int64 `json:"AutoRenewFlag"`
	// 创建时间
	CreateTime time.Time `json:"CreateTime"`
	// 实例ID
	InstanceId string `json:"InstanceId"`
	// 实例名称
	InstanceName string `json:"InstanceName"`
	// 隔离时间
	IsolatedTimestamp time.Time `json:"IsolatedTimestamp"`
	// 内存大小，单位 GB
	Memory int64 `json:"Memory"`
	// 到期时间
	PeriodEndTime time.Time `json:"PeriodEndTime"`
	// 项目ID
	ProjectId int64 `json:"ProjectId"`
	// 地域
	Region string `json:"Region"`
	// 分片个数
	ShardCount int64 `json:"ShardCount"`
	// 分片详情
	ShardDetail []*ShardInfo `json:"ShardDetail"`
	// 状态
	Status int64 `json:"Status"`
	// 状态中文描述
	StatusDesc string `json:"StatusDesc"`
	// 存储大小，单位 GB
	Storage int64 `json:"Storage"`
	// Subnet数字ID
	SubnetId int64 `json:"SubnetId"`
	// UIN
	Uin string `json:"Uin"`
	// 内网IP
	Vip string `json:"Vip"`
	// VPC数字ID
	VpcId int64 `json:"VpcId"`
	// 内网端口
	Vport int64 `json:"Vport"`
	// 可用区
	Zone string `json:"Zone"`
}

// 分片信息
type ShardInfo struct {
	// 创建时间
	Createtime string `json:"Createtime"`
	// 内存大小，单位 GB
	Memory int64 `json:"Memory"`
	// 分片数字ID
	ShardId int64 `json:"ShardId"`
	// 分片ID
	ShardInstanceId string `json:"ShardInstanceId"`
	// 分片Set ID
	ShardSerialId string `json:"ShardSerialId"`
	// 状态
	Status int64 `json:"Status"`
	// 存储大小，单位 GB
	Storage int64 `json:"Storage"`
}

// 按机型分类的规格配置
type SpecConfig struct {
	// 规格机型
	Machine string `json:"Machine"`
	// 规格列表
	SpecConfigInfos []*SpecConfigInfo `json:"SpecConfigInfos"`
}

// 实例可售卖规格详细信息，创建实例和扩容实例时 NodeCount、Memory 确定售卖规格，硬盘大小可用区间为[MinStorage,MaxStorage]
type SpecConfigInfo struct {
	// 数据盘规格最大值，单位 GB
	MaxStorage int64 `json:"MaxStorage"`
	// 内存大小，单位 GB
	Memory int64 `json:"Memory"`
	// 数据盘规格最小值，单位 GB
	MinStorage int64 `json:"MinStorage"`
	// 节点个数，2 表示一主一从，3 表示一主二从
	NodeCount int64 `json:"NodeCount"`
	// 产品类型 Id
	Pid int64 `json:"Pid"`
	// 最大 Qps 值
	Qps int64 `json:"Qps"`
	// 推荐的使用场景
	SuitInfo string `json:"SuitInfo"`
}

// 升级实例 -- 切分分片类型
type SplitShardConfig struct {
	// 分片ID数组
	ShardInstanceIds []*string `json:"ShardInstanceIds"`
	// 分片内存大小，单位 GB
	ShardMemory int64 `json:"ShardMemory"`
	// 分片存储大小，单位 GB
	ShardStorage int64 `json:"ShardStorage"`
	// 数据切分比例
	SplitRate int64 `json:"SplitRate"`
}

// 可用区信息
type ZonesInfo struct {
	// 可用区英文ID
	Zone string `json:"Zone"`
	// 可用区数字ID
	ZoneId int64 `json:"ZoneId"`
	// 可用区中文名
	ZoneName string `json:"ZoneName"`
}

// 升级实例 -- 新增分片类型
type AddShardConfig struct {
	// 新增分片的数量
	ShardCount int64 `json:"ShardCount"`
	// 分片内存大小，单位 GB
	ShardMemory int64 `json:"ShardMemory"`
	// 分片存储大小，单位 GB
	ShardStorage int64 `json:"ShardStorage"`
}

// 订单信息
type Deal struct {
	// 商品数量
	Count int64 `json:"Count"`
	// 订单号
	DealName string `json:"DealName"`
	// 关联的流程 Id，可用于查询流程执行状态
	FlowId int64 `json:"FlowId"`
	// 只有创建实例的订单会填充该字段，表示该订单创建的实例的 ID。
	InstanceIds []*string `json:"InstanceIds"`
	// 所属账号
	OwnerUin string `json:"OwnerUin"`
}

// 升级实例 -- 扩容分片类型
type ExpandShardConfig struct {
	// 分片ID数组
	ShardInstanceIds []*string `json:"ShardInstanceIds"`
	// 分片内存大小，单位 GB
	ShardMemory int64 `json:"ShardMemory"`
	// 分片存储大小，单位 GB
	ShardStorage int64 `json:"ShardStorage"`
}

// 售卖可用区信息
type RegionInfo struct {
	// 可选择的主可用区和从可用区
	AvailableChoice []*ShardZoneChooseInfo `json:"AvailableChoice"`
	// 地域英文ID
	Region string `json:"Region"`
	// 地域数字ID
	RegionId int64 `json:"RegionId"`
	// 地域中文名
	RegionName string `json:"RegionName"`
	// 可用区列表
	ZoneList []*ZonesInfo `json:"ZoneList"`
}

// 分片节点可用区选择
type ShardZoneChooseInfo struct {
	// 主可用区
	MasterZone ZonesInfo `json:"MasterZone"`
	// 可选的从可用区
	SlaveZones []*ZonesInfo `json:"SlaveZones"`
}
