package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	tests := map[string]struct {
		input_file string
		want       *Config
	}{
		"valid-config": {input_file: "testdata/valid-config.yaml", want: &Config{
			Grafana: GrafanaConfig{Url: "http://test.grafana.com", ApiKey: "superapikey"},
			Git:     GitConfig{Username: "gitusername", Email: "username@username.com", Url: "git@github.com:Jdavid77/grafana-exporter.git", Directory: "kubernetes/grafana/dashboards", PrivateKeyFile: "~/.ssh/id-rsa"}}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := LoadConfig(tc.input_file)
			assert.Equal(t, tc.want.Git.Url, got.Git.Url)
			assert.Equal(t, tc.want.Git.Username, got.Git.Username)
			assert.Equal(t, tc.want.Git.Directory, got.Git.Directory)
			assert.Equal(t, tc.want.Git.Email, got.Git.Email)
			assert.Equal(t, tc.want.Grafana.Url, got.Grafana.Url)
			assert.Equal(t, tc.want.Grafana.ApiKey, got.Grafana.ApiKey)
		})
	}
}
