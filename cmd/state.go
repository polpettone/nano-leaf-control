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
		state, err := out.GetState()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s", state), nil
	}

	switch args[0] {
	case "on":
		return out.SetState(args[0])
	case "off":
		return out.SetState(args[0])
	case "brightness":
		return handleStateBrightnessCommand(args)
	case "hue":
		return "hue setting comming soon", nil
	case "saturation":
		return "saturation setting comming soon", nil
	case "temperature":
		return "color temeratute setting comming soon", nil
	default:
		return "no valid command", nil
	}

	return "no valid command", nil
}

func handleStateBrightnessCommand(args []string) (string, error) {

	switch len(args) {
	case 2:
		value, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			return "brightness needs numeric value", nil
		}
		return out.SetBrightness(value, 0)
	case 3:
		brightnessValue, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			return "brightness needs numeric value", nil
		}
		durationValue, err := strconv.ParseInt(args[2], 10, 64)
		if err != nil {
			return "brightness durations needs numeric value", nil
		}
		return out.SetBrightness(brightnessValue, durationValue)
	default:
		return "brightness needs one or two numeric values", nil

	}
}

func init() {
	c := StateCmd()
	rootCmd.AddCommand(c)
}
