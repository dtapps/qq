package qq

import (
	"go.dtapp.net/golog"
)

// Client 实例
type Client struct {
	key     string
	gormLog struct {
		status bool           // 状态
		client *golog.ApiGorm // 日志服务
	}
	mongoLog struct {
		status bool            // 状态
		client *golog.ApiMongo // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(key string) (*Client, error) {
	return &Client{key: key}, nil
}
