package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var HookKeyCmd = &cobra.Command{
	Use:  "hookKey",
	Long: "Manage HookKeyAPI",
}

func init() {
	rootCmd.AddCommand(HookKeyCmd)
}

var (
	CreateHookKeydata string

	CreateHookKeyBackup bool
)

func NewCreateHookKeyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create",
		Long: "Create a key",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.HookKeyAPI.CreateHookKey(apiClient.GetConfig().Context)

			if CreateHookKeydata != "" {
				req = req.Data(CreateHookKeydata)
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
			if CreateHookKeyBackup {

				err := utils.BackupObject(d, "HookKey", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateHookKeydata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateHookKeyBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateHookKeyCmd := NewCreateHookKeyCmd()
	HookKeyCmd.AddCommand(CreateHookKeyCmd)
}

var ListHookKeysBackup bool

func NewListHookKeysCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "lists",
		Long: "List all keys",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.HookKeyAPI.ListHookKeys(apiClient.GetConfig().Context)

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
			if ListHookKeysBackup {

				err := utils.BackupObject(d, "HookKey", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListHookKeysBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListHookKeysCmd := NewListHookKeysCmd()
	HookKeyCmd.AddCommand(ListHookKeysCmd)
}

var (
	GetPublicKeypublicKeyId string

	GetPublicKeyBackup bool
)

func NewGetPublicKeyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getPublicKey",
		Long: "Retrieve a public key",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.HookKeyAPI.GetPublicKey(apiClient.GetConfig().Context, GetPublicKeypublicKeyId)

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
			if GetPublicKeyBackup {

				idParam := GetPublicKeypublicKeyId
				err := utils.BackupObject(d, "HookKey", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetPublicKeypublicKeyId, "publicKeyId", "", "", "")
	cmd.MarkFlagRequired("publicKeyId")

	cmd.Flags().BoolVarP(&GetPublicKeyBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetPublicKeyCmd := NewGetPublicKeyCmd()
	HookKeyCmd.AddCommand(GetPublicKeyCmd)
}

var (
	GetHookKeyhookKeyId string

	GetHookKeyBackup bool
)

func NewGetHookKeyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "get",
		Long: "Retrieve a key",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.HookKeyAPI.GetHookKey(apiClient.GetConfig().Context, GetHookKeyhookKeyId)

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
			if GetHookKeyBackup {

				idParam := GetHookKeyhookKeyId
				err := utils.BackupObject(d, "HookKey", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetHookKeyhookKeyId, "hookKeyId", "", "", "")
	cmd.MarkFlagRequired("hookKeyId")

	cmd.Flags().BoolVarP(&GetHookKeyBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetHookKeyCmd := NewGetHookKeyCmd()
	HookKeyCmd.AddCommand(GetHookKeyCmd)
}

var (
	ReplaceHookKeyhookKeyId string

	ReplaceHookKeydata string

	ReplaceHookKeyBackup bool
)

func NewReplaceHookKeyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replace",
		Long: "Replace a key",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.HookKeyAPI.ReplaceHookKey(apiClient.GetConfig().Context, ReplaceHookKeyhookKeyId)

			if ReplaceHookKeydata != "" {
				req = req.Data(ReplaceHookKeydata)
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
			if ReplaceHookKeyBackup {

				idParam := ReplaceHookKeyhookKeyId
				err := utils.BackupObject(d, "HookKey", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceHookKeyhookKeyId, "hookKeyId", "", "", "")
	cmd.MarkFlagRequired("hookKeyId")

	cmd.Flags().StringVarP(&ReplaceHookKeydata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceHookKeyBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceHookKeyCmd := NewReplaceHookKeyCmd()
	HookKeyCmd.AddCommand(ReplaceHookKeyCmd)
}

var (
	DeleteHookKeyhookKeyId string

	DeleteHookKeyBackup bool
)

func NewDeleteHookKeyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete",
		Long: "Delete a key",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.HookKeyAPI.DeleteHookKey(apiClient.GetConfig().Context, DeleteHookKeyhookKeyId)

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
			if DeleteHookKeyBackup {

				idParam := DeleteHookKeyhookKeyId
				err := utils.BackupObject(d, "HookKey", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteHookKeyhookKeyId, "hookKeyId", "", "", "")
	cmd.MarkFlagRequired("hookKeyId")

	cmd.Flags().BoolVarP(&DeleteHookKeyBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteHookKeyCmd := NewDeleteHookKeyCmd()
	HookKeyCmd.AddCommand(DeleteHookKeyCmd)
}
