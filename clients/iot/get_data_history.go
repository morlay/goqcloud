package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 获取数据历史
// https://cloud.tencent.com/document/api/568/16521
type GetDataHistoryRequest struct {
	// 设备名称列表，允许最多一次100台
	DeviceNames []*string `name:"DeviceNames"`
	// 查询结束时间
	EndTime time.Time `name:"EndTime"`
	// 时间排序（desc/asc）
	Order *string `name:"Order,omitempty"`
	// 产品Id
	ProductId string `name:"ProductId"`
	// 区域
	Region string `name:"Region"`
	// 查询游标
	ScrollId *string `name:"ScrollId,omitempty"`
	// 查询数据量
	Size []*int64 `name:"Size,omitempty"`
	// 查询开始时间
	StartTime time.Time `name:"StartTime"`
}

func (req *GetDataHistoryRequest) Invoke(client github_com_morlay_goqcloud.Client) (*GetDataHistoryResponse, error) {
	resp := &GetDataHistoryResponse{}
	err := client.Request("iot", "GetDataHistory", "2018-01-23").Do(req, resp)
	return resp, err
}

type GetDataHistoryResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 数据历史
	DataHistory []*Object `json:"DataHistory"`
	// 查询游标
	ScrollId string `json:"ScrollId"`
}
