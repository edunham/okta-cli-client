package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var AuthenticatorCmd = &cobra.Command{
	Use:  "authenticator",
	Long: "Manage AuthenticatorAPI",
}

func init() {
	rootCmd.AddCommand(AuthenticatorCmd)
}

var GetWellKnownAppAuthenticatorConfigurationBackup bool

func NewGetWellKnownAppAuthenticatorConfigurationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getWellKnownAppConfiguration",
		Long: "Retrieve the Well-Known App Authenticator Configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthenticatorAPI.GetWellKnownAppAuthenticatorConfiguration(apiClient.GetConfig().Context)

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
			if GetWellKnownAppAuthenticatorConfigurationBackup {

				err := utils.BackupObject(d, "Authenticator", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&GetWellKnownAppAuthenticatorConfigurationBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetWellKnownAppAuthenticatorConfigurationCmd := NewGetWellKnownAppAuthenticatorConfigurationCmd()
	AuthenticatorCmd.AddCommand(GetWellKnownAppAuthenticatorConfigurationCmd)
}

var (
	CreateAuthenticatordata string

	CreateAuthenticatorBackup bool
)

func NewCreateAuthenticatorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create",
		Long: "Create an Authenticator",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthenticatorAPI.CreateAuthenticator(apiClient.GetConfig().Context)

			if CreateAuthenticatordata != "" {
				req = req.Data(CreateAuthenticatordata)
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
			if CreateAuthenticatorBackup {

				err := utils.BackupObject(d, "Authenticator", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateAuthenticatordata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateAuthenticatorBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateAuthenticatorCmd := NewCreateAuthenticatorCmd()
	AuthenticatorCmd.AddCommand(CreateAuthenticatorCmd)
}

var ListAuthenticatorsBackup bool

func NewListAuthenticatorsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "lists",
		Long: "List all Authenticators",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthenticatorAPI.ListAuthenticators(apiClient.GetConfig().Context)

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
			if ListAuthenticatorsBackup {

				err := utils.BackupObject(d, "Authenticator", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListAuthenticatorsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListAuthenticatorsCmd := NewListAuthenticatorsCmd()
	AuthenticatorCmd.AddCommand(ListAuthenticatorsCmd)
}

var (
	GetAuthenticatorauthenticatorId string

	GetAuthenticatorBackup bool
)

func NewGetAuthenticatorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "get",
		Long: "Retrieve an Authenticator",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthenticatorAPI.GetAuthenticator(apiClient.GetConfig().Context, GetAuthenticatorauthenticatorId)

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
			if GetAuthenticatorBackup {

				idParam := GetAuthenticatorauthenticatorId
				err := utils.BackupObject(d, "Authenticator", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetAuthenticatorauthenticatorId, "authenticatorId", "", "", "")
	cmd.MarkFlagRequired("authenticatorId")

	cmd.Flags().BoolVarP(&GetAuthenticatorBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetAuthenticatorCmd := NewGetAuthenticatorCmd()
	AuthenticatorCmd.AddCommand(GetAuthenticatorCmd)
}

var (
	ReplaceAuthenticatorauthenticatorId string

	ReplaceAuthenticatordata string

	ReplaceAuthenticatorBackup bool
)

func NewReplaceAuthenticatorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replace",
		Long: "Replace an Authenticator",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthenticatorAPI.ReplaceAuthenticator(apiClient.GetConfig().Context, ReplaceAuthenticatorauthenticatorId)

			if ReplaceAuthenticatordata != "" {
				req = req.Data(ReplaceAuthenticatordata)
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
			if ReplaceAuthenticatorBackup {

				idParam := ReplaceAuthenticatorauthenticatorId
				err := utils.BackupObject(d, "Authenticator", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceAuthenticatorauthenticatorId, "authenticatorId", "", "", "")
	cmd.MarkFlagRequired("authenticatorId")

	cmd.Flags().StringVarP(&ReplaceAuthenticatordata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceAuthenticatorBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceAuthenticatorCmd := NewReplaceAuthenticatorCmd()
	AuthenticatorCmd.AddCommand(ReplaceAuthenticatorCmd)
}

var (
	ActivateAuthenticatorauthenticatorId string

	ActivateAuthenticatorBackup bool
)

func NewActivateAuthenticatorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "activate",
		Long: "Activate an Authenticator",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthenticatorAPI.ActivateAuthenticator(apiClient.GetConfig().Context, ActivateAuthenticatorauthenticatorId)

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
			if ActivateAuthenticatorBackup {

				idParam := ActivateAuthenticatorauthenticatorId
				err := utils.BackupObject(d, "Authenticator", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ActivateAuthenticatorauthenticatorId, "authenticatorId", "", "", "")
	cmd.MarkFlagRequired("authenticatorId")

	cmd.Flags().BoolVarP(&ActivateAuthenticatorBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ActivateAuthenticatorCmd := NewActivateAuthenticatorCmd()
	AuthenticatorCmd.AddCommand(ActivateAuthenticatorCmd)
}

var (
	DeactivateAuthenticatorauthenticatorId string

	DeactivateAuthenticatorBackup bool
)

func NewDeactivateAuthenticatorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deactivate",
		Long: "Deactivate an Authenticator",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthenticatorAPI.DeactivateAuthenticator(apiClient.GetConfig().Context, DeactivateAuthenticatorauthenticatorId)

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
			if DeactivateAuthenticatorBackup {

				idParam := DeactivateAuthenticatorauthenticatorId
				err := utils.BackupObject(d, "Authenticator", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeactivateAuthenticatorauthenticatorId, "authenticatorId", "", "", "")
	cmd.MarkFlagRequired("authenticatorId")

	cmd.Flags().BoolVarP(&DeactivateAuthenticatorBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeactivateAuthenticatorCmd := NewDeactivateAuthenticatorCmd()
	AuthenticatorCmd.AddCommand(DeactivateAuthenticatorCmd)
}

var (
	ListAuthenticatorMethodsauthenticatorId string

	ListAuthenticatorMethodsBackup bool
)

func NewListAuthenticatorMethodsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listMethods",
		Long: "List all Methods of an Authenticator",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthenticatorAPI.ListAuthenticatorMethods(apiClient.GetConfig().Context, ListAuthenticatorMethodsauthenticatorId)

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
			if ListAuthenticatorMethodsBackup {

				idParam := ListAuthenticatorMethodsauthenticatorId
				err := utils.BackupObject(d, "Authenticator", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListAuthenticatorMethodsauthenticatorId, "authenticatorId", "", "", "")
	cmd.MarkFlagRequired("authenticatorId")

	cmd.Flags().BoolVarP(&ListAuthenticatorMethodsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListAuthenticatorMethodsCmd := NewListAuthenticatorMethodsCmd()
	AuthenticatorCmd.AddCommand(ListAuthenticatorMethodsCmd)
}

var (
	GetAuthenticatorMethodauthenticatorId string

	GetAuthenticatorMethodmethodType string

	GetAuthenticatorMethodBackup bool
)

func NewGetAuthenticatorMethodCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getMethod",
		Long: "Retrieve a Method",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthenticatorAPI.GetAuthenticatorMethod(apiClient.GetConfig().Context, GetAuthenticatorMethodauthenticatorId, GetAuthenticatorMethodmethodType)

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
			if GetAuthenticatorMethodBackup {

				idParam := GetAuthenticatorMethodauthenticatorId
				err := utils.BackupObject(d, "Authenticator", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetAuthenticatorMethodauthenticatorId, "authenticatorId", "", "", "")
	cmd.MarkFlagRequired("authenticatorId")

	cmd.Flags().StringVarP(&GetAuthenticatorMethodmethodType, "methodType", "", "", "")
	cmd.MarkFlagRequired("methodType")

	cmd.Flags().BoolVarP(&GetAuthenticatorMethodBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetAuthenticatorMethodCmd := NewGetAuthenticatorMethodCmd()
	AuthenticatorCmd.AddCommand(GetAuthenticatorMethodCmd)
}

var (
	ReplaceAuthenticatorMethodauthenticatorId string

	ReplaceAuthenticatorMethodmethodType string

	ReplaceAuthenticatorMethoddata string

	ReplaceAuthenticatorMethodBackup bool
)

func NewReplaceAuthenticatorMethodCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replaceMethod",
		Long: "Replace a Method",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthenticatorAPI.ReplaceAuthenticatorMethod(apiClient.GetConfig().Context, ReplaceAuthenticatorMethodauthenticatorId, ReplaceAuthenticatorMethodmethodType)

			if ReplaceAuthenticatorMethoddata != "" {
				req = req.Data(ReplaceAuthenticatorMethoddata)
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
			if ReplaceAuthenticatorMethodBackup {

				idParam := ReplaceAuthenticatorMethodauthenticatorId
				err := utils.BackupObject(d, "Authenticator", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceAuthenticatorMethodauthenticatorId, "authenticatorId", "", "", "")
	cmd.MarkFlagRequired("authenticatorId")

	cmd.Flags().StringVarP(&ReplaceAuthenticatorMethodmethodType, "methodType", "", "", "")
	cmd.MarkFlagRequired("methodType")

	cmd.Flags().StringVarP(&ReplaceAuthenticatorMethoddata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceAuthenticatorMethodBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceAuthenticatorMethodCmd := NewReplaceAuthenticatorMethodCmd()
	AuthenticatorCmd.AddCommand(ReplaceAuthenticatorMethodCmd)
}

var (
	ActivateAuthenticatorMethodauthenticatorId string

	ActivateAuthenticatorMethodmethodType string

	ActivateAuthenticatorMethodBackup bool
)

func NewActivateAuthenticatorMethodCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "activateMethod",
		Long: "Activate an Authenticator Method",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthenticatorAPI.ActivateAuthenticatorMethod(apiClient.GetConfig().Context, ActivateAuthenticatorMethodauthenticatorId, ActivateAuthenticatorMethodmethodType)

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
			if ActivateAuthenticatorMethodBackup {

				idParam := ActivateAuthenticatorMethodauthenticatorId
				err := utils.BackupObject(d, "Authenticator", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ActivateAuthenticatorMethodauthenticatorId, "authenticatorId", "", "", "")
	cmd.MarkFlagRequired("authenticatorId")

	cmd.Flags().StringVarP(&ActivateAuthenticatorMethodmethodType, "methodType", "", "", "")
	cmd.MarkFlagRequired("methodType")

	cmd.Flags().BoolVarP(&ActivateAuthenticatorMethodBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ActivateAuthenticatorMethodCmd := NewActivateAuthenticatorMethodCmd()
	AuthenticatorCmd.AddCommand(ActivateAuthenticatorMethodCmd)
}

var (
	DeactivateAuthenticatorMethodauthenticatorId string

	DeactivateAuthenticatorMethodmethodType string

	DeactivateAuthenticatorMethodBackup bool
)

func NewDeactivateAuthenticatorMethodCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deactivateMethod",
		Long: "Deactivate an Authenticator Method",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AuthenticatorAPI.DeactivateAuthenticatorMethod(apiClient.GetConfig().Context, DeactivateAuthenticatorMethodauthenticatorId, DeactivateAuthenticatorMethodmethodType)

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
			if DeactivateAuthenticatorMethodBackup {

				idParam := DeactivateAuthenticatorMethodauthenticatorId
				err := utils.BackupObject(d, "Authenticator", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeactivateAuthenticatorMethodauthenticatorId, "authenticatorId", "", "", "")
	cmd.MarkFlagRequired("authenticatorId")

	cmd.Flags().StringVarP(&DeactivateAuthenticatorMethodmethodType, "methodType", "", "", "")
	cmd.MarkFlagRequired("methodType")

	cmd.Flags().BoolVarP(&DeactivateAuthenticatorMethodBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeactivateAuthenticatorMethodCmd := NewDeactivateAuthenticatorMethodCmd()
	AuthenticatorCmd.AddCommand(DeactivateAuthenticatorMethodCmd)
}
