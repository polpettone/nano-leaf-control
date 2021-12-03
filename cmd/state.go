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

	if args[0] == "on" || args[0] == "off" {
		return out.SetState(args[0])
	}

	if args[0] == "brightness" {

		if len(args) > 1 {

			value, err := strconv.ParseInt(args[1], 10, 64)

			if err != nil {
				return "brightness needs numeric value", nil
			}

			return out.SetBrightness(value, 0)

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
