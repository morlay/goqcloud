package as

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 创建定时任务
// https://cloud.tencent.com/document/api/377/20452

type CreateScheduledActionRequest struct {
	// 伸缩组ID
	AutoScalingGroupId string `name:"AutoScalingGroupId"`
	// 当定时任务触发时，设置的伸缩组期望实例数。
	DesiredCapacity int64 `name:"DesiredCapacity"`
	// 定时任务的结束时间，取值为北京时间（UTC+8），按照ISO8601标准，格式：YYYY-MM-DDThh:mm:ss+08:00。此参数与Recurrence需要同时指定，到达结束时间之后，定时任务将不再生效。
	EndTime *time.Time `name:"EndTime,omitempty"`
	// 当定时任务触发时，设置的伸缩组最大实例数。
	MaxSize int64 `name:"MaxSize"`
	// 当定时任务触发时，设置的伸缩组最小实例数。
	MinSize int64 `name:"MinSize"`
	// 定时任务的重复方式。为标准Cron格式此参数与EndTime需要同时指定。
	Recurrence *string `name:"Recurrence,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 定时任务名称。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超60个字节。同一伸缩组下必须唯一。
	ScheduledActionName string `name:"ScheduledActionName"`
	// 定时任务的首次触发时间，取值为北京时间（UTC+8），按照ISO8601标准，格式：YYYY-MM-DDThh:mm:ss+08:00。
	StartTime time.Time `name:"StartTime"`
}

func (req *CreateScheduledActionRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateScheduledActionResponse, error) {
	resp := &CreateScheduledActionResponse{}
	err := client.Request("as", "CreateScheduledAction", "2018-04-19").Do(req, resp)
	return resp, err
}

type CreateScheduledActionResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 定时任务ID
	ScheduledActionId string `json:"ScheduledActionId"`
}
