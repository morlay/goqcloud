package cdb

// 灾备实例信息

type DrInfo *DrInfo

// 导入任务记录

type ImportRecord *ImportRecord

// 实例预期重启时间

type InstanceRebootTime *InstanceRebootTime

// 数据库实例参数

type Parameter *Parameter

// 用于回档的数据库表名

type RollbackTableName *RollbackTableName

// 可用区售卖配置

type ZoneSellConf *ZoneSellConf

// ECDB第二个从库的配置信息，只有ECDB实例才有这个字段

type BackupConfig *BackupConfig

// 云数据库切换记录

type DBSwitchInfo *DBSwitchInfo

// 用于回档的实例详情

type RollbackInstancesInfo *RollbackInstancesInfo

// 多可用区信息

type ZoneConf *ZoneConf

// 表名

type TableName *TableName

// 列权限信息

type ColumnPrivilege *ColumnPrivilege

// 数据库表名

type DatabaseName *DatabaseName

// 数据库权限

type DatabasePrivilege *DatabasePrivilege

// 实例可回档时间范围

type InstanceRollbackRangeTime *InstanceRollbackRangeTime

// 实例参数信息

type ParamInfo *ParamInfo

// 从库的配置信息

type SlaveConfig *SlaveConfig

// 备机信息

type SlaveInstanceInfo *SlaveInstanceInfo

// 数据库账号信息

type Account *Account

// 账号详细信息

type AccountInfo *AccountInfo

// 安全组入站规则

type Inbound *Inbound

// 数据库表权限

type TablePrivilege *TablePrivilege

// 备份详细信息

type BackupInfo *BackupInfo

// 主实例信息

type MasterInfo *MasterInfo

// 只读组参数

type RoGroup *RoGroup

// 只读vip信息

type RoVipInfo *RoVipInfo

// 可回档时间范围

type RollbackTimeRange *RollbackTimeRange

// 售卖实例类型

type SellType *SellType

// sql文件信息

type SqlFileInfo *SqlFileInfo

// 实例详细信息

type InstanceInfo *InstanceInfo

// 实例参数的详细描述

type ParameterDetail *ParameterDetail

// 用于回档的数据库表详情

type RollbackTables *RollbackTables

// 安全组详情

type SecurityGroup *SecurityGroup

// 售卖配置详情

type SellConfig *SellConfig

// 慢查询日志详情

type SlowLogInfo *SlowLogInfo

// 地域售卖配置

type RegionSellConf *RegionSellConf

// 备机信息

type SlaveInfo *SlaveInfo

// 文件上传描述

type UploadInfo *UploadInfo

// 二进制日志信息

type BinlogInfo *BinlogInfo

// 安全组出站规则

type Outbound *Outbound

// RO实例的详细信息

type RoInstanceInfo *RoInstanceInfo

// 用于回档的数据库名

type RollbackDBName *RollbackDBName
