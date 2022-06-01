package fatsecret_go

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// clientID, clientSecret
// return token

var ClientID string
var ClientSecret string
var AuthToken string
var Regional string
var Languages string
var Scope string

func InitFatSecret(clientID, clientSecret, Region, Language, scope string) {
	var err error

	ClientID = clientID
	ClientSecret = clientSecret
	Regional = Region
	Languages = Language
	Scope = scope
	AuthToken, err = GetFatSecretAuthorization(clientID, clientSecret, scope)
	if err != nil {
		fmt.Println(err)
	}
	go refreshToken()
}

func GetFatSecretAuthorization(clientID, clientSecret, scope string) (string, error) {
	endPoint := "https://oauth.fatsecret.com/connect/token"
	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("grant_type", "client_credentials")
	data.Set("scope", scope)
	data.Set("json", "true")
	req, err := http.NewRequest("POST", endPoint, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	resultMap := map[string]interface{}{}
	err = json.Unmarshal(body, &resultMap)
	if err != nil {
		return "", err
	}
	return resultMap["access_token"].(string), err
}

func refreshToken() {
	ticker := time.NewTicker(time.Hour * 23)
	failedCount := 0
	for {
		select {
		case <-ticker.C:
			var err error
			AuthToken, err = GetFatSecretAuthorization(ClientID, ClientSecret, Scope)
			if err != nil {
				failedCount++
				continue
			}
			if failedCount > 5 {
				failedCount = 0
			}
		}
	}
}

func FoodGetV2(foodId string) (map[string]interface{}, error) {
	resultMap := map[string]interface{}{}
	endPoint := fmt.Sprintf("https://platform.fatsecret.com/rest/server.api?method=food.get.v2&food_id=%s&format=json&region=%s",
		foodId, Regional)
	data := url.Values{}
	req, err := http.NewRequest("POST", endPoint, strings.NewReader(data.Encode()))
	if err != nil {
		return resultMap, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", AuthToken))
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return resultMap, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return resultMap, err
	}
	err = json.Unmarshal(body, &resultMap)
	return resultMap, err
}

func FoodCategoriesGet() (map[string]interface{}, error) {
	resultMap := map[string]interface{}{}
	endPoint := fmt.Sprintf("https://platform.fatsecret.com/rest/server.api?method=food_categories.get&format=json&region=%s", Regional)
	data := url.Values{}
	req, err := http.NewRequest("POST", endPoint, strings.NewReader(data.Encode()))
	if err != nil {
		return resultMap, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", AuthToken))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return resultMap, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return resultMap, err
	}
	err = json.Unmarshal(body, &resultMap)
	return resultMap, err
}

func FoodSubCategoriesGet(categoryId string) (map[string]interface{}, error) {
	resultMap := map[string]interface{}{}
	method := "food_sub_categories.get"
	endPoint := fmt.Sprintf("https://platform.fatsecret.com/rest/server.api?method=%s&food_category_id=%s&format=json&region=%s", method, categoryId, Regional)
	data := url.Values{}
	req, err := http.NewRequest("POST", endPoint, strings.NewReader(data.Encode()))
	if err != nil {
		return resultMap, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", AuthToken))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return resultMap, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return resultMap, err
	}
	err = json.Unmarshal(body, &resultMap)
	return resultMap, err
}

func FoodsSearch(keywords string, page, offset int) (map[string]interface{}, error) {
	resultMap := map[string]interface{}{}
	method := "foods.search"
	endPoint := fmt.Sprintf("https://platform.fatsecret.com/rest/server.api?method=%s&search_expression=%s&page_number=%d&max_results=%d&format=json&region=KR", method, keywords, page, offset)
	data := url.Values{}
	req, err := http.NewRequest("POST", endPoint, strings.NewReader(data.Encode()))
	if err != nil {
		return resultMap, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", AuthToken))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return resultMap, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return resultMap, err
	}
	err = json.Unmarshal(body, &resultMap)
	return resultMap, err
}
