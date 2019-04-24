package telegrambot

import (
	"bytes"
	"car-quiz-v3/entities"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// Converts arbitrary data type to target struct
func toStruct(data interface{}, target interface{}) error {
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(payload, &target)
}

// Marshals payload to JSON and sends it to Telegram API server. Unmarshals response to target data struct.
// todo add proxy
func makePostRequest(timeout int, url string, payload interface{}, target interface{}) error {
	// marshal payload to JSON
	jsonPayload, marshalErr := json.Marshal(payload)
	if marshalErr != nil {
		fmt.Printf("failed to marshal POST JSON")
		return marshalErr
	}

	// make HTTP request
	var httpCli = &http.Client{Timeout: time.Duration(timeout) * time.Second}
	r, errMakePost := httpCli.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	defer func() {
		errRespBodyClose := r.Body.Close()
		if errRespBodyClose != nil {
			panic(errRespBodyClose)
		}
	}()
	if errMakePost != nil {
		return errMakePost
	}
	if r.StatusCode != 200 {
		return errors.New("POST status code != 200")
	}

	// Decode response to ApiResponse struct and verify it
	var apiResponse entities.ApiResponse
	if errDecode := json.NewDecoder(r.Body).Decode(&apiResponse); errDecode != nil {
		return errDecode
	}
	if !apiResponse.OK {
		return errors.New("ApiResponse.OK == false")
	}

	// unmarshal response to target
	if target != nil { // target may be nil if bool is expected as a result
		if errConv := toStruct(apiResponse.Result, &target); errConv != nil {
			return errConv
		}
	}

	// OTHER POSSIBLE SOLUTIONS:
	// 0. errMapstruct := mapstructure.Decode(apiResponse.Result, target)
	// POSSIBLE PROBLEM
	// https://stackoverflow.com/questions/51813411/golang-mapstructure-not-working-as-expected
	// 1. resultMsg, ok := targetResponse.Result.(*Message) // doesn't work because map[string]interface{} != Message
	// 2. toStruct(apiResponse.Result, target) // works, but that's a HACK
	// 3. switch t := apiResponse.Result.(type) {  // doesn't work
	//			case entities.Message:
	//				target = t
	//			case *entities.Message:
	//				target = t
	//			case map[string]interface{}:
	//				// that's it :(
	//			}
	// 4. Multiple ApiResponse structs for different response options  // will probably work. ugly.
	// 5. func (ar *ApiResponse) UnmarshalJSON(b []byte) error { ??

	return nil
}

// for getUpdates
func getJson(url string, target interface{}) error {
	// long polling; no timeout here because it is set in url query string
	var httpCli = &http.Client{}

	r, err := httpCli.Get(url)
	defer func() {
		errRespBodyClose := r.Body.Close()
		if errRespBodyClose != nil {
			panic(errRespBodyClose)
		}
	}()
	if err != nil {
		return err
	}

	if r.StatusCode != 200 {
		return errors.New("> wrong getUpdates response status code")
	}
	return json.NewDecoder(r.Body).Decode(target)
}
