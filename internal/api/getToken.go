package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/Resolution-hash/s1-report/config"
	httpclient "github.com/Resolution-hash/s1-report/internal/httpClient"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func getToken() (string, error) {
	url := "https://s1.detmir.ru"
	loginURL := "/v1/auth/login?language=ru"
	headers := map[string]string{
		"Accept":     "application/json, text/plain, */*",
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36",
	}

	userData, err := config.LoadUserConfig(false)
	if err != nil {
		return "", err
	}

	loginData := &LoginRequest{
		Username: userData.Login,
		Password: userData.Password,
	}

	log.Println(userData)

	client := httpclient.NewClient(url)

	statusCode, body, err := client.Post(loginURL, loginData, headers)
	if err != nil {
		return "", err
	}

	if statusCode == http.StatusOK {
		log.Println(statusCode)
		authKey, err := getAuthKey(body)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(authKey)
		return authKey, nil
	}

	return "", nil
}

func getAuthKey(body []byte) (string, error) {
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	if dataMap, ok := data["data"].(map[string]interface{}); ok {
		if authKey, exists := dataMap["auth_key"].(string); exists {
			return authKey, nil
		}
	}

	return "", fmt.Errorf("auth_key not found")
}
