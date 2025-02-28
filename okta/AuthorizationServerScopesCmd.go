package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var AuthorizationServerScopesCmd = &cobra.Command{
	Use:  "authorizationServerScopes",
	Long: "Manage AuthorizationServerScopesAPI",
}

func init() {
	rootCmd.AddCommand(AuthorizationServerScopesCmd)
}

var (
	CreateOAuth2ScopeauthServerId string

	CreateOAuth2Scopedata string

	CreateOAuth2ScopeBackup bool
)

func NewCreateOAuth2ScopeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "createOAuth2Scope",
		Long: "Create a Custom Token Scope",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerScopesAPI.CreateOAuth2Scope(apiClient.GetConfig().Context, CreateOAuth2ScopeauthServerId)

			if CreateOAuth2Scopedata != "" {
				req = req.Data(CreateOAuth2Scopedata)
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
			if CreateOAuth2ScopeBackup {

				idParam := CreateOAuth2ScopeauthServerId
				err := utils.BackupObject(d, "AuthorizationServerScopes", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateOAuth2ScopeauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().StringVarP(&CreateOAuth2Scopedata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateOAuth2ScopeBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateOAuth2ScopeCmd := NewCreateOAuth2ScopeCmd()
	AuthorizationServerScopesCmd.AddCommand(CreateOAuth2ScopeCmd)
}

var (
	ListOAuth2ScopesauthServerId string

	ListOAuth2ScopesBackup bool
)

func NewListOAuth2ScopesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listOAuth2Scopes",
		Long: "List all Custom Token Scopes",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerScopesAPI.ListOAuth2Scopes(apiClient.GetConfig().Context, ListOAuth2ScopesauthServerId)

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
			if ListOAuth2ScopesBackup {

				idParam := ListOAuth2ScopesauthServerId
				err := utils.BackupObject(d, "AuthorizationServerScopes", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListOAuth2ScopesauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().BoolVarP(&ListOAuth2ScopesBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListOAuth2ScopesCmd := NewListOAuth2ScopesCmd()
	AuthorizationServerScopesCmd.AddCommand(ListOAuth2ScopesCmd)
}

var (
	GetOAuth2ScopeauthServerId string

	GetOAuth2ScopescopeId string

	GetOAuth2ScopeBackup bool
)

func NewGetOAuth2ScopeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getOAuth2Scope",
		Long: "Retrieve a Custom Token Scope",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerScopesAPI.GetOAuth2Scope(apiClient.GetConfig().Context, GetOAuth2ScopeauthServerId, GetOAuth2ScopescopeId)

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
			if GetOAuth2ScopeBackup {

				idParam := GetOAuth2ScopeauthServerId
				err := utils.BackupObject(d, "AuthorizationServerScopes", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetOAuth2ScopeauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().StringVarP(&GetOAuth2ScopescopeId, "scopeId", "", "", "")
	cmd.MarkFlagRequired("scopeId")

	cmd.Flags().BoolVarP(&GetOAuth2ScopeBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetOAuth2ScopeCmd := NewGetOAuth2ScopeCmd()
	AuthorizationServerScopesCmd.AddCommand(GetOAuth2ScopeCmd)
}

var (
	ReplaceOAuth2ScopeauthServerId string

	ReplaceOAuth2ScopescopeId string

	ReplaceOAuth2Scopedata string

	ReplaceOAuth2ScopeBackup bool
)

func NewReplaceOAuth2ScopeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replaceOAuth2Scope",
		Long: "Replace a Custom Token Scope",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerScopesAPI.ReplaceOAuth2Scope(apiClient.GetConfig().Context, ReplaceOAuth2ScopeauthServerId, ReplaceOAuth2ScopescopeId)

			if ReplaceOAuth2Scopedata != "" {
				req = req.Data(ReplaceOAuth2Scopedata)
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
			if ReplaceOAuth2ScopeBackup {

				idParam := ReplaceOAuth2ScopeauthServerId
				err := utils.BackupObject(d, "AuthorizationServerScopes", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceOAuth2ScopeauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().StringVarP(&ReplaceOAuth2ScopescopeId, "scopeId", "", "", "")
	cmd.MarkFlagRequired("scopeId")

	cmd.Flags().StringVarP(&ReplaceOAuth2Scopedata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceOAuth2ScopeBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceOAuth2ScopeCmd := NewReplaceOAuth2ScopeCmd()
	AuthorizationServerScopesCmd.AddCommand(ReplaceOAuth2ScopeCmd)
}

var (
	DeleteOAuth2ScopeauthServerId string

	DeleteOAuth2ScopescopeId string

	DeleteOAuth2ScopeBackup bool
)

func NewDeleteOAuth2ScopeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deleteOAuth2Scope",
		Long: "Delete a Custom Token Scope",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerScopesAPI.DeleteOAuth2Scope(apiClient.GetConfig().Context, DeleteOAuth2ScopeauthServerId, DeleteOAuth2ScopescopeId)

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
			if DeleteOAuth2ScopeBackup {

				idParam := DeleteOAuth2ScopeauthServerId
				err := utils.BackupObject(d, "AuthorizationServerScopes", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteOAuth2ScopeauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().StringVarP(&DeleteOAuth2ScopescopeId, "scopeId", "", "", "")
	cmd.MarkFlagRequired("scopeId")

	cmd.Flags().BoolVarP(&DeleteOAuth2ScopeBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteOAuth2ScopeCmd := NewDeleteOAuth2ScopeCmd()
	AuthorizationServerScopesCmd.AddCommand(DeleteOAuth2ScopeCmd)
}
