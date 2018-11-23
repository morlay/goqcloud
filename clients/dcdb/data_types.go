package dcdb

// 数据库视图信息

type DatabaseView *DatabaseView

// 订单信息

type Deal *Deal

// 参数约束

type ParamConstraint *ParamConstraint

// 分片节点可用区选择

type ShardZoneChooseInfo *ShardZoneChooseInfo

// 实例可售卖规格详细信息，创建实例和扩容实例时 NodeCount、Memory 确定售卖规格，硬盘大小可用区间为[MinStorage,MaxStorage]

type SpecConfigInfo *SpecConfigInfo

// 描述一条sql日志的详细信息。

type SqlLogItem *SqlLogItem

// 数据库列信息

type TableColumn *TableColumn

// 可用区信息

type ZonesInfo *ZonesInfo

// 约束类型值的范围

type ConstraintRange *ConstraintRange

// 分布式数据库实例信息

type DCDBInstanceInfo *DCDBInstanceInfo

// 数据库表信息

type DatabaseTable *DatabaseTable

// DB参数描述

type ParamDesc *ParamDesc

// 售卖可用区信息

type RegionInfo *RegionInfo

// 按机型分类的规格配置

type SpecConfig *SpecConfig

// 云数据库账号信息

type DBAccount *DBAccount

// 云数据库参数信息。

type DBParamValue *DBParamValue

// 描述分布式数据库分片信息。

type DCDBShardInfo *DCDBShardInfo

// 数据库函数信息

type DatabaseFunction *DatabaseFunction

// 数据库存储过程信息

type DatabaseProcedure *DatabaseProcedure

// 升级实例 -- 扩容分片类型

type ExpandShardConfig *ExpandShardConfig

// 分片信息

type ShardInfo *ShardInfo

// 升级实例 -- 切分分片类型

type SplitShardConfig *SplitShardConfig

// 升级实例 -- 新增分片类型

type AddShardConfig *AddShardConfig

// 数据库信息

type Database *Database

// 拉取的日志信息

type LogFileInfo *LogFileInfo

// 修改参数结果

type ParamModifyResult *ParamModifyResult
