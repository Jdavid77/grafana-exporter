package main

import (
	"internal/configuration"
	"internal/logger"

	grafana "github.com/grafana/grafana-api-golang-client"
)

func main() {
	log := logger.InitZapLog()
	config, err := configuration.LoadConfig("config.yaml")
	log.Info("Starting application..")
	if err != nil {
		log.Error(err.Error())
	}

	grafanaConfig := grafana.Config{

		APIKey:      config.Grafana.ApiKey,
		BasicAuth:   nil,
		HTTPHeaders: nil,
		Client:      nil,
		OrgID:       1,
		NumRetries:  2,
	}

	client, err := grafana.New(config.Grafana.Url, grafanaConfig)

	if err != nil {
		log.Error("Error connecting to Grafana API")
	}

	dashboards, err := client.Dashboards()

	if err != nil {
		log.Error("Error retrieving Dashboards")
	}
	log.Infof("dashboard %v", dashboards)

}
