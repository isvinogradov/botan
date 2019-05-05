package botan

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/isvinogradov/botan/entities"
)

// Converts arbitrary data type to target struct
func toStruct(data interface{}, target interface{}) error {
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(payload, &target)
}

// close HTTP response body
func closeBody(r *http.Response) {
	errRespBodyClose := r.Body.Close()
	if errRespBodyClose != nil {
		panic(errRespBodyClose)
	}
}

type requestGate struct {
	postTimeoutSeconds     int
	getTimeoutSeconds      int
	socks5ConnectionString string
	postClient             *http.Client
	getClient              *http.Client
}

func (rg *requestGate) checkAndInit() error {
	var transport http.Transport
	if rg.socks5ConnectionString == "" {
		transport = http.Transport{Proxy: nil} // no proxy
	} else {
		proxyUrl, err := url.Parse(rg.socks5ConnectionString)
		if err != nil {
			return errors.New("failed to parse SOCKS5 connection string")
		}
		transport = http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	}

	rg.postClient = &http.Client{Transport: &transport, Timeout: time.Duration(rg.postTimeoutSeconds) * time.Second}
	rg.getClient = &http.Client{Transport: &transport, Timeout: time.Duration(rg.getTimeoutSeconds) * time.Second}

	return nil
}

// Marshals payload to JSON and sends it to Telegram API server. Unmarshals response to target data struct.
func (rg *requestGate) makePostRequest(url string, payload interface{}, target interface{}) error {
	// marshal payload to JSON
	jsonPayload, marshalErr := json.Marshal(payload)
	if marshalErr != nil {
		fmt.Printf("failed to marshal POST JSON")
		return marshalErr
	}

	// make HTTP request
	r, errMakePost := rg.postClient.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	defer closeBody(r)
	if errMakePost != nil {
		return errMakePost
	}
	if r.StatusCode != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bodyBytes))

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
func (rg *requestGate) makeGetRequest(url string, target interface{}) error {
	r, err := rg.getClient.Get(url) // long polling
	defer closeBody(r)
	if err != nil {
		return err
	}

	if r.StatusCode != http.StatusOK {
		return errors.New("> wrong getUpdates response status code")
	}
	return json.NewDecoder(r.Body).Decode(target)
}
