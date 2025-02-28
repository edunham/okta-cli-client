package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var UISchemaCmd = &cobra.Command{
	Use:  "uISchema",
	Long: "Manage UISchemaAPI",
}

func init() {
	rootCmd.AddCommand(UISchemaCmd)
}

var (
	CreateUISchemadata string

	CreateUISchemaBackup bool
)

func NewCreateUISchemaCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create",
		Long: "Create a UI Schema",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.UISchemaAPI.CreateUISchema(apiClient.GetConfig().Context)

			if CreateUISchemadata != "" {
				req = req.Data(CreateUISchemadata)
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
			if CreateUISchemaBackup {

				err := utils.BackupObject(d, "UISchema", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateUISchemadata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateUISchemaBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateUISchemaCmd := NewCreateUISchemaCmd()
	UISchemaCmd.AddCommand(CreateUISchemaCmd)
}

var ListUISchemasBackup bool

func NewListUISchemasCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "lists",
		Long: "List all UI Schemas",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.UISchemaAPI.ListUISchemas(apiClient.GetConfig().Context)

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
			if ListUISchemasBackup {

				err := utils.BackupObject(d, "UISchema", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListUISchemasBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListUISchemasCmd := NewListUISchemasCmd()
	UISchemaCmd.AddCommand(ListUISchemasCmd)
}

var (
	GetUISchemaid string

	GetUISchemaBackup bool
)

func NewGetUISchemaCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "get",
		Long: "Retrieve a UI Schema",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.UISchemaAPI.GetUISchema(apiClient.GetConfig().Context, GetUISchemaid)

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
			if GetUISchemaBackup {

				idParam := GetUISchemaid
				err := utils.BackupObject(d, "UISchema", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetUISchemaid, "id", "", "", "")
	cmd.MarkFlagRequired("id")

	cmd.Flags().BoolVarP(&GetUISchemaBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetUISchemaCmd := NewGetUISchemaCmd()
	UISchemaCmd.AddCommand(GetUISchemaCmd)
}

var (
	ReplaceUISchemasid string

	ReplaceUISchemasdata string

	ReplaceUISchemasBackup bool
)

func NewReplaceUISchemasCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replaces",
		Long: "Replace a UI Schema",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.UISchemaAPI.ReplaceUISchemas(apiClient.GetConfig().Context, ReplaceUISchemasid)

			if ReplaceUISchemasdata != "" {
				req = req.Data(ReplaceUISchemasdata)
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
			if ReplaceUISchemasBackup {

				idParam := ReplaceUISchemasid
				err := utils.BackupObject(d, "UISchema", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceUISchemasid, "id", "", "", "")
	cmd.MarkFlagRequired("id")

	cmd.Flags().StringVarP(&ReplaceUISchemasdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceUISchemasBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceUISchemasCmd := NewReplaceUISchemasCmd()
	UISchemaCmd.AddCommand(ReplaceUISchemasCmd)
}

var (
	DeleteUISchemasid string

	DeleteUISchemasBackup bool
)

func NewDeleteUISchemasCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deletes",
		Long: "Delete a UI Schema",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.UISchemaAPI.DeleteUISchemas(apiClient.GetConfig().Context, DeleteUISchemasid)

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
			if DeleteUISchemasBackup {

				idParam := DeleteUISchemasid
				err := utils.BackupObject(d, "UISchema", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteUISchemasid, "id", "", "", "")
	cmd.MarkFlagRequired("id")

	cmd.Flags().BoolVarP(&DeleteUISchemasBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteUISchemasCmd := NewDeleteUISchemasCmd()
	UISchemaCmd.AddCommand(DeleteUISchemasCmd)
}
