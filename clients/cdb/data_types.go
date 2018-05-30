package cdb

import (
	time "time"
)

// 售卖实例类型
type SellType struct {
	Configs       []*SellConfig `json:"Configs"`
	EngineVersion []*string     `json:"EngineVersion"`
	TypeName      string        `json:"TypeName"`
}

// 数据库表权限
type TablePrivilege struct {
	// 数据库名
	Database string `json:"Database"`
	// 权限信息
	Privileges []*string `json:"Privileges"`
	// 数据库表名
	Table string `json:"Table"`
}

// 数据库表名
type DatabaseName struct {
	DatabaseName string `json:"DatabaseName"`
}

// 导入任务记录
type ImportRecord struct {
	AsyncRequestId string    `json:"AsyncRequestId"`
	Code           int64     `json:"Code"`
	CostTime       int64     `json:"CostTime"`
	CreateTime     time.Time `json:"CreateTime"`
	DbName         string    `json:"DbName"`
	FileName       string    `json:"FileName"`
	FileSize       string    `json:"FileSize"`
	InstanceId     string    `json:"InstanceId"`
	JobId          int64     `json:"JobId"`
	Message        string    `json:"Message"`
	Process        int64     `json:"Process"`
	Status         int64     `json:"Status"`
	WorkId         string    `json:"WorkId"`
}

// 数据库账号信息
type Account struct {
	// 新账户的域名
	Host string `json:"Host"`
	// 新账户的名称
	User string `json:"User"`
}

// 数据库表列表
type DatabaseTableList struct {
	// 数据库名
	DatabaseName string `json:"DatabaseName"`
	// 数据表数组
	TableList []*string `json:"TableList"`
}

// 实例详细信息
type InstanceInfo struct {
	AutoRenew     int64        `json:"AutoRenew"`
	CdbError      int64        `json:"CdbError"`
	CreateTime    string       `json:"CreateTime"`
	DeadlineTime  time.Time    `json:"DeadlineTime"`
	DeployMode    int64        `json:"DeployMode"`
	DeviceType    string       `json:"DeviceType"`
	DrInfo        []*DrInfo    `json:"DrInfo"`
	EngineVersion string       `json:"EngineVersion"`
	InitFlag      int64        `json:"InitFlag"`
	InstanceId    string       `json:"InstanceId"`
	InstanceName  string       `json:"InstanceName"`
	InstanceType  int64        `json:"InstanceType"`
	MasterInfo    MasterInfo   `json:"MasterInfo"`
	Memory        int64        `json:"Memory"`
	PayType       int64        `json:"PayType"`
	ProjectId     int64        `json:"ProjectId"`
	ProtectMode   int64        `json:"ProtectMode"`
	Region        string       `json:"Region"`
	RoGroups      []*RoGroup   `json:"RoGroups"`
	RoVipInfo     []*RoVipInfo `json:"RoVipInfo"`
	SlaveInfo     SlaveInfo    `json:"SlaveInfo"`
	Status        int64        `json:"Status"`
	SubnetId      int64        `json:"SubnetId"`
	TaskStatus    int64        `json:"TaskStatus"`
	UniqSubnetId  string       `json:"UniqSubnetId"`
	UniqVpcId     string       `json:"UniqVpcId"`
	Vip           string       `json:"Vip"`
	Volume        int64        `json:"Volume"`
	VpcId         int64        `json:"VpcId"`
	Vport         int64        `json:"Vport"`
	WanDomain     string       `json:"WanDomain"`
	WanPort       int64        `json:"WanPort"`
	WanStatus     int64        `json:"WanStatus"`
	Zone          string       `json:"Zone"`
}

// 实例参数信息
type ParamInfo struct {
	// 参数名
	Name string `json:"Name"`
	// 参数值
	Value string `json:"Value"`
}

// 多可用区信息
type ZoneConf struct {
	// 实例为多可用区部署时，备库2所在的可用区
	BackupZone []*string `json:"BackupZone"`
	// 可用区部署方式，可能的值为：0-单可用区；1-多可用区
	DeployMode []*int64 `json:"DeployMode"`
	// 主实例所在的可用区
	MasterZone []*string `json:"MasterZone"`
	// 实例为多可用区部署时，备库1所在的可用区
	SlaveZone []*string `json:"SlaveZone"`
}

// 备份详细信息
type BackupInfo struct {
	BackupId    int64  `json:"BackupId"`
	Creator     string `json:"Creator"`
	Date        string `json:"Date"`
	FinishTime  string `json:"FinishTime"`
	InternetUrl string `json:"InternetUrl"`
	IntranetUrl string `json:"IntranetUrl"`
	Name        string `json:"Name"`
	Size        int64  `json:"Size"`
	Status      string `json:"Status"`
	Type        string `json:"Type"`
}

// 二进制日志信息
type BinlogInfo struct {
	Date        time.Time `json:"Date"`
	InternetUrl string    `json:"InternetUrl"`
	IntranetUrl string    `json:"IntranetUrl"`
	Name        string    `json:"Name"`
	Size        int64     `json:"Size"`
	Type        string    `json:"Type"`
}

// 云数据库切换记录
type DBSwitchInfo struct {
	SwitchTime time.Time `json:"SwitchTime"`
	SwitchType string    `json:"SwitchType"`
}

// 主实例信息
type MasterInfo struct {
	DeviceType    string `json:"DeviceType"`
	ExClusterId   string `json:"ExClusterId"`
	ExClusterName string `json:"ExClusterName"`
	InstanceId    string `json:"InstanceId"`
	InstanceName  string `json:"InstanceName"`
	InstanceType  int64  `json:"InstanceType"`
	Memory        int64  `json:"Memory"`
	Qps           int64  `json:"Qps"`
	Region        string `json:"Region"`
	RegionId      int64  `json:"RegionId"`
	ResourceId    string `json:"ResourceId"`
	Status        int64  `json:"Status"`
	SubnetId      int64  `json:"SubnetId"`
	TaskStatus    int64  `json:"TaskStatus"`
	Volume        int64  `json:"Volume"`
	VpcId         int64  `json:"VpcId"`
	Zone          string `json:"Zone"`
	ZoneId        int64  `json:"ZoneId"`
}

// 安全组出站规则
type Outbound struct {
	Action     string `json:"Action"`
	CidrIp     string `json:"CidrIp"`
	Dir        string `json:"Dir"`
	IpProtocol string `json:"IpProtocol"`
	PortRange  string `json:"PortRange"`
}

// 账号详细信息
type AccountInfo struct {
	CreateTime         time.Time `json:"CreateTime"`
	Host               string    `json:"Host"`
	ModifyPasswordTime time.Time `json:"ModifyPasswordTime"`
	ModifyTime         time.Time `json:"ModifyTime"`
	Notes              string    `json:"Notes"`
	User               string    `json:"User"`
}

// 实例预期重启时间
type InstanceRebootTime struct {
	InstanceId    string `json:"InstanceId"`
	TimeInSeconds int64  `json:"TimeInSeconds"`
}

// 数据库实例参数
type Parameter struct {
	// 参数值
	CurrentValue string `json:"CurrentValue"`
	// 参数名称
	Name string `json:"Name"`
}

// 慢查询日志详情
type SlowLogInfo struct {
	Date        time.Time `json:"Date"`
	InternetUrl string    `json:"InternetUrl"`
	IntranetUrl string    `json:"IntranetUrl"`
	Name        string    `json:"Name"`
	Size        int64     `json:"Size"`
	Type        string    `json:"Type"`
}

// ECDB第二个从库的配置信息，只有ECDB实例才有这个字段
type BackupConfig struct {
	ReplicationMode string `json:"ReplicationMode"`
	Vip             string `json:"Vip"`
	Vport           string `json:"Vport"`
	Zone            string `json:"Zone"`
}

// 列权限信息
type ColumnPrivilege struct {
	// 数据库列名
	Column string `json:"Column"`
	// 数据库名
	Database string `json:"Database"`
	// 权限信息
	Privileges []*string `json:"Privileges"`
	// 数据库表名
	Table string `json:"Table"`
}

// 只读vip信息
type RoVipInfo struct {
	RoSubnetId  int64  `json:"RoSubnetId"`
	RoVip       string `json:"RoVip"`
	RoVipStatus int64  `json:"RoVipStatus"`
	RoVpcId     int64  `json:"RoVpcId"`
	RoVport     int64  `json:"RoVport"`
}

// 安全组详情
type SecurityGroup struct {
	CreateTime          string      `json:"CreateTime"`
	Inbound             []*Inbound  `json:"Inbound"`
	Outbound            []*Outbound `json:"Outbound"`
	ProjectId           int64       `json:"ProjectId"`
	SecurityGroupId     string      `json:"SecurityGroupId"`
	SecurityGroupName   string      `json:"SecurityGroupName"`
	SecurityGroupRemark string      `json:"SecurityGroupRemark"`
}

// 表名
type TableName struct {
	TableName string `json:"TableName"`
}

// 可用区售卖配置
type ZoneSellConf struct {
	HourInstanceSaleMaxNum int64       `json:"HourInstanceSaleMaxNum"`
	IsBm                   bool        `json:"IsBm"`
	IsCustom               bool        `json:"IsCustom"`
	IsDefaultZone          bool        `json:"IsDefaultZone"`
	IsSupportDr            bool        `json:"IsSupportDr"`
	IsSupportVpc           bool        `json:"IsSupportVpc"`
	PayType                []*string   `json:"PayType"`
	ProtectMode            []*string   `json:"ProtectMode"`
	SellType               []*SellType `json:"SellType"`
	Status                 int64       `json:"Status"`
	Zone                   string      `json:"Zone"`
	ZoneConf               ZoneConf    `json:"ZoneConf"`
	ZoneName               string      `json:"ZoneName"`
}

// 第一备机信息
type First struct {
	Region string `json:"Region"`
	Vip    string `json:"Vip"`
	Vport  int64  `json:"Vport"`
	Zone   string `json:"Zone"`
}

// 地域售卖配置
type RegionSellConf struct {
	Area            string          `json:"Area"`
	IsDefaultRegion int64           `json:"IsDefaultRegion"`
	Region          string          `json:"Region"`
	RegionName      string          `json:"RegionName"`
	ZonesConf       []*ZoneSellConf `json:"ZonesConf"`
}

// 从库的配置信息
type SlaveConfig struct {
	ReplicationMode string `json:"ReplicationMode"`
	Zone            string `json:"Zone"`
}

// 数据库权限
type DatabasePrivilege struct {
	// 数据库名
	Database string `json:"Database"`
	// 权限信息
	Privileges []*string `json:"Privileges"`
}

// 灾备实例信息
type DrInfo struct {
	InstanceId   string `json:"InstanceId"`
	InstanceName string `json:"InstanceName"`
	InstanceType int64  `json:"InstanceType"`
	Region       string `json:"Region"`
	Status       int64  `json:"Status"`
	SyncStatus   int64  `json:"SyncStatus"`
	Zone         string `json:"Zone"`
}

// 安全组入站规则
type Inbound struct {
	Action     string `json:"Action"`
	CidrIp     string `json:"CidrIp"`
	Dir        string `json:"Dir"`
	IpProtocol string `json:"IpProtocol"`
	PortRange  string `json:"PortRange"`
}

// 只读组参数
type RoGroup struct {
	// 最少实例保留个数，若购买只读实例数量小于设置数量将不做剔除
	MinRoInGroup *int64 `json:"MinRoInGroup,omitempty"`
	// 只读组ID
	RoGroupId *string `json:"RoGroupId,omitempty"`
	// 只读组模式，可选值为：alone-系统自动分配只读组；allinone-新建只读组；join-使用现有只读组
	RoGroupMode *string `json:"RoGroupMode,omitempty"`
	// 只读组名称
	RoGroupName *string `json:"RoGroupName,omitempty"`
	// 延迟阀值
	RoMaxDelayTime *int64 `json:"RoMaxDelayTime,omitempty"`
	// 是否启用延迟超限剔除功能，启用该功能后，只读实例与主实例的延迟超过延迟阈值值，只读实例将被隔离。可选值：1-启用；0-不启用
	RoOfflineDelay *int64 `json:"RoOfflineDelay,omitempty"`
	// 权重值
	Weight *int64 `json:"Weight,omitempty"`
	// 读写权重分配模式，可选值：system-系统自动分配；custom-自定义
	WeightMode *string `json:"WeightMode,omitempty"`
}

// 售卖配置详情
type SellConfig struct {
	CdbType    string `json:"CdbType"`
	Connection int64  `json:"Connection"`
	Cpu        int64  `json:"Cpu"`
	Device     string `json:"Device"`
	Info       string `json:"Info"`
	Iops       int64  `json:"Iops"`
	Memory     int64  `json:"Memory"`
	Qps        int64  `json:"Qps"`
	Status     int64  `json:"Status"`
	Type       string `json:"Type"`
	VolumeMax  int64  `json:"VolumeMax"`
	VolumeMin  int64  `json:"VolumeMin"`
	VolumeStep int64  `json:"VolumeStep"`
}

// 备机信息
type SlaveInfo struct {
	First  First `json:"First"`
	Second First `json:"Second"`
}
