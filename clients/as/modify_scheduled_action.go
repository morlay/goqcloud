package as

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 修改定时任务
// https://cloud.tencent.com/document/api/377/20449

type ModifyScheduledActionRequest struct {
	// 当定时任务触发时，设置的伸缩组期望实例数。
	DesiredCapacity *int64 `name:"DesiredCapacity,omitempty"`
	// 定时任务的结束时间，取值为北京时间（UTC+8），按照ISO8601标准，格式：YYYY-MM-DDThh:mm:ss+08:00。此参数与Recurrence需要同时指定，到达结束时间之后，定时任务将不再生效。
	EndTime *time.Time `name:"EndTime,omitempty"`
	// 当定时任务触发时，设置的伸缩组最大实例数。
	MaxSize *int64 `name:"MaxSize,omitempty"`
	// 当定时任务触发时，设置的伸缩组最小实例数。
	MinSize *int64 `name:"MinSize,omitempty"`
	// 定时任务的重复方式。为标准Cron格式此参数与EndTime需要同时指定。
	Recurrence *string `name:"Recurrence,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 待修改的定时任务ID
	ScheduledActionId string `name:"ScheduledActionId"`
	// 定时任务名称。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超60个字节。同一伸缩组下必须唯一。
	ScheduledActionName *string `name:"ScheduledActionName,omitempty"`
	// 定时任务的首次触发时间，取值为北京时间（UTC+8），按照ISO8601标准，格式：YYYY-MM-DDThh:mm:ss+08:00。
	StartTime *time.Time `name:"StartTime,omitempty"`
}

func (req *ModifyScheduledActionRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyScheduledActionResponse, error) {
	resp := &ModifyScheduledActionResponse{}
	err := client.Request("as", "ModifyScheduledAction", "2018-04-19").Do(req, resp)
	return resp, err
}

type ModifyScheduledActionResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
