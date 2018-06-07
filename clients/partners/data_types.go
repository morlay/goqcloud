package partners

import (
	time "time"
)

// 业务信息定义

type AgentBillElem struct {
	// 支付金额，单位分
	Amt int64 `json:"Amt"`
	// 代客备注名称
	ClientRemark string `json:"ClientRemark"`
	// 代客账号ID
	ClientUin string `json:"ClientUin"`
	// 云产品名称
	GoodsType string `json:"GoodsType"`
	// 订单号，仅对预付费账单有意义
	OrderId string `json:"OrderId"`
	// 预付费/后付费
	PayMode string `json:"PayMode"`
	// 支付时间
	PayTime time.Time `json:"PayTime"`
	// agentpay：代付；selfpay：自付
	PayerMode string `json:"PayerMode"`
	// 支付月份
	SettleMonth string `json:"SettleMonth"`
	// 代理商账号ID
	Uin string `json:"Uin"`
}

// 描述代客信息

type AgentClientElem struct {
	// 代客申请时间戳
	ApplyTime int64 `json:"ApplyTime"`
	// 代客类型，可能值为a/b
	ClientFlag string `json:"ClientFlag"`
	// 代客账号ID
	ClientUin string `json:"ClientUin"`
	// 0表示不欠费，1表示欠费
	HasOverdueBill int64 `json:"HasOverdueBill"`
	// 代客邮箱，打码显示
	Mail string `json:"Mail"`
	// 代客手机，打码显示
	Phone string `json:"Phone"`
	// 代理商账号ID
	Uin string `json:"Uin"`
}

// 返佣信息定义

type RebateInfoElem struct {
	// 返佣金额，单位分
	Amt int64 `json:"Amt"`
	// NORMAL(正常)/HAS_OVERDUE_BILL(欠费)/NO_CONTRACT(缺合同)
	ExceptionFlag string `json:"ExceptionFlag"`
	// 月度业绩，单位分
	MonthSales int64 `json:"MonthSales"`
	// 季度业绩，单位分
	QuarterSales int64 `json:"QuarterSales"`
	// 返佣月份，如2018-02
	RebateMonth string `json:"RebateMonth"`
	// 代理商账号ID
	Uin string `json:"Uin"`
}
