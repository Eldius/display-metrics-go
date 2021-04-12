package config

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

const (
	EndpointEmptyError = "Endpoint is empty"
)

func GetMetricsEndpoint() (string, error) {
	endpoint := viper.GetString("client.metrics.endpoint")
	//log.Printf("Endpoint: '%s'\n", endpoint)
	if endpoint == "" {
		log.Printf(EndpointEmptyError)
		return "", fmt.Errorf("Endpoint not configured")
	}
	return endpoint, nil
}

func GetMetricsRefreshInterval() time.Duration {
	return viper.GetDuration("client.metrics.refres_interval")
}
