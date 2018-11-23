package ms

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 将应用绑定到资源
// https://cloud.tencent.com/document/api/283/18579

type CreateBindInstanceRequest struct {
	// app的icon的url
	AppIconUrl string `name:"AppIconUrl"`
	// app的名称
	AppName string `name:"AppName"`
	// app的包名
	AppPkgName string `name:"AppPkgName"`
	// 区域
	Region string `name:"Region"`
	// 资源id，全局唯一
	ResourceId string `name:"ResourceId"`
}

func (req *CreateBindInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateBindInstanceResponse, error) {
	resp := &CreateBindInstanceResponse{}
	err := client.Request("ms", "CreateBindInstance", "2018-04-08").Do(req, resp)
	return resp, err
}

type CreateBindInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	Progress int64 `json:"Progress"`
}
