package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var SystemLogCmd = &cobra.Command{
	Use:  "systemLog",
	Long: "Manage SystemLogAPI",
}

func init() {
	rootCmd.AddCommand(SystemLogCmd)
}

var ListLogEventsBackup bool

func NewListLogEventsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listLogEvents",
		Long: "List all System Log Events",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.SystemLogAPI.ListLogEvents(apiClient.GetConfig().Context)

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
			if ListLogEventsBackup {

				err := utils.BackupObject(d, "SystemLog", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListLogEventsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListLogEventsCmd := NewListLogEventsCmd()
	SystemLogCmd.AddCommand(ListLogEventsCmd)
}
