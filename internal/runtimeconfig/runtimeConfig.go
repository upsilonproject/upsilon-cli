package runtimeconfig

var RuntimeConfig = defaultConfig()

type runtimeConfig struct {
	OutputFormat string
	LogLevel string

	AmqpHost string
	AmqpUser string
	AmqpPass string
	AmqpPort int
}

func defaultConfig() *runtimeConfig {
	return &runtimeConfig{
		OutputFormat: "table",
		LogLevel: "info",
		AmqpHost: "upsilon",
		AmqpUser: "guest",
		AmqpPass: "guest",
		AmqpPort: 5672,
	}
}
