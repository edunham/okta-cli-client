package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var ApplicationTokensCmd = &cobra.Command{
	Use:  "applicationTokens",
	Long: "Manage ApplicationTokensAPI",
}

func init() {
	rootCmd.AddCommand(ApplicationTokensCmd)
}

var (
	ListOAuth2TokensForApplicationappId string

	ListOAuth2TokensForApplicationBackup bool
)

func NewListOAuth2TokensForApplicationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listOAuth2TokensForApplication",
		Long: "List all application refresh Tokens",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationTokensAPI.ListOAuth2TokensForApplication(apiClient.GetConfig().Context, ListOAuth2TokensForApplicationappId)

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
			if ListOAuth2TokensForApplicationBackup {

				idParam := ListOAuth2TokensForApplicationappId
				err := utils.BackupObject(d, "ApplicationTokens", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListOAuth2TokensForApplicationappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().BoolVarP(&ListOAuth2TokensForApplicationBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListOAuth2TokensForApplicationCmd := NewListOAuth2TokensForApplicationCmd()
	ApplicationTokensCmd.AddCommand(ListOAuth2TokensForApplicationCmd)
}

var (
	RevokeOAuth2TokensForApplicationappId string

	RevokeOAuth2TokensForApplicationBackup bool
)

func NewRevokeOAuth2TokensForApplicationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "revokeOAuth2TokensForApplication",
		Long: "Revoke all application Tokens",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationTokensAPI.RevokeOAuth2TokensForApplication(apiClient.GetConfig().Context, RevokeOAuth2TokensForApplicationappId)

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
			if RevokeOAuth2TokensForApplicationBackup {

				idParam := RevokeOAuth2TokensForApplicationappId
				err := utils.BackupObject(d, "ApplicationTokens", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&RevokeOAuth2TokensForApplicationappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().BoolVarP(&RevokeOAuth2TokensForApplicationBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	RevokeOAuth2TokensForApplicationCmd := NewRevokeOAuth2TokensForApplicationCmd()
	ApplicationTokensCmd.AddCommand(RevokeOAuth2TokensForApplicationCmd)
}

var (
	GetOAuth2TokenForApplicationappId string

	GetOAuth2TokenForApplicationtokenId string

	GetOAuth2TokenForApplicationBackup bool
)

func NewGetOAuth2TokenForApplicationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getOAuth2TokenForApplication",
		Long: "Retrieve an application Token",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationTokensAPI.GetOAuth2TokenForApplication(apiClient.GetConfig().Context, GetOAuth2TokenForApplicationappId, GetOAuth2TokenForApplicationtokenId)

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
			if GetOAuth2TokenForApplicationBackup {

				idParam := GetOAuth2TokenForApplicationappId
				err := utils.BackupObject(d, "ApplicationTokens", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetOAuth2TokenForApplicationappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().StringVarP(&GetOAuth2TokenForApplicationtokenId, "tokenId", "", "", "")
	cmd.MarkFlagRequired("tokenId")

	cmd.Flags().BoolVarP(&GetOAuth2TokenForApplicationBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetOAuth2TokenForApplicationCmd := NewGetOAuth2TokenForApplicationCmd()
	ApplicationTokensCmd.AddCommand(GetOAuth2TokenForApplicationCmd)
}

var (
	RevokeOAuth2TokenForApplicationappId string

	RevokeOAuth2TokenForApplicationtokenId string

	RevokeOAuth2TokenForApplicationBackup bool
)

func NewRevokeOAuth2TokenForApplicationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "revokeOAuth2TokenForApplication",
		Long: "Revoke an application Token",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationTokensAPI.RevokeOAuth2TokenForApplication(apiClient.GetConfig().Context, RevokeOAuth2TokenForApplicationappId, RevokeOAuth2TokenForApplicationtokenId)

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
			if RevokeOAuth2TokenForApplicationBackup {

				idParam := RevokeOAuth2TokenForApplicationappId
				err := utils.BackupObject(d, "ApplicationTokens", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&RevokeOAuth2TokenForApplicationappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().StringVarP(&RevokeOAuth2TokenForApplicationtokenId, "tokenId", "", "", "")
	cmd.MarkFlagRequired("tokenId")

	cmd.Flags().BoolVarP(&RevokeOAuth2TokenForApplicationBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	RevokeOAuth2TokenForApplicationCmd := NewRevokeOAuth2TokenForApplicationCmd()
	ApplicationTokensCmd.AddCommand(RevokeOAuth2TokenForApplicationCmd)
}
