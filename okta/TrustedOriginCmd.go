package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var TrustedOriginCmd = &cobra.Command{
	Use:  "trustedOrigin",
	Long: "Manage TrustedOriginAPI",
}

func init() {
	rootCmd.AddCommand(TrustedOriginCmd)
}

var (
	CreateTrustedOrigindata string

	CreateTrustedOriginBackup bool
)

func NewCreateTrustedOriginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create",
		Long: "Create a Trusted Origin",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.TrustedOriginAPI.CreateTrustedOrigin(apiClient.GetConfig().Context)

			if CreateTrustedOrigindata != "" {
				req = req.Data(CreateTrustedOrigindata)
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
			if CreateTrustedOriginBackup {

				err := utils.BackupObject(d, "TrustedOrigin", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateTrustedOrigindata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateTrustedOriginBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateTrustedOriginCmd := NewCreateTrustedOriginCmd()
	TrustedOriginCmd.AddCommand(CreateTrustedOriginCmd)
}

var ListTrustedOriginsBackup bool

func NewListTrustedOriginsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "lists",
		Long: "List all Trusted Origins",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.TrustedOriginAPI.ListTrustedOrigins(apiClient.GetConfig().Context)

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
			if ListTrustedOriginsBackup {

				err := utils.BackupObject(d, "TrustedOrigin", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListTrustedOriginsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListTrustedOriginsCmd := NewListTrustedOriginsCmd()
	TrustedOriginCmd.AddCommand(ListTrustedOriginsCmd)
}

var (
	GetTrustedOrigintrustedOriginId string

	GetTrustedOriginBackup bool
)

func NewGetTrustedOriginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "get",
		Long: "Retrieve a Trusted Origin",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.TrustedOriginAPI.GetTrustedOrigin(apiClient.GetConfig().Context, GetTrustedOrigintrustedOriginId)

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
			if GetTrustedOriginBackup {

				idParam := GetTrustedOrigintrustedOriginId
				err := utils.BackupObject(d, "TrustedOrigin", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetTrustedOrigintrustedOriginId, "trustedOriginId", "", "", "")
	cmd.MarkFlagRequired("trustedOriginId")

	cmd.Flags().BoolVarP(&GetTrustedOriginBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetTrustedOriginCmd := NewGetTrustedOriginCmd()
	TrustedOriginCmd.AddCommand(GetTrustedOriginCmd)
}

var (
	ReplaceTrustedOrigintrustedOriginId string

	ReplaceTrustedOrigindata string

	ReplaceTrustedOriginBackup bool
)

func NewReplaceTrustedOriginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replace",
		Long: "Replace a Trusted Origin",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.TrustedOriginAPI.ReplaceTrustedOrigin(apiClient.GetConfig().Context, ReplaceTrustedOrigintrustedOriginId)

			if ReplaceTrustedOrigindata != "" {
				req = req.Data(ReplaceTrustedOrigindata)
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
			if ReplaceTrustedOriginBackup {

				idParam := ReplaceTrustedOrigintrustedOriginId
				err := utils.BackupObject(d, "TrustedOrigin", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceTrustedOrigintrustedOriginId, "trustedOriginId", "", "", "")
	cmd.MarkFlagRequired("trustedOriginId")

	cmd.Flags().StringVarP(&ReplaceTrustedOrigindata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceTrustedOriginBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceTrustedOriginCmd := NewReplaceTrustedOriginCmd()
	TrustedOriginCmd.AddCommand(ReplaceTrustedOriginCmd)
}

var (
	DeleteTrustedOrigintrustedOriginId string

	DeleteTrustedOriginBackup bool
)

func NewDeleteTrustedOriginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete",
		Long: "Delete a Trusted Origin",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.TrustedOriginAPI.DeleteTrustedOrigin(apiClient.GetConfig().Context, DeleteTrustedOrigintrustedOriginId)

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
			if DeleteTrustedOriginBackup {

				idParam := DeleteTrustedOrigintrustedOriginId
				err := utils.BackupObject(d, "TrustedOrigin", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteTrustedOrigintrustedOriginId, "trustedOriginId", "", "", "")
	cmd.MarkFlagRequired("trustedOriginId")

	cmd.Flags().BoolVarP(&DeleteTrustedOriginBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteTrustedOriginCmd := NewDeleteTrustedOriginCmd()
	TrustedOriginCmd.AddCommand(DeleteTrustedOriginCmd)
}

var (
	ActivateTrustedOrigintrustedOriginId string

	ActivateTrustedOriginBackup bool
)

func NewActivateTrustedOriginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "activate",
		Long: "Activate a Trusted Origin",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.TrustedOriginAPI.ActivateTrustedOrigin(apiClient.GetConfig().Context, ActivateTrustedOrigintrustedOriginId)

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
			if ActivateTrustedOriginBackup {

				idParam := ActivateTrustedOrigintrustedOriginId
				err := utils.BackupObject(d, "TrustedOrigin", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ActivateTrustedOrigintrustedOriginId, "trustedOriginId", "", "", "")
	cmd.MarkFlagRequired("trustedOriginId")

	cmd.Flags().BoolVarP(&ActivateTrustedOriginBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ActivateTrustedOriginCmd := NewActivateTrustedOriginCmd()
	TrustedOriginCmd.AddCommand(ActivateTrustedOriginCmd)
}

var (
	DeactivateTrustedOrigintrustedOriginId string

	DeactivateTrustedOriginBackup bool
)

func NewDeactivateTrustedOriginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deactivate",
		Long: "Deactivate a Trusted Origin",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.TrustedOriginAPI.DeactivateTrustedOrigin(apiClient.GetConfig().Context, DeactivateTrustedOrigintrustedOriginId)

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
			if DeactivateTrustedOriginBackup {

				idParam := DeactivateTrustedOrigintrustedOriginId
				err := utils.BackupObject(d, "TrustedOrigin", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeactivateTrustedOrigintrustedOriginId, "trustedOriginId", "", "", "")
	cmd.MarkFlagRequired("trustedOriginId")

	cmd.Flags().BoolVarP(&DeactivateTrustedOriginBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeactivateTrustedOriginCmd := NewDeactivateTrustedOriginCmd()
	TrustedOriginCmd.AddCommand(DeactivateTrustedOriginCmd)
}
