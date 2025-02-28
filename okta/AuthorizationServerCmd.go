package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var AuthorizationServerCmd = &cobra.Command{
	Use:  "authorizationServer",
	Long: "Manage AuthorizationServerAPI",
}

func init() {
	rootCmd.AddCommand(AuthorizationServerCmd)
}

var (
	CreateAuthorizationServerdata string

	CreateAuthorizationServerBackup bool
)

func NewCreateAuthorizationServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create",
		Long: "Create an Authorization Server",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerAPI.CreateAuthorizationServer(apiClient.GetConfig().Context)

			if CreateAuthorizationServerdata != "" {
				req = req.Data(CreateAuthorizationServerdata)
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
			if CreateAuthorizationServerBackup {

				err := utils.BackupObject(d, "AuthorizationServer", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateAuthorizationServerdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateAuthorizationServerBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateAuthorizationServerCmd := NewCreateAuthorizationServerCmd()
	AuthorizationServerCmd.AddCommand(CreateAuthorizationServerCmd)
}

var ListAuthorizationServersBackup bool

func NewListAuthorizationServersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "lists",
		Long: "List all Authorization Servers",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerAPI.ListAuthorizationServers(apiClient.GetConfig().Context)

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
			if ListAuthorizationServersBackup {

				err := utils.BackupObject(d, "AuthorizationServer", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListAuthorizationServersBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListAuthorizationServersCmd := NewListAuthorizationServersCmd()
	AuthorizationServerCmd.AddCommand(ListAuthorizationServersCmd)
}

var (
	GetAuthorizationServerauthServerId string

	GetAuthorizationServerBackup bool
)

func NewGetAuthorizationServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "get",
		Long: "Retrieve an Authorization Server",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerAPI.GetAuthorizationServer(apiClient.GetConfig().Context, GetAuthorizationServerauthServerId)

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
			if GetAuthorizationServerBackup {

				idParam := GetAuthorizationServerauthServerId
				err := utils.BackupObject(d, "AuthorizationServer", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetAuthorizationServerauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().BoolVarP(&GetAuthorizationServerBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetAuthorizationServerCmd := NewGetAuthorizationServerCmd()
	AuthorizationServerCmd.AddCommand(GetAuthorizationServerCmd)
}

var (
	ReplaceAuthorizationServerauthServerId string

	ReplaceAuthorizationServerdata string

	ReplaceAuthorizationServerBackup bool
)

func NewReplaceAuthorizationServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replace",
		Long: "Replace an Authorization Server",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerAPI.ReplaceAuthorizationServer(apiClient.GetConfig().Context, ReplaceAuthorizationServerauthServerId)

			if ReplaceAuthorizationServerdata != "" {
				req = req.Data(ReplaceAuthorizationServerdata)
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
			if ReplaceAuthorizationServerBackup {

				idParam := ReplaceAuthorizationServerauthServerId
				err := utils.BackupObject(d, "AuthorizationServer", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceAuthorizationServerauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().StringVarP(&ReplaceAuthorizationServerdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceAuthorizationServerBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceAuthorizationServerCmd := NewReplaceAuthorizationServerCmd()
	AuthorizationServerCmd.AddCommand(ReplaceAuthorizationServerCmd)
}

var (
	DeleteAuthorizationServerauthServerId string

	DeleteAuthorizationServerBackup bool
)

func NewDeleteAuthorizationServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete",
		Long: "Delete an Authorization Server",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerAPI.DeleteAuthorizationServer(apiClient.GetConfig().Context, DeleteAuthorizationServerauthServerId)

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
			if DeleteAuthorizationServerBackup {

				idParam := DeleteAuthorizationServerauthServerId
				err := utils.BackupObject(d, "AuthorizationServer", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteAuthorizationServerauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().BoolVarP(&DeleteAuthorizationServerBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteAuthorizationServerCmd := NewDeleteAuthorizationServerCmd()
	AuthorizationServerCmd.AddCommand(DeleteAuthorizationServerCmd)
}

var (
	ActivateAuthorizationServerauthServerId string

	ActivateAuthorizationServerBackup bool
)

func NewActivateAuthorizationServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "activate",
		Long: "Activate an Authorization Server",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerAPI.ActivateAuthorizationServer(apiClient.GetConfig().Context, ActivateAuthorizationServerauthServerId)

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
			if ActivateAuthorizationServerBackup {

				idParam := ActivateAuthorizationServerauthServerId
				err := utils.BackupObject(d, "AuthorizationServer", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ActivateAuthorizationServerauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().BoolVarP(&ActivateAuthorizationServerBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ActivateAuthorizationServerCmd := NewActivateAuthorizationServerCmd()
	AuthorizationServerCmd.AddCommand(ActivateAuthorizationServerCmd)
}

var (
	DeactivateAuthorizationServerauthServerId string

	DeactivateAuthorizationServerBackup bool
)

func NewDeactivateAuthorizationServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deactivate",
		Long: "Deactivate an Authorization Server",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerAPI.DeactivateAuthorizationServer(apiClient.GetConfig().Context, DeactivateAuthorizationServerauthServerId)

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
			if DeactivateAuthorizationServerBackup {

				idParam := DeactivateAuthorizationServerauthServerId
				err := utils.BackupObject(d, "AuthorizationServer", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeactivateAuthorizationServerauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().BoolVarP(&DeactivateAuthorizationServerBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeactivateAuthorizationServerCmd := NewDeactivateAuthorizationServerCmd()
	AuthorizationServerCmd.AddCommand(DeactivateAuthorizationServerCmd)
}
