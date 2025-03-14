package okta

import (
	"fmt"
	"os"

	"github.com/okta/okta-cli-client/iostream"
	"github.com/okta/okta-cli-client/sdk"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:  "okta-cli-client",
	Long: "A command line tool for management API\n\nhttps://github.com/okta/okta-cli-client",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		prepareInteractivity(cmd)
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var apiClient *sdk.APIClient

func init() {
	configuration, err := sdk.NewConfiguration(sdk.WithCache(false))
	if err != nil {
		fmt.Printf("Create new config should not be error %v", err)
	}
	configuration.Debug = false

	apiClient = sdk.NewAPIClient(configuration)
}

func canPrompt(cmd *cobra.Command) bool {
	res := iostream.IsInputTerminal() && iostream.IsOutputTerminal()
	return res
}

func prepareInteractivity(cmd *cobra.Command) {
	if canPrompt(cmd) || !iostream.IsInputTerminal() {
		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
			_ = cmd.Flags().SetAnnotation(flag.Name, cobra.BashCompOneRequiredFlag, []string{"false"})
		})
	}
}
