package cws

// 站点数据

type Site *Site

// 站点验证数据

type SitesVerification *SitesVerification

// 漏洞数据

type Vul *Vul

// 用户漏洞数随时间变化统计数据

type VulsTimeline *VulsTimeline

// 描述键值对过滤器，用于条件过滤查询。例如过滤ID、名称、状态等

type Filter *Filter

// 监控任务基础数据

type Monitor *Monitor

// 监控任务中的站点信息。

type MonitorMiniSite *MonitorMiniSite

// 监控任务详细数据

type MonitorsDetail *MonitorsDetail
