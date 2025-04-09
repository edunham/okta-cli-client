package okta

import (
	"fmt"
	"os"

	"github.com/okta/okta-cli-client/iostream"
	"github.com/okta/okta-cli-client/sdk"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v3"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:  "okta-cli-client",
	Long: "A command line tool for management API\n\nhttps://github.com/okta/okta-cli-client",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var apiClient *sdk.APIClient

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.okta/okta.yaml)")

	originalPreRunE := rootCmd.PersistentPreRunE

	// can't just read a flag and then use it immediately, have to read the flag then use it later
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		var configSetters []sdk.ConfigSetter
		configSetters = append(configSetters, sdk.WithCache(false))

		//this is probably not the right way to pass config settings to the sdk!
		if cfgFile != "" {
			fmt.Printf("Using custom config file: %s\n", cfgFile)
			yamlConfig, err := os.ReadFile(cfgFile)
			if err != nil {
				fmt.Printf("Error reading config file %s: %v\n", cfgFile, err)
			} else {
				var config struct {
					Okta struct {
						Client struct {
							OrgUrl            string   `yaml:"orgUrl"`
							Token             string   `yaml:"token"`
							AuthorizationMode string   `yaml:"authorizationMode"`
							ClientId          string   `yaml:"clientId"`
							PrivateKey        string   `yaml:"privateKey"`
							PrivateKeyId      string   `yaml:"privateKeyId"`
							Scopes            []string `yaml:"scopes"`
						} `yaml:"client"`
					} `yaml:"okta"`
				}

				if err := yaml.Unmarshal(yamlConfig, &config); err != nil {
					fmt.Printf("Error parsing config file: %v\n", err)
				} else {
					if config.Okta.Client.OrgUrl != "" {
						configSetters = append(configSetters, sdk.WithOrgUrl(config.Okta.Client.OrgUrl))
					}

					if config.Okta.Client.AuthorizationMode != "" {
						configSetters = append(configSetters, sdk.WithAuthorizationMode(config.Okta.Client.AuthorizationMode))

						// API Token auth
						if config.Okta.Client.AuthorizationMode == "SSWS" && config.Okta.Client.Token != "" {
							configSetters = append(configSetters, sdk.WithToken(config.Okta.Client.Token))
						}

						// OAuth 2.0 auth
						if config.Okta.Client.AuthorizationMode == "PrivateKey" {
							if config.Okta.Client.ClientId != "" {
								configSetters = append(configSetters, sdk.WithClientId(config.Okta.Client.ClientId))
							}
							if config.Okta.Client.PrivateKey != "" {
								configSetters = append(configSetters, sdk.WithPrivateKey(config.Okta.Client.PrivateKey))
							}
							if config.Okta.Client.PrivateKeyId != "" {
								configSetters = append(configSetters, sdk.WithPrivateKeyId(config.Okta.Client.PrivateKeyId))
							}
							if len(config.Okta.Client.Scopes) > 0 {
								configSetters = append(configSetters, sdk.WithScopes(config.Okta.Client.Scopes))
							}
						}
					} else if config.Okta.Client.Token != "" {
						// Default to SSWS if authorizationMode not specified but token is present
						configSetters = append(configSetters, sdk.WithToken(config.Okta.Client.Token))
					}
				}
			}
		}

		configuration, err := sdk.NewConfiguration(configSetters...)
		if err != nil {
			return fmt.Errorf("error creating configuration: %v", err)
		}
		configuration.Debug = false

		apiClient = sdk.NewAPIClient(configuration)

		if originalPreRunE != nil {
			return originalPreRunE(cmd, args)
		}

		prepareInteractivity(cmd)
		return nil
	}
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
