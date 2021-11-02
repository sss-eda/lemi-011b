package main

import "context"

type Environment string

const (
	Development Environment = "Development"
	Testing                 = "Testing"
	Staging                 = "Staging"
	Production              = "Production"
)

// Config TODO
type Config struct {
	Environment Environment `json:"environment", yaml:"environment", env:"ENVIRONMENT"`
	// API
	// API         string      `json:"api", yaml:"api", env:"API_TYPE"`
	// Storage     string      `json:"storage", yaml:"storage", env:"STORAGE_TYPE"`
}

// ConfigParser TODO
type ConfigParser func(context.Context, *Config) error

// Parse TODO
func (config *Config) Parse(
	ctx context.Context,
	parser ConfigParser,
) error {
	return parser(ctx, config)
}

// // StorageType TODO
// type StorageType string

// const (
// 	TimescaleDB StorageType = "TimescaleDB"
// 	InfluxDB    StorageType = "InfluxDB"
// )

// type APIType string

// const (
// 	RestHttp  StorageType = "REST HTTP"
// 	RestHttps StorageType = "REST HTTPS"
// )

// // TimescaleDBConfig TODO
// type TimescaleDBConfig struct {
// 	URL string `json:"url", yaml:"url", env:"TIMESCALEDB_URL"`
// }

// type InfluxDBConfig struct {
// 	Host     string `json:"host", yaml:"host", env:"INFLUXDB_HOST"`
// 	Port     int    `json:"port", yaml:"port", env:"INFLUXDB_PORT"`
// 	Database string `json:"database", yaml:"database", env:"INFLUXDB_DATABASE"`
// }

// // SSLConfig TODO
// type SSLConfig struct {
// 	Public  string `json:"public", yaml:"public", env:"SSL_PUBLIC_CERT"`
// 	Private string `json:"private", yaml:"private", env:"SSL_PRIVATE_CERT"`
// }
