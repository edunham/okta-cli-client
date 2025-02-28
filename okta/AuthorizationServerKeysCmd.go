package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var AuthorizationServerKeysCmd = &cobra.Command{
	Use:  "authorizationServerKeys",
	Long: "Manage AuthorizationServerKeysAPI",
}

func init() {
	rootCmd.AddCommand(AuthorizationServerKeysCmd)
}

var (
	ListAuthorizationServerKeysauthServerId string

	ListAuthorizationServerKeysBackup bool
)

func NewListAuthorizationServerKeysCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "list",
		Long: "List all Credential Keys",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerKeysAPI.ListAuthorizationServerKeys(apiClient.GetConfig().Context, ListAuthorizationServerKeysauthServerId)

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
			if ListAuthorizationServerKeysBackup {

				idParam := ListAuthorizationServerKeysauthServerId
				err := utils.BackupObject(d, "AuthorizationServerKeys", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListAuthorizationServerKeysauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().BoolVarP(&ListAuthorizationServerKeysBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListAuthorizationServerKeysCmd := NewListAuthorizationServerKeysCmd()
	AuthorizationServerKeysCmd.AddCommand(ListAuthorizationServerKeysCmd)
}

var (
	RotateAuthorizationServerKeysauthServerId string

	RotateAuthorizationServerKeysdata string

	RotateAuthorizationServerKeysBackup bool
)

func NewRotateAuthorizationServerKeysCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "rotate",
		Long: "Rotate all Credential Keys",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerKeysAPI.RotateAuthorizationServerKeys(apiClient.GetConfig().Context, RotateAuthorizationServerKeysauthServerId)

			if RotateAuthorizationServerKeysdata != "" {
				req = req.Data(RotateAuthorizationServerKeysdata)
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
			if RotateAuthorizationServerKeysBackup {

				idParam := RotateAuthorizationServerKeysauthServerId
				err := utils.BackupObject(d, "AuthorizationServerKeys", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&RotateAuthorizationServerKeysauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().StringVarP(&RotateAuthorizationServerKeysdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&RotateAuthorizationServerKeysBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	RotateAuthorizationServerKeysCmd := NewRotateAuthorizationServerKeysCmd()
	AuthorizationServerKeysCmd.AddCommand(RotateAuthorizationServerKeysCmd)
}
