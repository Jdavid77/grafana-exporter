package main

import (
	grafana "github.com/grafana/grafana-api-golang-client"
	"go.uber.org/zap"

	env "grafana-exporter/pkg/env"
)

func main() {
	log := zap.S()
	envConfig, err := env.LoadConfig("config.yaml")

	if err != nil {
		log.Error()
	}

	config := grafana.Config{
		APIKey:      envConfig.Grafana.ApiKey,
		BasicAuth:   nil,
		HTTPHeaders: nil,
		Client:      nil,
		OrgID:       1,
		NumRetries:  2,
	}

	client, err := grafana.New(envConfig.Grafana.Url, config)

	if err != nil {
		log.Error("Error connecting to Grafana API")
	}

	dashboards, err := client.Dashboards()

	if err != nil {
		log.Error("Error retrieving Dashboards")
	}
	log.Infof("%v", dashboards)

}
