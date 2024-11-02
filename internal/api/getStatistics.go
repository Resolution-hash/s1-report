package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	httpclient "github.com/Resolution-hash/s1-report/internal/httpClient"
)

const TICKET_COUNT_URL = "/v1/list-filter/itsm_task?condition=(stateNOT%20IN7%4010%4012%5Eassignment_groupDYNAMIC159679003220074662%5Esys_db_table_id\u0021%3D156950639316968592)%5EORDERBYDESCsys_created_at"
const BASE_URL = "https://s1.detmir.ru"

func GetStatistics() (float64, error) {

	authKey, err := getToken()
	if err != nil {
		return 0, err
	}

	AuthData := "Bearer " + authKey

	headers := map[string]string{
		"Authorization": AuthData,
		"Accept":        "application/json, text/plain, */*",
	}

	log.Println("headers:", headers)

	client := httpclient.NewClient(BASE_URL)

	statusCode, body, err := client.Get(TICKET_COUNT_URL, headers)
	if err != nil {
		return 0, err
	}

	if statusCode == http.StatusOK {
		totalTicket, err := parseStatistics(body)
		if err != nil {
			return 0, err
		}
		return totalTicket, nil
	}

	// fmt.Println("Response status:", statusCode)
	// fmt.Println("Response body:", string(body))
	return 0, nil
}

func parseStatistics(body []byte) (float64, error) {
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}

	if data, ok := result["data"].(map[string]interface{}); ok {
		if pagination, ok := data["pagination"].(map[string]interface{}); ok {
			fmt.Println(pagination)
			if total, ok := pagination["total"].(float64); ok {
				fmt.Println("Total:", total)
				return total, nil
			}
		}
	}

	return 0, nil
}
