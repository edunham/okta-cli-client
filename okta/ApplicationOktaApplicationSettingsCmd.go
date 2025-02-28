package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var ApplicationOktaApplicationSettingsCmd = &cobra.Command{
	Use:  "applicationOktaApplicationSettings",
	Long: "Manage ApplicationOktaApplicationSettingsAPI",
}

func init() {
	rootCmd.AddCommand(ApplicationOktaApplicationSettingsCmd)
}

var (
	GetFirstPartyAppSettingsappName string

	GetFirstPartyAppSettingsBackup bool
)

func NewGetFirstPartyAppSettingsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getFirstPartyAppSettings",
		Long: "Retrieve the Okta app settings",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationOktaApplicationSettingsAPI.GetFirstPartyAppSettings(apiClient.GetConfig().Context, GetFirstPartyAppSettingsappName)

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
			if GetFirstPartyAppSettingsBackup {

				idParam := GetFirstPartyAppSettingsappName
				err := utils.BackupObject(d, "ApplicationOktaApplicationSettings", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetFirstPartyAppSettingsappName, "appName", "", "", "")
	cmd.MarkFlagRequired("appName")

	cmd.Flags().BoolVarP(&GetFirstPartyAppSettingsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetFirstPartyAppSettingsCmd := NewGetFirstPartyAppSettingsCmd()
	ApplicationOktaApplicationSettingsCmd.AddCommand(GetFirstPartyAppSettingsCmd)
}

var (
	ReplaceFirstPartyAppSettingsappName string

	ReplaceFirstPartyAppSettingsdata string

	ReplaceFirstPartyAppSettingsBackup bool
)

func NewReplaceFirstPartyAppSettingsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replaceFirstPartyAppSettings",
		Long: "Replace the Okta app settings",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationOktaApplicationSettingsAPI.ReplaceFirstPartyAppSettings(apiClient.GetConfig().Context, ReplaceFirstPartyAppSettingsappName)

			if ReplaceFirstPartyAppSettingsdata != "" {
				req = req.Data(ReplaceFirstPartyAppSettingsdata)
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
			if ReplaceFirstPartyAppSettingsBackup {

				idParam := ReplaceFirstPartyAppSettingsappName
				err := utils.BackupObject(d, "ApplicationOktaApplicationSettings", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceFirstPartyAppSettingsappName, "appName", "", "", "")
	cmd.MarkFlagRequired("appName")

	cmd.Flags().StringVarP(&ReplaceFirstPartyAppSettingsdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceFirstPartyAppSettingsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceFirstPartyAppSettingsCmd := NewReplaceFirstPartyAppSettingsCmd()
	ApplicationOktaApplicationSettingsCmd.AddCommand(ReplaceFirstPartyAppSettingsCmd)
}
