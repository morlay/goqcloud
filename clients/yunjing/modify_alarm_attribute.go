package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改告警设置
// https://cloud.tencent.com/document/api/296/19864

type ModifyAlarmAttributeRequest struct {
	// 告警项目。Offline：防护软件离线Malware：发现木马文件NonlocalLogin：发现异地登录行为CrackSuccess：被暴力破解成功
	Attribute string `name:"Attribute"`
	// 区域
	Region string `name:"Region"`
	// 告警项目属性。CLOSE：关闭OPEN：打开
	Value string `name:"Value"`
}

func (req *ModifyAlarmAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyAlarmAttributeResponse, error) {
	resp := &ModifyAlarmAttributeResponse{}
	err := client.Request("yunjing", "ModifyAlarmAttribute", "2018-02-28").Do(req, resp)
	return resp, err
}

type ModifyAlarmAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
