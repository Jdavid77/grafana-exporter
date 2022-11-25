package main

import (
	grafana "github.com/grafana/grafana-api-golang-client"
	"fmt"
	env "grafana-exporter/pkg/env"

)


func main() {

	envConfig , err := env.LoadConfig()

	if err != nil {
		fmt.Println(err.Error())
	}

	config := grafana.Config{
		APIKey: envConfig.ApiKey,
		BasicAuth: nil,
		HTTPHeaders: nil,
		Client: nil,
		OrgID: 1,
		NumRetries: 2,
	}

	
	client , err := grafana.New(envConfig.GranafaURL,config)

	if err != nil {
		fmt.Println("Error connecting to Grafana API")
	}

	dashboards , err := client.Dashboards()

	if err != nil {
		fmt.Println("Error retrieving Dashboards")
	}

	fmt.Println(dashboards)
	

}