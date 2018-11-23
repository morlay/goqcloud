package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建数据导入任务
// https://cloud.tencent.com/document/api/236/15858

type CreateDbImportJobRequest struct {
	// 导入的目标数据库名，不传表示不指定数据库。
	DbName *string `name:"DbName,omitempty"`
	// 文件名称。该文件是指用户已上传到腾讯云的文件。
	FileName string `name:"FileName"`
	// 实例的ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId string `name:"InstanceId"`
	// 云数据库实例User账号的密码。
	Password *string `name:"Password,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 云数据库的用户名。
	User string `name:"User"`
}

func (req *CreateDbImportJobRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateDbImportJobResponse, error) {
	resp := &CreateDbImportJobResponse{}
	err := client.Request("cdb", "CreateDBImportJob", "2017-03-20").Do(req, resp)
	return resp, err
}

type CreateDbImportJobResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步任务的请求ID，可使用此ID查询异步任务的执行结果
	AsyncRequestId string `json:"AsyncRequestId"`
}
