package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
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
	haveFun()
	return "", nil
}

func init() {
	runCmd := RunCmd()
	rootCmd.AddCommand(runCmd)
}

func haveFun() {
	url := viper.GetString("nano_leaf_url")
	token := viper.GetString("token")


	fmt.Printf("%s %s", url, token )

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Printf("%s", req)
	}

	fmt.Printf("%s", req)

}

