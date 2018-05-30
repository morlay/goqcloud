package batch

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 提交作业
// https://cloud.tencent.com/document/api/599/15907
type SubmitJobRequest struct {
	// 用于保证请求幂等性的字符串。该字符串由用户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `name:"ClientToken,omitempty"`
	// 作业信息
	Job Job `name:"Job"`
	// 作业所提交的位置信息。通过该参数可以指定作业关联CVM所属可用区等信息。
	Placement []*Placement `name:"Placement"`
	// 区域
	Region string `name:"Region"`
}

func (req *SubmitJobRequest) Invoke(client github_com_morlay_goqcloud.Client) (*SubmitJobResponse, error) {
	resp := &SubmitJobResponse{}
	err := client.Request("batch", "SubmitJob", "2017-03-12").Do(req, resp)
	return resp, err
}

type SubmitJobResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 当通过本接口来提交作业时会返回该参数，表示一个作业ID。返回作业ID列表并不代表作业解析/运行成功，可根据 DescribeJob 接口查询其状态。
	JobId string `json:"JobId"`
}
