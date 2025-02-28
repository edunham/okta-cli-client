package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var ApiTokenCmd = &cobra.Command{
	Use:  "apiToken",
	Long: "Manage ApiTokenAPI",
}

func init() {
	rootCmd.AddCommand(ApiTokenCmd)
}

var ListApiTokensBackup bool

func NewListApiTokensCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "lists",
		Long: "List all API Token Metadata",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApiTokenAPI.ListApiTokens(apiClient.GetConfig().Context)

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
			if ListApiTokensBackup {

				err := utils.BackupObject(d, "ApiToken", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListApiTokensBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListApiTokensCmd := NewListApiTokensCmd()
	ApiTokenCmd.AddCommand(ListApiTokensCmd)
}

var RevokeCurrentApiTokenBackup bool

func NewRevokeCurrentApiTokenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "revokeCurrent",
		Long: "Revoke the Current API Token",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApiTokenAPI.RevokeCurrentApiToken(apiClient.GetConfig().Context)

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
			if RevokeCurrentApiTokenBackup {

				err := utils.BackupObject(d, "ApiToken", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&RevokeCurrentApiTokenBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	RevokeCurrentApiTokenCmd := NewRevokeCurrentApiTokenCmd()
	ApiTokenCmd.AddCommand(RevokeCurrentApiTokenCmd)
}

var (
	GetApiTokenapiTokenId string

	GetApiTokenBackup bool
)

func NewGetApiTokenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "get",
		Long: "Retrieve an API Token's Metadata",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApiTokenAPI.GetApiToken(apiClient.GetConfig().Context, GetApiTokenapiTokenId)

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
			if GetApiTokenBackup {

				idParam := GetApiTokenapiTokenId
				err := utils.BackupObject(d, "ApiToken", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetApiTokenapiTokenId, "apiTokenId", "", "", "")
	cmd.MarkFlagRequired("apiTokenId")

	cmd.Flags().BoolVarP(&GetApiTokenBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetApiTokenCmd := NewGetApiTokenCmd()
	ApiTokenCmd.AddCommand(GetApiTokenCmd)
}

var (
	RevokeApiTokenapiTokenId string

	RevokeApiTokenBackup bool
)

func NewRevokeApiTokenCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "revoke",
		Long: "Revoke an API Token",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApiTokenAPI.RevokeApiToken(apiClient.GetConfig().Context, RevokeApiTokenapiTokenId)

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
			if RevokeApiTokenBackup {

				idParam := RevokeApiTokenapiTokenId
				err := utils.BackupObject(d, "ApiToken", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&RevokeApiTokenapiTokenId, "apiTokenId", "", "", "")
	cmd.MarkFlagRequired("apiTokenId")

	cmd.Flags().BoolVarP(&RevokeApiTokenBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	RevokeApiTokenCmd := NewRevokeApiTokenCmd()
	ApiTokenCmd.AddCommand(RevokeApiTokenCmd)
}
