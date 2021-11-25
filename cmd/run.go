package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
			fmt.Fprintf(cmd.OutOrStdout(), stdout)
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


func GetURL() string {
	nanoLeafURL := viper.GetString("nano_leaf_url")
	token := viper.GetString("token")
	apiPath := "api/v1"
	url := fmt.Sprintf("http://%s/%s/%s", nanoLeafURL, apiPath, token)
	return url
}





