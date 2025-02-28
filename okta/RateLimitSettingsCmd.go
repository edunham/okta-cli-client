package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var RateLimitSettingsCmd = &cobra.Command{
	Use:  "rateLimitSettings",
	Long: "Manage RateLimitSettingsAPI",
}

func init() {
	rootCmd.AddCommand(RateLimitSettingsCmd)
}

var GetRateLimitSettingsAdminNotificationsBackup bool

func NewGetRateLimitSettingsAdminNotificationsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getAdminNotifications",
		Long: "Retrieve the Rate Limit Admin Notification Settings",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RateLimitSettingsAPI.GetRateLimitSettingsAdminNotifications(apiClient.GetConfig().Context)

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
			if GetRateLimitSettingsAdminNotificationsBackup {

				err := utils.BackupObject(d, "RateLimitSettings", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&GetRateLimitSettingsAdminNotificationsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetRateLimitSettingsAdminNotificationsCmd := NewGetRateLimitSettingsAdminNotificationsCmd()
	RateLimitSettingsCmd.AddCommand(GetRateLimitSettingsAdminNotificationsCmd)
}

var (
	ReplaceRateLimitSettingsAdminNotificationsdata string

	ReplaceRateLimitSettingsAdminNotificationsBackup bool
)

func NewReplaceRateLimitSettingsAdminNotificationsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replaceAdminNotifications",
		Long: "Replace the Rate Limit Admin Notification Settings",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RateLimitSettingsAPI.ReplaceRateLimitSettingsAdminNotifications(apiClient.GetConfig().Context)

			if ReplaceRateLimitSettingsAdminNotificationsdata != "" {
				req = req.Data(ReplaceRateLimitSettingsAdminNotificationsdata)
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
			if ReplaceRateLimitSettingsAdminNotificationsBackup {

				err := utils.BackupObject(d, "RateLimitSettings", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceRateLimitSettingsAdminNotificationsdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceRateLimitSettingsAdminNotificationsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceRateLimitSettingsAdminNotificationsCmd := NewReplaceRateLimitSettingsAdminNotificationsCmd()
	RateLimitSettingsCmd.AddCommand(ReplaceRateLimitSettingsAdminNotificationsCmd)
}

var GetRateLimitSettingsPerClientBackup bool

func NewGetRateLimitSettingsPerClientCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getPerClient",
		Long: "Retrieve the Per-Client Rate Limit Settings",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RateLimitSettingsAPI.GetRateLimitSettingsPerClient(apiClient.GetConfig().Context)

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
			if GetRateLimitSettingsPerClientBackup {

				err := utils.BackupObject(d, "RateLimitSettings", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&GetRateLimitSettingsPerClientBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetRateLimitSettingsPerClientCmd := NewGetRateLimitSettingsPerClientCmd()
	RateLimitSettingsCmd.AddCommand(GetRateLimitSettingsPerClientCmd)
}

var (
	ReplaceRateLimitSettingsPerClientdata string

	ReplaceRateLimitSettingsPerClientBackup bool
)

func NewReplaceRateLimitSettingsPerClientCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replacePerClient",
		Long: "Replace the Per-Client Rate Limit Settings",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RateLimitSettingsAPI.ReplaceRateLimitSettingsPerClient(apiClient.GetConfig().Context)

			if ReplaceRateLimitSettingsPerClientdata != "" {
				req = req.Data(ReplaceRateLimitSettingsPerClientdata)
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
			if ReplaceRateLimitSettingsPerClientBackup {

				err := utils.BackupObject(d, "RateLimitSettings", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceRateLimitSettingsPerClientdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceRateLimitSettingsPerClientBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceRateLimitSettingsPerClientCmd := NewReplaceRateLimitSettingsPerClientCmd()
	RateLimitSettingsCmd.AddCommand(ReplaceRateLimitSettingsPerClientCmd)
}

var GetRateLimitSettingsWarningThresholdBackup bool

func NewGetRateLimitSettingsWarningThresholdCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getWarningThreshold",
		Long: "Retrieve the Rate Limit Warning Threshold Percentage",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RateLimitSettingsAPI.GetRateLimitSettingsWarningThreshold(apiClient.GetConfig().Context)

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
			if GetRateLimitSettingsWarningThresholdBackup {

				err := utils.BackupObject(d, "RateLimitSettings", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&GetRateLimitSettingsWarningThresholdBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetRateLimitSettingsWarningThresholdCmd := NewGetRateLimitSettingsWarningThresholdCmd()
	RateLimitSettingsCmd.AddCommand(GetRateLimitSettingsWarningThresholdCmd)
}

var (
	ReplaceRateLimitSettingsWarningThresholddata string

	ReplaceRateLimitSettingsWarningThresholdBackup bool
)

func NewReplaceRateLimitSettingsWarningThresholdCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replaceWarningThreshold",
		Long: "Replace the Rate Limit Warning Threshold Percentage",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RateLimitSettingsAPI.ReplaceRateLimitSettingsWarningThreshold(apiClient.GetConfig().Context)

			if ReplaceRateLimitSettingsWarningThresholddata != "" {
				req = req.Data(ReplaceRateLimitSettingsWarningThresholddata)
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
			if ReplaceRateLimitSettingsWarningThresholdBackup {

				err := utils.BackupObject(d, "RateLimitSettings", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceRateLimitSettingsWarningThresholddata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceRateLimitSettingsWarningThresholdBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceRateLimitSettingsWarningThresholdCmd := NewReplaceRateLimitSettingsWarningThresholdCmd()
	RateLimitSettingsCmd.AddCommand(ReplaceRateLimitSettingsWarningThresholdCmd)
}
