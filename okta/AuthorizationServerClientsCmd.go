package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var AuthorizationServerClientsCmd = &cobra.Command{
	Use:  "authorizationServerClients",
	Long: "Manage AuthorizationServerClientsAPI",
}

func init() {
	rootCmd.AddCommand(AuthorizationServerClientsCmd)
}

var (
	ListOAuth2ClientsForAuthorizationServerauthServerId string

	ListOAuth2ClientsForAuthorizationServerBackup bool
)

func NewListOAuth2ClientsForAuthorizationServerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listOAuth2ClientsForAuthorizationServer",
		Long: "List all Client resources for an authorization server",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerClientsAPI.ListOAuth2ClientsForAuthorizationServer(apiClient.GetConfig().Context, ListOAuth2ClientsForAuthorizationServerauthServerId)

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
			if ListOAuth2ClientsForAuthorizationServerBackup {

				idParam := ListOAuth2ClientsForAuthorizationServerauthServerId
				err := utils.BackupObject(d, "AuthorizationServerClients", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListOAuth2ClientsForAuthorizationServerauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().BoolVarP(&ListOAuth2ClientsForAuthorizationServerBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListOAuth2ClientsForAuthorizationServerCmd := NewListOAuth2ClientsForAuthorizationServerCmd()
	AuthorizationServerClientsCmd.AddCommand(ListOAuth2ClientsForAuthorizationServerCmd)
}

var (
	ListRefreshTokensForAuthorizationServerAndClientauthServerId string

	ListRefreshTokensForAuthorizationServerAndClientclientId string

	ListRefreshTokensForAuthorizationServerAndClientBackup bool
)

func NewListRefreshTokensForAuthorizationServerAndClientCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listRefreshTokensForAuthorizationServerAndClient",
		Long: "List all refresh tokens for a Client",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerClientsAPI.ListRefreshTokensForAuthorizationServerAndClient(apiClient.GetConfig().Context, ListRefreshTokensForAuthorizationServerAndClientauthServerId, ListRefreshTokensForAuthorizationServerAndClientclientId)

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
			if ListRefreshTokensForAuthorizationServerAndClientBackup {

				idParam := ListRefreshTokensForAuthorizationServerAndClientauthServerId
				err := utils.BackupObject(d, "AuthorizationServerClients", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListRefreshTokensForAuthorizationServerAndClientauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().StringVarP(&ListRefreshTokensForAuthorizationServerAndClientclientId, "clientId", "", "", "")
	cmd.MarkFlagRequired("clientId")

	cmd.Flags().BoolVarP(&ListRefreshTokensForAuthorizationServerAndClientBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListRefreshTokensForAuthorizationServerAndClientCmd := NewListRefreshTokensForAuthorizationServerAndClientCmd()
	AuthorizationServerClientsCmd.AddCommand(ListRefreshTokensForAuthorizationServerAndClientCmd)
}

var (
	RevokeRefreshTokensForAuthorizationServerAndClientauthServerId string

	RevokeRefreshTokensForAuthorizationServerAndClientclientId string

	RevokeRefreshTokensForAuthorizationServerAndClientBackup bool
)

func NewRevokeRefreshTokensForAuthorizationServerAndClientCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "revokeRefreshTokensForAuthorizationServerAndClient",
		Long: "Revoke all refresh tokens for a Client",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerClientsAPI.RevokeRefreshTokensForAuthorizationServerAndClient(apiClient.GetConfig().Context, RevokeRefreshTokensForAuthorizationServerAndClientauthServerId, RevokeRefreshTokensForAuthorizationServerAndClientclientId)

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
			if RevokeRefreshTokensForAuthorizationServerAndClientBackup {

				idParam := RevokeRefreshTokensForAuthorizationServerAndClientauthServerId
				err := utils.BackupObject(d, "AuthorizationServerClients", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&RevokeRefreshTokensForAuthorizationServerAndClientauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().StringVarP(&RevokeRefreshTokensForAuthorizationServerAndClientclientId, "clientId", "", "", "")
	cmd.MarkFlagRequired("clientId")

	cmd.Flags().BoolVarP(&RevokeRefreshTokensForAuthorizationServerAndClientBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	RevokeRefreshTokensForAuthorizationServerAndClientCmd := NewRevokeRefreshTokensForAuthorizationServerAndClientCmd()
	AuthorizationServerClientsCmd.AddCommand(RevokeRefreshTokensForAuthorizationServerAndClientCmd)
}

var (
	GetRefreshTokenForAuthorizationServerAndClientauthServerId string

	GetRefreshTokenForAuthorizationServerAndClientclientId string

	GetRefreshTokenForAuthorizationServerAndClienttokenId string

	GetRefreshTokenForAuthorizationServerAndClientBackup bool
)

func NewGetRefreshTokenForAuthorizationServerAndClientCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getRefreshTokenForAuthorizationServerAndClient",
		Long: "Retrieve a refresh token for a Client",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerClientsAPI.GetRefreshTokenForAuthorizationServerAndClient(apiClient.GetConfig().Context, GetRefreshTokenForAuthorizationServerAndClientauthServerId, GetRefreshTokenForAuthorizationServerAndClientclientId, GetRefreshTokenForAuthorizationServerAndClienttokenId)

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
			if GetRefreshTokenForAuthorizationServerAndClientBackup {

				idParam := GetRefreshTokenForAuthorizationServerAndClientauthServerId
				err := utils.BackupObject(d, "AuthorizationServerClients", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetRefreshTokenForAuthorizationServerAndClientauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().StringVarP(&GetRefreshTokenForAuthorizationServerAndClientclientId, "clientId", "", "", "")
	cmd.MarkFlagRequired("clientId")

	cmd.Flags().StringVarP(&GetRefreshTokenForAuthorizationServerAndClienttokenId, "tokenId", "", "", "")
	cmd.MarkFlagRequired("tokenId")

	cmd.Flags().BoolVarP(&GetRefreshTokenForAuthorizationServerAndClientBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetRefreshTokenForAuthorizationServerAndClientCmd := NewGetRefreshTokenForAuthorizationServerAndClientCmd()
	AuthorizationServerClientsCmd.AddCommand(GetRefreshTokenForAuthorizationServerAndClientCmd)
}

var (
	RevokeRefreshTokenForAuthorizationServerAndClientauthServerId string

	RevokeRefreshTokenForAuthorizationServerAndClientclientId string

	RevokeRefreshTokenForAuthorizationServerAndClienttokenId string

	RevokeRefreshTokenForAuthorizationServerAndClientBackup bool
)

func NewRevokeRefreshTokenForAuthorizationServerAndClientCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "revokeRefreshTokenForAuthorizationServerAndClient",
		Long: "Revoke a refresh token for a Client",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerClientsAPI.RevokeRefreshTokenForAuthorizationServerAndClient(apiClient.GetConfig().Context, RevokeRefreshTokenForAuthorizationServerAndClientauthServerId, RevokeRefreshTokenForAuthorizationServerAndClientclientId, RevokeRefreshTokenForAuthorizationServerAndClienttokenId)

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
			if RevokeRefreshTokenForAuthorizationServerAndClientBackup {

				idParam := RevokeRefreshTokenForAuthorizationServerAndClientauthServerId
				err := utils.BackupObject(d, "AuthorizationServerClients", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&RevokeRefreshTokenForAuthorizationServerAndClientauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().StringVarP(&RevokeRefreshTokenForAuthorizationServerAndClientclientId, "clientId", "", "", "")
	cmd.MarkFlagRequired("clientId")

	cmd.Flags().StringVarP(&RevokeRefreshTokenForAuthorizationServerAndClienttokenId, "tokenId", "", "", "")
	cmd.MarkFlagRequired("tokenId")

	cmd.Flags().BoolVarP(&RevokeRefreshTokenForAuthorizationServerAndClientBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	RevokeRefreshTokenForAuthorizationServerAndClientCmd := NewRevokeRefreshTokenForAuthorizationServerAndClientCmd()
	AuthorizationServerClientsCmd.AddCommand(RevokeRefreshTokenForAuthorizationServerAndClientCmd)
}
