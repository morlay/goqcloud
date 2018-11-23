package es

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 更新ES集群实例
// https://cloud.tencent.com/document/api/845/30629

type UpdateInstanceRequest struct {
	// 磁盘大小,单位GB
	DiskSize *int64 `name:"DiskSize,omitempty"`
	// 修改后的访问控制列表
	EsAcl *EsAcl `name:"EsAcl,omitempty"`
	// 修改后的配置项, JSON格式字符串
	EsConfig *string `name:"EsConfig,omitempty"`
	// 要操作的实例ID
	InstanceId string `name:"InstanceId"`
	// 修改后的实例名称, 1-50 个英文、汉字、数字、连接线-或下划线_
	InstanceName *string `name:"InstanceName,omitempty"`
	// 横向扩缩容后的节点个数
	NodeNum *int64 `name:"NodeNum,omitempty"`
	// 节点规格: ES.S1.SMALL2: 1 核 2GES.S1.MEDIUM4: 2 核 4G ES.S1.MEDIUM8: 2 核 8G ES.S1.LARGE16: 4 核 16G ES.S1.2XLARGE32: 8 核 32G ES.S1.4XLARGE64: 16 核 64G
	NodeType *string `name:"NodeType,omitempty"`
	// 重置后的Kibana密码, 8到16位，至少包括两项（[a-z,A-Z],[0-9]和[-!@#$%&^*+=_:;,.?]的特殊符号
	Password *string `name:"Password,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *UpdateInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*UpdateInstanceResponse, error) {
	resp := &UpdateInstanceResponse{}
	err := client.Request("es", "UpdateInstance", "2018-04-16").Do(req, resp)
	return resp, err
}

type UpdateInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
