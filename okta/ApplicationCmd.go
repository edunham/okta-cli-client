package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var ApplicationCmd = &cobra.Command{
	Use:  "application",
	Long: "Manage ApplicationAPI",
}

func init() {
	rootCmd.AddCommand(ApplicationCmd)
}

var (
	CreateApplicationdata string

	CreateApplicationBackup bool
)

func NewCreateApplicationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create",
		Long: "Create an Application",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationAPI.CreateApplication(apiClient.GetConfig().Context)

			if CreateApplicationdata != "" {
				req = req.Data(CreateApplicationdata)
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
			if CreateApplicationBackup {

				err := utils.BackupObject(d, "Application", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateApplicationdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateApplicationBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateApplicationCmd := NewCreateApplicationCmd()
	ApplicationCmd.AddCommand(CreateApplicationCmd)
}

var ListApplicationsBackup bool

func NewListApplicationsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "lists",
		Long: "List all Applications",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationAPI.ListApplications(apiClient.GetConfig().Context)

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
			if ListApplicationsBackup {

				err := utils.BackupObject(d, "Application", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListApplicationsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListApplicationsCmd := NewListApplicationsCmd()
	ApplicationCmd.AddCommand(ListApplicationsCmd)
}

var (
	GetApplicationappId string

	GetApplicationBackup bool
)

func NewGetApplicationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "get",
		Long: "Retrieve an Application",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationAPI.GetApplication(apiClient.GetConfig().Context, GetApplicationappId)

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
			if GetApplicationBackup {

				idParam := GetApplicationappId
				err := utils.BackupObject(d, "Application", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetApplicationappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().BoolVarP(&GetApplicationBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetApplicationCmd := NewGetApplicationCmd()
	ApplicationCmd.AddCommand(GetApplicationCmd)
}

var (
	ReplaceApplicationappId string

	ReplaceApplicationdata string

	ReplaceApplicationBackup bool
)

func NewReplaceApplicationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replace",
		Long: "Replace an Application",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationAPI.ReplaceApplication(apiClient.GetConfig().Context, ReplaceApplicationappId)

			if ReplaceApplicationdata != "" {
				req = req.Data(ReplaceApplicationdata)
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
			if ReplaceApplicationBackup {

				idParam := ReplaceApplicationappId
				err := utils.BackupObject(d, "Application", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceApplicationappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().StringVarP(&ReplaceApplicationdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceApplicationBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceApplicationCmd := NewReplaceApplicationCmd()
	ApplicationCmd.AddCommand(ReplaceApplicationCmd)
}

var (
	DeleteApplicationappId string

	DeleteApplicationBackup bool
)

func NewDeleteApplicationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete",
		Long: "Delete an Application",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationAPI.DeleteApplication(apiClient.GetConfig().Context, DeleteApplicationappId)

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
			if DeleteApplicationBackup {

				idParam := DeleteApplicationappId
				err := utils.BackupObject(d, "Application", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteApplicationappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().BoolVarP(&DeleteApplicationBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteApplicationCmd := NewDeleteApplicationCmd()
	ApplicationCmd.AddCommand(DeleteApplicationCmd)
}

var (
	ActivateApplicationappId string

	ActivateApplicationBackup bool
)

func NewActivateApplicationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "activate",
		Long: "Activate an Application",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationAPI.ActivateApplication(apiClient.GetConfig().Context, ActivateApplicationappId)

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
			if ActivateApplicationBackup {

				idParam := ActivateApplicationappId
				err := utils.BackupObject(d, "Application", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ActivateApplicationappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().BoolVarP(&ActivateApplicationBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ActivateApplicationCmd := NewActivateApplicationCmd()
	ApplicationCmd.AddCommand(ActivateApplicationCmd)
}

var (
	DeactivateApplicationappId string

	DeactivateApplicationBackup bool
)

func NewDeactivateApplicationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deactivate",
		Long: "Deactivate an Application",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationAPI.DeactivateApplication(apiClient.GetConfig().Context, DeactivateApplicationappId)

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
			if DeactivateApplicationBackup {

				idParam := DeactivateApplicationappId
				err := utils.BackupObject(d, "Application", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeactivateApplicationappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().BoolVarP(&DeactivateApplicationBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeactivateApplicationCmd := NewDeactivateApplicationCmd()
	ApplicationCmd.AddCommand(DeactivateApplicationCmd)
}
