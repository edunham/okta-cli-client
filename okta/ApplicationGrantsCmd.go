package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var ApplicationGrantsCmd = &cobra.Command{
	Use:  "applicationGrants",
	Long: "Manage ApplicationGrantsAPI",
}

func init() {
	rootCmd.AddCommand(ApplicationGrantsCmd)
}

var (
	GrantConsentToScopeappId string

	GrantConsentToScopedata string

	GrantConsentToScopeBackup bool
)

func NewGrantConsentToScopeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "grantConsentToScope",
		Long: "Grant consent to scope",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationGrantsAPI.GrantConsentToScope(apiClient.GetConfig().Context, GrantConsentToScopeappId)

			if GrantConsentToScopedata != "" {
				req = req.Data(GrantConsentToScopedata)
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
			if GrantConsentToScopeBackup {

				idParam := GrantConsentToScopeappId
				err := utils.BackupObject(d, "ApplicationGrants", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GrantConsentToScopeappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().StringVarP(&GrantConsentToScopedata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&GrantConsentToScopeBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GrantConsentToScopeCmd := NewGrantConsentToScopeCmd()
	ApplicationGrantsCmd.AddCommand(GrantConsentToScopeCmd)
}

var (
	ListScopeConsentGrantsappId string

	ListScopeConsentGrantsBackup bool
)

func NewListScopeConsentGrantsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listScopeConsentGrants",
		Long: "List all app Grants",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationGrantsAPI.ListScopeConsentGrants(apiClient.GetConfig().Context, ListScopeConsentGrantsappId)

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
			if ListScopeConsentGrantsBackup {

				idParam := ListScopeConsentGrantsappId
				err := utils.BackupObject(d, "ApplicationGrants", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListScopeConsentGrantsappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().BoolVarP(&ListScopeConsentGrantsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListScopeConsentGrantsCmd := NewListScopeConsentGrantsCmd()
	ApplicationGrantsCmd.AddCommand(ListScopeConsentGrantsCmd)
}

var (
	GetScopeConsentGrantappId string

	GetScopeConsentGrantgrantId string

	GetScopeConsentGrantBackup bool
)

func NewGetScopeConsentGrantCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getScopeConsentGrant",
		Long: "Retrieve an app Grant",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationGrantsAPI.GetScopeConsentGrant(apiClient.GetConfig().Context, GetScopeConsentGrantappId, GetScopeConsentGrantgrantId)

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
			if GetScopeConsentGrantBackup {

				idParam := GetScopeConsentGrantappId
				err := utils.BackupObject(d, "ApplicationGrants", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetScopeConsentGrantappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().StringVarP(&GetScopeConsentGrantgrantId, "grantId", "", "", "")
	cmd.MarkFlagRequired("grantId")

	cmd.Flags().BoolVarP(&GetScopeConsentGrantBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetScopeConsentGrantCmd := NewGetScopeConsentGrantCmd()
	ApplicationGrantsCmd.AddCommand(GetScopeConsentGrantCmd)
}

var (
	RevokeScopeConsentGrantappId string

	RevokeScopeConsentGrantgrantId string

	RevokeScopeConsentGrantBackup bool
)

func NewRevokeScopeConsentGrantCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "revokeScopeConsentGrant",
		Long: "Revoke an app Grant",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationGrantsAPI.RevokeScopeConsentGrant(apiClient.GetConfig().Context, RevokeScopeConsentGrantappId, RevokeScopeConsentGrantgrantId)

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
			if RevokeScopeConsentGrantBackup {

				idParam := RevokeScopeConsentGrantappId
				err := utils.BackupObject(d, "ApplicationGrants", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&RevokeScopeConsentGrantappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().StringVarP(&RevokeScopeConsentGrantgrantId, "grantId", "", "", "")
	cmd.MarkFlagRequired("grantId")

	cmd.Flags().BoolVarP(&RevokeScopeConsentGrantBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	RevokeScopeConsentGrantCmd := NewRevokeScopeConsentGrantCmd()
	ApplicationGrantsCmd.AddCommand(RevokeScopeConsentGrantCmd)
}
