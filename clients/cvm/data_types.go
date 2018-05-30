package cvm

import (
	time "time"
)

// 定时任务
type ActionTimer struct {
	// 执行时间，格式形如：2018/5/29 11:26:40,执行时间必须大于当前时间5分钟。
	ActionTime *string `json:"ActionTime,omitempty"`
	// 扩展数据
	Externals Externals `json:"Externals"`
	// 定时器名称，目前仅支持销毁一个值：TerminateInstances。
	TimerAction *string `json:"TimerAction,omitempty"`
}

// 描述了实例的增强服务启用情况与其设置，如云安全，云监控等实例 Agent
type EnhancedService struct {
	// 开启云安全服务。若不指定该参数，则默认开启云监控服务。
	MonitorService *RunMonitorServiceEnabled `json:"MonitorService,omitempty"`
	// 开启云安全服务。若不指定该参数，则默认开启云安全服务。
	SecurityService *RunSecurityServiceEnabled `json:"SecurityService,omitempty"`
}

// 描述实例的信息
type Instance struct {
	CPU                 int64               `json:"CPU"`
	CreatedTime         time.Time           `json:"CreatedTime"`
	DataDisks           []*DataDisk         `json:"DataDisks"`
	ExpiredTime         time.Time           `json:"ExpiredTime"`
	ImageId             string              `json:"ImageId"`
	InstanceChargeType  string              `json:"InstanceChargeType"`
	InstanceId          string              `json:"InstanceId"`
	InstanceName        string              `json:"InstanceName"`
	InstanceType        string              `json:"InstanceType"`
	InternetAccessible  InternetAccessible  `json:"InternetAccessible"`
	LoginSettings       LoginSettings       `json:"LoginSettings"`
	Memory              int64               `json:"Memory"`
	OsName              string              `json:"OsName"`
	Placement           Placement           `json:"Placement"`
	PrivateIpAddresses  []*string           `json:"PrivateIpAddresses"`
	PublicIpAddresses   []*string           `json:"PublicIpAddresses"`
	RenewFlag           string              `json:"RenewFlag"`
	RestrictState       string              `json:"RestrictState"`
	SecurityGroupIds    []*string           `json:"SecurityGroupIds"`
	SystemDisk          SystemDisk          `json:"SystemDisk"`
	VirtualPrivateCloud VirtualPrivateCloud `json:"VirtualPrivateCloud"`
}

// 描述了 “云监控” 服务相关的信息
type RunMonitorServiceEnabled struct {
	// 是否开启云监控服务。取值范围：TRUE：表示开启云监控服务FALSE：表示不开启云监控服务默认取值：TRUE。
	Enabled *bool `json:"Enabled,omitempty"`
}

// 描述预付费模式，即包年包月相关参数。包括购买时长和自动续费逻辑等。
type ChargePrepaid struct {
	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。
	Period int64 `json:"Period"`
	// 自动续费标识。取值范围：NOTIFY_AND_AUTO_RENEW：通知过期且自动续费NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费默认取值：NOTIFY_AND_AUTO_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitempty"`
}

// 描述键值对过滤器，用于条件过滤查询。例如过滤ID、名称、状态等
// 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
// 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。
// 以DescribeInstances接口的Filter为例。若我们需要查询可用区（zone）为广州一区 并且 实例计费模式（instance-charge-type）为包年包月 或者 按量计费的实例时，可如下实现：
// Filters.0.Name=zone
// &Filters.0.Values.0=ap-guangzhou-1
// &Filters.1.Name=instance-charge-type
// &Filters.1.Values.0=PREPAID
// &Filters.1.Values.1=POSTPAID_BY_HOUR
type Filter struct {
	// 需要过滤的字段。
	Name string `json:"Name"`
	// 字段的过滤值。
	Values []*string `json:"Values"`
}

// 描述实例的状态。状态类型详见实例状态表
type InstanceStatus struct {
	InstanceId    string `json:"InstanceId"`
	InstanceState string `json:"InstanceState"`
}

// 描述实例机型配额信息。
type InstanceTypeQuotaItem struct {
	Cpu                int64            `json:"Cpu"`
	Externals          Externals        `json:"Externals"`
	InstanceChargeType string           `json:"InstanceChargeType"`
	InstanceFamily     string           `json:"InstanceFamily"`
	InstanceType       string           `json:"InstanceType"`
	LocalDiskTypeList  []*LocalDiskType `json:"LocalDiskTypeList"`
	Memory             int64            `json:"Memory"`
	NetworkCard        int64            `json:"NetworkCard"`
	Price              ItemPrice        `json:"Price"`
	Status             string           `json:"Status"`
	TypeName           string           `json:"TypeName"`
	Zone               string           `json:"Zone"`
}

// 描述了单项的价格信息
type ItemPrice struct {
	ChargeUnit    string  `json:"ChargeUnit"`
	DiscountPrice float64 `json:"DiscountPrice"`
	OriginalPrice float64 `json:"OriginalPrice"`
	UnitPrice     float64 `json:"UnitPrice"`
}

// 创建云主机实例时同时绑定的标签对说明
type TagSpecification struct {
	// 标签绑定的资源类型
	ResourceType string `json:"ResourceType"`
	// 标签对列表
	Tags []*Tag `json:"Tags"`
}

// 描述了数据盘的信息
type DataDisk struct {
	// 系统盘ID。LOCAL_BASIC 和 LOCAL_SSD 类型没有ID。暂时不支持该参数。
	DiskId *string `json:"DiskId,omitempty"`
	// 数据盘大小，单位：GB。最小调整步长为10G，不同数据盘类型取值范围不同，具体限制详见：CVM实例配置。默认值为0，表示不购买数据盘。更多限制详见产品文档。
	DiskSize int64 `json:"DiskSize"`
	// 数据盘类型。数据盘类型限制详见CVM实例配置。取值范围：LOCAL_BASIC：本地硬盘LOCAL_SSD：本地SSD硬盘CLOUD_BASIC：普通云硬盘CLOUD_PREMIUM：高性能云硬盘CLOUD_SSD：SSD云硬盘默认取值：LOCAL_BASIC。该参数对ResizeInstanceDisk接口无效。
	DiskType *string `json:"DiskType,omitempty"`
}

// 描述了按带宽计费的相关信息
type InternetBandwidthConfig struct {
	EndTime            time.Time          `json:"EndTime"`
	InternetAccessible InternetAccessible `json:"InternetAccessible"`
	StartTime          time.Time          `json:"StartTime"`
}

// 可用区信息
type ZoneInfo struct {
	Zone      string `json:"Zone"`
	ZoneId    string `json:"ZoneId"`
	ZoneName  string `json:"ZoneName"`
	ZoneState string `json:"ZoneState"`
}

// cdh实例的资源信息
type HostResource struct {
	CpuAvailable  int64   `json:"CpuAvailable"`
	CpuTotal      int64   `json:"CpuTotal"`
	DiskAvailable int64   `json:"DiskAvailable"`
	DiskTotal     int64   `json:"DiskTotal"`
	MemAvailable  float64 `json:"MemAvailable"`
	MemTotal      float64 `json:"MemTotal"`
}

// 描述实例机型配置信息
type InstanceTypeConfig struct {
	CPU               int64  `json:"CPU"`
	CbsSupport        string `json:"CbsSupport"`
	GPU               int64  `json:"GPU"`
	InstanceFamily    string `json:"InstanceFamily"`
	InstanceType      string `json:"InstanceType"`
	InstanceTypeState string `json:"InstanceTypeState"`
	Memory            int64  `json:"Memory"`
	Zone              string `json:"Zone"`
}

// 描述了实例的公网可访问性，声明了实例的公网使用计费模式，最大带宽等
type InternetAccessible struct {
	// 网络计费类型。取值范围：BANDWIDTH_PREPAID：预付费按带宽结算TRAFFIC_POSTPAID_BY_HOUR：流量按小时后付费BANDWIDTH_POSTPAID_BY_HOUR：带宽按小时后付费BANDWIDTH_PACKAGE：带宽包用户默认取值：TRAFFIC_POSTPAID_BY_HOUR。
	InternetChargeType *string `json:"InternetChargeType,omitempty"`
	// 公网出带宽上限，单位：Mbps。默认值：0Mbps。不同机型带宽上限范围不一致，具体限制详见购买网络带宽。
	InternetMaxBandwidthOut *int64 `json:"InternetMaxBandwidthOut,omitempty"`
	// 是否分配公网IP。取值范围：TRUE：表示分配公网IPFALSE：表示不分配公网IP当公网带宽大于0Mbps时，可自由选择开通与否，默认开通公网IP；当公网带宽为0，则不允许分配公网IP。
	PublicIpAssigned *bool `json:"PublicIpAssigned,omitempty"`
}

// 描述了实例登录相关配置与信息。
type LoginSettings struct {
	// 保持镜像的原始设置。该参数与Password或KeyIds.N不能同时指定。只有使用自定义镜像、共享镜像或外部导入镜像创建实例时才能指定该参数为TRUE。取值范围：TRUE：表示保持镜像的登录设置FALSE：表示不保持镜像的登录设置默认取值：FALSE。
	KeepImageLogin *string `json:"KeepImageLogin,omitempty"`
	// 密钥ID列表。关联密钥后，就可以通过对应的私钥来访问实例；KeyId可通过接口DescribeKeyPairs获取，密钥与密码不能同时指定，同时Windows操作系统不支持指定密钥。当前仅支持购买的时候指定一个密钥。
	KeyIds []*string `json:"KeyIds,omitempty"`
	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：Linux实例密码必须8到16位，至少包括两项[a-z，A-Z]、[0-9] 和 [( ) ~ ! @ # $ % ^ & * - + = &#124; { } [ ] : ; ' , . ? / ]中的特殊符号。<br><li>Windows实例密码必须12到16位，至少包括三项[a-z]，[A-Z]，[0-9] 和 [( ) ~ ! @ # $ % ^ & * - + = { } [ ] : ; ' , . ? /]中的特殊符号。若不指定该参数，则由系统随机生成密码，并通过站内信方式通知到用户。
	Password *string `json:"Password,omitempty"`
}

// 地域信息
type RegionInfo struct {
	Region      string `json:"Region"`
	RegionName  string `json:"RegionName"`
	RegionState string `json:"RegionState"`
}

// 镜像分享信息结构
type SharePermission struct {
	Account    string    `json:"Account"`
	CreateTime time.Time `json:"CreateTime"`
}

// 描述了操作系统所在块设备即系统盘的信息
type SystemDisk struct {
	// 系统盘ID。LOCAL_BASIC 和 LOCAL_SSD 类型没有ID。暂时不支持该参数。
	DiskId *string `json:"DiskId,omitempty"`
	// 系统盘大小，单位：GB。默认值为 50
	DiskSize *int64 `json:"DiskSize,omitempty"`
	// 系统盘类型。系统盘类型限制详见CVM实例配置。取值范围：LOCAL_BASIC：本地硬盘LOCAL_SSD：本地SSD硬盘CLOUD_BASIC：普通云硬盘CLOUD_SSD：SSD云硬盘默认取值：LOCAL_BASIC。
	DiskType *string `json:"DiskType,omitempty"`
}

// cdh实例详细信息
type HostItem struct {
	CreatedTime    time.Time    `json:"CreatedTime"`
	ExpiredTime    time.Time    `json:"ExpiredTime"`
	HostChargeType string       `json:"HostChargeType"`
	HostId         string       `json:"HostId"`
	HostIp         string       `json:"HostIp"`
	HostName       string       `json:"HostName"`
	HostResource   HostResource `json:"HostResource"`
	HostState      string       `json:"HostState"`
	HostType       string       `json:"HostType"`
	InstanceIds    string       `json:"InstanceIds"`
	Placement      Placement    `json:"Placement"`
	RenewFlag      string       `json:"RenewFlag"`
}

// 描述实例的机型族配置信息
// 形如：{'InstanceFamilyName': '标准型S1', 'InstanceFamily': 'S1'}、{'InstanceFamilyName': '网络优化型N1', 'InstanceFamily': 'N1'}、{'InstanceFamilyName': '高IO型I1', 'InstanceFamily': 'I1'}等。
type InstanceFamilyConfig struct {
	InstanceFamily     string `json:"InstanceFamily"`
	InstanceFamilyName string `json:"InstanceFamilyName"`
}

// 竞价请求相关选项
type InstanceMarketOptionsRequest struct {
	// 市场选项类型，当前只支持取值：spot
	MarketType *string `json:"MarketType,omitempty"`
	// 竞价相关选项
	SpotOptions *SpotMarketOptions `json:"SpotOptions,omitempty"`
}

// 本地磁盘规格
type LocalDiskType struct {
	MaxSize       int64  `json:"MaxSize"`
	MinSize       int64  `json:"MinSize"`
	PartitionType string `json:"PartitionType"`
	Type          string `json:"Type"`
}

// 价格
type Price struct {
	BandwidthPrice ItemPrice `json:"BandwidthPrice"`
	InstancePrice  ItemPrice `json:"InstancePrice"`
}

// 描述了 “云安全” 服务相关的信息
type RunSecurityServiceEnabled struct {
	// 是否开启云安全服务。取值范围：TRUE：表示开启云安全服务FALSE：表示不开启云安全服务默认取值：TRUE。
	Enabled *bool `json:"Enabled,omitempty"`
}

// 描述了网络计费
type InternetChargeTypeConfig struct {
	// 网络计费模式描述信息。
	Description *string `json:"Description,omitempty"`
	// 网络计费模式。
	InternetChargeType *string `json:"InternetChargeType,omitempty"`
}

// 标签键值对
type Tag struct {
	// 标签键
	Key string `json:"Key"`
	// 标签值
	Value string `json:"Value"`
}

// 扩展数据
type Externals struct {
	// 释放地址
	ReleaseAddress *bool `json:"ReleaseAddress,omitempty"`
	// HDD本地存储属性
	StorageBlockAttr *StorageBlock `json:"StorageBlockAttr,omitempty"`
	// 不支持的网络类型
	UnsupportNetworks []*string `json:"UnsupportNetworks,omitempty"`
}

// 描述了实例的计费模式
type InstanceChargePrepaid struct {
	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。
	Period int64 `json:"Period"`
	// 自动续费标识。取值范围：NOTIFY_AND_AUTO_RENEW：通知过期且自动续费NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费默认取值：NOTIFY_AND_AUTO_RENEW。若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `json:"RenewFlag,omitempty"`
}

// HDD的本地存储信息
type StorageBlock struct {
	MaxSize int64  `json:"MaxSize"`
	MinSize int64  `json:"MinSize"`
	Type    string `json:"Type"`
}

// 一个关于镜像详细信息的结构体，主要包括镜像的主要状态与属性。
type Image struct {
	Architecture     string    `json:"Architecture"`
	CreatedTime      time.Time `json:"CreatedTime"`
	ImageCreator     string    `json:"ImageCreator"`
	ImageDescription string    `json:"ImageDescription"`
	ImageId          string    `json:"ImageId"`
	ImageName        string    `json:"ImageName"`
	ImageSize        int64     `json:"ImageSize"`
	ImageSource      string    `json:"ImageSource"`
	ImageState       string    `json:"ImageState"`
	ImageType        string    `json:"ImageType"`
	OsName           string    `json:"OsName"`
	Platform         string    `json:"Platform"`
}

// 描述密钥对信息
type KeyPair struct {
	AssociatedInstanceIds []*string `json:"AssociatedInstanceIds"`
	CreatedTime           time.Time `json:"CreatedTime"`
	Description           string    `json:"Description"`
	KeyId                 string    `json:"KeyId"`
	KeyName               string    `json:"KeyName"`
	PrivateKey            string    `json:"PrivateKey"`
	ProjectId             string    `json:"ProjectId"`
	PublicKey             string    `json:"PublicKey"`
}

// 描述了实例的抽象位置，包括其所在的可用区，所属的项目，宿主机等（仅CDH产品可用）
type Placement struct {
	// 实例所属的专用宿主机ID列表。如果您有购买专用宿主机并且指定了该参数，则您购买的实例就会随机的部署在这些专用宿主机上。当前暂不支持。
	HostIds []*string `json:"HostIds,omitempty"`
	// 实例所属项目ID。该参数可以通过调用 DescribeProject 的返回值中的 projectId 字段来获取。不填为默认项目。
	ProjectId *int64 `json:"ProjectId,omitempty"`
	// 实例所属的可用区ID。该参数也可以通过调用  DescribeZones 的返回值中的Zone字段来获取。
	Zone string `json:"Zone"`
}

// 竞价相关选项
type SpotMarketOptions struct {
	// 竞价出价
	MaxPrice string `json:"MaxPrice"`
	// 竞价请求类型
	SpotInstanceType *string `json:"SpotInstanceType,omitempty"`
}

// 描述了VPC相关信息，包括子网，IP信息等
type VirtualPrivateCloud struct {
	// 是否用作公网网关。公网网关只有在实例拥有公网IP以及处于私有网络下时才能正常使用。取值范围：TRUE：表示用作公网网关FALSE：表示不用作公网网关默认取值：FALSE。
	AsVpcGateway *bool `json:"AsVpcGateway,omitempty"`
	// 私有网络子网 IP 数组，在创建实例、修改实例vpc属性操作中可使用此参数。当前仅批量创建多台实例时支持传入相同子网的多个 IP。
	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty"`
	// 私有网络子网ID，形如subnet-xxx。有效的私有网络子网ID可通过登录控制台查询；也可以调用接口  DescribeSubnetEx ，从接口返回中的unSubnetId字段获取。
	SubnetId string `json:"SubnetId"`
	// 私有网络ID，形如vpc-xxx。有效的VpcId可通过登录控制台查询；也可以调用接口 DescribeVpcEx ，从接口返回中的unVpcId字段获取。
	VpcId string `json:"VpcId"`
}
