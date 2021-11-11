package runtimeconfig

var RuntimeConfig = defaultConfig()

type runtimeConfig struct {
	OutputFormat string
}

func defaultConfig() *runtimeConfig {
	return &runtimeConfig{
		OutputFormat: "table",
	}
}
