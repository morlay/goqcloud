package soe

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 发音数据传输接口
// https://cloud.tencent.com/document/api/884/19318

type TransmitOralProcessRequest struct {
	// 是否传输完毕标志，若为0表示未完毕，若为1则传输完毕开始评估，非流式模式下无意义。
	IsEnd int64 `name:"IsEnd"`
	// 长效session标识，当该参数为1时，session的持续时间为300s，但会一定程度上影响第一个数据包的返回速度。当InitOralProcess接口调用时此项为1时，此项必填1才可生效。
	IsLongLifeSession *int64 `name:"IsLongLifeSession,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 流式数据包的序号，从1开始，当IsEnd字段为1后后续序号无意义，当IsLongLifeSession不为1且为非流式模式时无意义。
	SeqId int64 `name:"SeqId"`
	// 语音段唯一标识，一个完整语音一个SessionId。
	SessionId string `name:"SessionId"`
	// 业务应用ID，与账号应用APPID无关，是用来方便客户管理服务的参数，需要结合控制台使用。
	SoeAppId *string `name:"SoeAppId,omitempty"`
	// 当前数据包数据, 流式模式下数据包大小可以按需设置，数据包大小必须 >= 4K，编码格式要求为BASE64。
	UserVoiceData string `name:"UserVoiceData"`
	// 语音编码类型    1:pcm。
	VoiceEncodeType int64 `name:"VoiceEncodeType"`
	// 语音文件类型  1:raw, 2:wav, 3:mp3(三种格式目前仅支持16k采样率16bit编码单声道，如有不一致可能导致评估不准确或失败)。
	VoiceFileType int64 `name:"VoiceFileType"`
}

func (req *TransmitOralProcessRequest) Invoke(client github_com_morlay_goqcloud.Client) (*TransmitOralProcessResponse, error) {
	resp := &TransmitOralProcessResponse{}
	err := client.Request("soe", "TransmitOralProcess", "2018-07-24").Do(req, resp)
	return resp, err
}

type TransmitOralProcessResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 发音精准度，取值范围[-1, 100]，当取-1时指完全不匹配，当为句子模式时，是所有已识别单词准确度的加权平均值。当为流式模式且请求中IsEnd未置1时，取值无意义
	PronAccuracy float64 `json:"PronAccuracy"`
	// 发音完整度，取值范围[0, 1]，当为词模式时，取值无意义；当为流式模式且请求中IsEnd未置1时，取值无意义
	PronCompletion float64 `json:"PronCompletion"`
	// 发音流利度，取值范围[0, 1]，当为词模式时，取值无意义；当为流式模式且请求中IsEnd未置1时，取值无意义
	PronFluency float64 `json:"PronFluency"`
	// 语音段唯一标识，一段语音一个SessionId
	SessionId string `json:"SessionId"`
	// 详细发音评估结果
	Words []*WordRsp `json:"Words"`
}
