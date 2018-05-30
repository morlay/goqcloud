package postgres

import (
	time "time"
)

// 描述地域的编码和状态等信息
type RegionInfo struct {
	// 该地域对应的英文名称
	Region string `json:"Region"`
	// 该地域对应的数字编号
	RegionId int64 `json:"RegionId"`
	// 该地域对应的中文名称
	RegionName string `json:"RegionName"`
	// 可用状态，UNAVAILABLE表示不可用，AVAILABLE表示可用
	RegionState string `json:"RegionState"`
}

// 描述某个地域下某个可用区的可售卖规格详细信息。
type SpecInfo struct {
	// 地域英文编码，对应RegionSet的Region字段
	Region string `json:"Region"`
	// 规格详细信息列表
	SpecItemInfoList []*SpecItemInfo `json:"SpecItemInfoList"`
	// 区域英文编码，对应ZoneSet的Zone字段
	Zone string `json:"Zone"`
}

// 描述一种规格的信息信息
type SpecItemInfo struct {
	// CPU核数
	Cpu []*int64 `json:"Cpu"`
	// 该规格所支持最大存储容量，单位：GB
	MaxStorage int64 `json:"MaxStorage"`
	// 内存大小，单位：MB
	Memory []*int64 `json:"Memory"`
	// 该规格所支持最小存储容量，单位：GB
	MinStorage int64 `json:"MinStorage"`
	// 该规格对应的计费ID
	Pid int64 `json:"Pid"`
	// 该规格的预估QPS
	Qps int64 `json:"Qps"`
	// 规格ID
	SpecCode string `json:"SpecCode"`
	// PostgreSQL的内核版本编号
	Version string `json:"Version"`
	// 内核编号对应的完整版本名称
	VersionName string `json:"VersionName"`
}

// 描述可用区的编码和状态信息
type ZoneInfo struct {
	// 该可用区的英文名称
	Zone string `json:"Zone"`
	// 该可用区对应的数字编号
	ZoneId int64 `json:"ZoneId"`
	// 该可用区的中文名称
	ZoneName string `json:"ZoneName"`
	// 可用状态，UNAVAILABLE表示不可用，AVAILABLE表示可用
	ZoneState string `json:"ZoneState"`
}

// 描述实例的详细信息
type DBInstance struct {
	// 是否自动续费，1：自动续费，0：不自动续费
	AutoRenew int64 `json:"AutoRenew"`
	// 实例创建时间
	CreateTime time.Time `json:"CreateTime"`
	// 实例DB字符集
	DBCharset string `json:"DBCharset"`
	// 售卖规格ID
	DBInstanceClass string `json:"DBInstanceClass"`
	// 实例分配的CPU数量，单位：个
	DBInstanceCpu int64 `json:"DBInstanceCpu"`
	// 实例ID
	DBInstanceId string `json:"DBInstanceId"`
	// 实例分配的内存大小，单位：GB
	DBInstanceMemory int64 `json:"DBInstanceMemory"`
	// 实例名称
	DBInstanceName string `json:"DBInstanceName"`
	// 实例网络连接信息
	DBInstanceNetInfo []*DBInstanceNetInfo `json:"DBInstanceNetInfo"`
	// 实例状态
	DBInstanceStatus string `json:"DBInstanceStatus"`
	// 实例分配的存储空间大小，单位：GB
	DBInstanceStorage int64 `json:"DBInstanceStorage"`
	// 实例类型，类型有：1、primary（主实例）；2、readonly（只读实例）；3、guard（灾备实例）；4、temp（临时实例）
	DBInstanceType string `json:"DBInstanceType"`
	// 实例版本，目前只支持standard（双机高可用版, 一主一从）
	DBInstanceVersion string `json:"DBInstanceVersion"`
	// PostgreSQL内核版本
	DBVersion string `json:"DBVersion"`
	// 实例到期时间
	ExpireTime time.Time `json:"ExpireTime"`
	// 实例隔离时间
	IsolatedTime time.Time `json:"IsolatedTime"`
	// 计费模式，1、prepaid（包年包月,预付费）；2、postpaid（按量计费，后付费）
	PayType string `json:"PayType"`
	// 项目ID
	ProjectId int64 `json:"ProjectId"`
	// 实例所属地域，如: ap-guangzhou，对应RegionSet的Region字段
	Region string `json:"Region"`
	// SubnetId
	SubnetId string `json:"SubnetId"`
	// 实例执行最后一次更新的时间
	UpdateTime time.Time `json:"UpdateTime"`
	// 私有网络ID
	VpcId string `json:"VpcId"`
	// 实例所属可用区， 如：ap-guangzhou-3，对应ZoneSet的Zone字段
	Zone string `json:"Zone"`
}

// 描述实例的网络连接信息
type DBInstanceNetInfo struct {
	// DNS域名
	Address string `json:"Address"`
	// Ip
	Ip string `json:"Ip"`
	// 网络类型，1、inner（内网地址）；2、public（外网地址）
	NetType string `json:"NetType"`
	// 连接Port地址
	Port int64 `json:"Port"`
	// 网络连接状态
	Status string `json:"Status"`
}

// 描述键值对过滤器，用于条件过滤查询。例如过滤ID、名称等
type Filter *Filter
