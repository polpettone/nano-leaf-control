package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func RunCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			stdout, err := handleRunCommand()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Fprint(cmd.OutOrStdout(), stdout)
		},
	}
}

func handleRunCommand() (string, error) {
	return "", nil
}

func init() {
	runCmd := RunCmd()
	rootCmd.AddCommand(runCmd)
}
