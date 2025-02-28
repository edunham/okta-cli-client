package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var LinkedObjectCmd = &cobra.Command{
	Use:  "linkedObject",
	Long: "Manage LinkedObjectAPI",
}

func init() {
	rootCmd.AddCommand(LinkedObjectCmd)
}

var (
	CreateLinkedObjectDefinitiondata string

	CreateLinkedObjectDefinitionBackup bool
)

func NewCreateLinkedObjectDefinitionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "createDefinition",
		Long: "Create a Linked Object Definition",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.LinkedObjectAPI.CreateLinkedObjectDefinition(apiClient.GetConfig().Context)

			if CreateLinkedObjectDefinitiondata != "" {
				req = req.Data(CreateLinkedObjectDefinitiondata)
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
			if CreateLinkedObjectDefinitionBackup {

				err := utils.BackupObject(d, "LinkedObject", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateLinkedObjectDefinitiondata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateLinkedObjectDefinitionBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateLinkedObjectDefinitionCmd := NewCreateLinkedObjectDefinitionCmd()
	LinkedObjectCmd.AddCommand(CreateLinkedObjectDefinitionCmd)
}

var ListLinkedObjectDefinitionsBackup bool

func NewListLinkedObjectDefinitionsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listDefinitions",
		Long: "List all Linked Object Definitions",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.LinkedObjectAPI.ListLinkedObjectDefinitions(apiClient.GetConfig().Context)

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
			if ListLinkedObjectDefinitionsBackup {

				err := utils.BackupObject(d, "LinkedObject", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListLinkedObjectDefinitionsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListLinkedObjectDefinitionsCmd := NewListLinkedObjectDefinitionsCmd()
	LinkedObjectCmd.AddCommand(ListLinkedObjectDefinitionsCmd)
}

var (
	GetLinkedObjectDefinitionlinkedObjectName string

	GetLinkedObjectDefinitionBackup bool
)

func NewGetLinkedObjectDefinitionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getDefinition",
		Long: "Retrieve a Linked Object Definition",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.LinkedObjectAPI.GetLinkedObjectDefinition(apiClient.GetConfig().Context, GetLinkedObjectDefinitionlinkedObjectName)

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
			if GetLinkedObjectDefinitionBackup {

				idParam := GetLinkedObjectDefinitionlinkedObjectName
				err := utils.BackupObject(d, "LinkedObject", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetLinkedObjectDefinitionlinkedObjectName, "linkedObjectName", "", "", "")
	cmd.MarkFlagRequired("linkedObjectName")

	cmd.Flags().BoolVarP(&GetLinkedObjectDefinitionBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetLinkedObjectDefinitionCmd := NewGetLinkedObjectDefinitionCmd()
	LinkedObjectCmd.AddCommand(GetLinkedObjectDefinitionCmd)
}

var (
	DeleteLinkedObjectDefinitionlinkedObjectName string

	DeleteLinkedObjectDefinitionBackup bool
)

func NewDeleteLinkedObjectDefinitionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deleteDefinition",
		Long: "Delete a Linked Object Definition",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.LinkedObjectAPI.DeleteLinkedObjectDefinition(apiClient.GetConfig().Context, DeleteLinkedObjectDefinitionlinkedObjectName)

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
			if DeleteLinkedObjectDefinitionBackup {

				idParam := DeleteLinkedObjectDefinitionlinkedObjectName
				err := utils.BackupObject(d, "LinkedObject", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteLinkedObjectDefinitionlinkedObjectName, "linkedObjectName", "", "", "")
	cmd.MarkFlagRequired("linkedObjectName")

	cmd.Flags().BoolVarP(&DeleteLinkedObjectDefinitionBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteLinkedObjectDefinitionCmd := NewDeleteLinkedObjectDefinitionCmd()
	LinkedObjectCmd.AddCommand(DeleteLinkedObjectDefinitionCmd)
}
