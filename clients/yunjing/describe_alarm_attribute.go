package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取告警设置
// https://cloud.tencent.com/document/api/296/19866

type DescribeAlarmAttributeRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeAlarmAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeAlarmAttributeResponse, error) {
	resp := &DescribeAlarmAttributeResponse{}
	err := client.Request("yunjing", "DescribeAlarmAttribute", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeAlarmAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 被暴力破解成功告警状态：OPEN：告警已开启CLOSE： 告警已关闭
	CrackSuccess string `json:"CrackSuccess"`
	// 发现木马告警状态：OPEN：告警已开启CLOSE： 告警已关闭
	Malware string `json:"Malware"`
	// 发现异地登录告警状态：OPEN：告警已开启CLOSE： 告警已关闭
	NonlocalLogin string `json:"NonlocalLogin"`
	// 防护软件离线告警状态：OPEN：告警已开启CLOSE： 告警已关闭
	Offline string `json:"Offline"`
}
