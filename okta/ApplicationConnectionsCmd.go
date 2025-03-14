package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var ApplicationConnectionsCmd = &cobra.Command{
	Use:  "applicationConnections",
	Long: "Manage ApplicationConnectionsAPI",
}

func init() {
	rootCmd.AddCommand(ApplicationConnectionsCmd)
}

var (
	UpdateDefaultProvisioningConnectionForApplicationappId string

	UpdateDefaultProvisioningConnectionForApplicationdata string

	UpdateDefaultProvisioningConnectionForApplicationBackup bool
)

func NewUpdateDefaultProvisioningConnectionForApplicationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "updateDefaultProvisioningConnectionForApplication",
		Long: "Update the default Provisioning Connection",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationConnectionsAPI.UpdateDefaultProvisioningConnectionForApplication(apiClient.GetConfig().Context, UpdateDefaultProvisioningConnectionForApplicationappId)

			if UpdateDefaultProvisioningConnectionForApplicationdata != "" {
				req = req.Data(UpdateDefaultProvisioningConnectionForApplicationdata)
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
			if UpdateDefaultProvisioningConnectionForApplicationBackup {

				idParam := UpdateDefaultProvisioningConnectionForApplicationappId
				err := utils.BackupObject(d, "ApplicationConnections", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UpdateDefaultProvisioningConnectionForApplicationappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().StringVarP(&UpdateDefaultProvisioningConnectionForApplicationdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&UpdateDefaultProvisioningConnectionForApplicationBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UpdateDefaultProvisioningConnectionForApplicationCmd := NewUpdateDefaultProvisioningConnectionForApplicationCmd()
	ApplicationConnectionsCmd.AddCommand(UpdateDefaultProvisioningConnectionForApplicationCmd)
}

var (
	GetDefaultProvisioningConnectionForApplicationappId string

	GetDefaultProvisioningConnectionForApplicationBackup bool
)

func NewGetDefaultProvisioningConnectionForApplicationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getDefaultProvisioningConnectionForApplication",
		Long: "Retrieve the default Provisioning Connection",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationConnectionsAPI.GetDefaultProvisioningConnectionForApplication(apiClient.GetConfig().Context, GetDefaultProvisioningConnectionForApplicationappId)

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
			if GetDefaultProvisioningConnectionForApplicationBackup {

				idParam := GetDefaultProvisioningConnectionForApplicationappId
				err := utils.BackupObject(d, "ApplicationConnections", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetDefaultProvisioningConnectionForApplicationappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().BoolVarP(&GetDefaultProvisioningConnectionForApplicationBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetDefaultProvisioningConnectionForApplicationCmd := NewGetDefaultProvisioningConnectionForApplicationCmd()
	ApplicationConnectionsCmd.AddCommand(GetDefaultProvisioningConnectionForApplicationCmd)
}

var (
	ActivateDefaultProvisioningConnectionForApplicationappId string

	ActivateDefaultProvisioningConnectionForApplicationBackup bool
)

func NewActivateDefaultProvisioningConnectionForApplicationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "activateDefaultProvisioningConnectionForApplication",
		Long: "Activate the default Provisioning Connection",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationConnectionsAPI.ActivateDefaultProvisioningConnectionForApplication(apiClient.GetConfig().Context, ActivateDefaultProvisioningConnectionForApplicationappId)

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
			if ActivateDefaultProvisioningConnectionForApplicationBackup {

				idParam := ActivateDefaultProvisioningConnectionForApplicationappId
				err := utils.BackupObject(d, "ApplicationConnections", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ActivateDefaultProvisioningConnectionForApplicationappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().BoolVarP(&ActivateDefaultProvisioningConnectionForApplicationBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ActivateDefaultProvisioningConnectionForApplicationCmd := NewActivateDefaultProvisioningConnectionForApplicationCmd()
	ApplicationConnectionsCmd.AddCommand(ActivateDefaultProvisioningConnectionForApplicationCmd)
}

var (
	DeactivateDefaultProvisioningConnectionForApplicationappId string

	DeactivateDefaultProvisioningConnectionForApplicationBackup bool
)

func NewDeactivateDefaultProvisioningConnectionForApplicationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deactivateDefaultProvisioningConnectionForApplication",
		Long: "Deactivate the default Provisioning Connection",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationConnectionsAPI.DeactivateDefaultProvisioningConnectionForApplication(apiClient.GetConfig().Context, DeactivateDefaultProvisioningConnectionForApplicationappId)

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
			if DeactivateDefaultProvisioningConnectionForApplicationBackup {

				idParam := DeactivateDefaultProvisioningConnectionForApplicationappId
				err := utils.BackupObject(d, "ApplicationConnections", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeactivateDefaultProvisioningConnectionForApplicationappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().BoolVarP(&DeactivateDefaultProvisioningConnectionForApplicationBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeactivateDefaultProvisioningConnectionForApplicationCmd := NewDeactivateDefaultProvisioningConnectionForApplicationCmd()
	ApplicationConnectionsCmd.AddCommand(DeactivateDefaultProvisioningConnectionForApplicationCmd)
}

var (
	VerifyProvisioningConnectionForApplicationappName string

	VerifyProvisioningConnectionForApplicationappId string

	VerifyProvisioningConnectionForApplicationBackup bool
)

func NewVerifyProvisioningConnectionForApplicationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "verifyProvisioningConnectionForApplication",
		Long: "Verify the Provisioning Connection",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationConnectionsAPI.VerifyProvisioningConnectionForApplication(apiClient.GetConfig().Context, VerifyProvisioningConnectionForApplicationappName, VerifyProvisioningConnectionForApplicationappId)

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
			if VerifyProvisioningConnectionForApplicationBackup {

				idParam := VerifyProvisioningConnectionForApplicationappName
				err := utils.BackupObject(d, "ApplicationConnections", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&VerifyProvisioningConnectionForApplicationappName, "appName", "", "", "")
	cmd.MarkFlagRequired("appName")

	cmd.Flags().StringVarP(&VerifyProvisioningConnectionForApplicationappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().BoolVarP(&VerifyProvisioningConnectionForApplicationBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	VerifyProvisioningConnectionForApplicationCmd := NewVerifyProvisioningConnectionForApplicationCmd()
	ApplicationConnectionsCmd.AddCommand(VerifyProvisioningConnectionForApplicationCmd)
}
