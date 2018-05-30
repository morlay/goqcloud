package mariadb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 设置账号权限
// https://cloud.tencent.com/document/api/237/16166
type GrantAccountPrivilegesRequest struct {
	// 当 Type=table 时，ColName 为 * 表示对表授权，如果为具体字段名，表示对字段授权
	ColName *string `name:"ColName,omitempty"`
	// 数据库名。如果为 *，表示设置全局权限（即 *.*），此时忽略 Type 和 Object 参数
	DbName string `name:"DbName"`
	// 用户允许的访问 host，用户名+host唯一确定一个账号。
	Host string `name:"Host"`
	// 实例 ID，形如：tdsql-ow728lmc，可以通过 DescribeDBInstances 查询实例详情获得。
	InstanceId string `name:"InstanceId"`
	// 具体的 Type 的名称，比如 Type 为 table 时就是具体的表名。DbName 和 Type 都为具体名称，则 Object 表示具体对象名，不能为 * 或者为空
	Object *string `name:"Object,omitempty"`
	// 全局权限： SELECT，INSERT，UPDATE，DELETE，CREATE，DROP，REFERENCES，INDEX，ALTER，CREATE TEMPORARY TABLES，LOCK TABLES，EXECUTE，CREATE VIEW，SHOW VIEW，CREATE ROUTINE，ALTER ROUTINE，EVENT，TRIGGER，SHOW DATABASES 库权限： SELECT，INSERT，UPDATE，DELETE，CREATE，DROP，REFERENCES，INDEX，ALTER，CREATE TEMPORARY TABLES，LOCK TABLES，EXECUTE，CREATE VIEW，SHOW VIEW，CREATE ROUTINE，ALTER ROUTINE，EVENT，TRIGGER 表/视图权限： SELECT，INSERT，UPDATE，DELETE，CREATE，DROP，REFERENCES，INDEX，ALTER，CREATE VIEW，SHOW VIEW，TRIGGER 存储过程/函数权限： ALTER ROUTINE，EXECUTE 字段权限： INSERT，REFERENCES，SELECT，UPDATE
	Privileges []*string `name:"Privileges"`
	// 区域
	Region string `name:"Region"`
	// 类型,可以填入 table 、 view 、 proc 、 func 和 *。当 DbName 为具体数据库名，Type为 * 时，表示设置该数据库权限（即db.*），此时忽略 Object 参数
	Type *string `name:"Type,omitempty"`
	// 登录用户名。
	UserName string `name:"UserName"`
}

func (req *GrantAccountPrivilegesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*GrantAccountPrivilegesResponse, error) {
	resp := &GrantAccountPrivilegesResponse{}
	err := client.Request("mariadb", "GrantAccountPrivileges", "2017-03-12").Do(req, resp)
	return resp, err
}

type GrantAccountPrivilegesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}
