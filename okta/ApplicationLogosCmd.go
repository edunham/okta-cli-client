package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var ApplicationLogosCmd = &cobra.Command{
	Use:  "applicationLogos",
	Long: "Manage ApplicationLogosAPI",
}

func init() {
	rootCmd.AddCommand(ApplicationLogosCmd)
}

var (
	UploadApplicationLogoappId string

	UploadApplicationLogodata string

	UploadApplicationLogoBackup bool
)

func NewUploadApplicationLogoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "uploadApplicationLogo",
		Long: "Upload an application Logo",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationLogosAPI.UploadApplicationLogo(apiClient.GetConfig().Context, UploadApplicationLogoappId)

			if UploadApplicationLogodata != "" {
				req = req.Data(UploadApplicationLogodata)
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
			if UploadApplicationLogoBackup {

				idParam := UploadApplicationLogoappId
				err := utils.BackupObject(d, "ApplicationLogos", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UploadApplicationLogoappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().StringVarP(&UploadApplicationLogodata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&UploadApplicationLogoBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UploadApplicationLogoCmd := NewUploadApplicationLogoCmd()
	ApplicationLogosCmd.AddCommand(UploadApplicationLogoCmd)
}
