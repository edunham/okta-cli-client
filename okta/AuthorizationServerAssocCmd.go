package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var AuthorizationServerAssocCmd = &cobra.Command{
	Use:  "authorizationServerAssoc",
	Long: "Manage AuthorizationServerAssocAPI",
}

func init() {
	rootCmd.AddCommand(AuthorizationServerAssocCmd)
}

var (
	CreateAssociatedServersauthServerId string

	CreateAssociatedServersdata string

	CreateAssociatedServersBackup bool
)

func NewCreateAssociatedServersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "createAssociatedServers",
		Long: "Create an associated Authorization Server",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerAssocAPI.CreateAssociatedServers(apiClient.GetConfig().Context, CreateAssociatedServersauthServerId)

			if CreateAssociatedServersdata != "" {
				req = req.Data(CreateAssociatedServersdata)
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
			if CreateAssociatedServersBackup {

				idParam := CreateAssociatedServersauthServerId
				err := utils.BackupObject(d, "AuthorizationServerAssoc", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateAssociatedServersauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().StringVarP(&CreateAssociatedServersdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateAssociatedServersBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateAssociatedServersCmd := NewCreateAssociatedServersCmd()
	AuthorizationServerAssocCmd.AddCommand(CreateAssociatedServersCmd)
}

var (
	ListAssociatedServersByTrustedTypeauthServerId string

	ListAssociatedServersByTrustedTypeBackup bool
)

func NewListAssociatedServersByTrustedTypeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listAssociatedServersByTrustedType",
		Long: "List all associated Authorization Servers",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerAssocAPI.ListAssociatedServersByTrustedType(apiClient.GetConfig().Context, ListAssociatedServersByTrustedTypeauthServerId)

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
			if ListAssociatedServersByTrustedTypeBackup {

				idParam := ListAssociatedServersByTrustedTypeauthServerId
				err := utils.BackupObject(d, "AuthorizationServerAssoc", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListAssociatedServersByTrustedTypeauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().BoolVarP(&ListAssociatedServersByTrustedTypeBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListAssociatedServersByTrustedTypeCmd := NewListAssociatedServersByTrustedTypeCmd()
	AuthorizationServerAssocCmd.AddCommand(ListAssociatedServersByTrustedTypeCmd)
}

var (
	DeleteAssociatedServerauthServerId string

	DeleteAssociatedServerassociatedServerId string

	DeleteAssociatedServerBackup bool
)

func NewDeleteAssociatedServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deleteAssociatedServer",
		Long: "Delete an associated Authorization Server",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerAssocAPI.DeleteAssociatedServer(apiClient.GetConfig().Context, DeleteAssociatedServerauthServerId, DeleteAssociatedServerassociatedServerId)

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
			if DeleteAssociatedServerBackup {

				idParam := DeleteAssociatedServerauthServerId
				err := utils.BackupObject(d, "AuthorizationServerAssoc", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteAssociatedServerauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().StringVarP(&DeleteAssociatedServerassociatedServerId, "associatedServerId", "", "", "")
	cmd.MarkFlagRequired("associatedServerId")

	cmd.Flags().BoolVarP(&DeleteAssociatedServerBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteAssociatedServerCmd := NewDeleteAssociatedServerCmd()
	AuthorizationServerAssocCmd.AddCommand(DeleteAssociatedServerCmd)
}
