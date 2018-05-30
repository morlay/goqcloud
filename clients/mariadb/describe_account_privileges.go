package mariadb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询账号权限
// https://cloud.tencent.com/document/api/237/16164
type DescribeAccountPrivilegesRequest struct {
	// 当 Type=table 时，ColName 为 * 表示查询表的权限，如果为具体字段名，表示查询对应字段的权限
	ColName *string `name:"ColName,omitempty"`
	// 数据库名。如果为 *，表示查询全局权限（即 *.*），此时忽略 Type 和 Object 参数
	DbName string `name:"DbName"`
	// 用户允许的访问 host，用户名+host唯一确定一个账号。
	Host string `name:"Host"`
	// 实例 ID，形如：tdsql-ow728lmc，可以通过 DescribeDBInstances 查询实例详情获得。
	InstanceId string `name:"InstanceId"`
	// 具体的 Type 的名称，比如 Type 为 table 时就是具体的表名。DbName 和 Type 都为具体名称，则 Object 表示具体对象名，不能为 * 或者为空
	Object *string `name:"Object,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 类型,可以填入 table 、 view 、 proc 、 func 和 *。当 DbName 为具体数据库名，Type为 * 时，表示查询该数据库权限（即db.*），此时忽略 Object 参数
	Type *string `name:"Type,omitempty"`
	// 登录用户名。
	UserName string `name:"UserName"`
}

func (req *DescribeAccountPrivilegesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeAccountPrivilegesResponse, error) {
	resp := &DescribeAccountPrivilegesResponse{}
	err := client.Request("mariadb", "DescribeAccountPrivileges", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeAccountPrivilegesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 数据库账号Host
	Host string `json:"Host"`
	// 实例Id
	InstanceId string `json:"InstanceId"`
	// 权限列表。
	Privileges []*string `json:"Privileges"`
	// 数据库账号用户名
	UserName string `json:"UserName"`
}
