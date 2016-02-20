//日志相关功能
package config

import (
	"github.com/lunny/log"
)

// 初始化日志配置
func InitLogger() {
	w := log.NewFileWriter(log.FileOptions{
		ByType: log.ByDay,
		Dir:    "./logs",
	})
	log.Std.SetOutput(w)
}
