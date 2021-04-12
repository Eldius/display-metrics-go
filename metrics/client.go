package metrics

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Eldius/display-metrics-go/config"
)

type SummaryResponse struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	Data   SummaryData `json:"data"`
}

type SummaryData struct {
	Nodes      int     `json:"nodes"`
	CPU        float64 `json:"cpu"`
	Memory     float64 `json:"memory"`
	Pods       int     `json:"pods"`
	Containers int     `json:"containers"`
}

func GetSummary() (*SummaryResponse, error) {
	endpoint, err := config.GetMetricsEndpoint()
	//log.Printf("Endpoint: '%s'\n", endpoint)
	if endpoint == "" {
		log.Printf(err.Error())
		return nil, fmt.Errorf("Endpoint not configured")
	}
	c := http.DefaultClient
	res, err := c.Get(endpoint)
	if err != nil {
		log.Println("Failed to fetch metrics data")
		return nil, err
	}
	defer res.Body.Close()
	var response *SummaryResponse
	json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Println("Failed to parse metrics data")
		return nil, err
	}
	return response, nil
}
