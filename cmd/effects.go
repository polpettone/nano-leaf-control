package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"time"
)

func EffectsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "effects",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			stdout, err := handleEffectsCommand(args[0])
			if err != nil {
				fmt.Println(err)
			}
			fmt.Fprintf(cmd.OutOrStdout(), stdout)
		},
	}
}

func handleEffectsCommand(effect string) (string, error) {

	if effect == "" {
		return getEffects()
	}
	return setEffect(effect)
}

func init() {
	c := EffectsCmd()
	rootCmd.AddCommand(c)
}

func getEffects() (string, error) {

	url := GetURL()

	req, err := http.NewRequest("GET", url+"/effects", nil)

	if err != nil {
		fmt.Printf("%s", err)
	}

	client := &http.Client{Timeout: time.Second * 1}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("%s", err)
		return "", nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("%s", err)
		return "", nil
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
		fmt.Printf("%s", err)
		return "", nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("%s", err)
		return "", nil
	}

	return string(body), nil
}
