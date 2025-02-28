package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var EmailServerCmd = &cobra.Command{
	Use:  "emailServer",
	Long: "Manage EmailServerAPI",
}

func init() {
	rootCmd.AddCommand(EmailServerCmd)
}

var (
	CreateEmailServerdata string

	CreateEmailServerBackup bool
)

func NewCreateEmailServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create",
		Long: "Create a custom SMTP server",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.EmailServerAPI.CreateEmailServer(apiClient.GetConfig().Context)

			if CreateEmailServerdata != "" {
				req = req.Data(CreateEmailServerdata)
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
			if CreateEmailServerBackup {

				err := utils.BackupObject(d, "EmailServer", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateEmailServerdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateEmailServerBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateEmailServerCmd := NewCreateEmailServerCmd()
	EmailServerCmd.AddCommand(CreateEmailServerCmd)
}

var ListEmailServersBackup bool

func NewListEmailServersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "lists",
		Long: "List all enrolled SMTP servers",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.EmailServerAPI.ListEmailServers(apiClient.GetConfig().Context)

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
			if ListEmailServersBackup {

				err := utils.BackupObject(d, "EmailServer", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListEmailServersBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListEmailServersCmd := NewListEmailServersCmd()
	EmailServerCmd.AddCommand(ListEmailServersCmd)
}

var (
	GetEmailServeremailServerId string

	GetEmailServerBackup bool
)

func NewGetEmailServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "get",
		Long: "Retrieve an SMTP Server configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.EmailServerAPI.GetEmailServer(apiClient.GetConfig().Context, GetEmailServeremailServerId)

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
			if GetEmailServerBackup {

				idParam := GetEmailServeremailServerId
				err := utils.BackupObject(d, "EmailServer", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetEmailServeremailServerId, "emailServerId", "", "", "")
	cmd.MarkFlagRequired("emailServerId")

	cmd.Flags().BoolVarP(&GetEmailServerBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetEmailServerCmd := NewGetEmailServerCmd()
	EmailServerCmd.AddCommand(GetEmailServerCmd)
}

var (
	DeleteEmailServeremailServerId string

	DeleteEmailServerBackup bool
)

func NewDeleteEmailServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete",
		Long: "Delete an SMTP Server configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.EmailServerAPI.DeleteEmailServer(apiClient.GetConfig().Context, DeleteEmailServeremailServerId)

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
			if DeleteEmailServerBackup {

				idParam := DeleteEmailServeremailServerId
				err := utils.BackupObject(d, "EmailServer", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteEmailServeremailServerId, "emailServerId", "", "", "")
	cmd.MarkFlagRequired("emailServerId")

	cmd.Flags().BoolVarP(&DeleteEmailServerBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteEmailServerCmd := NewDeleteEmailServerCmd()
	EmailServerCmd.AddCommand(DeleteEmailServerCmd)
}

var (
	UpdateEmailServeremailServerId string

	UpdateEmailServerdata string

	UpdateEmailServerBackup bool
)

func NewUpdateEmailServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "update",
		Long: "Update an SMTP Server configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.EmailServerAPI.UpdateEmailServer(apiClient.GetConfig().Context, UpdateEmailServeremailServerId)

			if UpdateEmailServerdata != "" {
				req = req.Data(UpdateEmailServerdata)
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
			if UpdateEmailServerBackup {

				idParam := UpdateEmailServeremailServerId
				err := utils.BackupObject(d, "EmailServer", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UpdateEmailServeremailServerId, "emailServerId", "", "", "")
	cmd.MarkFlagRequired("emailServerId")

	cmd.Flags().StringVarP(&UpdateEmailServerdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&UpdateEmailServerBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UpdateEmailServerCmd := NewUpdateEmailServerCmd()
	EmailServerCmd.AddCommand(UpdateEmailServerCmd)
}

var (
	TestEmailServeremailServerId string

	TestEmailServerdata string

	TestEmailServerBackup bool
)

func NewTestEmailServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "test",
		Long: "Test an SMTP Server configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.EmailServerAPI.TestEmailServer(apiClient.GetConfig().Context, TestEmailServeremailServerId)

			if TestEmailServerdata != "" {
				req = req.Data(TestEmailServerdata)
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
			if TestEmailServerBackup {

				idParam := TestEmailServeremailServerId
				err := utils.BackupObject(d, "EmailServer", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&TestEmailServeremailServerId, "emailServerId", "", "", "")
	cmd.MarkFlagRequired("emailServerId")

	cmd.Flags().StringVarP(&TestEmailServerdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&TestEmailServerBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	TestEmailServerCmd := NewTestEmailServerCmd()
	EmailServerCmd.AddCommand(TestEmailServerCmd)
}
