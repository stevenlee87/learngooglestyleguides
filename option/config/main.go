package main

import "fmt"

type Config struct {
	Host   string
	Port   int
	MaxNum int
}

type Option func(cfg *Config)

func WithHost(host string) Option {
	return func(cfg *Config) {
		cfg.Host = host
	}
}

func WithPort(port int) Option {
	return func(cfg *Config) {
		cfg.Port = port
	}
}

func WithMaxNum(maxNum int) Option {
	return func(cfg *Config) {
		cfg.MaxNum = maxNum
	}
}

func NewConfig(options ...Option) Config {
	cfg := Config{}
	for _, option := range options {
		option(&cfg)
	}
	return cfg
}

func main() {
	cfg := NewConfig(
		WithHost("127.0.0.1"),
		WithPort(80),
		WithMaxNum(10),
	)
	fmt.Printf("cfg:%+v", cfg)
}
