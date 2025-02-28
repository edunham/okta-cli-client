package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var IdentityProviderCmd = &cobra.Command{
	Use:  "identityProvider",
	Long: "Manage IdentityProviderAPI",
}

func init() {
	rootCmd.AddCommand(IdentityProviderCmd)
}

var (
	CreateIdentityProviderdata string

	CreateIdentityProviderBackup bool
)

func NewCreateIdentityProviderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create",
		Long: "Create an Identity Provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentityProviderAPI.CreateIdentityProvider(apiClient.GetConfig().Context)

			if CreateIdentityProviderdata != "" {
				req = req.Data(CreateIdentityProviderdata)
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
			if CreateIdentityProviderBackup {

				err := utils.BackupObject(d, "IdentityProvider", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateIdentityProviderdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateIdentityProviderBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateIdentityProviderCmd := NewCreateIdentityProviderCmd()
	IdentityProviderCmd.AddCommand(CreateIdentityProviderCmd)
}

var ListIdentityProvidersBackup bool

func NewListIdentityProvidersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "lists",
		Long: "List all Identity Providers",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentityProviderAPI.ListIdentityProviders(apiClient.GetConfig().Context)

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
			if ListIdentityProvidersBackup {

				err := utils.BackupObject(d, "IdentityProvider", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListIdentityProvidersBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListIdentityProvidersCmd := NewListIdentityProvidersCmd()
	IdentityProviderCmd.AddCommand(ListIdentityProvidersCmd)
}

var (
	CreateIdentityProviderKeydata string

	CreateIdentityProviderKeyBackup bool
)

func NewCreateIdentityProviderKeyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "createKey",
		Long: "Create an X.509 Certificate Public Key",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentityProviderAPI.CreateIdentityProviderKey(apiClient.GetConfig().Context)

			if CreateIdentityProviderKeydata != "" {
				req = req.Data(CreateIdentityProviderKeydata)
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
			if CreateIdentityProviderKeyBackup {

				err := utils.BackupObject(d, "IdentityProvider", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateIdentityProviderKeydata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateIdentityProviderKeyBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateIdentityProviderKeyCmd := NewCreateIdentityProviderKeyCmd()
	IdentityProviderCmd.AddCommand(CreateIdentityProviderKeyCmd)
}

var ListIdentityProviderKeysBackup bool

func NewListIdentityProviderKeysCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listKeys",
		Long: "List all Credential Keys",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentityProviderAPI.ListIdentityProviderKeys(apiClient.GetConfig().Context)

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
			if ListIdentityProviderKeysBackup {

				err := utils.BackupObject(d, "IdentityProvider", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListIdentityProviderKeysBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListIdentityProviderKeysCmd := NewListIdentityProviderKeysCmd()
	IdentityProviderCmd.AddCommand(ListIdentityProviderKeysCmd)
}

var (
	GetIdentityProviderKeyidpKeyId string

	GetIdentityProviderKeyBackup bool
)

func NewGetIdentityProviderKeyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getKey",
		Long: "Retrieve an Credential Key",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentityProviderAPI.GetIdentityProviderKey(apiClient.GetConfig().Context, GetIdentityProviderKeyidpKeyId)

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
			if GetIdentityProviderKeyBackup {

				idParam := GetIdentityProviderKeyidpKeyId
				err := utils.BackupObject(d, "IdentityProvider", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetIdentityProviderKeyidpKeyId, "idpKeyId", "", "", "")
	cmd.MarkFlagRequired("idpKeyId")

	cmd.Flags().BoolVarP(&GetIdentityProviderKeyBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetIdentityProviderKeyCmd := NewGetIdentityProviderKeyCmd()
	IdentityProviderCmd.AddCommand(GetIdentityProviderKeyCmd)
}

var (
	DeleteIdentityProviderKeyidpKeyId string

	DeleteIdentityProviderKeyBackup bool
)

func NewDeleteIdentityProviderKeyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deleteKey",
		Long: "Delete a Signing Credential Key",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentityProviderAPI.DeleteIdentityProviderKey(apiClient.GetConfig().Context, DeleteIdentityProviderKeyidpKeyId)

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
			if DeleteIdentityProviderKeyBackup {

				idParam := DeleteIdentityProviderKeyidpKeyId
				err := utils.BackupObject(d, "IdentityProvider", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteIdentityProviderKeyidpKeyId, "idpKeyId", "", "", "")
	cmd.MarkFlagRequired("idpKeyId")

	cmd.Flags().BoolVarP(&DeleteIdentityProviderKeyBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteIdentityProviderKeyCmd := NewDeleteIdentityProviderKeyCmd()
	IdentityProviderCmd.AddCommand(DeleteIdentityProviderKeyCmd)
}

var (
	GetIdentityProvideridpId string

	GetIdentityProviderBackup bool
)

func NewGetIdentityProviderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "get",
		Long: "Retrieve an Identity Provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentityProviderAPI.GetIdentityProvider(apiClient.GetConfig().Context, GetIdentityProvideridpId)

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
			if GetIdentityProviderBackup {

				idParam := GetIdentityProvideridpId
				err := utils.BackupObject(d, "IdentityProvider", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetIdentityProvideridpId, "idpId", "", "", "")
	cmd.MarkFlagRequired("idpId")

	cmd.Flags().BoolVarP(&GetIdentityProviderBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetIdentityProviderCmd := NewGetIdentityProviderCmd()
	IdentityProviderCmd.AddCommand(GetIdentityProviderCmd)
}

var (
	ReplaceIdentityProvideridpId string

	ReplaceIdentityProviderdata string

	ReplaceIdentityProviderBackup bool
)

func NewReplaceIdentityProviderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replace",
		Long: "Replace an Identity Provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentityProviderAPI.ReplaceIdentityProvider(apiClient.GetConfig().Context, ReplaceIdentityProvideridpId)

			if ReplaceIdentityProviderdata != "" {
				req = req.Data(ReplaceIdentityProviderdata)
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
			if ReplaceIdentityProviderBackup {

				idParam := ReplaceIdentityProvideridpId
				err := utils.BackupObject(d, "IdentityProvider", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceIdentityProvideridpId, "idpId", "", "", "")
	cmd.MarkFlagRequired("idpId")

	cmd.Flags().StringVarP(&ReplaceIdentityProviderdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceIdentityProviderBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceIdentityProviderCmd := NewReplaceIdentityProviderCmd()
	IdentityProviderCmd.AddCommand(ReplaceIdentityProviderCmd)
}

var (
	DeleteIdentityProvideridpId string

	DeleteIdentityProviderBackup bool
)

func NewDeleteIdentityProviderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete",
		Long: "Delete an Identity Provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentityProviderAPI.DeleteIdentityProvider(apiClient.GetConfig().Context, DeleteIdentityProvideridpId)

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
			if DeleteIdentityProviderBackup {

				idParam := DeleteIdentityProvideridpId
				err := utils.BackupObject(d, "IdentityProvider", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteIdentityProvideridpId, "idpId", "", "", "")
	cmd.MarkFlagRequired("idpId")

	cmd.Flags().BoolVarP(&DeleteIdentityProviderBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteIdentityProviderCmd := NewDeleteIdentityProviderCmd()
	IdentityProviderCmd.AddCommand(DeleteIdentityProviderCmd)
}

var (
	GenerateCsrForIdentityProvideridpId string

	GenerateCsrForIdentityProviderdata string

	GenerateCsrForIdentityProviderBackup bool
)

func NewGenerateCsrForIdentityProviderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "generateCsrFor",
		Long: "Generate a Certificate Signing Request",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentityProviderAPI.GenerateCsrForIdentityProvider(apiClient.GetConfig().Context, GenerateCsrForIdentityProvideridpId)

			if GenerateCsrForIdentityProviderdata != "" {
				req = req.Data(GenerateCsrForIdentityProviderdata)
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
			if GenerateCsrForIdentityProviderBackup {

				idParam := GenerateCsrForIdentityProvideridpId
				err := utils.BackupObject(d, "IdentityProvider", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GenerateCsrForIdentityProvideridpId, "idpId", "", "", "")
	cmd.MarkFlagRequired("idpId")

	cmd.Flags().StringVarP(&GenerateCsrForIdentityProviderdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&GenerateCsrForIdentityProviderBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GenerateCsrForIdentityProviderCmd := NewGenerateCsrForIdentityProviderCmd()
	IdentityProviderCmd.AddCommand(GenerateCsrForIdentityProviderCmd)
}

var (
	ListCsrsForIdentityProvideridpId string

	ListCsrsForIdentityProviderBackup bool
)

func NewListCsrsForIdentityProviderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listCsrsFor",
		Long: "List all Certificate Signing Requests",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentityProviderAPI.ListCsrsForIdentityProvider(apiClient.GetConfig().Context, ListCsrsForIdentityProvideridpId)

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
			if ListCsrsForIdentityProviderBackup {

				idParam := ListCsrsForIdentityProvideridpId
				err := utils.BackupObject(d, "IdentityProvider", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListCsrsForIdentityProvideridpId, "idpId", "", "", "")
	cmd.MarkFlagRequired("idpId")

	cmd.Flags().BoolVarP(&ListCsrsForIdentityProviderBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListCsrsForIdentityProviderCmd := NewListCsrsForIdentityProviderCmd()
	IdentityProviderCmd.AddCommand(ListCsrsForIdentityProviderCmd)
}

var (
	GetCsrForIdentityProvideridpId string

	GetCsrForIdentityProvideridpCsrId string

	GetCsrForIdentityProviderBackup bool
)

func NewGetCsrForIdentityProviderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getCsrFor",
		Long: "Retrieve a Certificate Signing Request",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentityProviderAPI.GetCsrForIdentityProvider(apiClient.GetConfig().Context, GetCsrForIdentityProvideridpId, GetCsrForIdentityProvideridpCsrId)

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
			if GetCsrForIdentityProviderBackup {

				idParam := GetCsrForIdentityProvideridpId
				err := utils.BackupObject(d, "IdentityProvider", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetCsrForIdentityProvideridpId, "idpId", "", "", "")
	cmd.MarkFlagRequired("idpId")

	cmd.Flags().StringVarP(&GetCsrForIdentityProvideridpCsrId, "idpCsrId", "", "", "")
	cmd.MarkFlagRequired("idpCsrId")

	cmd.Flags().BoolVarP(&GetCsrForIdentityProviderBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetCsrForIdentityProviderCmd := NewGetCsrForIdentityProviderCmd()
	IdentityProviderCmd.AddCommand(GetCsrForIdentityProviderCmd)
}

var (
	RevokeCsrForIdentityProvideridpId string

	RevokeCsrForIdentityProvideridpCsrId string

	RevokeCsrForIdentityProviderBackup bool
)

func NewRevokeCsrForIdentityProviderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "revokeCsrFor",
		Long: "Revoke a Certificate Signing Request",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentityProviderAPI.RevokeCsrForIdentityProvider(apiClient.GetConfig().Context, RevokeCsrForIdentityProvideridpId, RevokeCsrForIdentityProvideridpCsrId)

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
			if RevokeCsrForIdentityProviderBackup {

				idParam := RevokeCsrForIdentityProvideridpId
				err := utils.BackupObject(d, "IdentityProvider", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&RevokeCsrForIdentityProvideridpId, "idpId", "", "", "")
	cmd.MarkFlagRequired("idpId")

	cmd.Flags().StringVarP(&RevokeCsrForIdentityProvideridpCsrId, "idpCsrId", "", "", "")
	cmd.MarkFlagRequired("idpCsrId")

	cmd.Flags().BoolVarP(&RevokeCsrForIdentityProviderBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	RevokeCsrForIdentityProviderCmd := NewRevokeCsrForIdentityProviderCmd()
	IdentityProviderCmd.AddCommand(RevokeCsrForIdentityProviderCmd)
}

var (
	PublishCsrForIdentityProvideridpId string

	PublishCsrForIdentityProvideridpCsrId string

	PublishCsrForIdentityProviderdata string

	PublishCsrForIdentityProviderBackup bool
)

func NewPublishCsrForIdentityProviderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "publishCsrFor",
		Long: "Publish a Certificate Signing Request",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentityProviderAPI.PublishCsrForIdentityProvider(apiClient.GetConfig().Context, PublishCsrForIdentityProvideridpId, PublishCsrForIdentityProvideridpCsrId)

			if PublishCsrForIdentityProviderdata != "" {
				req = req.Data(PublishCsrForIdentityProviderdata)
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
			if PublishCsrForIdentityProviderBackup {

				idParam := PublishCsrForIdentityProvideridpId
				err := utils.BackupObject(d, "IdentityProvider", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&PublishCsrForIdentityProvideridpId, "idpId", "", "", "")
	cmd.MarkFlagRequired("idpId")

	cmd.Flags().StringVarP(&PublishCsrForIdentityProvideridpCsrId, "idpCsrId", "", "", "")
	cmd.MarkFlagRequired("idpCsrId")

	cmd.Flags().StringVarP(&PublishCsrForIdentityProviderdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&PublishCsrForIdentityProviderBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	PublishCsrForIdentityProviderCmd := NewPublishCsrForIdentityProviderCmd()
	IdentityProviderCmd.AddCommand(PublishCsrForIdentityProviderCmd)
}

var (
	ListIdentityProviderSigningKeysidpId string

	ListIdentityProviderSigningKeysBackup bool
)

func NewListIdentityProviderSigningKeysCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listSigningKeys",
		Long: "List all Signing Credential Keys",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentityProviderAPI.ListIdentityProviderSigningKeys(apiClient.GetConfig().Context, ListIdentityProviderSigningKeysidpId)

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
			if ListIdentityProviderSigningKeysBackup {

				idParam := ListIdentityProviderSigningKeysidpId
				err := utils.BackupObject(d, "IdentityProvider", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListIdentityProviderSigningKeysidpId, "idpId", "", "", "")
	cmd.MarkFlagRequired("idpId")

	cmd.Flags().BoolVarP(&ListIdentityProviderSigningKeysBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListIdentityProviderSigningKeysCmd := NewListIdentityProviderSigningKeysCmd()
	IdentityProviderCmd.AddCommand(ListIdentityProviderSigningKeysCmd)
}

var (
	GenerateIdentityProviderSigningKeyidpId string

	GenerateIdentityProviderSigningKeyBackup bool
)

func NewGenerateIdentityProviderSigningKeyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "generateSigningKey",
		Long: "Generate a new Signing Credential Key",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentityProviderAPI.GenerateIdentityProviderSigningKey(apiClient.GetConfig().Context, GenerateIdentityProviderSigningKeyidpId)

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
			if GenerateIdentityProviderSigningKeyBackup {

				idParam := GenerateIdentityProviderSigningKeyidpId
				err := utils.BackupObject(d, "IdentityProvider", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GenerateIdentityProviderSigningKeyidpId, "idpId", "", "", "")
	cmd.MarkFlagRequired("idpId")

	cmd.Flags().BoolVarP(&GenerateIdentityProviderSigningKeyBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GenerateIdentityProviderSigningKeyCmd := NewGenerateIdentityProviderSigningKeyCmd()
	IdentityProviderCmd.AddCommand(GenerateIdentityProviderSigningKeyCmd)
}

var (
	GetIdentityProviderSigningKeyidpId string

	GetIdentityProviderSigningKeyidpKeyId string

	GetIdentityProviderSigningKeyBackup bool
)

func NewGetIdentityProviderSigningKeyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getSigningKey",
		Long: "Retrieve a Signing Credential Key",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentityProviderAPI.GetIdentityProviderSigningKey(apiClient.GetConfig().Context, GetIdentityProviderSigningKeyidpId, GetIdentityProviderSigningKeyidpKeyId)

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
			if GetIdentityProviderSigningKeyBackup {

				idParam := GetIdentityProviderSigningKeyidpId
				err := utils.BackupObject(d, "IdentityProvider", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetIdentityProviderSigningKeyidpId, "idpId", "", "", "")
	cmd.MarkFlagRequired("idpId")

	cmd.Flags().StringVarP(&GetIdentityProviderSigningKeyidpKeyId, "idpKeyId", "", "", "")
	cmd.MarkFlagRequired("idpKeyId")

	cmd.Flags().BoolVarP(&GetIdentityProviderSigningKeyBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetIdentityProviderSigningKeyCmd := NewGetIdentityProviderSigningKeyCmd()
	IdentityProviderCmd.AddCommand(GetIdentityProviderSigningKeyCmd)
}

var (
	CloneIdentityProviderKeyidpId string

	CloneIdentityProviderKeyidpKeyId string

	CloneIdentityProviderKeyBackup bool
)

func NewCloneIdentityProviderKeyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "cloneKey",
		Long: "Clone a Signing Credential Key",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentityProviderAPI.CloneIdentityProviderKey(apiClient.GetConfig().Context, CloneIdentityProviderKeyidpId, CloneIdentityProviderKeyidpKeyId)

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
			if CloneIdentityProviderKeyBackup {

				idParam := CloneIdentityProviderKeyidpId
				err := utils.BackupObject(d, "IdentityProvider", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CloneIdentityProviderKeyidpId, "idpId", "", "", "")
	cmd.MarkFlagRequired("idpId")

	cmd.Flags().StringVarP(&CloneIdentityProviderKeyidpKeyId, "idpKeyId", "", "", "")
	cmd.MarkFlagRequired("idpKeyId")

	cmd.Flags().BoolVarP(&CloneIdentityProviderKeyBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CloneIdentityProviderKeyCmd := NewCloneIdentityProviderKeyCmd()
	IdentityProviderCmd.AddCommand(CloneIdentityProviderKeyCmd)
}

var (
	ActivateIdentityProvideridpId string

	ActivateIdentityProviderBackup bool
)

func NewActivateIdentityProviderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "activate",
		Long: "Activate an Identity Provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentityProviderAPI.ActivateIdentityProvider(apiClient.GetConfig().Context, ActivateIdentityProvideridpId)

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
			if ActivateIdentityProviderBackup {

				idParam := ActivateIdentityProvideridpId
				err := utils.BackupObject(d, "IdentityProvider", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ActivateIdentityProvideridpId, "idpId", "", "", "")
	cmd.MarkFlagRequired("idpId")

	cmd.Flags().BoolVarP(&ActivateIdentityProviderBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ActivateIdentityProviderCmd := NewActivateIdentityProviderCmd()
	IdentityProviderCmd.AddCommand(ActivateIdentityProviderCmd)
}

var (
	DeactivateIdentityProvideridpId string

	DeactivateIdentityProviderBackup bool
)

func NewDeactivateIdentityProviderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deactivate",
		Long: "Deactivate an Identity Provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentityProviderAPI.DeactivateIdentityProvider(apiClient.GetConfig().Context, DeactivateIdentityProvideridpId)

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
			if DeactivateIdentityProviderBackup {

				idParam := DeactivateIdentityProvideridpId
				err := utils.BackupObject(d, "IdentityProvider", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeactivateIdentityProvideridpId, "idpId", "", "", "")
	cmd.MarkFlagRequired("idpId")

	cmd.Flags().BoolVarP(&DeactivateIdentityProviderBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeactivateIdentityProviderCmd := NewDeactivateIdentityProviderCmd()
	IdentityProviderCmd.AddCommand(DeactivateIdentityProviderCmd)
}

var (
	ListIdentityProviderApplicationUsersidpId string

	ListIdentityProviderApplicationUsersBackup bool
)

func NewListIdentityProviderApplicationUsersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listApplicationUsers",
		Long: "List all Users",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentityProviderAPI.ListIdentityProviderApplicationUsers(apiClient.GetConfig().Context, ListIdentityProviderApplicationUsersidpId)

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
			if ListIdentityProviderApplicationUsersBackup {

				idParam := ListIdentityProviderApplicationUsersidpId
				err := utils.BackupObject(d, "IdentityProvider", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListIdentityProviderApplicationUsersidpId, "idpId", "", "", "")
	cmd.MarkFlagRequired("idpId")

	cmd.Flags().BoolVarP(&ListIdentityProviderApplicationUsersBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListIdentityProviderApplicationUsersCmd := NewListIdentityProviderApplicationUsersCmd()
	IdentityProviderCmd.AddCommand(ListIdentityProviderApplicationUsersCmd)
}

var (
	LinkUserToIdentityProvideridpId string

	LinkUserToIdentityProvideruserId string

	LinkUserToIdentityProviderdata string

	LinkUserToIdentityProviderBackup bool
)

func NewLinkUserToIdentityProviderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "linkUserTo",
		Long: "Link a User to a Social IdP",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentityProviderAPI.LinkUserToIdentityProvider(apiClient.GetConfig().Context, LinkUserToIdentityProvideridpId, LinkUserToIdentityProvideruserId)

			if LinkUserToIdentityProviderdata != "" {
				req = req.Data(LinkUserToIdentityProviderdata)
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
			if LinkUserToIdentityProviderBackup {

				idParam := LinkUserToIdentityProvideridpId
				err := utils.BackupObject(d, "IdentityProvider", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&LinkUserToIdentityProvideridpId, "idpId", "", "", "")
	cmd.MarkFlagRequired("idpId")

	cmd.Flags().StringVarP(&LinkUserToIdentityProvideruserId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	cmd.Flags().StringVarP(&LinkUserToIdentityProviderdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&LinkUserToIdentityProviderBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	LinkUserToIdentityProviderCmd := NewLinkUserToIdentityProviderCmd()
	IdentityProviderCmd.AddCommand(LinkUserToIdentityProviderCmd)
}

var (
	GetIdentityProviderApplicationUseridpId string

	GetIdentityProviderApplicationUseruserId string

	GetIdentityProviderApplicationUserBackup bool
)

func NewGetIdentityProviderApplicationUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getApplicationUser",
		Long: "Retrieve a User",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentityProviderAPI.GetIdentityProviderApplicationUser(apiClient.GetConfig().Context, GetIdentityProviderApplicationUseridpId, GetIdentityProviderApplicationUseruserId)

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
			if GetIdentityProviderApplicationUserBackup {

				idParam := GetIdentityProviderApplicationUseridpId
				err := utils.BackupObject(d, "IdentityProvider", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetIdentityProviderApplicationUseridpId, "idpId", "", "", "")
	cmd.MarkFlagRequired("idpId")

	cmd.Flags().StringVarP(&GetIdentityProviderApplicationUseruserId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	cmd.Flags().BoolVarP(&GetIdentityProviderApplicationUserBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetIdentityProviderApplicationUserCmd := NewGetIdentityProviderApplicationUserCmd()
	IdentityProviderCmd.AddCommand(GetIdentityProviderApplicationUserCmd)
}

var (
	UnlinkUserFromIdentityProvideridpId string

	UnlinkUserFromIdentityProvideruserId string

	UnlinkUserFromIdentityProviderBackup bool
)

func NewUnlinkUserFromIdentityProviderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "unlinkUserFrom",
		Long: "Unlink a User from IdP",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentityProviderAPI.UnlinkUserFromIdentityProvider(apiClient.GetConfig().Context, UnlinkUserFromIdentityProvideridpId, UnlinkUserFromIdentityProvideruserId)

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
			if UnlinkUserFromIdentityProviderBackup {

				idParam := UnlinkUserFromIdentityProvideridpId
				err := utils.BackupObject(d, "IdentityProvider", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UnlinkUserFromIdentityProvideridpId, "idpId", "", "", "")
	cmd.MarkFlagRequired("idpId")

	cmd.Flags().StringVarP(&UnlinkUserFromIdentityProvideruserId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	cmd.Flags().BoolVarP(&UnlinkUserFromIdentityProviderBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UnlinkUserFromIdentityProviderCmd := NewUnlinkUserFromIdentityProviderCmd()
	IdentityProviderCmd.AddCommand(UnlinkUserFromIdentityProviderCmd)
}

var (
	ListSocialAuthTokensidpId string

	ListSocialAuthTokensuserId string

	ListSocialAuthTokensBackup bool
)

func NewListSocialAuthTokensCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listSocialAuthTokens",
		Long: "List all Tokens from a OIDC Identity Provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentityProviderAPI.ListSocialAuthTokens(apiClient.GetConfig().Context, ListSocialAuthTokensidpId, ListSocialAuthTokensuserId)

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
			if ListSocialAuthTokensBackup {

				idParam := ListSocialAuthTokensidpId
				err := utils.BackupObject(d, "IdentityProvider", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListSocialAuthTokensidpId, "idpId", "", "", "")
	cmd.MarkFlagRequired("idpId")

	cmd.Flags().StringVarP(&ListSocialAuthTokensuserId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	cmd.Flags().BoolVarP(&ListSocialAuthTokensBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListSocialAuthTokensCmd := NewListSocialAuthTokensCmd()
	IdentityProviderCmd.AddCommand(ListSocialAuthTokensCmd)
}
