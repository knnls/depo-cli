package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func (a *API) SubmitMetadata(metadata map[string]string) {
	var url = fmt.Sprintf("%s/device/add", GetUrl())
	requestBodyBytes, err := json.Marshal(metadata)
	if err != nil {
		log.Fatalf("Error encoding JSON: %v\n", err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		log.Fatalf("Error creating request: %v\n", err)
	}

	req.Header.Set("Content-Type", "application/json")
	tokenValue := viper.GetString("token")
	accessToken := viper.GetString("api_key")

	if viper.IsSet("token") {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tokenValue))
	} else {
		log.Fatal("authentication failed")
	}

	if viper.IsSet("api_key") {
		req.Header.Set("x-access-token", accessToken)
	} else {
		log.Fatal("link cli with new access token from the dashboard")
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v\n", err)
	}
	defer resp.Body.Close()

	// if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
	// 	log.Fatalf("Request failed with status code: %d\n", resp.StatusCode)
	// }
	var responseMap map[string]interface{}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&responseMap); err != nil {
		log.Fatalf("Error decoding response: %v\n", err)
	}
	fmt.Println(responseMap, "asdfasdf")
}

func (a *API) MakeAuthRequest(accessToken string) map[string]interface{} {
	var responseMap map[string]interface{}
	url := fmt.Sprintf("%s/cli/link", GetUrl())
	requestBody := map[string]string{
		"access_token": accessToken,
	}
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Printf("Error encoding JSON: %v\n", err)
		return responseMap
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return responseMap
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return responseMap
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Request failed with status code: %d\n", resp.StatusCode)
		return responseMap
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&responseMap); err != nil {
		fmt.Printf("Error decoding response: %v\n", err)
		return responseMap
	}

	return responseMap
}
