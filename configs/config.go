package configs

import (
	"github.com/spf13/viper"
	"os"
)

var (
	config *Config
)

type option struct {
	configFolder string
	configFile   string
	configType   string
}

func Init(opts ...Option) error {
	opt := &option{
		configFolder: getDefaultConfigFolder(),
		configFile:   getDefaultConfigFile(),
		configType:   getDefaultConfigType(),
	}

	for _, optFunc := range opts {
		optFunc(opt)
	}

	// Config File Path
	viper.AddConfigPath(opt.configFolder)
	// Config File Name
	viper.SetConfigName(opt.configFile)
	// Config File Type
	viper.SetConfigType(opt.configType)
	viper.AutomaticEnv()

	// Application config --> config/config.yaml
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	if err = viper.Unmarshal(&config); err != nil {
		return err
	}

	// Infrastructure --> Env vars
	if err = config.setConfigFromENV(); err != nil {
		return err
	}

	return nil
}

type Option func(*option)

func WithConfigFolder(configFolder string) Option {
	return func(opt *option) {
		opt.configFolder = configFolder
	}
}

func WithConfigFile(configFile string) Option {
	return func(opt *option) {
		env := os.Getenv("GO_ENV")
		if env == "" {
			opt.configFile = configFile
			return
		}

		opt.configFile = configFile + "." + env
	}
}

func WithConfigType(configType string) Option {
	return func(opt *option) {
		opt.configType = configType
	}
}

func getDefaultConfigFolder() string {
	configPath := "./configs/"

	return configPath
}

func getDefaultConfigFile() string {
	env := os.Getenv("GO_ENV")
	if env == "" {
		return "config"
	}

	return "config." + env
}

func getDefaultConfigType() string {
	return "yaml"
}

func Get() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}

func (cfg *Config) setConfigFromENV() (err error) {
	if pgTelkomReadConnstring := os.Getenv("PG_TELKOM_READ_CONNSTRING"); pgTelkomReadConnstring != "" {
		cfg.Postgre.Telkom.Read = pgTelkomReadConnstring
	}

	if pgTelkomWriteConnstring := os.Getenv("PG_TELKOM_WRITE_CONNSTRING"); pgTelkomWriteConnstring != "" {
		cfg.Postgre.Telkom.Write = pgTelkomWriteConnstring
	}

	return nil
}
