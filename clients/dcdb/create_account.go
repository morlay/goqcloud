package dcdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建账号
// https://cloud.tencent.com/document/api/557/19996

type CreateAccountRequest struct {
	// 如果备机延迟超过本参数设置值，系统将认为备机发生故障建议该参数值大于10。当ReadOnly选择1、2时该参数生效。
	DelayThresh *int64 `name:"DelayThresh,omitempty"`
	// 账号备注，可以包含中文、英文字符、常见符号和数字，长度为0~256字符
	Description *string `name:"Description,omitempty"`
	// 可以登录的主机，与mysql 账号的 host 格式一致，可以支持通配符，例如 %，10.%，10.20.%。
	Host string `name:"Host"`
	// 实例 ID，形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId string `name:"InstanceId"`
	// 账号密码，由字母、数字或常见符号组成，不能包含分号、单引号和双引号，长度为6~32位。
	Password string `name:"Password"`
	// 是否创建为只读账号，0：否， 1：该账号的sql请求优先选择备机执行，备机不可用时选择主机执行，2：优先选择备机执行，备机不可用时操作失败，3：只从备机读取。
	ReadOnly *int64 `name:"ReadOnly,omitempty"`
	// 区域
	Region string `name:"Region"`
	// AccountName
	UserName string `name:"UserName"`
}

func (req *CreateAccountRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateAccountResponse, error) {
	resp := &CreateAccountResponse{}
	err := client.Request("dcdb", "CreateAccount", "2018-04-11").Do(req, resp)
	return resp, err
}

type CreateAccountResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 允许访问的 host，透传入参。
	Host string `json:"Host"`
	// 实例Id，透传入参。
	InstanceId string `json:"InstanceId"`
	// 透传入参。
	ReadOnly int64 `json:"ReadOnly"`
	// 用户名，透传入参。
	UserName string `json:"UserName"`
}
