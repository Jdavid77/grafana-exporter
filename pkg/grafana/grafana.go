package grafana

import (
	"fmt"
	"sync"

	grafana "github.com/grafana/grafana-api-golang-client"
	"go.uber.org/zap"
)

type Grafana struct {
	Client *grafana.Client
	L      *zap.SugaredLogger
}

func NewGrafana(apiKey, url string) *Grafana {
	grafanaConfig := grafana.Config{
		APIKey:      apiKey,
		BasicAuth:   nil,
		HTTPHeaders: nil,
		Client:      nil,
		OrgID:       1,
		NumRetries:  2,
	}

	client, err := grafana.New(url, grafanaConfig)
	if err != nil {
		zap.S().Fatal("Cannot construct grafana client " + err.Error())
	}

	return &Grafana{
		Client: client,
		L:      zap.S(),
	}
}
func (g Grafana) DownloadDashboards() {

	dashboards, err := g.Client.Dashboards()
	if err != nil {
		g.L.Error("Error retrieving Dashboards")
	}

	sliceLength := len(dashboards)
	var wg sync.WaitGroup
	wg.Add(sliceLength)
	g.L.Debugf("Running for loop…")
	for i := 0; i < sliceLength; i++ {
		go func(i int) {
			defer wg.Done()
			val := dashboards[i]
			g.L.Infof("i: %v, val: %v\n", i, val)
		}(i)
	}
	wg.Wait()
	fmt.Println("Finished for loop")
}
