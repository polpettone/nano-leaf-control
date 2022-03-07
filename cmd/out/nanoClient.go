package out

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/polpettone/nano-leaf-control/cmd/models"
)

func GetState(nanoLeafURL string) (string, error) {

	req, err := http.NewRequest("GET", nanoLeafURL+"/state", nil)

	if err != nil {
		return "", err
	}

	client := &http.Client{Timeout: time.Second * 1}
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func SetState(nanoLeafURL string, state string) (string, error) {

	var jsonValue []byte

	if state == "on" {
		jsonValue = stateBody(true)
	} else {
		jsonValue = stateBody(false)
	}

	fmt.Printf("set state %s", jsonValue)

	return changeStateAPICall(nanoLeafURL, jsonValue)
}

func SetBrightness(nanoLeafURL string, value, duration int64) (string, error) {

	jsonValue := brightnessBody(value, duration)

	fmt.Printf("set brightness %s", jsonValue)

	return changeStateAPICall(nanoLeafURL, jsonValue)
}

func GetBrightness(nanoLeafURL string) (*models.CurrentBrightness, error) {
	req, err := http.NewRequest("GET", nanoLeafURL+"/state/brightness", bytes.NewBuffer([]byte("")))
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: time.Second * 1}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var b models.CurrentBrightness
	err = json.Unmarshal(body, &b)

	if err != nil {
		return nil, err
	}

	return &b, nil
}

func changeStateAPICall(nanoLeafURL string, stateJsonBody []byte) (string, error) {
	req, err := http.NewRequest("PUT", nanoLeafURL+"/state", bytes.NewBuffer(stateJsonBody))
	if err != nil {
		return "", err
	}
	client := &http.Client{Timeout: time.Second * 1}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func brightnessBody(value int64, duration int64) []byte {
	b := models.Brightness{
		BrightnessValue: models.BrightnessValue{
			Value:    value,
			Duration: duration,
		},
	}
	jsonValue, _ := json.Marshal(b)
	return jsonValue
}

func stateBody(onValue bool) []byte {
	onState := models.OnState{
		On: models.OnStateValue{
			Value: onValue,
		},
	}
	jsonValue, _ := json.Marshal(onState)
	return jsonValue
}
