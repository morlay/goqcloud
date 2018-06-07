package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 切换访问新实例
// https://cloud.tencent.com/document/api/236/15864

type SwitchForUpgradeRequest struct {
	// 实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *SwitchForUpgradeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*SwitchForUpgradeResponse, error) {
	resp := &SwitchForUpgradeResponse{}
	err := client.Request("cdb", "SwitchForUpgrade", "2017-03-20").Do(req, resp)
	return resp, err
}

type SwitchForUpgradeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
