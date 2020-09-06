package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List valid options for gitignore.io",
	Long: `List valid options for gitignore.io

Example:

	gitignorio list

Search through the list:

	gitignorio list | grep node
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := http.Get("https://www.toptal.com/developers/gitignore/api/list")
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil
		}

		fmt.Fprint(cmd.OutOrStdout(), string(body))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
