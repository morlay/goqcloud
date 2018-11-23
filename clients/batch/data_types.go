package batch

// 描述了实例登录相关配置与信息。

type LoginSettings *LoginSettings

// 描述了操作系统所在块设备即系统盘的信息

type SystemDisk *SystemDisk

// agent运行模式

type AgentRunningMode *AgentRunningMode

// Docker容器信息

type Docker *Docker

// 描述实例机型配额信息。

type InstanceTypeQuotaItem *InstanceTypeQuotaItem

// 描述了数据盘的信息

type DataDisk *DataDisk

// 输入映射

type InputMapping *InputMapping

// 批量计算可用的InstanceTypeConfig信息

type InstanceTypeConfig *InstanceTypeConfig

// 描述了实例的公网可访问性，声明了实例的公网使用计费模式，最大带宽等

type InternetAccessible *InternetAccessible

// 描述了实例的抽象位置，包括其所在的可用区，所属的项目，宿主机等（仅CDH产品可用）

type Placement *Placement

// 授权认证信息

type Authentication *Authentication

// 计算环境创建信息。

type ComputeEnvCreateInfo *ComputeEnvCreateInfo

// 计算节点

type ComputeNode *ComputeNode

// 计算环境数据

type EnvData *EnvData

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

// 描述了 “云监控” 服务相关的信息

type RunMonitorServiceEnabled *RunMonitorServiceEnabled

// 任务

type Task *Task

// 任务实例日志详情。

type TaskInstanceLog *TaskInstanceLog

// 任务实例视图信息

type TaskInstanceView *TaskInstanceView

// 扩展数据

type Externals *Externals

// 作业

type Job *Job

// 输出映射

type OutputMapping *OutputMapping

// 自定义键值对

type EventVar *EventVar

// 竞价请求相关选项

type InstanceMarketOptionsRequest *InstanceMarketOptionsRequest

// 描述了单项的价格信息

type ItemPrice *ItemPrice

// 重定向信息

type RedirectInfo *RedirectInfo

// 任务实例统计指标

type TaskInstanceMetrics *TaskInstanceMetrics

// 计算环境属性数据

type ComputeEnvData *ComputeEnvData

// 计算节点统计指标

type ComputeNodeMetrics *ComputeNodeMetrics

// 描述了实例的增强服务启用情况与其设置，如云安全，云监控等实例 Agent

type EnhancedService *EnhancedService

// 任务统计指标

type TaskMetrics *TaskMetrics

// 本地重定向信息

type RedirectLocalInfo *RedirectLocalInfo

// 任务模板信息

type TaskTemplateView *TaskTemplateView

// 计算环境的创建或销毁活动

type Activity *Activity

// 计算环境

type AnonymousComputeEnv *AnonymousComputeEnv

// 计算环境信息

type ComputeEnvView *ComputeEnvView

// 本地磁盘规格

type LocalDiskType *LocalDiskType

// 描述了 “云安全” 服务相关的信息

type RunSecurityServiceEnabled *RunSecurityServiceEnabled

// 竞价相关选项

type SpotMarketOptions *SpotMarketOptions

// HDD的本地存储信息

type StorageBlock *StorageBlock

// 任务视图信息

type TaskView *TaskView

// 应用程序信息

type Application *Application

// 环境变量

type EnvVar *EnvVar

// 事件配置

type EventConfig *EventConfig

// 描述了VPC相关信息，包括子网，IP信息等

type VirtualPrivateCloud *VirtualPrivateCloud

// 计算环境

type NamedComputeEnv *NamedComputeEnv

// 通知信息

type Notification *Notification

// 输出映射配置

type OutputMappingConfig *OutputMappingConfig

// 依赖关系

type Dependence *Dependence

// 作业信息

type JobView *JobView

// 数据盘挂载选项

type MountDataDisk *MountDataDisk
