package main

import "fmt"

// ServerConfig 配置结构体
type ServerConfig struct {
	Host string
	Port int
	TLS  bool
}

// Option 是一个函数类型, 用于修改 ServerConfig
type Option func(*ServerConfig)

// NewServer 使用默认参数初始化ServerConfig, 并使用Option函数定制参数
func NewServer(options ...Option) *ServerConfig {
	// 设置默认值
	config := &ServerConfig{
		Host: "localhost",
		Port: 8080,
		TLS:  false,
	}

	// 应用所有选项
	for _, option := range options {
		option(config)
	}
	return config
}

// WithHost 选配函数
func WithHost(host string) Option {
	return func(config *ServerConfig) {
		config.Host = host
	}
}

// WithPort 选配函数
func WithPort(port int) Option {
	return func(config *ServerConfig) {
		config.Port = port
	}
}

// WithTLS 选配函数
func WithTLS(tls bool) Option {
	return func(config *ServerConfig) {
		config.TLS = tls
	}
}

func main() {
	// 创建 Server 实例，使用默认配置
	server1 := NewServer()
	fmt.Printf("Server1: %+v\n", server1)

	// 创建 Server 实例，使用部分可选配置
	server2 := NewServer(
		WithHost("example.com"),
		WithTLS(true),
	)
	fmt.Printf("Server2: %+v\n", server2)

	// 创建 Server 实例，使用全部可选配置
	server3 := NewServer(
		WithHost("example.com"),
		WithPort(9090),
		WithTLS(true),
	)
	fmt.Printf("Server3: %+v\n", server3)
}
