package env

import (
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type GrafanaConfig struct {
	Url    string `yaml:url`
	ApiKey string `yaml:apiKey`
}
type GitConfig struct {
	Username       string `yaml:"username"`
	Email          string `yaml:"email"`
	Url            string `yaml:"url"`
	Directory      string `yaml:"directory"`
	PrivateKeyFile string `yaml:"privateKeyFile"`
}

type Config struct {
	Grafana GrafanaConfig `yaml:"grafana"`
	Git     GitConfig     `yaml:"git"`
}

func LoadConfig(path string) (config *Config, err error) {
	// Allows env var to override the config file variable
	//https://github.com/spf13/viper/issues/188#issuecomment-1273983955
	if path != "" {
		viper.SetConfigFile(path)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
	}

	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		zap.S().Fatal("config: error reading config file: " + err.Error())
	}

	for _, key := range viper.AllKeys() {
		envKey := strings.ToUpper(strings.ReplaceAll(key, ".", "_"))
		err := viper.BindEnv(key, envKey)
		if err != nil {
			zap.S().Fatal("config: unable to bind env: " + err.Error())
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		zap.S().Fatal("config: unable to decode into struct: " + err.Error())
	}
	return config, nil
}
