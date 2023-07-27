package main

import (
	"grafana_exporter/pkg/configuration"
	"grafana_exporter/pkg/grafana"
	"grafana_exporter/pkg/logger"
	"grafana_exporter/pkg/kustomization"

)

func main() {
	log := logger.InitZapLog()
	config, err := configuration.LoadConfig("config.yaml")
	log.Info("Starting application..")
	if err != nil {
		log.Error(err.Error())
	}

	grafana := grafana.NewGrafana(config.Grafana.ApiKey, config.Grafana.Url)
	grafana.DownloadDashboards()

	kustomization.GenerateKustomizations("backup/")
	

}
