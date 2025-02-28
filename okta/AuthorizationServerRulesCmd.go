package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var AuthorizationServerRulesCmd = &cobra.Command{
	Use:  "authorizationServerRules",
	Long: "Manage AuthorizationServerRulesAPI",
}

func init() {
	rootCmd.AddCommand(AuthorizationServerRulesCmd)
}

var (
	CreateAuthorizationServerPolicyRuleauthServerId string

	CreateAuthorizationServerPolicyRulepolicyId string

	CreateAuthorizationServerPolicyRuledata string

	CreateAuthorizationServerPolicyRuleBackup bool
)

func NewCreateAuthorizationServerPolicyRuleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "createAuthorizationServerPolicyRule",
		Long: "Create a Policy Rule",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerRulesAPI.CreateAuthorizationServerPolicyRule(apiClient.GetConfig().Context, CreateAuthorizationServerPolicyRuleauthServerId, CreateAuthorizationServerPolicyRulepolicyId)

			if CreateAuthorizationServerPolicyRuledata != "" {
				req = req.Data(CreateAuthorizationServerPolicyRuledata)
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
			if CreateAuthorizationServerPolicyRuleBackup {

				idParam := CreateAuthorizationServerPolicyRuleauthServerId
				err := utils.BackupObject(d, "AuthorizationServerRules", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateAuthorizationServerPolicyRuleauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().StringVarP(&CreateAuthorizationServerPolicyRulepolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().StringVarP(&CreateAuthorizationServerPolicyRuledata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateAuthorizationServerPolicyRuleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateAuthorizationServerPolicyRuleCmd := NewCreateAuthorizationServerPolicyRuleCmd()
	AuthorizationServerRulesCmd.AddCommand(CreateAuthorizationServerPolicyRuleCmd)
}

var (
	ListAuthorizationServerPolicyRulesauthServerId string

	ListAuthorizationServerPolicyRulespolicyId string

	ListAuthorizationServerPolicyRulesBackup bool
)

func NewListAuthorizationServerPolicyRulesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listAuthorizationServerPolicyRules",
		Long: "List all Policy Rules",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerRulesAPI.ListAuthorizationServerPolicyRules(apiClient.GetConfig().Context, ListAuthorizationServerPolicyRulesauthServerId, ListAuthorizationServerPolicyRulespolicyId)

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
			if ListAuthorizationServerPolicyRulesBackup {

				idParam := ListAuthorizationServerPolicyRulesauthServerId
				err := utils.BackupObject(d, "AuthorizationServerRules", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListAuthorizationServerPolicyRulesauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().StringVarP(&ListAuthorizationServerPolicyRulespolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().BoolVarP(&ListAuthorizationServerPolicyRulesBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListAuthorizationServerPolicyRulesCmd := NewListAuthorizationServerPolicyRulesCmd()
	AuthorizationServerRulesCmd.AddCommand(ListAuthorizationServerPolicyRulesCmd)
}

var (
	GetAuthorizationServerPolicyRuleauthServerId string

	GetAuthorizationServerPolicyRulepolicyId string

	GetAuthorizationServerPolicyRuleruleId string

	GetAuthorizationServerPolicyRuleBackup bool
)

func NewGetAuthorizationServerPolicyRuleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getAuthorizationServerPolicyRule",
		Long: "Retrieve a Policy Rule",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerRulesAPI.GetAuthorizationServerPolicyRule(apiClient.GetConfig().Context, GetAuthorizationServerPolicyRuleauthServerId, GetAuthorizationServerPolicyRulepolicyId, GetAuthorizationServerPolicyRuleruleId)

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
			if GetAuthorizationServerPolicyRuleBackup {

				idParam := GetAuthorizationServerPolicyRuleauthServerId
				err := utils.BackupObject(d, "AuthorizationServerRules", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetAuthorizationServerPolicyRuleauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().StringVarP(&GetAuthorizationServerPolicyRulepolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().StringVarP(&GetAuthorizationServerPolicyRuleruleId, "ruleId", "", "", "")
	cmd.MarkFlagRequired("ruleId")

	cmd.Flags().BoolVarP(&GetAuthorizationServerPolicyRuleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetAuthorizationServerPolicyRuleCmd := NewGetAuthorizationServerPolicyRuleCmd()
	AuthorizationServerRulesCmd.AddCommand(GetAuthorizationServerPolicyRuleCmd)
}

var (
	ReplaceAuthorizationServerPolicyRuleauthServerId string

	ReplaceAuthorizationServerPolicyRulepolicyId string

	ReplaceAuthorizationServerPolicyRuleruleId string

	ReplaceAuthorizationServerPolicyRuledata string

	ReplaceAuthorizationServerPolicyRuleBackup bool
)

func NewReplaceAuthorizationServerPolicyRuleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replaceAuthorizationServerPolicyRule",
		Long: "Replace a Policy Rule",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerRulesAPI.ReplaceAuthorizationServerPolicyRule(apiClient.GetConfig().Context, ReplaceAuthorizationServerPolicyRuleauthServerId, ReplaceAuthorizationServerPolicyRulepolicyId, ReplaceAuthorizationServerPolicyRuleruleId)

			if ReplaceAuthorizationServerPolicyRuledata != "" {
				req = req.Data(ReplaceAuthorizationServerPolicyRuledata)
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
			if ReplaceAuthorizationServerPolicyRuleBackup {

				idParam := ReplaceAuthorizationServerPolicyRuleauthServerId
				err := utils.BackupObject(d, "AuthorizationServerRules", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceAuthorizationServerPolicyRuleauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().StringVarP(&ReplaceAuthorizationServerPolicyRulepolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().StringVarP(&ReplaceAuthorizationServerPolicyRuleruleId, "ruleId", "", "", "")
	cmd.MarkFlagRequired("ruleId")

	cmd.Flags().StringVarP(&ReplaceAuthorizationServerPolicyRuledata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceAuthorizationServerPolicyRuleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceAuthorizationServerPolicyRuleCmd := NewReplaceAuthorizationServerPolicyRuleCmd()
	AuthorizationServerRulesCmd.AddCommand(ReplaceAuthorizationServerPolicyRuleCmd)
}

var (
	DeleteAuthorizationServerPolicyRuleauthServerId string

	DeleteAuthorizationServerPolicyRulepolicyId string

	DeleteAuthorizationServerPolicyRuleruleId string

	DeleteAuthorizationServerPolicyRuleBackup bool
)

func NewDeleteAuthorizationServerPolicyRuleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deleteAuthorizationServerPolicyRule",
		Long: "Delete a Policy Rule",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerRulesAPI.DeleteAuthorizationServerPolicyRule(apiClient.GetConfig().Context, DeleteAuthorizationServerPolicyRuleauthServerId, DeleteAuthorizationServerPolicyRulepolicyId, DeleteAuthorizationServerPolicyRuleruleId)

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
			if DeleteAuthorizationServerPolicyRuleBackup {

				idParam := DeleteAuthorizationServerPolicyRuleauthServerId
				err := utils.BackupObject(d, "AuthorizationServerRules", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteAuthorizationServerPolicyRuleauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().StringVarP(&DeleteAuthorizationServerPolicyRulepolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().StringVarP(&DeleteAuthorizationServerPolicyRuleruleId, "ruleId", "", "", "")
	cmd.MarkFlagRequired("ruleId")

	cmd.Flags().BoolVarP(&DeleteAuthorizationServerPolicyRuleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteAuthorizationServerPolicyRuleCmd := NewDeleteAuthorizationServerPolicyRuleCmd()
	AuthorizationServerRulesCmd.AddCommand(DeleteAuthorizationServerPolicyRuleCmd)
}

var (
	ActivateAuthorizationServerPolicyRuleauthServerId string

	ActivateAuthorizationServerPolicyRulepolicyId string

	ActivateAuthorizationServerPolicyRuleruleId string

	ActivateAuthorizationServerPolicyRuleBackup bool
)

func NewActivateAuthorizationServerPolicyRuleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "activateAuthorizationServerPolicyRule",
		Long: "Activate a Policy Rule",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerRulesAPI.ActivateAuthorizationServerPolicyRule(apiClient.GetConfig().Context, ActivateAuthorizationServerPolicyRuleauthServerId, ActivateAuthorizationServerPolicyRulepolicyId, ActivateAuthorizationServerPolicyRuleruleId)

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
			if ActivateAuthorizationServerPolicyRuleBackup {

				idParam := ActivateAuthorizationServerPolicyRuleauthServerId
				err := utils.BackupObject(d, "AuthorizationServerRules", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ActivateAuthorizationServerPolicyRuleauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().StringVarP(&ActivateAuthorizationServerPolicyRulepolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().StringVarP(&ActivateAuthorizationServerPolicyRuleruleId, "ruleId", "", "", "")
	cmd.MarkFlagRequired("ruleId")

	cmd.Flags().BoolVarP(&ActivateAuthorizationServerPolicyRuleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ActivateAuthorizationServerPolicyRuleCmd := NewActivateAuthorizationServerPolicyRuleCmd()
	AuthorizationServerRulesCmd.AddCommand(ActivateAuthorizationServerPolicyRuleCmd)
}

var (
	DeactivateAuthorizationServerPolicyRuleauthServerId string

	DeactivateAuthorizationServerPolicyRulepolicyId string

	DeactivateAuthorizationServerPolicyRuleruleId string

	DeactivateAuthorizationServerPolicyRuleBackup bool
)

func NewDeactivateAuthorizationServerPolicyRuleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deactivateAuthorizationServerPolicyRule",
		Long: "Deactivate a Policy Rule",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthorizationServerRulesAPI.DeactivateAuthorizationServerPolicyRule(apiClient.GetConfig().Context, DeactivateAuthorizationServerPolicyRuleauthServerId, DeactivateAuthorizationServerPolicyRulepolicyId, DeactivateAuthorizationServerPolicyRuleruleId)

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
			if DeactivateAuthorizationServerPolicyRuleBackup {

				idParam := DeactivateAuthorizationServerPolicyRuleauthServerId
				err := utils.BackupObject(d, "AuthorizationServerRules", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeactivateAuthorizationServerPolicyRuleauthServerId, "authServerId", "", "", "")
	cmd.MarkFlagRequired("authServerId")

	cmd.Flags().StringVarP(&DeactivateAuthorizationServerPolicyRulepolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().StringVarP(&DeactivateAuthorizationServerPolicyRuleruleId, "ruleId", "", "", "")
	cmd.MarkFlagRequired("ruleId")

	cmd.Flags().BoolVarP(&DeactivateAuthorizationServerPolicyRuleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeactivateAuthorizationServerPolicyRuleCmd := NewDeactivateAuthorizationServerPolicyRuleCmd()
	AuthorizationServerRulesCmd.AddCommand(DeactivateAuthorizationServerPolicyRuleCmd)
}
