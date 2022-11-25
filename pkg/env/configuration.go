package env

import (
	"errors"

	"github.com/spf13/viper"
)

type Config struct {
	GranafaURL string
	ApiKey     string

}

func LoadConfig() (config *Config, err error) {
	viper.AutomaticEnv()
	grafanaURL := viper.GetString("GRAFANA_URL")
	if len(grafanaURL) == 0 {
		return nil, errors.New("Grafana URL not set")
	}
	apiKey := viper.GetString("API_KEY")
	if len(apiKey) == 0 {
		return nil, errors.New("API_KEY not set")
	}

	result := &Config{
		GranafaURL: grafanaURL,
		ApiKey: apiKey,
	}

	return result, nil
}


