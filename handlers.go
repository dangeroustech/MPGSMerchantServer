package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//set your Gateway Information
var mid string = os.Getenv("GATEWAY_MERCHANT_ID")
var apiVer string = os.Getenv("GATEWAY_API_VERSION") //>39
var region string = os.Getenv("GATEWAY_REGION")      //ap, eu, na, in, mtf
var user string = "merchant." + mid
var pass string = os.Getenv("GATEWAY_API_PASSWORD")

//Index Function - Expects: GET Request - Returns: HTTP 200
func Index(w http.ResponseWriter, r *http.Request) {
	Logger("The Index Page Has Been Accessed")
}

//StartPayment Function - Expects: Empty POST Request - Returns: SessionID and Operation Result
func StartPayment(w http.ResponseWriter, r *http.Request) {

	if region == "MTF" {
		region = "test"
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	//build request JSON
	jsonData := map[string]interface{}{
		"correlationId": "001",
		"session": map[string]string{
			"authenticationLimit": "3",
		},
	}
	jsonValue, _ := json.Marshal(jsonData)
	//make request
	request, _ := http.NewRequest("POST", "https://"+region+"-gateway.mastercard.com/api/rest/version/"+apiVer+"/merchant/"+mid+"/session", bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(user, pass)
	client := &http.Client{}
	response, err := client.Do(request)

	//read the response
	if err != nil {
		//empty response
		Logger("An Error Occurred: Nil Response")

		response := map[string]string{
			"id":     "NONE",
			"result": "FAILURE",
		}
		json.NewEncoder(w).Encode(response)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		//response doesn't parse properly with json, use strings
		id := strings.SplitAfter(string(data), "SESSION")[1]
		id = strings.SplitAfter(id, ",")[0]
		id = "SESSION" + strings.Trim(id, "\",")
		Logger("Session " + id + " Obtained")

		//send response back to App
		response := map[string]string{
			"id":     string(id),
			"result": "SUCCESS",
		}
		json.NewEncoder(w).Encode(response)
	}
}

//FinishPayment Function - Expects: SessionID - Returns: SessionID and Operation Result
func FinishPayment(w http.ResponseWriter, r *http.Request) {

	if region == "MTF" {
		region = "test"
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	//read the SessionID coming from the App
	body, _ := ioutil.ReadAll(r.Body)
	id := strings.SplitAfter(string(body), "SESSION")[1]
	id = "SESSION" + id[:len(id)-3]

	//build request JSON
	jsonData := map[string]interface{}{
		"apiOperation": "PAY",
		"order": map[string]string{
			"currency": "GBP",
			"amount":   "10",
		},
		"session": map[string]string{
			"id": id,
		},
		"sourceOfFunds": map[string]string{
			"type": "CARD",
		},
	}
	jsonValue, _ := json.Marshal(jsonData)

	//make request
	request, _ := http.NewRequest("PUT", "https://"+region+"-gateway.mastercard.com/api/rest/version/"+apiVer+"/merchant/"+mid+"/order/"+id+"/transaction/"+id, bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(user, pass)
	client := &http.Client{}
	response, err := client.Do(request)
	data, _ := ioutil.ReadAll(response.Body)

	//read the response
	if err != nil {
		//empty response
		Logger("An Error Occurred: Nil Response")
		finalResponse := map[string]string{
			"id":     string(id),
			"result": "FAILURE"}
		json.NewEncoder(w).Encode(finalResponse)
	} else if strings.Contains(string(data), "SUCCESS") {
		//SUCCESS response
		Logger("PAY RESPONSE: " + string(data))
		finalResponse := map[string]string{
			"id":     string(id),
			"result": "SUCCESS"}
		json.NewEncoder(w).Encode(finalResponse)
	} else {
		//Error Response
		Logger("An Error Occurred: Bad Request\n" + string(data))
		finalResponse := map[string]string{
			"id":     string(id),
			"result": "FAILURE"}
		json.NewEncoder(w).Encode(finalResponse)
	}
}

//Auth Function: Handles incoming authentication from the App
func Auth(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey := r.Header.Get("APIKEY")
		if apikey != "TESTSDK" {
			http.Error(w, "Unauthorized.", http.StatusUnauthorized)
			return
		}
		fn(w, r)
	}
}
