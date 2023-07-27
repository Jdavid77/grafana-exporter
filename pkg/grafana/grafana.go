package grafana

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"grafana_exporter/pkg/logger"

	grafana "github.com/grafana/grafana-api-golang-client"
	"go.uber.org/zap"
)

type Grafana struct {
	Client *grafana.Client
	L      *zap.SugaredLogger
	Url    string
	ApiKey string
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
		L:      logger.InitZapLog(),
		Url:    url,
		ApiKey: apiKey,
	}
}

func (g Grafana) DownloadDashboards() {

	dashboards, err := g.Client.Dashboards()
	g.L.Info(len(dashboards))
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

			dashboard, err := g.Client.DashboardByUID(dashboards[i].UID)

			if err != nil {
				g.L.Errorf("Dashboard %s Not Found", dashboards[i].Slug)
				return
			}

			folder, err := g.Client.Folder(dashboard.FolderID)

			if err != nil {
				g.L.Errorf("Folder %s Not Found", folder)
				return
			}

			if folder.Title != "General" {

				folder := "backup/" + folder.Title
				dashboardPath := folder + "/" + dashboard.Meta.Slug

				if _, err := os.Stat(folder); os.IsNotExist(err) {
					err := os.MkdirAll(folder, 0755)
					if err != nil {
						panic(err)
					}
				}

				if err := exportDashboard(dashboardPath, dashboard.Model, dashboard.Meta.Slug); err != nil {
					g.L.Infof("DASHBOARD %s COULDN'T BE EXPORTED: %s\n", dashboard.Meta.Slug, err)
				} else {
					g.L.Infof("DASHBOARD SUCCESSFULLY EXPORTED: %s\n", dashboard.Meta.Slug)
				}

			}

			
		}(i)
	}
	wg.Wait()
	fmt.Println("Finished for loop")
}


func exportDashboard(path string, dashboard interface{}, slug string) error {

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(dashboard)
	if err != nil {
		return err
	}

	return nil
}