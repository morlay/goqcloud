package tia

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除Job
// https://cloud.tencent.com/document/api/851/18307

type DeleteJobRequest struct {
	// 运行任务的集群
	Cluster string `name:"Cluster"`
	// 任务名称
	Name string `name:"Name"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteJobRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteJobResponse, error) {
	resp := &DeleteJobResponse{}
	err := client.Request("tia", "DeleteJob", "2018-02-26").Do(req, resp)
	return resp, err
}

type DeleteJobResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
