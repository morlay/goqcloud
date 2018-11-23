package as

// 启动配置的系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。

type SystemDisk *SystemDisk

// 负载均衡器目标属性

type TargetAttribute *TargetAttribute

// 描述了实例登录相关配置与信息，出于安全性考虑，不会描述敏感信息。

type LimitedLoginSettings *LimitedLoginSettings

// 应用型负载均衡器

type ForwardLoadBalancer *ForwardLoadBalancer

// 实例信息

type Instance *Instance

// 描述了 “云安全” 服务相关的信息

type RunSecurityServiceEnabled *RunSecurityServiceEnabled

// 描述定时任务的信息

type ScheduledAction *ScheduledAction

// 伸缩组简明信息。

type AutoScalingGroupAbstract *AutoScalingGroupAbstract

// 符合条件的启动配置信息的集合。

type LaunchConfiguration *LaunchConfiguration

// 描述了 “云监控” 服务相关的信息。

type RunMonitorServiceEnabled *RunMonitorServiceEnabled

// 描述了启动配置创建实例的公网可访问性，声明了实例的公网使用计费模式，最大带宽等

type InternetAccessible *InternetAccessible

// 启动配置的数据盘配置信息。若不指定该参数，则默认不购买数据盘，当前仅支持购买的时候指定一个数据盘。

type DataDisk *DataDisk

// 描述了实例的增强服务启用情况与其设置，如云安全，云监控等实例 Agent。

type EnhancedService *EnhancedService

//
// 描述键值对过滤器，用于条件过滤查询。例如过滤ID、名称、状态等
//
// 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
// 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。
//
// 以DescribeInstances接口的Filter为例。若我们需要查询可用区（zone）为广州一区 并且 实例计费模式（instance-charge-type）为包年包月 或者 按量计费的实例时，可如下实现：
// Filters.1.Name=zone
// &Filters.1.Values.1=ap-guangzhou-1
// &Filters.2.Name=instance-charge-type
// &Filters.2.Values.1=PREPAID
// &Filters.3.Values.2=POSTPAID_BY_HOUR
//

type Filter *Filter

// CVM竞价请求相关选项

type InstanceMarketOptionsRequest *InstanceMarketOptionsRequest

// 描述了实例登录相关配置与信息。

type LoginSettings *LoginSettings

// 竞价相关选项

type SpotMarketOptions *SpotMarketOptions

// 伸缩组

type AutoScalingGroup *AutoScalingGroup
