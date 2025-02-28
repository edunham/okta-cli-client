package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var AuthorizationServerClaimsCmd = &cobra.Command{
	Use:  "authorizationServerClaims",
	Long: "Manage AuthorizationServerClaimsAPI",
}

func init() {
	rootCmd.AddCommand(AuthorizationServerClaimsCmd)
}

var (
	CreateOAuth2ClaimauthServerId string

	CreateOAuth2Claimdata string

	CreateOAuth2ClaimBackup bool
)

func NewCreateOAuth2ClaimCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "createOAuth2Claim",
		Long: "Create a custom token Claim",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerClaimsAPI.CreateOAuth2Claim(apiClient.GetConfig().Context, CreateOAuth2ClaimauthServerId)

			if CreateOAuth2Claimdata != "" {
				req = req.Data(CreateOAuth2Claimdata)
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
			if CreateOAuth2ClaimBackup {

				idParam := CreateOAuth2ClaimauthServerId
				err := utils.BackupObject(d, "AuthorizationServerClaims", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateOAuth2ClaimauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().StringVarP(&CreateOAuth2Claimdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateOAuth2ClaimBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateOAuth2ClaimCmd := NewCreateOAuth2ClaimCmd()
	AuthorizationServerClaimsCmd.AddCommand(CreateOAuth2ClaimCmd)
}

var (
	ListOAuth2ClaimsauthServerId string

	ListOAuth2ClaimsBackup bool
)

func NewListOAuth2ClaimsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listOAuth2Claims",
		Long: "List all custom token Claims",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerClaimsAPI.ListOAuth2Claims(apiClient.GetConfig().Context, ListOAuth2ClaimsauthServerId)

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
			if ListOAuth2ClaimsBackup {

				idParam := ListOAuth2ClaimsauthServerId
				err := utils.BackupObject(d, "AuthorizationServerClaims", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListOAuth2ClaimsauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().BoolVarP(&ListOAuth2ClaimsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListOAuth2ClaimsCmd := NewListOAuth2ClaimsCmd()
	AuthorizationServerClaimsCmd.AddCommand(ListOAuth2ClaimsCmd)
}

var (
	GetOAuth2ClaimauthServerId string

	GetOAuth2ClaimclaimId string

	GetOAuth2ClaimBackup bool
)

func NewGetOAuth2ClaimCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getOAuth2Claim",
		Long: "Retrieve a custom token Claim",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerClaimsAPI.GetOAuth2Claim(apiClient.GetConfig().Context, GetOAuth2ClaimauthServerId, GetOAuth2ClaimclaimId)

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
			if GetOAuth2ClaimBackup {

				idParam := GetOAuth2ClaimauthServerId
				err := utils.BackupObject(d, "AuthorizationServerClaims", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetOAuth2ClaimauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().StringVarP(&GetOAuth2ClaimclaimId, "claimId", "", "", "")
	cmd.MarkFlagRequired("claimId")

	cmd.Flags().BoolVarP(&GetOAuth2ClaimBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetOAuth2ClaimCmd := NewGetOAuth2ClaimCmd()
	AuthorizationServerClaimsCmd.AddCommand(GetOAuth2ClaimCmd)
}

var (
	ReplaceOAuth2ClaimauthServerId string

	ReplaceOAuth2ClaimclaimId string

	ReplaceOAuth2Claimdata string

	ReplaceOAuth2ClaimBackup bool
)

func NewReplaceOAuth2ClaimCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replaceOAuth2Claim",
		Long: "Replace a custom token Claim",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerClaimsAPI.ReplaceOAuth2Claim(apiClient.GetConfig().Context, ReplaceOAuth2ClaimauthServerId, ReplaceOAuth2ClaimclaimId)

			if ReplaceOAuth2Claimdata != "" {
				req = req.Data(ReplaceOAuth2Claimdata)
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
			if ReplaceOAuth2ClaimBackup {

				idParam := ReplaceOAuth2ClaimauthServerId
				err := utils.BackupObject(d, "AuthorizationServerClaims", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceOAuth2ClaimauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().StringVarP(&ReplaceOAuth2ClaimclaimId, "claimId", "", "", "")
	cmd.MarkFlagRequired("claimId")

	cmd.Flags().StringVarP(&ReplaceOAuth2Claimdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceOAuth2ClaimBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceOAuth2ClaimCmd := NewReplaceOAuth2ClaimCmd()
	AuthorizationServerClaimsCmd.AddCommand(ReplaceOAuth2ClaimCmd)
}

var (
	DeleteOAuth2ClaimauthServerId string

	DeleteOAuth2ClaimclaimId string

	DeleteOAuth2ClaimBackup bool
)

func NewDeleteOAuth2ClaimCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deleteOAuth2Claim",
		Long: "Delete a custom token Claim",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerClaimsAPI.DeleteOAuth2Claim(apiClient.GetConfig().Context, DeleteOAuth2ClaimauthServerId, DeleteOAuth2ClaimclaimId)

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
			if DeleteOAuth2ClaimBackup {

				idParam := DeleteOAuth2ClaimauthServerId
				err := utils.BackupObject(d, "AuthorizationServerClaims", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteOAuth2ClaimauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().StringVarP(&DeleteOAuth2ClaimclaimId, "claimId", "", "", "")
	cmd.MarkFlagRequired("claimId")

	cmd.Flags().BoolVarP(&DeleteOAuth2ClaimBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteOAuth2ClaimCmd := NewDeleteOAuth2ClaimCmd()
	AuthorizationServerClaimsCmd.AddCommand(DeleteOAuth2ClaimCmd)
}
