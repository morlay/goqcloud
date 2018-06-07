package batch

// 授权认证信息

type Authentication struct {
	// 授权场景，例如COS
	Scene string `json:"Scene"`
	// SecretId
	SecretId string `json:"SecretId"`
	// SecretKey
	SecretKey string `json:"SecretKey"`
}

// 计算环境信息

type ComputeEnvView struct {
	// 计算节点统计指标
	ComputeNodeMetrics ComputeNodeMetrics `json:"ComputeNodeMetrics"`
	// 创建时间
	CreateTime string `json:"CreateTime"`
	// 计算节点期望个数
	DesiredComputeNodeCount int64 `json:"DesiredComputeNodeCount"`
	// 计算环境ID
	EnvId string `json:"EnvId"`
	// 计算环境名称
	EnvName string `json:"EnvName"`
	// 计算环境类型
	EnvType string `json:"EnvType"`
	// 位置信息
	Placement Placement `json:"Placement"`
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

// 作业信息

type JobView struct {
	// 创建时间
	CreateTime string `json:"CreateTime"`
	// 结束时间
	EndTime *string `json:"EndTime,omitempty"`
	// 作业ID
	JobId string `json:"JobId"`
	// 作业名称
	JobName string `json:"JobName"`
	// 作业状态
	JobState string `json:"JobState"`
	// 位置信息
	Placement *Placement `json:"Placement,omitempty"`
	// 作业优先级
	Priority int64 `json:"Priority"`
	// 任务统计指标
	TaskMetrics TaskMetrics `json:"TaskMetrics"`
}

// 重定向信息

type RedirectInfo struct {
	// 标准错误重定向文件名，支持三个占位符${BATCH_JOB_ID}、${BATCH_TASK_NAME}、${BATCH_TASK_INSTANCE_INDEX}
	StderrRedirectFileName *string `json:"StderrRedirectFileName,omitempty"`
	// 标准错误重定向路径
	StderrRedirectPath *string `json:"StderrRedirectPath,omitempty"`
	// 标准输出重定向文件名，支持三个占位符${BATCH_JOB_ID}、${BATCH_TASK_NAME}、${BATCH_TASK_INSTANCE_INDEX}
	StdoutRedirectFileName *string `json:"StdoutRedirectFileName,omitempty"`
	// 标准输出重定向路径
	StdoutRedirectPath *string `json:"StdoutRedirectPath,omitempty"`
}

// 计算环境

type AnonymousComputeEnv struct {
	// agent运行模式，适用于Windows系统
	AgentRunningMode *AgentRunningMode `json:"AgentRunningMode,omitempty"`
	// 计算环境具体参数
	EnvData EnvData `json:"EnvData"`
	// 计算环境管理类型
	EnvType string `json:"EnvType"`
	// 数据盘挂载选项
	MountDataDisks []*MountDataDisk `json:"MountDataDisks,omitempty"`
}

// 应用程序信息

type Application struct {
	// 任务执行命令
	Command string `json:"Command"`
	// 应用程序的交付方式，包括PACKAGE、LOCAL 两种取值，分别指远程存储的软件包、计算环境本地。
	DeliveryForm string `json:"DeliveryForm"`
	// 应用使用Docker的相关配置。在使用Docker配置的情况下，DeliveryForm 为 LOCAL 表示直接使用Docker镜像内部的应用软件，通过Docker方式运行；DeliveryForm 为 PACKAGE，表示将远程应用包注入到Docker镜像后，通过Docker方式运行。为避免Docker不同版本的兼容性问题，Docker安装包及相关依赖由Batch统一负责，对于已安装Docker的自定义镜像，请卸载后再使用Docker特性。
	Docker *Docker `json:"Docker,omitempty"`
	// 应用程序软件包的远程存储路径
	PackagePath *string `json:"PackagePath,omitempty"`
}

// 任务实例视图信息

type TaskInstanceView struct {
	// 任务实例运行时所在计算节点（例如CVM）的InstanceId。任务实例未运行或者完结时，本字段为空。任务实例重试时，本字段会随之变化
	ComputeNodeInstanceId *string `json:"ComputeNodeInstanceId,omitempty"`
	// 创建时间
	CreateTime string `json:"CreateTime"`
	// 结束时间
	EndTime *string `json:"EndTime,omitempty"`
	// 应用程序执行结束的exit code
	ExitCode *int64 `json:"ExitCode,omitempty"`
	// 启动时间
	LaunchTime *string `json:"LaunchTime,omitempty"`
	// 重定向信息
	RedirectInfo RedirectInfo `json:"RedirectInfo"`
	// 开始运行时间
	RunningTime *string `json:"RunningTime,omitempty"`
	// 任务实例状态原因详情，任务实例失败时，会记录失败原因
	StateDetailedReason string `json:"StateDetailedReason"`
	// 任务实例状态原因，任务实例失败时，会记录失败原因
	StateReason string `json:"StateReason"`
	// 任务实例索引
	TaskInstanceIndex int64 `json:"TaskInstanceIndex"`
	// 任务实例状态
	TaskInstanceState string `json:"TaskInstanceState"`
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

// 环境变量

type EnvVar struct {
	// 环境变量名称
	Name string `json:"Name"`
	// 环境变量取值
	Value string `json:"Value"`
}

// 计算环境

type NamedComputeEnv struct {
	// 非活跃节点处理策略，默认“RECREATE”，即对于实例创建失败或异常退还的计算节点，定期重新创建实例资源。
	ActionIfComputeNodeInactive *string `json:"ActionIfComputeNodeInactive,omitempty"`
	// agent运行模式，适用于Windows系统
	AgentRunningMode *AgentRunningMode `json:"AgentRunningMode,omitempty"`
	// 授权信息
	Authentications []*Authentication `json:"Authentications,omitempty"`
	// 计算节点期望个数
	DesiredComputeNodeCount int64 `json:"DesiredComputeNodeCount"`
	// 计算环境具体参数
	EnvData EnvData `json:"EnvData"`
	// 计算环境描述
	EnvDescription *string `json:"EnvDescription,omitempty"`
	// 计算环境名称
	EnvName string `json:"EnvName"`
	// 计算环境管理类型
	EnvType string `json:"EnvType"`
	// 输入映射信息
	InputMappings []*InputMapping `json:"InputMappings,omitempty"`
	// 数据盘挂载选项
	MountDataDisks []*MountDataDisk `json:"MountDataDisks,omitempty"`
	// 通知信息
	Notifications *Notification `json:"Notifications,omitempty"`
}

// 任务

type Task struct {
	// 应用程序信息
	Application Application `json:"Application"`
	// 授权信息
	Authentications []*EnvVar `json:"Authentications,omitempty"`
	// 运行环境信息，ComputeEnv 和 EnvId 必须指定一个（且只有一个）参数。
	ComputeEnv *AnonymousComputeEnv `json:"ComputeEnv,omitempty"`
	// 计算环境ID，ComputeEnv 和 EnvId 必须指定一个（且只有一个）参数。
	EnvId *string `json:"EnvId,omitempty"`
	// 自定义环境变量
	EnvVars []*Authentication `json:"EnvVars,omitempty"`
	// TaskInstance失败后处理方式，取值包括TERMINATE（默认）、INTERRUPT、FAST_INTERRUPT。
	FailedAction *string `json:"FailedAction,omitempty"`
	// 输入映射
	InputMappings []*InputMapping `json:"InputMappings,omitempty"`
	// 任务失败后的最大重试次数，默认为0
	MaxRetryCount int64 `json:"MaxRetryCount"`
	// 输出映射配置
	OutputMappingConfigs []*OutputMappingConfig `json:"OutputMappingConfigs,omitempty"`
	// 输出映射
	OutputMappings []*OutputMapping `json:"OutputMappings,omitempty"`
	// 重定向信息
	RedirectInfo RedirectInfo `json:"RedirectInfo"`
	// 重定向本地信息
	RedirectLocalInfo *RedirectLocalInfo `json:"RedirectLocalInfo,omitempty"`
	// 任务实例运行个数
	TaskInstanceNum int64 `json:"TaskInstanceNum"`
	// 任务名称，在一个作业内部唯一
	TaskName string `json:"TaskName"`
	// 任务启动后的超时时间，单位秒，默认为3600秒
	Timeout int64 `json:"Timeout"`
}

// 任务视图信息

type TaskView struct {
	// 开始时间
	CreateTime string `json:"CreateTime"`
	// 结束时间
	EndTime *string `json:"EndTime,omitempty"`
	// 任务名称
	TaskName string `json:"TaskName"`
	// 任务状态
	TaskState string `json:"TaskState"`
}

// agent运行模式

type AgentRunningMode struct {
	// 场景类型，支持WINDOWS
	Scene string `json:"Scene"`
	// 运行Agent的Session
	Session string `json:"Session"`
	// 运行Agent的User
	User string `json:"User"`
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

// 计算节点统计指标

type ComputeNodeMetrics struct {
	// 异常的计算节点数量
	AbnormalCount string `json:"AbnormalCount"`
	// 完成创建的计算节点数量
	CreatedCount string `json:"CreatedCount"`
	// 创建中的计算节点数量
	CreatingCount string `json:"CreatingCount"`
	// 创建失败的计算节点数量
	CreationFailedCount string `json:"CreationFailedCount"`
	// 销毁中的计算节点数量
	DeletingCount string `json:"DeletingCount"`
	// 运行中的计算节点数量
	RunningCount string `json:"RunningCount"`
	// 已经完成提交的计算节点数量
	SubmittedCount string `json:"SubmittedCount"`
}

// 描述了实例的增强服务启用情况与其设置，如云安全，云监控等实例 Agent

type EnhancedService struct {
	// 开启云安全服务。若不指定该参数，则默认开启云监控服务。
	MonitorService *RunMonitorServiceEnabled `json:"MonitorService,omitempty"`
	// 开启云安全服务。若不指定该参数，则默认开启云安全服务。
	SecurityService *RunSecurityServiceEnabled `json:"SecurityService,omitempty"`
}

// 作业

type Job struct {
	// 依赖信息
	Dependences []*Dependence `json:"Dependences,omitempty"`
	// 作业描述
	JobDescription *string `json:"JobDescription,omitempty"`
	// 作业名称
	JobName string `json:"JobName"`
	// 通知信息
	Notifications []*Notification `json:"Notifications,omitempty"`
	// 作业优先级，任务（Task）和任务实例（TaskInstance）会继承作业优先级
	Priority int64 `json:"Priority"`
	// 表示创建 CVM 失败按照何种策略处理。取值范围包括 FAILED，RUNNABLE。FAILED 表示创建 CVM 失败按照一次执行失败处理，RUNNABLE 表示创建 CVM 失败按照继续等待处理。默认值为FAILED。StateIfCreateCvmFailed对于提交的指定计算环境的作业无效。
	StateIfCreateCvmFailed *string `json:"StateIfCreateCvmFailed,omitempty"`
	// 对于存在依赖关系的任务中，后序任务执行对于前序任务的依赖条件。取值范围包括 PRE_TASK_SUCCEED，PRE_TASK_AT_LEAST_PARTLY_SUCCEED，PRE_TASK_FINISHED，默认值为PRE_TASK_SUCCEED。
	TaskExecutionDependOn *string `json:"TaskExecutionDependOn,omitempty"`
	// 任务信息
	Tasks []*Task `json:"Tasks"`
}

// 通知信息

type Notification struct {
	// 事件配置
	EventConfigs []*EventConfig `json:"EventConfigs"`
	// CMQ主题名字，要求主题名有效且关联订阅
	TopicName string `json:"TopicName"`
}

// 输出映射配置

type OutputMappingConfig struct {
	// 存储类型，仅支持COS
	Scene string `json:"Scene"`
	// 并行worker数量
	WorkerNum int64 `json:"WorkerNum"`
	// worker分块大小
	WorkerPartSize int64 `json:"WorkerPartSize"`
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

// 计算环境的创建或销毁活动

type Activity struct {
	// 活动ID
	ActivityId string `json:"ActivityId"`
	// 活动状态
	ActivityState string `json:"ActivityState"`
	// 起因
	Cause string `json:"Cause"`
	// 计算节点活动类型，创建或者销毁
	ComputeNodeActivityType string `json:"ComputeNodeActivityType"`
	// 计算节点ID
	ComputeNodeId string `json:"ComputeNodeId"`
	// 活动结束时间
	EndTime *string `json:"EndTime,omitempty"`
	// 计算环境ID
	EnvId string `json:"EnvId"`
	// 活动开始时间
	StartTime string `json:"StartTime"`
	// 状态原因
	StateReason string `json:"StateReason"`
}

// 计算节点

type ComputeNode struct {
	// Batch Agent 版本
	AgentVersion *string `json:"AgentVersion,omitempty"`
	// 计算节点ID
	ComputeNodeId string `json:"ComputeNodeId"`
	// 计算节点实例ID，对于CVM场景，即为CVM的InstanceId
	ComputeNodeInstanceId *string `json:"ComputeNodeInstanceId,omitempty"`
	// 计算节点状态
	ComputeNodeState string `json:"ComputeNodeState"`
	// CPU核数
	Cpu int64 `json:"Cpu"`
	// 内存容量，单位GiB
	Mem int64 `json:"Mem"`
	// 实例内网IP
	PrivateIpAddresses []*string `json:"PrivateIpAddresses"`
	// 实例公网IP
	PublicIpAddresses []*string `json:"PublicIpAddresses"`
	// 资源创建完成时间
	ResourceCreatedTime *string `json:"ResourceCreatedTime,omitempty"`
	// 计算节点运行  TaskInstance 可用容量。0表示计算节点忙碌。
	TaskInstanceNumAvailable int64 `json:"TaskInstanceNumAvailable"`
}

// 本地重定向信息

type RedirectLocalInfo struct {
	// 标准错误重定向本地文件名，支持三个占位符${BATCH_JOB_ID}、${BATCH_TASK_NAME}、${BATCH_TASK_INSTANCE_INDEX}
	StderrLocalFileName *string `json:"StderrLocalFileName,omitempty"`
	// 标准错误重定向本地路径
	StderrLocalPath *string `json:"StderrLocalPath,omitempty"`
	// 标准输出重定向本地文件名，支持三个占位符${BATCH_JOB_ID}、${BATCH_TASK_NAME}、${BATCH_TASK_INSTANCE_INDEX}
	StdoutLocalFileName *string `json:"StdoutLocalFileName,omitempty"`
	// 标准输出重定向本地路径
	StdoutLocalPath *string `json:"StdoutLocalPath,omitempty"`
}

// 描述了 “云安全” 服务相关的信息

type RunSecurityServiceEnabled struct {
	// 是否开启云安全服务。取值范围：TRUE：表示开启云安全服务FALSE：表示不开启云安全服务默认取值：TRUE。
	Enabled *bool `json:"Enabled,omitempty"`
}

// 计算环境数据

type EnvData struct {
	// 实例数据盘配置信息
	DataDisks []*DataDisk `json:"DataDisks,omitempty"`
	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `json:"EnhancedService,omitempty"`
	// CVM镜像ID
	ImageId string `json:"ImageId"`
	// CVM实例显示名称
	InstanceName *string `json:"InstanceName,omitempty"`
	// CVM实例类型
	InstanceType string `json:"InstanceType"`
	// 公网带宽相关信息设置
	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty"`
	// 实例登录设置
	LoginSettings *LoginSettings `json:"LoginSettings,omitempty"`
	// 实例所属安全组
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty"`
	// 实例系统盘配置信息
	SystemDisk *SystemDisk `json:"SystemDisk,omitempty"`
	// 私有网络相关信息配置
	VirtualPrivateCloud *VirtualPrivateCloud `json:"VirtualPrivateCloud,omitempty"`
}

// 事件配置

type EventConfig struct {
	// 事件类型
	EventName string `json:"EventName"`
	// 自定义键值对
	EventVars []*EventVar `json:"EventVars"`
}

// 描述了 “云监控” 服务相关的信息

type RunMonitorServiceEnabled struct {
	// 是否开启云监控服务。取值范围：TRUE：表示开启云监控服务FALSE：表示不开启云监控服务默认取值：TRUE。
	Enabled *bool `json:"Enabled,omitempty"`
}

// 任务统计指标

type TaskMetrics struct {
	// Failed个数
	FailedCount int64 `json:"FailedCount"`
	// FailedInterrupted个数
	FailedInterruptedCount int64 `json:"FailedInterruptedCount"`
	// Pending个数
	PendingCount int64 `json:"PendingCount"`
	// Runnable个数
	RunnableCount int64 `json:"RunnableCount"`
	// Running个数
	RunningCount int64 `json:"RunningCount"`
	// Starting个数
	StartingCount int64 `json:"StartingCount"`
	// Submitted个数
	SubmittedCount int64 `json:"SubmittedCount"`
	// Succeed个数
	SucceedCount int64 `json:"SucceedCount"`
}

// 依赖关系

type Dependence struct {
	// 依赖关系的终点任务名称
	EndTask string `json:"EndTask"`
	// 依赖关系的起点任务名称
	StartTask string `json:"StartTask"`
}

// Docker容器信息

type Docker struct {
	// Docker Hub填写“[user/repo]:[tag]”，Tencent Registry填写“ccr.ccs.tencentyun.com/[namespace/repo]:[tag]”
	Image string `json:"Image"`
	// Docker Hub 密码或 Tencent Registry 密码
	Password string `json:"Password"`
	// Docker Hub 可以不填，但确保具有公网访问能力。或者是 Tencent Registry 服务地址“ccr.ccs.tencentyun.com”
	Server *string `json:"Server,omitempty"`
	// Docker Hub 用户名或 Tencent Registry 用户名
	User string `json:"User"`
}

// 任务模板信息

type TaskTemplateView struct {
	// 创建时间
	CreateTime string `json:"CreateTime"`
	// 任务模板描述
	TaskTemplateDescription string `json:"TaskTemplateDescription"`
	// 任务模板ID
	TaskTemplateId string `json:"TaskTemplateId"`
	// 任务模板信息
	TaskTemplateInfo Task `json:"TaskTemplateInfo"`
	// 任务模板名称
	TaskTemplateName string `json:"TaskTemplateName"`
}

// 数据盘挂载选项

type MountDataDisk struct {
	// 文件系统类型，Linux系统下支持"EXT3"和"EXT4"两种，默认"EXT3"；Windows系统下仅支持"NTFS"
	FileSystemType *string `json:"FileSystemType,omitempty"`
	// 挂载点，Linux系统合法路径，或Windows系统盘符,比如"H:\"
	LocalPath string `json:"LocalPath"`
}

// 输出映射

type OutputMapping struct {
	// 目的端路径
	DestinationPath string `json:"DestinationPath"`
	// 源端路径
	SourcePath string `json:"SourcePath"`
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

// 任务实例统计指标

type TaskInstanceMetrics struct {
	// Failed个数
	FailedCount int64 `json:"FailedCount"`
	// FailedInterrupted个数
	FailedInterruptedCount int64 `json:"FailedInterruptedCount"`
	// Pending个数
	PendingCount int64 `json:"PendingCount"`
	// Runnable个数
	RunnableCount int64 `json:"RunnableCount"`
	// Running个数
	RunningCount int64 `json:"RunningCount"`
	// Starting个数
	StartingCount int64 `json:"StartingCount"`
	// Submitted个数
	SubmittedCount int64 `json:"SubmittedCount"`
	// Succeed个数
	SucceedCount int64 `json:"SucceedCount"`
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

// 自定义键值对

type EventVar struct {
	// 自定义键
	Name string `json:"Name"`
	// 自定义值
	Value string `json:"Value"`
}

// 输入映射

type InputMapping struct {
	// 目的端路径
	DestinationPath string `json:"DestinationPath"`
	// 挂载配置项参数
	MountOptionParameter *string `json:"MountOptionParameter,omitempty"`
	// 源端路径
	SourcePath string `json:"SourcePath"`
}

// 描述实例机型配置信息

type InstanceTypeConfig struct {
	// CPU核数，单位：核。
	CPU *int64 `json:"CPU,omitempty"`
	// 是否支持云硬盘。取值范围：TRUE：表示支持云硬盘；FALSE：表示不支持云硬盘。
	CbsSupport *string `json:"CbsSupport,omitempty"`
	// GPU核数，单位：核。
	GPU *int64 `json:"GPU,omitempty"`
	// 实例机型系列。
	InstanceFamily *string `json:"InstanceFamily,omitempty"`
	// 实例机型。
	InstanceType *string `json:"InstanceType,omitempty"`
	// 机型状态。取值范围：AVAILABLE：表示机型可用；UNAVAILABLE：表示机型不可用。
	InstanceTypeState *string `json:"InstanceTypeState,omitempty"`
	// 内存容量，单位：GB。
	Memory *int64 `json:"Memory,omitempty"`
	// 可用区。
	Zone *string `json:"Zone,omitempty"`
}

// 计算环境创建信息。

type ComputeEnvCreateInfo struct {
	// 授权信息
	Authentications []*Authentication `json:"Authentications"`
	// 计算节点期望个数
	DesiredComputeNodeCount int64 `json:"DesiredComputeNodeCount"`
	// 计算环境参数
	EnvData EnvData `json:"EnvData"`
	// 计算环境描述
	EnvDescription string `json:"EnvDescription"`
	// 计算环境 ID
	EnvId string `json:"EnvId"`
	// 计算环境名称
	EnvName string `json:"EnvName"`
	// 计算环境类型，仅支持“MANAGED”类型
	EnvType string `json:"EnvType"`
	// 输入映射
	InputMappings []*InputMapping `json:"InputMappings"`
	// 数据盘挂载选项
	MountDataDisks []*MountDataDisk `json:"MountDataDisks"`
	// 通知信息
	Notifications []*Notification `json:"Notifications"`
}

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

type Filter struct {
	// 需要过滤的字段。
	Name string `json:"Name"`
	// 字段的过滤值。
	Values []*string `json:"Values"`
}
