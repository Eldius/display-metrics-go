package metrics

import (
	"encoding/json"
	"log"
	"net/http"
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
	c := http.DefaultClient
	res, err := c.Get("http://192.168.100.195/dashboard/summary")
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
