package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

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

	return setState(args[0])
}

func init() {
	c := StateCmd()
	rootCmd.AddCommand(c)
}

func getState() (string, error) {

	url := GetURL()

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

	url := GetURL()

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
