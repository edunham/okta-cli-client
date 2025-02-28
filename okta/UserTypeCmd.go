package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var UserTypeCmd = &cobra.Command{
	Use:  "userType",
	Long: "Manage UserTypeAPI",
}

func init() {
	rootCmd.AddCommand(UserTypeCmd)
}

var (
	CreateUserTypedata string

	CreateUserTypeBackup bool
)

func NewCreateUserTypeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create",
		Long: "Create a User Type",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.UserTypeAPI.CreateUserType(apiClient.GetConfig().Context)

			if CreateUserTypedata != "" {
				req = req.Data(CreateUserTypedata)
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
			if CreateUserTypeBackup {

				err := utils.BackupObject(d, "UserType", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateUserTypedata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateUserTypeBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateUserTypeCmd := NewCreateUserTypeCmd()
	UserTypeCmd.AddCommand(CreateUserTypeCmd)
}

var ListUserTypesBackup bool

func NewListUserTypesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "lists",
		Long: "List all User Types",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.UserTypeAPI.ListUserTypes(apiClient.GetConfig().Context)

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
			if ListUserTypesBackup {

				err := utils.BackupObject(d, "UserType", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListUserTypesBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListUserTypesCmd := NewListUserTypesCmd()
	UserTypeCmd.AddCommand(ListUserTypesCmd)
}

var (
	UpdateUserTypetypeId string

	UpdateUserTypedata string

	UpdateUserTypeBackup bool
)

func NewUpdateUserTypeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "update",
		Long: "Update a User Type",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.UserTypeAPI.UpdateUserType(apiClient.GetConfig().Context, UpdateUserTypetypeId)

			if UpdateUserTypedata != "" {
				req = req.Data(UpdateUserTypedata)
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
			if UpdateUserTypeBackup {

				idParam := UpdateUserTypetypeId
				err := utils.BackupObject(d, "UserType", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UpdateUserTypetypeId, "typeId", "", "", "")
	cmd.MarkFlagRequired("typeId")

	cmd.Flags().StringVarP(&UpdateUserTypedata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&UpdateUserTypeBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UpdateUserTypeCmd := NewUpdateUserTypeCmd()
	UserTypeCmd.AddCommand(UpdateUserTypeCmd)
}

var (
	GetUserTypetypeId string

	GetUserTypeBackup bool
)

func NewGetUserTypeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "get",
		Long: "Retrieve a User Type",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.UserTypeAPI.GetUserType(apiClient.GetConfig().Context, GetUserTypetypeId)

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
			if GetUserTypeBackup {

				idParam := GetUserTypetypeId
				err := utils.BackupObject(d, "UserType", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetUserTypetypeId, "typeId", "", "", "")
	cmd.MarkFlagRequired("typeId")

	cmd.Flags().BoolVarP(&GetUserTypeBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetUserTypeCmd := NewGetUserTypeCmd()
	UserTypeCmd.AddCommand(GetUserTypeCmd)
}

var (
	ReplaceUserTypetypeId string

	ReplaceUserTypedata string

	ReplaceUserTypeBackup bool
)

func NewReplaceUserTypeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replace",
		Long: "Replace a User Type",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.UserTypeAPI.ReplaceUserType(apiClient.GetConfig().Context, ReplaceUserTypetypeId)

			if ReplaceUserTypedata != "" {
				req = req.Data(ReplaceUserTypedata)
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
			if ReplaceUserTypeBackup {

				idParam := ReplaceUserTypetypeId
				err := utils.BackupObject(d, "UserType", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceUserTypetypeId, "typeId", "", "", "")
	cmd.MarkFlagRequired("typeId")

	cmd.Flags().StringVarP(&ReplaceUserTypedata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceUserTypeBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceUserTypeCmd := NewReplaceUserTypeCmd()
	UserTypeCmd.AddCommand(ReplaceUserTypeCmd)
}

var (
	DeleteUserTypetypeId string

	DeleteUserTypeBackup bool
)

func NewDeleteUserTypeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete",
		Long: "Delete a User Type",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.UserTypeAPI.DeleteUserType(apiClient.GetConfig().Context, DeleteUserTypetypeId)

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
			if DeleteUserTypeBackup {

				idParam := DeleteUserTypetypeId
				err := utils.BackupObject(d, "UserType", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteUserTypetypeId, "typeId", "", "", "")
	cmd.MarkFlagRequired("typeId")

	cmd.Flags().BoolVarP(&DeleteUserTypeBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteUserTypeCmd := NewDeleteUserTypeCmd()
	UserTypeCmd.AddCommand(DeleteUserTypeCmd)
}
