package cws

import (
	time "time"
)

// 监控任务中的站点信息。

type MonitorMiniSite struct {
	// 站点ID。
	SiteId *int64 `json:"SiteId,omitempty"`
	// 站点Url。
	Url *string `json:"Url,omitempty"`
}

// 监控任务详细数据

type MonitorsDetail struct {
	// 监控任务基础信息。
	Basic *Monitor `json:"Basic,omitempty"`
	// 监控任务包含的受漏洞威胁的站点列表数量。
	ImpactSiteNumber *int64 `json:"ImpactSiteNumber,omitempty"`
	// 监控任务包含的受漏洞威胁的站点列表。
	ImpactSites []*MonitorMiniSite `json:"ImpactSites,omitempty"`
	// 扫描页面总数。
	PageCount int64 `json:"PageCount"`
	// 监控任务包含的站点列表的平均扫描进度。
	Progress int64 `json:"Progress"`
	// 监控任务包含的站点列表数量。
	SiteNumber *int64 `json:"SiteNumber,omitempty"`
	// 监控任务包含的站点列表。
	Sites []*MonitorMiniSite `json:"Sites,omitempty"`
	// 高风险漏洞数量。
	VulsHighNumber *int64 `json:"VulsHighNumber,omitempty"`
	// 低风险漏洞数量。
	VulsLowNumber *int64 `json:"VulsLowNumber,omitempty"`
	// 中风险漏洞数量。
	VulsMiddleNumber *int64 `json:"VulsMiddleNumber,omitempty"`
	// 提示数量。
	VulsNoticeNumber *int64 `json:"VulsNoticeNumber,omitempty"`
}

// 站点数据

type Site struct {
	// 云用户appid。
	Appid int64 `json:"Appid"`
	// 记录添加时间。
	CreatedAt *time.Time `json:"CreatedAt,omitempty"`
	// 站点ID。
	Id *int64 `json:"Id,omitempty"`
	// 最近一次取消时间，取消即使用上一次扫描结果。
	LastScanCancelTime *time.Time `json:"LastScanCancelTime,omitempty"`
	// 最近一次扫描各个类型风险漏洞数量，存储的是json对象
	LastScanExtsCount string `json:"LastScanExtsCount"`
	// 最近一次扫描完成时间。
	LastScanFinishTime *time.Time `json:"LastScanFinishTime,omitempty"`
	// 最近一次扫描提示总数量
	LastScanNoticeNum *int64 `json:"LastScanNoticeNum,omitempty"`
	// 最近一次扫描扫描的页面数。
	LastScanPageCount *int64 `json:"LastScanPageCount,omitempty"`
	// 速率限制，每秒发送X个HTTP请求。
	LastScanRateLimit *int64 `json:"LastScanRateLimit,omitempty"`
	// normal-正常扫描；deep-深度扫描。
	LastScanScannerType *string `json:"LastScanScannerType,omitempty"`
	// 最近一次扫描开始时间。
	LastScanStartTime *time.Time `json:"LastScanStartTime,omitempty"`
	// 最近一次的AIScanner的扫描任务id，注意取消的情况。
	LastScanTaskId *int64 `json:"LastScanTaskId,omitempty"`
	// 最近一次扫描高风险漏洞数量。
	LastScanVulsHighNum *int64 `json:"LastScanVulsHighNum,omitempty"`
	// 最近一次扫描低风险漏洞数量。
	LastScanVulsLowNum *int64 `json:"LastScanVulsLowNum,omitempty"`
	// 最近一次扫描中风险漏洞数量。
	LastScanVulsMiddleNum *int64 `json:"LastScanVulsMiddleNum,omitempty"`
	// 最近一次扫描提示信息数量。
	LastScanVulsNoticeNum *int64 `json:"LastScanVulsNoticeNum,omitempty"`
	// 最近一次扫描漏洞总数量。
	LastScanVulsNum *int64 `json:"LastScanVulsNum,omitempty"`
	// 监控任务ID，为0时表示未加入监控任务。
	MonitorId *int64 `json:"MonitorId,omitempty"`
	// 监测状态：0-未监测；1-监测中；2-暂停监测。
	MonitorStatus *int64 `json:"MonitorStatus,omitempty"`
	// 站点名称。
	Name *string `json:"Name,omitempty"`
	// 扫描进度，百分比整数
	Progress int64 `json:"Progress"`
	// 扫描状态：0-待扫描（无任何扫描结果）；1-扫描中（正在进行扫描）；2-已扫描（有扫描结果且不正在扫描）；3-扫描完成待同步结果。
	ScanStatus *int64 `json:"ScanStatus,omitempty"`
	// 云用户标识。
	Uin string `json:"Uin"`
	// 记录最近修改时间。
	UpdatedAt *time.Time `json:"UpdatedAt,omitempty"`
	// 站点url。
	Url *string `json:"Url,omitempty"`
	// 验证状态：0-未验证；1-已验证；2-验证失效，待重新验证。
	VerifyStatus *int64 `json:"VerifyStatus,omitempty"`
}

// 站点验证数据

type SitesVerification struct {
	// 云用户appid
	Appid int64 `json:"Appid"`
	// CreatedAt。
	CreatedAt *time.Time `json:"CreatedAt,omitempty"`
	// 根域名。
	Domain *string `json:"Domain,omitempty"`
	// ID。
	Id int64 `json:"Id"`
	// txt解析域名验证的name。
	TxtName *string `json:"TxtName,omitempty"`
	// txt解析域名验证的text。
	TxtText *string `json:"TxtText,omitempty"`
	// UpdatedAt。
	UpdatedAt *time.Time `json:"UpdatedAt,omitempty"`
	// 验证有效期，在此之前有效。
	ValidTo *time.Time `json:"ValidTo,omitempty"`
	// 验证状态：0-未验证；1-已验证；2-验证失效，待重新验证。
	VerifyStatus *int64 `json:"VerifyStatus,omitempty"`
}

// 漏洞数据

type Vul struct {
	// CreatedAt。
	CreatedAt *time.Time `json:"CreatedAt,omitempty"`
	// 漏洞描述。
	Describe *string `json:"Describe,omitempty"`
	// 漏洞参考。
	From *string `json:"From,omitempty"`
	// 危害说明。
	Harm *string `json:"Harm,omitempty"`
	// 网址/细节。
	Html *string `json:"Html,omitempty"`
	// 漏洞ID。
	Id *int64 `json:"Id,omitempty"`
	// 是否已经添加误报，0-否，1-是。
	IsReported int64 `json:"IsReported"`
	// 漏洞级别：high、middle、low、notice。
	Level *string `json:"Level,omitempty"`
	// 漏洞名称。
	Name *string `json:"Name,omitempty"`
	// 漏洞类型。
	Nickname *string `json:"Nickname,omitempty"`
	// 漏洞通过该参数攻击。
	Parameter *string `json:"Parameter,omitempty"`
	// 站点ID。
	SiteId *int64 `json:"SiteId,omitempty"`
	// 解决方案。
	Solution *string `json:"Solution,omitempty"`
	// 扫描引擎的扫描任务ID。
	TaskId *int64 `json:"TaskId,omitempty"`
	// UpdatedAt。
	UpdatedAt *time.Time `json:"UpdatedAt,omitempty"`
	// 出现漏洞的url。
	Url *string `json:"Url,omitempty"`
}

// 描述键值对过滤器，用于条件过滤查询。例如过滤ID、名称、状态等

type Filter *Filter

// 监控任务基础数据

type Monitor struct {
	// 云用户appid。
	Appid int64 `json:"Appid"`
	// CreatedAt。
	CreatedAt *time.Time `json:"CreatedAt,omitempty"`
	// 扫描周期，单位小时，每X小时执行一次。
	Crontab *int64 `json:"Crontab,omitempty"`
	// 当前扫描开始时间，如扫描完成则为上一次扫描的开始时间。
	CurrentScanStartTime *time.Time `json:"CurrentScanStartTime,omitempty"`
	// 首次扫描开始时间。
	FirstScanStartTime *time.Time `json:"FirstScanStartTime,omitempty"`
	// 监控任务ID。
	Id *int64 `json:"Id,omitempty"`
	// 指定扫描类型，3位数每位依次表示：扫描Web漏洞、扫描系统漏洞、扫描系统端口。
	IncludedVulsTypes *string `json:"IncludedVulsTypes,omitempty"`
	// 上一次扫描完成时间。
	LastScanFinishTime *time.Time `json:"LastScanFinishTime,omitempty"`
	// 监测状态：1-监测中；2-暂停监测。
	MonitorStatus *int64 `json:"MonitorStatus,omitempty"`
	// 监控名称。
	Name *string `json:"Name,omitempty"`
	// 速率限制，每秒发送X个HTTP请求。
	RateLimit *int64 `json:"RateLimit,omitempty"`
	// 扫描状态：0-待扫描（无任何扫描结果）；1-扫描中（正在进行扫描）；2-已扫描（有扫描结果且不正在扫描）；3-扫描完成待同步结果。
	ScanStatus *int64 `json:"ScanStatus,omitempty"`
	// 监测模式，normal-正常扫描；deep-深度扫描。
	ScannerType *string `json:"ScannerType,omitempty"`
	// UpdatedAt。
	UpdatedAt *time.Time `json:"UpdatedAt,omitempty"`
}
