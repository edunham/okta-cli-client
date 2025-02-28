package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var RoleAssignmentCmd = &cobra.Command{
	Use:  "roleAssignment",
	Long: "Manage RoleAssignmentAPI",
}

func init() {
	rootCmd.AddCommand(RoleAssignmentCmd)
}

var (
	AssignRoleToGroupgroupId string

	AssignRoleToGroupdata string

	AssignRoleToGroupBackup bool
)

func NewAssignRoleToGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "assignRoleToGroup",
		Long: "Assign a Role to a Group",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RoleAssignmentAPI.AssignRoleToGroup(apiClient.GetConfig().Context, AssignRoleToGroupgroupId)

			if AssignRoleToGroupdata != "" {
				req = req.Data(AssignRoleToGroupdata)
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
			if AssignRoleToGroupBackup {

				idParam := AssignRoleToGroupgroupId
				err := utils.BackupObject(d, "RoleAssignment", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&AssignRoleToGroupgroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	cmd.Flags().StringVarP(&AssignRoleToGroupdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&AssignRoleToGroupBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	AssignRoleToGroupCmd := NewAssignRoleToGroupCmd()
	RoleAssignmentCmd.AddCommand(AssignRoleToGroupCmd)
}

var (
	ListGroupAssignedRolesgroupId string

	ListGroupAssignedRolesBackup bool
)

func NewListGroupAssignedRolesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listGroupAssignedRoles",
		Long: "List all Assigned Roles of Group",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RoleAssignmentAPI.ListGroupAssignedRoles(apiClient.GetConfig().Context, ListGroupAssignedRolesgroupId)

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
			if ListGroupAssignedRolesBackup {

				idParam := ListGroupAssignedRolesgroupId
				err := utils.BackupObject(d, "RoleAssignment", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListGroupAssignedRolesgroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	cmd.Flags().BoolVarP(&ListGroupAssignedRolesBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListGroupAssignedRolesCmd := NewListGroupAssignedRolesCmd()
	RoleAssignmentCmd.AddCommand(ListGroupAssignedRolesCmd)
}

var (
	GetGroupAssignedRolegroupId string

	GetGroupAssignedRoleroleId string

	GetGroupAssignedRoleBackup bool
)

func NewGetGroupAssignedRoleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getGroupAssignedRole",
		Long: "Retrieve a Role assigned to Group",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RoleAssignmentAPI.GetGroupAssignedRole(apiClient.GetConfig().Context, GetGroupAssignedRolegroupId, GetGroupAssignedRoleroleId)

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
			if GetGroupAssignedRoleBackup {

				idParam := GetGroupAssignedRolegroupId
				err := utils.BackupObject(d, "RoleAssignment", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetGroupAssignedRolegroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	cmd.Flags().StringVarP(&GetGroupAssignedRoleroleId, "roleId", "", "", "")
	cmd.MarkFlagRequired("roleId")

	cmd.Flags().BoolVarP(&GetGroupAssignedRoleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetGroupAssignedRoleCmd := NewGetGroupAssignedRoleCmd()
	RoleAssignmentCmd.AddCommand(GetGroupAssignedRoleCmd)
}

var (
	UnassignRoleFromGroupgroupId string

	UnassignRoleFromGrouproleId string

	UnassignRoleFromGroupBackup bool
)

func NewUnassignRoleFromGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "unassignRoleFromGroup",
		Long: "Unassign a Role from a Group",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RoleAssignmentAPI.UnassignRoleFromGroup(apiClient.GetConfig().Context, UnassignRoleFromGroupgroupId, UnassignRoleFromGrouproleId)

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
			if UnassignRoleFromGroupBackup {

				idParam := UnassignRoleFromGroupgroupId
				err := utils.BackupObject(d, "RoleAssignment", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UnassignRoleFromGroupgroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	cmd.Flags().StringVarP(&UnassignRoleFromGrouproleId, "roleId", "", "", "")
	cmd.MarkFlagRequired("roleId")

	cmd.Flags().BoolVarP(&UnassignRoleFromGroupBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UnassignRoleFromGroupCmd := NewUnassignRoleFromGroupCmd()
	RoleAssignmentCmd.AddCommand(UnassignRoleFromGroupCmd)
}

var ListUsersWithRoleAssignmentsBackup bool

func NewListUsersWithRoleAssignmentsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listUsersWiths",
		Long: "List all Users with Role Assignments",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RoleAssignmentAPI.ListUsersWithRoleAssignments(apiClient.GetConfig().Context)

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
			if ListUsersWithRoleAssignmentsBackup {

				err := utils.BackupObject(d, "RoleAssignment", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListUsersWithRoleAssignmentsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListUsersWithRoleAssignmentsCmd := NewListUsersWithRoleAssignmentsCmd()
	RoleAssignmentCmd.AddCommand(ListUsersWithRoleAssignmentsCmd)
}

var (
	AssignRoleToUseruserId string

	AssignRoleToUserdata string

	AssignRoleToUserBackup bool
)

func NewAssignRoleToUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "assignRoleToUser",
		Long: "Assign a Role to a User",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RoleAssignmentAPI.AssignRoleToUser(apiClient.GetConfig().Context, AssignRoleToUseruserId)

			if AssignRoleToUserdata != "" {
				req = req.Data(AssignRoleToUserdata)
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
			if AssignRoleToUserBackup {

				idParam := AssignRoleToUseruserId
				err := utils.BackupObject(d, "RoleAssignment", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&AssignRoleToUseruserId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	cmd.Flags().StringVarP(&AssignRoleToUserdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&AssignRoleToUserBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	AssignRoleToUserCmd := NewAssignRoleToUserCmd()
	RoleAssignmentCmd.AddCommand(AssignRoleToUserCmd)
}

var (
	ListAssignedRolesForUseruserId string

	ListAssignedRolesForUserBackup bool
)

func NewListAssignedRolesForUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listAssignedRolesForUser",
		Long: "List all Roles assigned to a User",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RoleAssignmentAPI.ListAssignedRolesForUser(apiClient.GetConfig().Context, ListAssignedRolesForUseruserId)

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
			if ListAssignedRolesForUserBackup {

				idParam := ListAssignedRolesForUseruserId
				err := utils.BackupObject(d, "RoleAssignment", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListAssignedRolesForUseruserId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	cmd.Flags().BoolVarP(&ListAssignedRolesForUserBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListAssignedRolesForUserCmd := NewListAssignedRolesForUserCmd()
	RoleAssignmentCmd.AddCommand(ListAssignedRolesForUserCmd)
}

var (
	GetUserAssignedRoleuserId string

	GetUserAssignedRoleroleId string

	GetUserAssignedRoleBackup bool
)

func NewGetUserAssignedRoleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getUserAssignedRole",
		Long: "Retrieve a Role assigned to a User",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RoleAssignmentAPI.GetUserAssignedRole(apiClient.GetConfig().Context, GetUserAssignedRoleuserId, GetUserAssignedRoleroleId)

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
			if GetUserAssignedRoleBackup {

				idParam := GetUserAssignedRoleuserId
				err := utils.BackupObject(d, "RoleAssignment", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetUserAssignedRoleuserId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	cmd.Flags().StringVarP(&GetUserAssignedRoleroleId, "roleId", "", "", "")
	cmd.MarkFlagRequired("roleId")

	cmd.Flags().BoolVarP(&GetUserAssignedRoleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetUserAssignedRoleCmd := NewGetUserAssignedRoleCmd()
	RoleAssignmentCmd.AddCommand(GetUserAssignedRoleCmd)
}

var (
	UnassignRoleFromUseruserId string

	UnassignRoleFromUserroleId string

	UnassignRoleFromUserBackup bool
)

func NewUnassignRoleFromUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "unassignRoleFromUser",
		Long: "Unassign a Role from a User",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RoleAssignmentAPI.UnassignRoleFromUser(apiClient.GetConfig().Context, UnassignRoleFromUseruserId, UnassignRoleFromUserroleId)

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
			if UnassignRoleFromUserBackup {

				idParam := UnassignRoleFromUseruserId
				err := utils.BackupObject(d, "RoleAssignment", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UnassignRoleFromUseruserId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	cmd.Flags().StringVarP(&UnassignRoleFromUserroleId, "roleId", "", "", "")
	cmd.MarkFlagRequired("roleId")

	cmd.Flags().BoolVarP(&UnassignRoleFromUserBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UnassignRoleFromUserCmd := NewUnassignRoleFromUserCmd()
	RoleAssignmentCmd.AddCommand(UnassignRoleFromUserCmd)
}
