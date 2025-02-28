package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var ApplicationCredentialsCmd = &cobra.Command{
	Use:  "applicationCredentials",
	Long: "Manage ApplicationCredentialsAPI",
}

func init() {
	rootCmd.AddCommand(ApplicationCredentialsCmd)
}

var (
	GenerateCsrForApplicationappId string

	GenerateCsrForApplicationdata string

	GenerateCsrForApplicationBackup bool
)

func NewGenerateCsrForApplicationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "generateCsrForApplication",
		Long: "Generate a Certificate Signing Request",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationCredentialsAPI.GenerateCsrForApplication(apiClient.GetConfig().Context, GenerateCsrForApplicationappId)

			if GenerateCsrForApplicationdata != "" {
				req = req.Data(GenerateCsrForApplicationdata)
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
			if GenerateCsrForApplicationBackup {

				idParam := GenerateCsrForApplicationappId
				err := utils.BackupObject(d, "ApplicationCredentials", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GenerateCsrForApplicationappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().StringVarP(&GenerateCsrForApplicationdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&GenerateCsrForApplicationBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GenerateCsrForApplicationCmd := NewGenerateCsrForApplicationCmd()
	ApplicationCredentialsCmd.AddCommand(GenerateCsrForApplicationCmd)
}

var (
	ListCsrsForApplicationappId string

	ListCsrsForApplicationBackup bool
)

func NewListCsrsForApplicationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listCsrsForApplication",
		Long: "List all Certificate Signing Requests",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationCredentialsAPI.ListCsrsForApplication(apiClient.GetConfig().Context, ListCsrsForApplicationappId)

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
			if ListCsrsForApplicationBackup {

				idParam := ListCsrsForApplicationappId
				err := utils.BackupObject(d, "ApplicationCredentials", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListCsrsForApplicationappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().BoolVarP(&ListCsrsForApplicationBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListCsrsForApplicationCmd := NewListCsrsForApplicationCmd()
	ApplicationCredentialsCmd.AddCommand(ListCsrsForApplicationCmd)
}

var (
	GetCsrForApplicationappId string

	GetCsrForApplicationcsrId string

	GetCsrForApplicationBackup bool
)

func NewGetCsrForApplicationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getCsrForApplication",
		Long: "Retrieve a Certificate Signing Request",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationCredentialsAPI.GetCsrForApplication(apiClient.GetConfig().Context, GetCsrForApplicationappId, GetCsrForApplicationcsrId)

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
			if GetCsrForApplicationBackup {

				idParam := GetCsrForApplicationappId
				err := utils.BackupObject(d, "ApplicationCredentials", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetCsrForApplicationappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().StringVarP(&GetCsrForApplicationcsrId, "csrId", "", "", "")
	cmd.MarkFlagRequired("csrId")

	cmd.Flags().BoolVarP(&GetCsrForApplicationBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetCsrForApplicationCmd := NewGetCsrForApplicationCmd()
	ApplicationCredentialsCmd.AddCommand(GetCsrForApplicationCmd)
}

var (
	RevokeCsrFromApplicationappId string

	RevokeCsrFromApplicationcsrId string

	RevokeCsrFromApplicationBackup bool
)

func NewRevokeCsrFromApplicationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "revokeCsrFromApplication",
		Long: "Revoke a Certificate Signing Request",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationCredentialsAPI.RevokeCsrFromApplication(apiClient.GetConfig().Context, RevokeCsrFromApplicationappId, RevokeCsrFromApplicationcsrId)

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
			if RevokeCsrFromApplicationBackup {

				idParam := RevokeCsrFromApplicationappId
				err := utils.BackupObject(d, "ApplicationCredentials", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&RevokeCsrFromApplicationappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().StringVarP(&RevokeCsrFromApplicationcsrId, "csrId", "", "", "")
	cmd.MarkFlagRequired("csrId")

	cmd.Flags().BoolVarP(&RevokeCsrFromApplicationBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	RevokeCsrFromApplicationCmd := NewRevokeCsrFromApplicationCmd()
	ApplicationCredentialsCmd.AddCommand(RevokeCsrFromApplicationCmd)
}

var (
	PublishCsrFromApplicationappId string

	PublishCsrFromApplicationcsrId string

	PublishCsrFromApplicationdata string

	PublishCsrFromApplicationBackup bool
)

func NewPublishCsrFromApplicationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "publishCsrFromApplication",
		Long: "Publish a Certificate Signing Request",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationCredentialsAPI.PublishCsrFromApplication(apiClient.GetConfig().Context, PublishCsrFromApplicationappId, PublishCsrFromApplicationcsrId)

			if PublishCsrFromApplicationdata != "" {
				req = req.Data(PublishCsrFromApplicationdata)
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
			if PublishCsrFromApplicationBackup {

				idParam := PublishCsrFromApplicationappId
				err := utils.BackupObject(d, "ApplicationCredentials", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&PublishCsrFromApplicationappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().StringVarP(&PublishCsrFromApplicationcsrId, "csrId", "", "", "")
	cmd.MarkFlagRequired("csrId")

	cmd.Flags().StringVarP(&PublishCsrFromApplicationdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&PublishCsrFromApplicationBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	PublishCsrFromApplicationCmd := NewPublishCsrFromApplicationCmd()
	ApplicationCredentialsCmd.AddCommand(PublishCsrFromApplicationCmd)
}

var (
	ListApplicationKeysappId string

	ListApplicationKeysBackup bool
)

func NewListApplicationKeysCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listApplicationKeys",
		Long: "List all Key Credentials",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationCredentialsAPI.ListApplicationKeys(apiClient.GetConfig().Context, ListApplicationKeysappId)

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
			if ListApplicationKeysBackup {

				idParam := ListApplicationKeysappId
				err := utils.BackupObject(d, "ApplicationCredentials", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListApplicationKeysappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().BoolVarP(&ListApplicationKeysBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListApplicationKeysCmd := NewListApplicationKeysCmd()
	ApplicationCredentialsCmd.AddCommand(ListApplicationKeysCmd)
}

var (
	GenerateApplicationKeyappId string

	GenerateApplicationKeyBackup bool
)

func NewGenerateApplicationKeyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "generateApplicationKey",
		Long: "Generate a Key Credential",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationCredentialsAPI.GenerateApplicationKey(apiClient.GetConfig().Context, GenerateApplicationKeyappId)

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
			if GenerateApplicationKeyBackup {

				idParam := GenerateApplicationKeyappId
				err := utils.BackupObject(d, "ApplicationCredentials", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GenerateApplicationKeyappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().BoolVarP(&GenerateApplicationKeyBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GenerateApplicationKeyCmd := NewGenerateApplicationKeyCmd()
	ApplicationCredentialsCmd.AddCommand(GenerateApplicationKeyCmd)
}

var (
	GetApplicationKeyappId string

	GetApplicationKeykeyId string

	GetApplicationKeyBackup bool
)

func NewGetApplicationKeyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getApplicationKey",
		Long: "Retrieve a Key Credential",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationCredentialsAPI.GetApplicationKey(apiClient.GetConfig().Context, GetApplicationKeyappId, GetApplicationKeykeyId)

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
			if GetApplicationKeyBackup {

				idParam := GetApplicationKeyappId
				err := utils.BackupObject(d, "ApplicationCredentials", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetApplicationKeyappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().StringVarP(&GetApplicationKeykeyId, "keyId", "", "", "")
	cmd.MarkFlagRequired("keyId")

	cmd.Flags().BoolVarP(&GetApplicationKeyBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetApplicationKeyCmd := NewGetApplicationKeyCmd()
	ApplicationCredentialsCmd.AddCommand(GetApplicationKeyCmd)
}

var (
	CloneApplicationKeyappId string

	CloneApplicationKeykeyId string

	CloneApplicationKeyBackup bool
)

func NewCloneApplicationKeyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "cloneApplicationKey",
		Long: "Clone a Key Credential",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationCredentialsAPI.CloneApplicationKey(apiClient.GetConfig().Context, CloneApplicationKeyappId, CloneApplicationKeykeyId)

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
			if CloneApplicationKeyBackup {

				idParam := CloneApplicationKeyappId
				err := utils.BackupObject(d, "ApplicationCredentials", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CloneApplicationKeyappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().StringVarP(&CloneApplicationKeykeyId, "keyId", "", "", "")
	cmd.MarkFlagRequired("keyId")

	cmd.Flags().BoolVarP(&CloneApplicationKeyBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CloneApplicationKeyCmd := NewCloneApplicationKeyCmd()
	ApplicationCredentialsCmd.AddCommand(CloneApplicationKeyCmd)
}
