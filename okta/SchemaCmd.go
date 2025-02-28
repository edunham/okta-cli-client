package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var SchemaCmd = &cobra.Command{
	Use:  "schema",
	Long: "Manage SchemaAPI",
}

func init() {
	rootCmd.AddCommand(SchemaCmd)
}

var (
	UpdateApplicationUserProfileappId string

	UpdateApplicationUserProfiledata string

	UpdateApplicationUserProfileBackup bool
)

func NewUpdateApplicationUserProfileCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "updateApplicationUserProfile",
		Long: "Update the default Application User Schema for an Application",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.SchemaAPI.UpdateApplicationUserProfile(apiClient.GetConfig().Context, UpdateApplicationUserProfileappId)

			if UpdateApplicationUserProfiledata != "" {
				req = req.Data(UpdateApplicationUserProfiledata)
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
			if UpdateApplicationUserProfileBackup {

				idParam := UpdateApplicationUserProfileappId
				err := utils.BackupObject(d, "Schema", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UpdateApplicationUserProfileappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().StringVarP(&UpdateApplicationUserProfiledata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&UpdateApplicationUserProfileBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UpdateApplicationUserProfileCmd := NewUpdateApplicationUserProfileCmd()
	SchemaCmd.AddCommand(UpdateApplicationUserProfileCmd)
}

var (
	GetApplicationUserSchemaappId string

	GetApplicationUserSchemaBackup bool
)

func NewGetApplicationUserSchemaCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getApplicationUser",
		Long: "Retrieve the default Application User Schema for an Application",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.SchemaAPI.GetApplicationUserSchema(apiClient.GetConfig().Context, GetApplicationUserSchemaappId)

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
			if GetApplicationUserSchemaBackup {

				idParam := GetApplicationUserSchemaappId
				err := utils.BackupObject(d, "Schema", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetApplicationUserSchemaappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().BoolVarP(&GetApplicationUserSchemaBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetApplicationUserSchemaCmd := NewGetApplicationUserSchemaCmd()
	SchemaCmd.AddCommand(GetApplicationUserSchemaCmd)
}

var (
	UpdateGroupSchemadata string

	UpdateGroupSchemaBackup bool
)

func NewUpdateGroupSchemaCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "updateGroup",
		Long: "Update the default Group Schema",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.SchemaAPI.UpdateGroupSchema(apiClient.GetConfig().Context)

			if UpdateGroupSchemadata != "" {
				req = req.Data(UpdateGroupSchemadata)
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
			if UpdateGroupSchemaBackup {

				err := utils.BackupObject(d, "Schema", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UpdateGroupSchemadata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&UpdateGroupSchemaBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UpdateGroupSchemaCmd := NewUpdateGroupSchemaCmd()
	SchemaCmd.AddCommand(UpdateGroupSchemaCmd)
}

var GetGroupSchemaBackup bool

func NewGetGroupSchemaCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getGroup",
		Long: "Retrieve the default Group Schema",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.SchemaAPI.GetGroupSchema(apiClient.GetConfig().Context)

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
			if GetGroupSchemaBackup {

				err := utils.BackupObject(d, "Schema", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&GetGroupSchemaBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetGroupSchemaCmd := NewGetGroupSchemaCmd()
	SchemaCmd.AddCommand(GetGroupSchemaCmd)
}

var ListLogStreamSchemasBackup bool

func NewListLogStreamSchemasCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listLogStreams",
		Long: "List the Log Stream Schemas",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.SchemaAPI.ListLogStreamSchemas(apiClient.GetConfig().Context)

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
			if ListLogStreamSchemasBackup {

				err := utils.BackupObject(d, "Schema", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListLogStreamSchemasBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListLogStreamSchemasCmd := NewListLogStreamSchemasCmd()
	SchemaCmd.AddCommand(ListLogStreamSchemasCmd)
}

var (
	GetLogStreamSchemalogStreamType string

	GetLogStreamSchemaBackup bool
)

func NewGetLogStreamSchemaCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getLogStream",
		Long: "Retrieve the Log Stream Schema for the schema type",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.SchemaAPI.GetLogStreamSchema(apiClient.GetConfig().Context, GetLogStreamSchemalogStreamType)

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
			if GetLogStreamSchemaBackup {

				idParam := GetLogStreamSchemalogStreamType
				err := utils.BackupObject(d, "Schema", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetLogStreamSchemalogStreamType, "logStreamType", "", "", "")
	cmd.MarkFlagRequired("logStreamType")

	cmd.Flags().BoolVarP(&GetLogStreamSchemaBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetLogStreamSchemaCmd := NewGetLogStreamSchemaCmd()
	SchemaCmd.AddCommand(GetLogStreamSchemaCmd)
}

var (
	UpdateUserProfileschemaId string

	UpdateUserProfiledata string

	UpdateUserProfileBackup bool
)

func NewUpdateUserProfileCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "updateUserProfile",
		Long: "Update a User Schema",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.SchemaAPI.UpdateUserProfile(apiClient.GetConfig().Context, UpdateUserProfileschemaId)

			if UpdateUserProfiledata != "" {
				req = req.Data(UpdateUserProfiledata)
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
			if UpdateUserProfileBackup {

				idParam := UpdateUserProfileschemaId
				err := utils.BackupObject(d, "Schema", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UpdateUserProfileschemaId, "schemaId", "", "", "")
	cmd.MarkFlagRequired("schemaId")

	cmd.Flags().StringVarP(&UpdateUserProfiledata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&UpdateUserProfileBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UpdateUserProfileCmd := NewUpdateUserProfileCmd()
	SchemaCmd.AddCommand(UpdateUserProfileCmd)
}

var (
	GetUserSchemaschemaId string

	GetUserSchemaBackup bool
)

func NewGetUserSchemaCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getUser",
		Long: "Retrieve a User Schema",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.SchemaAPI.GetUserSchema(apiClient.GetConfig().Context, GetUserSchemaschemaId)

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
			if GetUserSchemaBackup {

				idParam := GetUserSchemaschemaId
				err := utils.BackupObject(d, "Schema", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetUserSchemaschemaId, "schemaId", "", "", "")
	cmd.MarkFlagRequired("schemaId")

	cmd.Flags().BoolVarP(&GetUserSchemaBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetUserSchemaCmd := NewGetUserSchemaCmd()
	SchemaCmd.AddCommand(GetUserSchemaCmd)
}
