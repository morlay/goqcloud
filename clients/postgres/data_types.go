package postgres

// 描述实例的详细信息

type DBInstance *DBInstance

// 错误日志详情

type ErrLogDetail *ErrLogDetail

// 描述某个地域下某个可用区的可售卖规格详细信息。

type SpecInfo *SpecInfo

// 数据库Xlog信息

type Xlog *Xlog

// 数据库备份信息

type DBBackup *DBBackup

// 描述实例的网络连接信息

type DBInstanceNetInfo *DBInstanceNetInfo

// 描述一种规格的信息信息

type SpecItemInfo *SpecItemInfo

// 单条SlowQuery信息

type NormalQueryItem *NormalQueryItem

// 订单详情

type PgDeal *PgDeal

// 描述地域的编码和状态等信息

type RegionInfo *RegionInfo

// 慢查询详情

type SlowlogDetail *SlowlogDetail

// 描述可用区的编码和状态信息

type ZoneInfo *ZoneInfo

// 账户信息

type AccountInfo *AccountInfo

// 描述键值对过滤器，用于条件过滤查询。例如过滤ID、名称等

type Filter *Filter
