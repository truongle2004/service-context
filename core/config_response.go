package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

type propertySource struct {
	Source map[string]interface{} `json:"source"`
}

type configResponse struct {
	PropertySources []propertySource `json:"propertySources"`
}

// LoadConfig loads Spring Cloud Config into a provided Viper instance
func LoadConfig(app, profile, configURL string, v *viper.Viper) error {
	url := fmt.Sprintf("%s/%s/%s", strings.TrimSuffix(configURL, "/"), app, profile)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch config: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("config server error %d: %s", resp.StatusCode, string(body))
	}

	var cfg configResponse
	if err := json.NewDecoder(resp.Body).Decode(&cfg); err != nil {
		return fmt.Errorf("decode error: %w", err)
	}

	for i := len(cfg.PropertySources) - 1; i >= 0; i-- {
		for k, val := range cfg.PropertySources[i].Source {
			v.Set(k, val)
		}
	}

	return nil
}
