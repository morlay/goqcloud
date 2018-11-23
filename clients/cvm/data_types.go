package cvm

// 支持的操作系统类型，根据windows和linux分类。

type ImageOsList *ImageOsList

// 描述实例的状态。状态类型详见实例状态表

type InstanceStatus *InstanceStatus

// 描述了数据盘的信息

type DataDisk *DataDisk

//
// 描述键值对过滤器，用于条件过滤查询。例如过滤ID、名称、状态等
//
// 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
// 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。
//
// 以DescribeInstances接口的Filter为例。若我们需要查询可用区（zone）为广州一区 并且 实例计费模式（instance-charge-type）为包年包月 或者 按量计费的实例时，可如下实现：
// Filters.0.Name=zone
// &Filters.0.Values.0=ap-guangzhou-1
// &Filters.1.Name=instance-charge-type
// &Filters.1.Values.0=PREPAID
// &Filters.1.Values.1=POSTPAID_BY_HOUR
//

type Filter *Filter

// 地域信息

type RegionInfo *RegionInfo

// 创建云主机实例时同时绑定的标签对说明

type TagSpecification *TagSpecification

// 可用区信息

type ZoneInfo *ZoneInfo

// 描述了 “云监控” 服务相关的信息

type RunMonitorServiceEnabled *RunMonitorServiceEnabled

// 竞价相关选项

type SpotMarketOptions *SpotMarketOptions

// 描述了操作系统所在块设备即系统盘的信息

type SystemDisk *SystemDisk

// 定时任务

type ActionTimer *ActionTimer

// 描述实例机型配额信息。

type InstanceTypeQuotaItem *InstanceTypeQuotaItem

// 描述了实例的公网可访问性，声明了实例的公网使用计费模式，最大带宽等

type InternetAccessible *InternetAccessible

// 描述了网络计费

type InternetChargeTypeConfig *InternetChargeTypeConfig

// 描述了单项的价格信息

type ItemPrice *ItemPrice

// cdh实例详细信息

type HostItem *HostItem

// 描述了实例的计费模式

type InstanceChargePrepaid *InstanceChargePrepaid

// 竞价请求相关选项

type InstanceMarketOptionsRequest *InstanceMarketOptionsRequest

// 描述实例机型配置信息

type InstanceTypeConfig *InstanceTypeConfig

// 描述了实例的增强服务启用情况与其设置，如云安全，云监控等实例 Agent

type EnhancedService *EnhancedService

// 扩展数据

type Externals *Externals

// 描述实例的信息

type Instance *Instance

// 描述了按带宽计费的相关信息

type InternetBandwidthConfig *InternetBandwidthConfig

// 操作系统支持的类型。

type OsVersion *OsVersion

// 标签键值对

type Tag *Tag

// 描述预付费模式，即包年包月相关参数。包括购买时长和自动续费逻辑等。

type ChargePrepaid *ChargePrepaid

// 容灾组信息

type DisasterRecoverGroup *DisasterRecoverGroup

// 描述实例的机型族配置信息
// 形如：{'InstanceFamilyName': '标准型S1', 'InstanceFamily': 'S1'}、{'InstanceFamilyName': '网络优化型N1', 'InstanceFamily': 'N1'}、{'InstanceFamilyName': '高IO型I1', 'InstanceFamily': 'I1'}等。

type InstanceFamilyConfig *InstanceFamilyConfig

// 描述了实例的抽象位置，包括其所在的可用区，所属的项目，宿主机等（仅CDH产品可用）

type Placement *Placement

// 价格

type Price *Price

// 描述了 “云安全” 服务相关的信息

type RunSecurityServiceEnabled *RunSecurityServiceEnabled

// 镜像分享信息结构

type SharePermission *SharePermission

// cdh实例的资源信息

type HostResource *HostResource

// 一个关于镜像详细信息的结构体，主要包括镜像的主要状态与属性。

type Image *Image

// 描述密钥对信息

type KeyPair *KeyPair

// 本地磁盘规格

type LocalDiskType *LocalDiskType

// 描述了实例登录相关配置与信息。

type LoginSettings *LoginSettings

// HDD的本地存储信息

type StorageBlock *StorageBlock

// 描述了VPC相关信息，包括子网，IP信息等

type VirtualPrivateCloud *VirtualPrivateCloud
