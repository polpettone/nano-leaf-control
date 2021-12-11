package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/polpettone/nano-leaf-control/cmd/config"
	"github.com/spf13/cobra"
)

func InfoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			stdout, err := handleInfoCommand()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Fprintf(cmd.OutOrStdout(), stdout)
		},
	}
}

func handleInfoCommand() (string, error) {
	nanoLeafID := int64(2)
	return getInfo(nanoLeafID)
}

func init() {
	i := InfoCmd()
	rootCmd.AddCommand(i)
}

func getInfo(nanoLeafID int64) (string, error) {
	url := config.GetURL(nanoLeafID)
	req, err := http.NewRequest("GET", url, nil)

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
