package out

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/polpettone/nano-leaf-control/cmd/config"
	"github.com/polpettone/nano-leaf-control/cmd/models"
)

func GetState() (string, error) {

	url := config.GetURL()

	req, err := http.NewRequest("GET", url+"/state", nil)

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

func SetState(state string) (string, error) {

	var jsonValue []byte

	if state == "on" {
		jsonValue = stateBody(true)
	} else {
		jsonValue = stateBody(false)
	}

	fmt.Printf("set state %s", jsonValue)

	return makeStateAPICall(jsonValue)
}

func SetBrightness(value int64, duration int64) (string, error) {

	jsonValue := brightnessBody(value, duration)

	fmt.Printf("set brightness %s", jsonValue)

	return makeStateAPICall(jsonValue)
}

func makeStateAPICall(stateJsonBody []byte) (string, error) {
	url := config.GetURL()
	req, err := http.NewRequest("PUT", url+"/state", bytes.NewBuffer(stateJsonBody))
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
