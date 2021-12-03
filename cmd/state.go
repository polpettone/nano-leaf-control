package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/polpettone/nano-leaf-control/cmd/config"
	"github.com/polpettone/nano-leaf-control/cmd/models"
	"github.com/spf13/cobra"
)

func StateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "state",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			stdout, err := handleStateCommand(args)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Fprintf(cmd.OutOrStdout(), stdout)
		},
	}
}

func handleStateCommand(args []string) (string, error) {
	if len(args) < 1 {
		state, err := getState()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s", state), nil
	}

	if args[0] == "on" || args[0] == "off" {
		return setState(args[0])
	}

	if args[0] == "brightness" {

		if len(args) > 1 {

			value, err := strconv.ParseInt(args[1], 10, 64)

			if err != nil {
				return "brightness needs numeric value", nil
			}

			return setBrightness(value, 0)

		} else {
			return "brightness needs at least one value", nil
		}
	}

	if args[0] == "hue" {
		return "hue setting comming soon", nil
	}

	if args[0] == "saturation" {
		return "saturation setting comming soon", nil
	}

	if args[0] == "temperature" {
		return "color temeratute setting comming soon", nil
	}

	return "no valid command", nil
}

func init() {
	c := StateCmd()
	rootCmd.AddCommand(c)
}

func getState() (string, error) {

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

	if err != nil {
		return "", err
	}

	return string(body), nil
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

func setState(state string) (string, error) {

	var jsonValue []byte

	if state == "on" {
		jsonValue = stateBody(true)
	} else {
		jsonValue = stateBody(false)
	}

	url := config.GetURL()

	req, err := http.NewRequest("PUT", url+"/state", bytes.NewBuffer(jsonValue))

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

func setBrightness(value int64, duration int64) (string, error) {

	jsonValue := brightnessBody(value, duration)

	fmt.Printf("set brightness %s", jsonValue)

	url := config.GetURL()

	req, err := http.NewRequest("PUT", url+"/state", bytes.NewBuffer(jsonValue))

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
