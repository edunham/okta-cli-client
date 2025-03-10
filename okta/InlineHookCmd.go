package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var InlineHookCmd = &cobra.Command{
	Use:  "inlineHook",
	Long: "Manage InlineHookAPI",
}

func init() {
	rootCmd.AddCommand(InlineHookCmd)
}

var (
	CreateInlineHookdata string

	CreateInlineHookBackup bool
)

func NewCreateInlineHookCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create",
		Long: "Create an Inline Hook",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.InlineHookAPI.CreateInlineHook(apiClient.GetConfig().Context)

			if CreateInlineHookdata != "" {
				req = req.Data(CreateInlineHookdata)
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
			if CreateInlineHookBackup {

				err := utils.BackupObject(d, "InlineHook", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateInlineHookdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateInlineHookBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateInlineHookCmd := NewCreateInlineHookCmd()
	InlineHookCmd.AddCommand(CreateInlineHookCmd)
}

var ListInlineHooksBackup bool

func NewListInlineHooksCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "lists",
		Long: "List all Inline Hooks",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.InlineHookAPI.ListInlineHooks(apiClient.GetConfig().Context)

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
			if ListInlineHooksBackup {

				err := utils.BackupObject(d, "InlineHook", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListInlineHooksBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListInlineHooksCmd := NewListInlineHooksCmd()
	InlineHookCmd.AddCommand(ListInlineHooksCmd)
}

var (
	GetInlineHookinlineHookId string

	GetInlineHookBackup bool
)

func NewGetInlineHookCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "get",
		Long: "Retrieve an Inline Hook",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.InlineHookAPI.GetInlineHook(apiClient.GetConfig().Context, GetInlineHookinlineHookId)

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
			if GetInlineHookBackup {

				idParam := GetInlineHookinlineHookId
				err := utils.BackupObject(d, "InlineHook", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetInlineHookinlineHookId, "inlineHookId", "", "", "")
	cmd.MarkFlagRequired("inlineHookId")

	cmd.Flags().BoolVarP(&GetInlineHookBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetInlineHookCmd := NewGetInlineHookCmd()
	InlineHookCmd.AddCommand(GetInlineHookCmd)
}

var (
	ReplaceInlineHookinlineHookId string

	ReplaceInlineHookdata string

	ReplaceInlineHookBackup bool
)

func NewReplaceInlineHookCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replace",
		Long: "Replace an Inline Hook",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.InlineHookAPI.ReplaceInlineHook(apiClient.GetConfig().Context, ReplaceInlineHookinlineHookId)

			if ReplaceInlineHookdata != "" {
				req = req.Data(ReplaceInlineHookdata)
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
			if ReplaceInlineHookBackup {

				idParam := ReplaceInlineHookinlineHookId
				err := utils.BackupObject(d, "InlineHook", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceInlineHookinlineHookId, "inlineHookId", "", "", "")
	cmd.MarkFlagRequired("inlineHookId")

	cmd.Flags().StringVarP(&ReplaceInlineHookdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceInlineHookBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceInlineHookCmd := NewReplaceInlineHookCmd()
	InlineHookCmd.AddCommand(ReplaceInlineHookCmd)
}

var (
	DeleteInlineHookinlineHookId string

	DeleteInlineHookBackup bool
)

func NewDeleteInlineHookCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete",
		Long: "Delete an Inline Hook",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.InlineHookAPI.DeleteInlineHook(apiClient.GetConfig().Context, DeleteInlineHookinlineHookId)

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
			if DeleteInlineHookBackup {

				idParam := DeleteInlineHookinlineHookId
				err := utils.BackupObject(d, "InlineHook", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteInlineHookinlineHookId, "inlineHookId", "", "", "")
	cmd.MarkFlagRequired("inlineHookId")

	cmd.Flags().BoolVarP(&DeleteInlineHookBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteInlineHookCmd := NewDeleteInlineHookCmd()
	InlineHookCmd.AddCommand(DeleteInlineHookCmd)
}

var (
	ExecuteInlineHookinlineHookId string

	ExecuteInlineHookdata string

	ExecuteInlineHookBackup bool
)

func NewExecuteInlineHookCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "execute",
		Long: "Execute an Inline Hook",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.InlineHookAPI.ExecuteInlineHook(apiClient.GetConfig().Context, ExecuteInlineHookinlineHookId)

			if ExecuteInlineHookdata != "" {
				req = req.Data(ExecuteInlineHookdata)
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
			if ExecuteInlineHookBackup {

				idParam := ExecuteInlineHookinlineHookId
				err := utils.BackupObject(d, "InlineHook", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ExecuteInlineHookinlineHookId, "inlineHookId", "", "", "")
	cmd.MarkFlagRequired("inlineHookId")

	cmd.Flags().StringVarP(&ExecuteInlineHookdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ExecuteInlineHookBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ExecuteInlineHookCmd := NewExecuteInlineHookCmd()
	InlineHookCmd.AddCommand(ExecuteInlineHookCmd)
}

var (
	ActivateInlineHookinlineHookId string

	ActivateInlineHookBackup bool
)

func NewActivateInlineHookCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "activate",
		Long: "Activate an Inline Hook",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.InlineHookAPI.ActivateInlineHook(apiClient.GetConfig().Context, ActivateInlineHookinlineHookId)

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
			if ActivateInlineHookBackup {

				idParam := ActivateInlineHookinlineHookId
				err := utils.BackupObject(d, "InlineHook", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ActivateInlineHookinlineHookId, "inlineHookId", "", "", "")
	cmd.MarkFlagRequired("inlineHookId")

	cmd.Flags().BoolVarP(&ActivateInlineHookBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ActivateInlineHookCmd := NewActivateInlineHookCmd()
	InlineHookCmd.AddCommand(ActivateInlineHookCmd)
}

var (
	DeactivateInlineHookinlineHookId string

	DeactivateInlineHookBackup bool
)

func NewDeactivateInlineHookCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deactivate",
		Long: "Deactivate an Inline Hook",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.InlineHookAPI.DeactivateInlineHook(apiClient.GetConfig().Context, DeactivateInlineHookinlineHookId)

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
			if DeactivateInlineHookBackup {

				idParam := DeactivateInlineHookinlineHookId
				err := utils.BackupObject(d, "InlineHook", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeactivateInlineHookinlineHookId, "inlineHookId", "", "", "")
	cmd.MarkFlagRequired("inlineHookId")

	cmd.Flags().BoolVarP(&DeactivateInlineHookBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeactivateInlineHookCmd := NewDeactivateInlineHookCmd()
	InlineHookCmd.AddCommand(DeactivateInlineHookCmd)
}
