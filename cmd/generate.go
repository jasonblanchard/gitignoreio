package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Print .gitignore file to std out",
	Long: `Print .gitignore file to std out based on arguments passed in.

Arguments must be valid options on gitignore.io

Example:

	gitignoreio generate osx go

Pipe it to a file:

	gitignoreio generate osx go > .gitignore
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		options := argsToURLPath(args)
		resp, err := http.Get(fmt.Sprintf("https://www.toptal.com/developers/gitignore/api/%s", options))
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		fmt.Fprintf(cmd.OutOrStdout(), string(body))

		if resp.StatusCode >= 400 {
			color.New(color.FgRed).Fprintln(cmd.ErrOrStderr(), "\nSomething went wrong. You probably selected an option unsupported by gitignore.io. Your .gitignore file might be fine, check the comments in the file for defaults\n")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func argsToURLPath(args []string) string {
	var buf strings.Builder
	for i, a := range args {
		buf.WriteString(a)
		if i < len(args)-1 {
			buf.WriteString(",")
		}
	}
	return buf.String()
}
