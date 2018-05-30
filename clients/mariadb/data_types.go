package mariadb

import (
	time "time"
)

// 慢查询条目信息
type SlowLogData struct {
	// 语句校验和，用于查询详情
	CheckSum string `json:"CheckSum"`
	// 数据库名称
	Db string `json:"Db"`
	// 抽象的SQL语句
	FingerPrint string `json:"FingerPrint"`
	// 平均的锁时间
	LockTimeAvg string `json:"LockTimeAvg"`
	// 最大锁时间
	LockTimeMax string `json:"LockTimeMax"`
	// 最小锁时间
	LockTimeMin string `json:"LockTimeMin"`
	// 锁时间总和
	LockTimeSum string `json:"LockTimeSum"`
	// 查询次数
	QueryCount string `json:"QueryCount"`
	// 平均查询时间
	QueryTimeAvg string `json:"QueryTimeAvg"`
	// 最大查询时间
	QueryTimeMax string `json:"QueryTimeMax"`
	// 最小查询时间
	QueryTimeMin string `json:"QueryTimeMin"`
	// 查询时间总和
	QueryTimeSum string `json:"QueryTimeSum"`
	// 扫描行数
	RowsExaminedSum string `json:"RowsExaminedSum"`
	// 发送行数
	RowsSentSum string `json:"RowsSentSum"`
	// 首次执行时间
	TsMax time.Time `json:"TsMax"`
	// 最后执行时间
	TsMin time.Time `json:"TsMin"`
	// 帐号
	User string `json:"User"`
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

// 描述云数据库实例的详细信息。
type DBInstance struct {
	// 实例所属应用 Id
	AppId int64 `json:"AppId"`
	// 自动续费标志：0 否，1 是
	AutoRenewFlag int64 `json:"AutoRenewFlag"`
	// 实例创建时间，格式为 2006-01-02 15:04:05
	CreateTime time.Time `json:"CreateTime"`
	// 实例 Id，唯一标识一个 TDSQL 实例
	InstanceId string `json:"InstanceId"`
	// 实例名称，用户可修改
	InstanceName string `json:"InstanceName"`
	// 实例内存大小，单位 GB
	Memory int64 `json:"Memory"`
	// 实例到期时间，格式为 2006-01-02 15:04:05
	PeriodEndTime time.Time `json:"PeriodEndTime"`
	// 产品类型 Id
	Pid int64 `json:"Pid"`
	// 实例所属项目 Id
	ProjectId int64 `json:"ProjectId"`
	// 实例所在地域名称，如 ap-shanghai
	Region string `json:"Region"`
	// 实例状态：0 创建中，1 流程处理中， 2 运行中，3 实例未初始化，-1 实例已隔离，-2 实例已删除
	Status int64 `json:"Status"`
	// 实例存储大小，单位 GB
	Storage int64 `json:"Storage"`
	// 子网 Id，基础网络时为 0
	SubnetId int64 `json:"SubnetId"`
	// TDSQL 版本信息，如 10.1.9
	TdsqlVersion string `json:"TdsqlVersion"`
	// 实例所属账号
	Uin string `json:"Uin"`
	// 实例最后更新时间，格式为 2006-01-02 15:04:05
	UpdateTime time.Time `json:"UpdateTime"`
	// 内网 IP 地址
	Vip string `json:"Vip"`
	// 私有网络 Id，基础网络时为 0
	VpcId int64 `json:"VpcId"`
	// 内网端口
	Vport int64 `json:"Vport"`
	// 外网访问的域名，公网可解析
	WanDomain string `json:"WanDomain"`
	// 外网端口
	WanPort int64 `json:"WanPort"`
	// 外网 IP 地址，公网可访问
	WanVip string `json:"WanVip"`
	// 实例所在可用区名称，如 ap-shanghai-1
	Zone string `json:"Zone"`
}

// 修改参数结果
type ParamModifyResult struct {
	// 参数修改结果。0表示修改成功；-1表示修改失败；-2表示该参数值非法
	Code int64 `json:"Code"`
	// 修改参数名字
	Param string `json:"Param"`
}

// DB资源使用情况监控指标集合
type ResourceUsageMonitorSet struct {
	// binlog日志磁盘可用空间,单位GB
	BinlogDiskAvailable MonitorData `json:"BinlogDiskAvailable"`
	// CPU利用率
	CpuUsageRate MonitorData `json:"CpuUsageRate"`
	// 磁盘可用空间,单位GB
	DataDiskAvailable MonitorIntData `json:"DataDiskAvailable"`
	// 内存可用空间,单位GB
	MemAvailable MonitorData `json:"MemAvailable"`
}

// 云数据库账号信息
type DBAccount struct {
	// 创建时间
	CreateTime time.Time `json:"CreateTime"`
	// 用户备注信息
	Description string `json:"Description"`
	// 用户可以从哪台主机登录（对应 MySQL 用户的 host 字段，UserName + Host 唯一标识一个用户，IP形式，IP段以%结尾；支持填入%；为空默认等于%）
	Host string `json:"Host"`
	// 只读标记，0：否， 1：该账号的sql请求优先选择备机执行，备机不可用时选择主机执行，2：优先选择备机执行，备机不可用时操作失败。
	ReadOnly int64 `json:"ReadOnly"`
	// 最后更新时间
	UpdateTime time.Time `json:"UpdateTime"`
	// 用户名
	UserName string `json:"UserName"`
}

// 按机型归类的实例可售卖规格信息
type InstanceSpec struct {
	// 设备型号
	Machine string `json:"Machine"`
	// 该机型对应的可售卖规格列表
	SpecInfos []*SpecConfigInfo `json:"SpecInfos"`
}

// 拉取的日志信息
type LogFileInfo struct {
	// 文件长度
	Length int64 `json:"Length"`
	// Log最后修改时间
	Mtime int64 `json:"Mtime"`
	// 下载Log时用到的统一资源标识符
	Uri int64 `json:"Uri"`
}

// DB参数描述
type ParamDesc struct {
	// 参数限制
	Constraint ParamConstraint `json:"Constraint"`
	// 系统默认值
	Default string `json:"Default"`
	// 参数名字
	Param string `json:"Param"`
	// 设置过的值，参数生效后，该值和value一样。未设置过就不返回该字段。
	SetValue string `json:"SetValue"`
	// 当前参数值
	Value string `json:"Value"`
}

// DB性能监控指标集合
type PerformanceMonitorSet struct {
	// 活跃连接数
	ConnActive MonitorData `json:"ConnActive"`
	// 删除操作数DELETE
	DeleteTotal MonitorData `json:"DeleteTotal"`
	// 磁盘每秒IO次数
	DiskIops MonitorData `json:"DiskIops"`
	// 插入操作数INSERT
	InsertTotal MonitorData `json:"InsertTotal"`
	// 是否发生主备切换，1为发生，0否
	IsMasterSwitched MonitorData `json:"IsMasterSwitched"`
	// 慢查询数
	LongQuery MonitorData `json:"LongQuery"`
	// 缓存命中率
	MemHitRate MonitorData `json:"MemHitRate"`
	// 查询操作数SELECT
	SelectTotal MonitorData `json:"SelectTotal"`
	// 主备延迟
	SlaveDelay MonitorData `json:"SlaveDelay"`
	// 更新操作数UPDATE
	UpdateTotal MonitorData `json:"UpdateTotal"`
}

// 分片节点可用区选择
type ZoneChooseInfo struct {
	// 主可用区
	MasterZone ZonesInfo `json:"MasterZone"`
	// 可选的从可用区
	SlaveZones []*ZonesInfo `json:"SlaveZones"`
}

// 约束类型值的范围
type ConstraintRange struct {
	// 约束类型为section时的最大值
	Max string `json:"Max"`
	// 约束类型为section时的最小值
	Min string `json:"Min"`
}

// 云数据库参数信息。
type DBParamValue struct {
	// 参数名称
	Param string `json:"Param"`
	// 参数值
	Value string `json:"Value"`
}

// 订单信息
type Deal struct {
	// 订单号
	DealName string `json:"DealName"`
	// 关联的流程 Id，可用于查询流程执行状态
	FlowId int64 `json:"FlowId"`
	// 只有创建实例的订单会填充该字段，表示该订单创建的实例的 ID。
	InstanceIds []*string `json:"InstanceIds"`
	// 所属账号
	OwnerUin string `json:"OwnerUin"`
	// 商品数量
	Quantity int64 `json:"Quantity"`
}

// 参数约束
type ParamConstraint struct {
	// 约束类型为enum时的可选值列表
	Enum *string `json:"Enum,omitempty"`
	// 约束类型为section时的范围
	Range *ConstraintRange `json:"Range,omitempty"`
	// 约束类型,如枚举enum，区间section
	Type string `json:"Type"`
}

// 售卖可用区信息
type RegionInfo struct {
	// 可选择的主可用区和从可用区
	AvailableChoice []*ZoneChooseInfo `json:"AvailableChoice"`
	// 地域英文ID
	Region string `json:"Region"`
	// 地域数字ID
	RegionId int64 `json:"RegionId"`
	// 地域中文名
	RegionName string `json:"RegionName"`
	// 可用区列表
	ZoneList []*ZonesInfo `json:"ZoneList"`
}

// 实例可售卖规格详细信息，创建实例和扩容实例时 Pid+MemSize 唯一确定一种售卖规格，磁盘大小可用区间为[MinDataDisk,MaxDataDisk]
type SpecConfigInfo struct {
	// 设备型号
	Machine string `json:"Machine"`
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

// 云数据库实例备份时间配置信息
type DBBackupTimeConfig struct {
	// 每天备份执行的区间的结束时间，格式 mm:ss，形如 23:00
	EndBackupTime string `json:"EndBackupTime"`
	// 实例 Id
	InstanceId string `json:"InstanceId"`
	// 每天备份执行的区间的开始时间，格式 mm:ss，形如 22:00
	StartBackupTime string `json:"StartBackupTime"`
}

// 监控数据
type MonitorData struct {
	// 监控数据
	Data []*float64 `json:"Data"`
	// 结束时间，形如 2018-03-24 23:59:59
	EndTime time.Time `json:"EndTime"`
	// 起始时间，形如 2018-03-24 23:59:59
	StartTime time.Time `json:"StartTime"`
}

// 整形监控数据
type MonitorIntData struct {
	// 监控数据
	Data int64 `json:"Data"`
	// 结束时间
	EndTime time.Time `json:"EndTime"`
	// 起始时间
	StartTime time.Time `json:"StartTime"`
}
