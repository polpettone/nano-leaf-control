package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

func EffectsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "effects",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			stdout, err := handleEffectsCommand(args)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Fprintf(cmd.OutOrStdout(), stdout)
		},
	}
}

func handleEffectsCommand(args []string) (string, error) {
	if len(args) < 1 {
		return getEffects()
	}

	return setEffect(args[0])
}

func init() {
	c := EffectsCmd()
	rootCmd.AddCommand(c)
}

func getEffects() (string, error) {

	url := GetURL()

	req, err := http.NewRequest("GET", url+"/effects", nil)

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

func setEffect(effect string) (string, error) {

	values := map[string]string{"select": effect}

	jsonValue, _ := json.Marshal(values)

	url := GetURL()

	req, err := http.NewRequest("PUT", url+"/effects", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Printf("%s", err)
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
