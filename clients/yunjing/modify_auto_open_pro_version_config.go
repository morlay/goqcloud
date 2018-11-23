package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 设置新增主机自动开通专业版配置
// https://cloud.tencent.com/document/api/296/19863

type ModifyAutoOpenProVersionConfigRequest struct {
	// 区域
	Region string `name:"Region"`
	// 设置自动开通状态。CLOSE：关闭OPEN：打开
	Status string `name:"Status"`
}

func (req *ModifyAutoOpenProVersionConfigRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyAutoOpenProVersionConfigResponse, error) {
	resp := &ModifyAutoOpenProVersionConfigResponse{}
	err := client.Request("yunjing", "ModifyAutoOpenProVersionConfig", "2018-02-28").Do(req, resp)
	return resp, err
}

type ModifyAutoOpenProVersionConfigResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
