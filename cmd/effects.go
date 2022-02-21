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

func EffectsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "effects",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			stdout, err := handleEffectsCommand(cmd, args)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Fprint(cmd.OutOrStdout(), stdout)
		},
	}
}

func handleEffectsCommand(command *cobra.Command, args []string) (string, error) {

	nanoLeafID := int64(2)
	id, err := command.Flags().GetString("id")
	if err != nil {
		return "", err
	}
	nanoLeafID, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		return "", err
	}

	nextFlag, err := command.Flags().GetBool("next")
	if err != nil {
		return "", err
	}

	previousFlag, err := command.Flags().GetBool("previous")
	if err != nil {
		return "", err
	}

	effects, err := getEffects(nanoLeafID)
	if err != nil {
		return "", err
	}

	if nextFlag {
		selectedFlagIndex := 0
		nextEffectIndex := 0
		for i, e := range effects.EffectsList {
			if e == effects.Selected {
				selectedFlagIndex = i
			}
		}
		if len(effects.EffectsList)-1 > selectedFlagIndex {
			nextEffectIndex = selectedFlagIndex + 1
		}
		return setEffect(nanoLeafID, effects.EffectsList[nextEffectIndex])
	}

	if previousFlag {
		selectedFlagIndex := 0
		nextEffectIndex := 0
		for i, e := range effects.EffectsList {
			if e == effects.Selected {
				selectedFlagIndex = i
			}
		}
		if selectedFlagIndex != 0 {
			nextEffectIndex = selectedFlagIndex - 1
		} else {
			nextEffectIndex = len(effects.EffectsList) - 1
		}
		return setEffect(nanoLeafID, effects.EffectsList[nextEffectIndex])
	}

	if len(args) < 1 {
		prettyPrint, err := json.MarshalIndent(effects, "", "    ")
		if err != nil {
			return "", err
		}
		return string(prettyPrint), nil
	}

	return setEffect(nanoLeafID, args[0])
}

func init() {
	c := EffectsCmd()
	rootCmd.AddCommand(c)

	c.Flags().StringP(
		"id",
		"",
		"",
		"use the nano leaf with the given id")

	c.Flags().BoolP(
		"next",
		"n",
		false,
		"select next effect",
	)

	c.Flags().BoolP(
		"previous",
		"p",
		false,
		"select previous effect",
	)

}

func getEffects(nanoLeafID int64) (*models.Effects, error) {

	url := config.GetURL(nanoLeafID)

	req, err := http.NewRequest("GET", url+"/effects", nil)

	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: time.Second * 1}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	effects, err := models.ConvertJsonToEffects(body)

	if err != nil {
		return nil, err
	}

	return effects, nil
}

func setEffect(nanoLeafID int64, effect string) (string, error) {

	values := map[string]string{"select": effect}

	jsonValue, _ := json.Marshal(values)

	url := config.GetURL(nanoLeafID)

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
