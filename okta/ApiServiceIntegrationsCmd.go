package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var ApiServiceIntegrationsCmd = &cobra.Command{
	Use:  "apiServiceIntegrations",
	Long: "Manage ApiServiceIntegrationsAPI",
}

func init() {
	rootCmd.AddCommand(ApiServiceIntegrationsCmd)
}

var (
	CreateApiServiceIntegrationInstancedata string

	CreateApiServiceIntegrationInstanceBackup bool
)

func NewCreateApiServiceIntegrationInstanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "createApiServiceIntegrationInstance",
		Long: "Create an API Service Integration instance",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApiServiceIntegrationsAPI.CreateApiServiceIntegrationInstance(apiClient.GetConfig().Context)

			if CreateApiServiceIntegrationInstancedata != "" {
				req = req.Data(CreateApiServiceIntegrationInstancedata)
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
			if CreateApiServiceIntegrationInstanceBackup {

				err := utils.BackupObject(d, "ApiServiceIntegrations", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateApiServiceIntegrationInstancedata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateApiServiceIntegrationInstanceBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateApiServiceIntegrationInstanceCmd := NewCreateApiServiceIntegrationInstanceCmd()
	ApiServiceIntegrationsCmd.AddCommand(CreateApiServiceIntegrationInstanceCmd)
}

var ListApiServiceIntegrationInstancesBackup bool

func NewListApiServiceIntegrationInstancesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listApiServiceIntegrationInstances",
		Long: "List all API Service Integration instances",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApiServiceIntegrationsAPI.ListApiServiceIntegrationInstances(apiClient.GetConfig().Context)

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
			if ListApiServiceIntegrationInstancesBackup {

				err := utils.BackupObject(d, "ApiServiceIntegrations", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListApiServiceIntegrationInstancesBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListApiServiceIntegrationInstancesCmd := NewListApiServiceIntegrationInstancesCmd()
	ApiServiceIntegrationsCmd.AddCommand(ListApiServiceIntegrationInstancesCmd)
}

var (
	GetApiServiceIntegrationInstanceapiServiceId string

	GetApiServiceIntegrationInstanceBackup bool
)

func NewGetApiServiceIntegrationInstanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getApiServiceIntegrationInstance",
		Long: "Retrieve an API Service Integration instance",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApiServiceIntegrationsAPI.GetApiServiceIntegrationInstance(apiClient.GetConfig().Context, GetApiServiceIntegrationInstanceapiServiceId)

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
			if GetApiServiceIntegrationInstanceBackup {

				idParam := GetApiServiceIntegrationInstanceapiServiceId
				err := utils.BackupObject(d, "ApiServiceIntegrations", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetApiServiceIntegrationInstanceapiServiceId, "apiServiceId", "", "", "")
	cmd.MarkFlagRequired("apiServiceId")

	cmd.Flags().BoolVarP(&GetApiServiceIntegrationInstanceBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetApiServiceIntegrationInstanceCmd := NewGetApiServiceIntegrationInstanceCmd()
	ApiServiceIntegrationsCmd.AddCommand(GetApiServiceIntegrationInstanceCmd)
}

var (
	DeleteApiServiceIntegrationInstanceapiServiceId string

	DeleteApiServiceIntegrationInstanceBackup bool
)

func NewDeleteApiServiceIntegrationInstanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deleteApiServiceIntegrationInstance",
		Long: "Delete an API Service Integration instance",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApiServiceIntegrationsAPI.DeleteApiServiceIntegrationInstance(apiClient.GetConfig().Context, DeleteApiServiceIntegrationInstanceapiServiceId)

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
			if DeleteApiServiceIntegrationInstanceBackup {

				idParam := DeleteApiServiceIntegrationInstanceapiServiceId
				err := utils.BackupObject(d, "ApiServiceIntegrations", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteApiServiceIntegrationInstanceapiServiceId, "apiServiceId", "", "", "")
	cmd.MarkFlagRequired("apiServiceId")

	cmd.Flags().BoolVarP(&DeleteApiServiceIntegrationInstanceBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteApiServiceIntegrationInstanceCmd := NewDeleteApiServiceIntegrationInstanceCmd()
	ApiServiceIntegrationsCmd.AddCommand(DeleteApiServiceIntegrationInstanceCmd)
}

var (
	CreateApiServiceIntegrationInstanceSecretapiServiceId string

	CreateApiServiceIntegrationInstanceSecretBackup bool
)

func NewCreateApiServiceIntegrationInstanceSecretCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "createApiServiceIntegrationInstanceSecret",
		Long: "Create an API Service Integration instance Secret",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApiServiceIntegrationsAPI.CreateApiServiceIntegrationInstanceSecret(apiClient.GetConfig().Context, CreateApiServiceIntegrationInstanceSecretapiServiceId)

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
			if CreateApiServiceIntegrationInstanceSecretBackup {

				idParam := CreateApiServiceIntegrationInstanceSecretapiServiceId
				err := utils.BackupObject(d, "ApiServiceIntegrations", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateApiServiceIntegrationInstanceSecretapiServiceId, "apiServiceId", "", "", "")
	cmd.MarkFlagRequired("apiServiceId")

	cmd.Flags().BoolVarP(&CreateApiServiceIntegrationInstanceSecretBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateApiServiceIntegrationInstanceSecretCmd := NewCreateApiServiceIntegrationInstanceSecretCmd()
	ApiServiceIntegrationsCmd.AddCommand(CreateApiServiceIntegrationInstanceSecretCmd)
}

var (
	ListApiServiceIntegrationInstanceSecretsapiServiceId string

	ListApiServiceIntegrationInstanceSecretsBackup bool
)

func NewListApiServiceIntegrationInstanceSecretsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listApiServiceIntegrationInstanceSecrets",
		Long: "List all API Service Integration instance Secrets",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApiServiceIntegrationsAPI.ListApiServiceIntegrationInstanceSecrets(apiClient.GetConfig().Context, ListApiServiceIntegrationInstanceSecretsapiServiceId)

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
			if ListApiServiceIntegrationInstanceSecretsBackup {

				idParam := ListApiServiceIntegrationInstanceSecretsapiServiceId
				err := utils.BackupObject(d, "ApiServiceIntegrations", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListApiServiceIntegrationInstanceSecretsapiServiceId, "apiServiceId", "", "", "")
	cmd.MarkFlagRequired("apiServiceId")

	cmd.Flags().BoolVarP(&ListApiServiceIntegrationInstanceSecretsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListApiServiceIntegrationInstanceSecretsCmd := NewListApiServiceIntegrationInstanceSecretsCmd()
	ApiServiceIntegrationsCmd.AddCommand(ListApiServiceIntegrationInstanceSecretsCmd)
}

var (
	DeleteApiServiceIntegrationInstanceSecretapiServiceId string

	DeleteApiServiceIntegrationInstanceSecretsecretId string

	DeleteApiServiceIntegrationInstanceSecretBackup bool
)

func NewDeleteApiServiceIntegrationInstanceSecretCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deleteApiServiceIntegrationInstanceSecret",
		Long: "Delete an API Service Integration instance Secret",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApiServiceIntegrationsAPI.DeleteApiServiceIntegrationInstanceSecret(apiClient.GetConfig().Context, DeleteApiServiceIntegrationInstanceSecretapiServiceId, DeleteApiServiceIntegrationInstanceSecretsecretId)

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
			if DeleteApiServiceIntegrationInstanceSecretBackup {

				idParam := DeleteApiServiceIntegrationInstanceSecretapiServiceId
				err := utils.BackupObject(d, "ApiServiceIntegrations", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteApiServiceIntegrationInstanceSecretapiServiceId, "apiServiceId", "", "", "")
	cmd.MarkFlagRequired("apiServiceId")

	cmd.Flags().StringVarP(&DeleteApiServiceIntegrationInstanceSecretsecretId, "secretId", "", "", "")
	cmd.MarkFlagRequired("secretId")

	cmd.Flags().BoolVarP(&DeleteApiServiceIntegrationInstanceSecretBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteApiServiceIntegrationInstanceSecretCmd := NewDeleteApiServiceIntegrationInstanceSecretCmd()
	ApiServiceIntegrationsCmd.AddCommand(DeleteApiServiceIntegrationInstanceSecretCmd)
}

var (
	ActivateApiServiceIntegrationInstanceSecretapiServiceId string

	ActivateApiServiceIntegrationInstanceSecretsecretId string

	ActivateApiServiceIntegrationInstanceSecretBackup bool
)

func NewActivateApiServiceIntegrationInstanceSecretCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "activateApiServiceIntegrationInstanceSecret",
		Long: "Activate an API Service Integration instance Secret",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApiServiceIntegrationsAPI.ActivateApiServiceIntegrationInstanceSecret(apiClient.GetConfig().Context, ActivateApiServiceIntegrationInstanceSecretapiServiceId, ActivateApiServiceIntegrationInstanceSecretsecretId)

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
			if ActivateApiServiceIntegrationInstanceSecretBackup {

				idParam := ActivateApiServiceIntegrationInstanceSecretapiServiceId
				err := utils.BackupObject(d, "ApiServiceIntegrations", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ActivateApiServiceIntegrationInstanceSecretapiServiceId, "apiServiceId", "", "", "")
	cmd.MarkFlagRequired("apiServiceId")

	cmd.Flags().StringVarP(&ActivateApiServiceIntegrationInstanceSecretsecretId, "secretId", "", "", "")
	cmd.MarkFlagRequired("secretId")

	cmd.Flags().BoolVarP(&ActivateApiServiceIntegrationInstanceSecretBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ActivateApiServiceIntegrationInstanceSecretCmd := NewActivateApiServiceIntegrationInstanceSecretCmd()
	ApiServiceIntegrationsCmd.AddCommand(ActivateApiServiceIntegrationInstanceSecretCmd)
}

var (
	DeactivateApiServiceIntegrationInstanceSecretapiServiceId string

	DeactivateApiServiceIntegrationInstanceSecretsecretId string

	DeactivateApiServiceIntegrationInstanceSecretBackup bool
)

func NewDeactivateApiServiceIntegrationInstanceSecretCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deactivateApiServiceIntegrationInstanceSecret",
		Long: "Deactivate an API Service Integration instance Secret",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApiServiceIntegrationsAPI.DeactivateApiServiceIntegrationInstanceSecret(apiClient.GetConfig().Context, DeactivateApiServiceIntegrationInstanceSecretapiServiceId, DeactivateApiServiceIntegrationInstanceSecretsecretId)

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
			if DeactivateApiServiceIntegrationInstanceSecretBackup {

				idParam := DeactivateApiServiceIntegrationInstanceSecretapiServiceId
				err := utils.BackupObject(d, "ApiServiceIntegrations", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeactivateApiServiceIntegrationInstanceSecretapiServiceId, "apiServiceId", "", "", "")
	cmd.MarkFlagRequired("apiServiceId")

	cmd.Flags().StringVarP(&DeactivateApiServiceIntegrationInstanceSecretsecretId, "secretId", "", "", "")
	cmd.MarkFlagRequired("secretId")

	cmd.Flags().BoolVarP(&DeactivateApiServiceIntegrationInstanceSecretBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeactivateApiServiceIntegrationInstanceSecretCmd := NewDeactivateApiServiceIntegrationInstanceSecretCmd()
	ApiServiceIntegrationsCmd.AddCommand(DeactivateApiServiceIntegrationInstanceSecretCmd)
}
