package ms

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取云COS临时密钥
// https://cloud.tencent.com/document/api/283/30640

type CreateCosSecKeyInstanceRequest struct {
	// 地域信息，例如广州：ap-guangzhou，上海：ap-shanghai，默认为广州。
	CosRegion *string `name:"CosRegion,omitempty"`
	// 密钥有效时间，默认为1小时。
	Duration *int64 `name:"Duration,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *CreateCosSecKeyInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateCosSecKeyInstanceResponse, error) {
	resp := &CreateCosSecKeyInstanceResponse{}
	err := client.Request("ms", "CreateCosSecKeyInstance", "2018-04-08").Do(req, resp)
	return resp, err
}

type CreateCosSecKeyInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// COS密钥对应的AppId
	CosAppid int64 `json:"CosAppid"`
	// COS密钥对应的存储桶名
	CosBucket string `json:"CosBucket"`
	// 密钥ID信息
	CosId string `json:"CosId"`
	// 密钥KEY信息
	CosKey string `json:"CosKey"`
	// 密钥可访问的文件前缀人。例如：CosPrefix=test/123/666，则该密钥只能操作test/123/666为前缀的文件，例如test/123/666/1.txt
	CosPrefix string `json:"CosPrefix"`
	// 存储桶对应的地域
	CosRegion string `json:"CosRegion"`
	// 密钥TOCKEN信息
	CosTocken string `json:"CosTocken"`
	// 密钥过期时间
	ExpireTime int64 `json:"ExpireTime"`
}
