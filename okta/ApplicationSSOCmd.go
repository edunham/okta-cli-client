package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var ApplicationSSOCmd = &cobra.Command{
	Use:  "applicationSSO",
	Long: "Manage ApplicationSSOAPI",
}

func init() {
	rootCmd.AddCommand(ApplicationSSOCmd)
}

var (
	PreviewSAMLmetadataForApplicationappId string

	PreviewSAMLmetadataForApplicationBackup bool
)

func NewPreviewSAMLmetadataForApplicationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "previewSAMLmetadataForApplication",
		Long: "Preview the application SAML metadata",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationSSOAPI.PreviewSAMLmetadataForApplication(apiClient.GetConfig().Context, PreviewSAMLmetadataForApplicationappId)

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
			if PreviewSAMLmetadataForApplicationBackup {

				idParam := PreviewSAMLmetadataForApplicationappId
				err := utils.BackupObject(d, "ApplicationSSO", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&PreviewSAMLmetadataForApplicationappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().BoolVarP(&PreviewSAMLmetadataForApplicationBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	PreviewSAMLmetadataForApplicationCmd := NewPreviewSAMLmetadataForApplicationCmd()
	ApplicationSSOCmd.AddCommand(PreviewSAMLmetadataForApplicationCmd)
}
