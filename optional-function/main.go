package main

import (
	"fmt"
)

type Option struct {
	Port int
	Mode string
}

type Server struct {
	Option
}

type ServerBuilder struct {
	option Option
}

func NewServerBuilder() *ServerBuilder {
	return &ServerBuilder{}
}

func (b *ServerBuilder) WithPort(port int) *ServerBuilder {
	b.option.Port = port
	return b
}

func (b *ServerBuilder) WithMode(mode string) *ServerBuilder {
	b.option.Mode = mode
	return b
}

func (b *ServerBuilder) Build() *Server {
	return &Server{Option: b.option}
}

func main() {
	serverBuilder := NewServerBuilder()

	server := serverBuilder.WithPort(8080).WithMode("release").Build()

	fmt.Printf("server is running on port %d with %s mode", server.Port, server.Mode)
}
