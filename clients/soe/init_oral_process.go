package soe

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 发音评估初始化
// https://cloud.tencent.com/document/api/884/19319

type InitOralProcessRequest struct {
	// 评估模式，0:词模式, 1:句子模式，当为词模式评估时，能够提供每个音节的评估信息，当为句子模式时，能够提供完整度和流利度信息。
	EvalMode int64 `name:"EvalMode"`
	// 长效session标识，当该参数为1时，session的持续时间为300s，但会一定程度上影响第一个数据包的返回速度，且TransmitOralProcess必须同时为1才可生效。
	IsLongLifeSession *int64 `name:"IsLongLifeSession,omitempty"`
	// 被评估语音对应的文本，不支持ascii大于128以上的字符，会统一替换成空格。
	RefText string `name:"RefText"`
	// 区域
	Region string `name:"Region"`
	// 评价苛刻指数，取值为[1.0 - 4.0]范围内的浮点数，用于平滑不同年龄段的分数，1.0为小年龄段，4.0为最高年龄段
	ScoreCoeff float64 `name:"ScoreCoeff"`
	// 语音段唯一标识，一段语音一个SessionId
	SessionId string `name:"SessionId"`
	// 业务应用ID，与账号应用APPID无关，是用来方便客户管理服务的参数，需要结合控制台使用。
	SoeAppId *string `name:"SoeAppId,omitempty"`
	// 语音输入模式，0流式分片，1非流式一次性评估
	WorkMode int64 `name:"WorkMode"`
}

func (req *InitOralProcessRequest) Invoke(client github_com_morlay_goqcloud.Client) (*InitOralProcessResponse, error) {
	resp := &InitOralProcessResponse{}
	err := client.Request("soe", "InitOralProcess", "2018-07-24").Do(req, resp)
	return resp, err
}

type InitOralProcessResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 语音段唯一标识，一个完整语音一个SessionId
	SessionId string `json:"SessionId"`
}
