package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var ThreatInsightCmd = &cobra.Command{
	Use:  "threatInsight",
	Long: "Manage ThreatInsightAPI",
}

func init() {
	rootCmd.AddCommand(ThreatInsightCmd)
}

var (
	UpdateConfigurationdata string

	UpdateConfigurationBackup bool
)

func NewUpdateConfigurationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "updateConfiguration",
		Long: "Update the ThreatInsight Configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ThreatInsightAPI.UpdateConfiguration(apiClient.GetConfig().Context)

			if UpdateConfigurationdata != "" {
				req = req.Data(UpdateConfigurationdata)
			}

			resp, err := req.Execute()
			if err != nil {
				if resp != nil && resp.Body != nil {
					d, err := io.ReadAll(resp.Body)
					if err == nil {
						utils.PrettyPrintByte(d)
					}
				}
				return err
			}
			d, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			if UpdateConfigurationBackup {

				err := utils.BackupObject(d, "ThreatInsight", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UpdateConfigurationdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&UpdateConfigurationBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UpdateConfigurationCmd := NewUpdateConfigurationCmd()
	ThreatInsightCmd.AddCommand(UpdateConfigurationCmd)
}

var GetCurrentConfigurationBackup bool

func NewGetCurrentConfigurationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getCurrentConfiguration",
		Long: "Retrieve the ThreatInsight Configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ThreatInsightAPI.GetCurrentConfiguration(apiClient.GetConfig().Context)

			resp, err := req.Execute()
			if err != nil {
				if resp != nil && resp.Body != nil {
					d, err := io.ReadAll(resp.Body)
					if err == nil {
						utils.PrettyPrintByte(d)
					}
				}
				return err
			}
			d, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			if GetCurrentConfigurationBackup {

				err := utils.BackupObject(d, "ThreatInsight", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&GetCurrentConfigurationBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetCurrentConfigurationCmd := NewGetCurrentConfigurationCmd()
	ThreatInsightCmd.AddCommand(GetCurrentConfigurationCmd)
}
