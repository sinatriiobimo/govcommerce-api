package configs

type (
	Config struct {
		App         AppConfig         `yaml:"app"`
		Postgre     PostgreList       `yaml:"postgre"`
		TimeoutHTTP TimeoutHTTPConfig `yaml:"timeout_http"`
	}

	AppConfig struct {
		HttpPort        string `yaml:"httpPort"`
		ShutdownTimeout int    `yaml:"shutdownTimeout"`
		LogLevel        string `yaml:"logLevel"`
	}

	PostgreList struct {
		Telkom PostgreConnString `yaml:"telkom"`
	}

	PostgreConnString struct {
		Read  string `yaml:"read"`
		Write string `yaml:"write"`
	}

	TimeoutHTTPConfig struct {
		Write int `yaml:"write"`
		Read  int `yaml:"read"`
	}
)
