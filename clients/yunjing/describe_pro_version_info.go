package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取专业版信息
// https://cloud.tencent.com/document/api/296/19865

type DescribeProVersionInfoRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeProVersionInfoRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeProVersionInfoResponse, error) {
	resp := &DescribeProVersionInfoResponse{}
	err := client.Request("yunjing", "DescribeProVersionInfo", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeProVersionInfoResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 新增主机是否自动开通专业版
	IsAutoOpenProVersion bool `json:"IsAutoOpenProVersion"`
	// 后付费昨日扣费
	PostPayCost int64 `json:"PostPayCost"`
	// 开通专业版主机数
	ProVersionNum int64 `json:"ProVersionNum"`
}
