package cmd

import (
	"fmt"
	"strconv"

	"github.com/polpettone/nano-leaf-control/cmd/out"
	"github.com/spf13/cobra"
)

func StateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "state",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			stdout, err := handleStateCommand(cmd, args)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Fprint(cmd.OutOrStdout(), stdout)
		},
	}
}

func handleStateCommand(command *cobra.Command, args []string) (string, error) {

	nanoLeafID := int64(2)
	id, err := command.Flags().GetString("id")
	if err != nil {
		return "", err
	}
	nanoLeafID, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		return "", err
	}

	if len(args) < 1 {
		state, err := out.GetState(nanoLeafID)
		if err != nil {
			return "", err
		}
		return state, nil
	}

	switch args[0] {
	case "on":
		return out.SetState(nanoLeafID, args[0])
	case "off":
		return out.SetState(nanoLeafID, args[0])
	case "brightness":
		return handleStateBrightnessCommand(nanoLeafID, args)
	case "increaseBrightness":
		return handleIncreaseBrightnessCommand(nanoLeafID, args)
	case "decreaseBrightness":
		return handleDecreaseBrightnessCommand(nanoLeafID, args)
	case "hue":
		return "hue setting comming soon", nil
	case "saturation":
		return "saturation setting comming soon", nil
	case "temperature":
		return "color temeratute setting comming soon", nil
	default:
		return "no valid command", nil
	}
}

func handleIncreaseBrightnessCommand(nanoLeafID int64, args []string) (string, error) {
	currentBrightness, err := out.GetBrightness(nanoLeafID)
	if err != nil {
		return "", err
	}
	fmt.Printf("currentBrightness: %v \n", currentBrightness)
	newValue := currentBrightness.Value + 10
	return out.SetBrightness(nanoLeafID, newValue, 0)
}

func handleDecreaseBrightnessCommand(nanoLeafID int64, args []string) (string, error) {
	currentBrightness, err := out.GetBrightness(nanoLeafID)
	if err != nil {
		return "", err
	}
	fmt.Printf("currentBrightness: %v \n", currentBrightness)
	newValue := currentBrightness.Value - 10
	return out.SetBrightness(nanoLeafID, newValue, 0)
}

func handleStateBrightnessCommand(nanoLeafID int64, args []string) (string, error) {

	switch len(args) {
	case 2:
		value, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			return "brightness needs numeric value", nil
		}
		return out.SetBrightness(nanoLeafID, value, 0)
	case 3:
		brightnessValue, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			return "brightness needs numeric value", nil
		}
		durationValue, err := strconv.ParseInt(args[2], 10, 64)
		if err != nil {
			return "brightness durations needs numeric value", nil
		}
		return out.SetBrightness(nanoLeafID, brightnessValue, durationValue)
	default:
		return "brightness needs one or two numeric values", nil

	}
}

func init() {
	c := StateCmd()
	rootCmd.AddCommand(c)

	c.Flags().StringP(
		"id",
		"",
		"",
		"use the nano leaf with the given id")

}
