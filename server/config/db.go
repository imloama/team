//数据库基础功能封装
package config

type DB struct {
	Url      string //连接地址
	Name     string //数据库名称
	Username string //连接用户名
	Password string //连接密码
}

//初始化数据库
func init() {

}
