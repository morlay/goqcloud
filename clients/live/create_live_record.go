package live

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建录制任务
// https://cloud.tencent.com/document/api/267/30148

type CreateLiveRecordRequest struct {
	// 直播流所属应用名称。
	AppName *string `name:"AppName,omitempty"`
	// 推流域名。多域名推流必须设置。
	DomainName *string `name:"DomainName,omitempty"`
	// 任务结束时间。若指定精彩视频录制，结束时间不超过当前时间+30分钟，如果超过或小于起始时间，则实际结束时间为当前时间+30分钟。
	EndTime *string `name:"EndTime,omitempty"`
	// 录制文件格式。不区分大小写。其值为：“flv”,“hls”,”mp4”,“aac”,”mp3”，默认“flv”。
	FileFormat *string `name:"FileFormat,omitempty"`
	// 精彩视频标志。0：普通视频【默认】；1：精彩视频。
	Highlight *int64 `name:"Highlight,omitempty"`
	// A+B=C混流标志。0：非A+B=C混流录制【默认】；1：标示为A+B=C混流录制。
	MixStream *int64 `name:"MixStream,omitempty"`
	// 录制类型。不区分大小写。“video” : 音视频录制【默认】。“audio” : 纯音频录制。
	RecordType *string `name:"RecordType,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 任务起始时间，录制视频为精彩视频时，忽略此字段。如 2017-01-01 10:10:01。
	StartTime *string `name:"StartTime,omitempty"`
	// 流名称。
	StreamName string `name:"StreamName"`
	// 录制流参数，当前支持以下参数： interval 录制分片时长，单位 秒，0 - 7200storage_time 录制文件存储时长，单位 秒eg. interval=3600&storage_time=7200注：参数需要url encode。
	StreamParam *string `name:"StreamParam,omitempty"`
}

func (req *CreateLiveRecordRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateLiveRecordResponse, error) {
	resp := &CreateLiveRecordResponse{}
	err := client.Request("live", "CreateLiveRecord", "2018-08-01").Do(req, resp)
	return resp, err
}

type CreateLiveRecordResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 任务ID，全局唯一标识录制任务。
	TaskId int64 `json:"TaskId"`
}
