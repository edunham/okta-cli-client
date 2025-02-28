package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var RoleTargetCmd = &cobra.Command{
	Use:  "roleTarget",
	Long: "Manage RoleTargetAPI",
}

func init() {
	rootCmd.AddCommand(RoleTargetCmd)
}

var (
	ListApplicationTargetsForApplicationAdministratorRoleForGroupgroupId string

	ListApplicationTargetsForApplicationAdministratorRoleForGrouproleId string

	ListApplicationTargetsForApplicationAdministratorRoleForGroupBackup bool
)

func NewListApplicationTargetsForApplicationAdministratorRoleForGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listApplicationTargetsForApplicationAdministratorRoleForGroup",
		Long: "List all Application Targets for an Application Administrator Role",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RoleTargetAPI.ListApplicationTargetsForApplicationAdministratorRoleForGroup(apiClient.GetConfig().Context, ListApplicationTargetsForApplicationAdministratorRoleForGroupgroupId, ListApplicationTargetsForApplicationAdministratorRoleForGrouproleId)

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
			if ListApplicationTargetsForApplicationAdministratorRoleForGroupBackup {

				idParam := ListApplicationTargetsForApplicationAdministratorRoleForGroupgroupId
				err := utils.BackupObject(d, "RoleTarget", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListApplicationTargetsForApplicationAdministratorRoleForGroupgroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	cmd.Flags().StringVarP(&ListApplicationTargetsForApplicationAdministratorRoleForGrouproleId, "roleId", "", "", "")
	cmd.MarkFlagRequired("roleId")

	cmd.Flags().BoolVarP(&ListApplicationTargetsForApplicationAdministratorRoleForGroupBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListApplicationTargetsForApplicationAdministratorRoleForGroupCmd := NewListApplicationTargetsForApplicationAdministratorRoleForGroupCmd()
	RoleTargetCmd.AddCommand(ListApplicationTargetsForApplicationAdministratorRoleForGroupCmd)
}

var (
	AssignAppTargetToAdminRoleForGroupgroupId string

	AssignAppTargetToAdminRoleForGrouproleId string

	AssignAppTargetToAdminRoleForGroupappName string

	AssignAppTargetToAdminRoleForGroupBackup bool
)

func NewAssignAppTargetToAdminRoleForGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "assignAppTargetToAdminRoleForGroup",
		Long: "Assign an Application Target to Administrator Role",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RoleTargetAPI.AssignAppTargetToAdminRoleForGroup(apiClient.GetConfig().Context, AssignAppTargetToAdminRoleForGroupgroupId, AssignAppTargetToAdminRoleForGrouproleId, AssignAppTargetToAdminRoleForGroupappName)

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
			if AssignAppTargetToAdminRoleForGroupBackup {

				idParam := AssignAppTargetToAdminRoleForGroupgroupId
				err := utils.BackupObject(d, "RoleTarget", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&AssignAppTargetToAdminRoleForGroupgroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	cmd.Flags().StringVarP(&AssignAppTargetToAdminRoleForGrouproleId, "roleId", "", "", "")
	cmd.MarkFlagRequired("roleId")

	cmd.Flags().StringVarP(&AssignAppTargetToAdminRoleForGroupappName, "appName", "", "", "")
	cmd.MarkFlagRequired("appName")

	cmd.Flags().BoolVarP(&AssignAppTargetToAdminRoleForGroupBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	AssignAppTargetToAdminRoleForGroupCmd := NewAssignAppTargetToAdminRoleForGroupCmd()
	RoleTargetCmd.AddCommand(AssignAppTargetToAdminRoleForGroupCmd)
}

var (
	UnassignAppTargetToAdminRoleForGroupgroupId string

	UnassignAppTargetToAdminRoleForGrouproleId string

	UnassignAppTargetToAdminRoleForGroupappName string

	UnassignAppTargetToAdminRoleForGroupBackup bool
)

func NewUnassignAppTargetToAdminRoleForGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "unassignAppTargetToAdminRoleForGroup",
		Long: "Unassign an Application Target from Application Administrator Role",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RoleTargetAPI.UnassignAppTargetToAdminRoleForGroup(apiClient.GetConfig().Context, UnassignAppTargetToAdminRoleForGroupgroupId, UnassignAppTargetToAdminRoleForGrouproleId, UnassignAppTargetToAdminRoleForGroupappName)

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
			if UnassignAppTargetToAdminRoleForGroupBackup {

				idParam := UnassignAppTargetToAdminRoleForGroupgroupId
				err := utils.BackupObject(d, "RoleTarget", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UnassignAppTargetToAdminRoleForGroupgroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	cmd.Flags().StringVarP(&UnassignAppTargetToAdminRoleForGrouproleId, "roleId", "", "", "")
	cmd.MarkFlagRequired("roleId")

	cmd.Flags().StringVarP(&UnassignAppTargetToAdminRoleForGroupappName, "appName", "", "", "")
	cmd.MarkFlagRequired("appName")

	cmd.Flags().BoolVarP(&UnassignAppTargetToAdminRoleForGroupBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UnassignAppTargetToAdminRoleForGroupCmd := NewUnassignAppTargetToAdminRoleForGroupCmd()
	RoleTargetCmd.AddCommand(UnassignAppTargetToAdminRoleForGroupCmd)
}

var (
	AssignAppInstanceTargetToAppAdminRoleForGroupgroupId string

	AssignAppInstanceTargetToAppAdminRoleForGrouproleId string

	AssignAppInstanceTargetToAppAdminRoleForGroupappName string

	AssignAppInstanceTargetToAppAdminRoleForGroupappId string

	AssignAppInstanceTargetToAppAdminRoleForGroupBackup bool
)

func NewAssignAppInstanceTargetToAppAdminRoleForGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "assignAppInstanceTargetToAppAdminRoleForGroup",
		Long: "Assign an Application Instance Target to Application Administrator Role",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RoleTargetAPI.AssignAppInstanceTargetToAppAdminRoleForGroup(apiClient.GetConfig().Context, AssignAppInstanceTargetToAppAdminRoleForGroupgroupId, AssignAppInstanceTargetToAppAdminRoleForGrouproleId, AssignAppInstanceTargetToAppAdminRoleForGroupappName, AssignAppInstanceTargetToAppAdminRoleForGroupappId)

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
			if AssignAppInstanceTargetToAppAdminRoleForGroupBackup {

				idParam := AssignAppInstanceTargetToAppAdminRoleForGroupgroupId
				err := utils.BackupObject(d, "RoleTarget", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&AssignAppInstanceTargetToAppAdminRoleForGroupgroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	cmd.Flags().StringVarP(&AssignAppInstanceTargetToAppAdminRoleForGrouproleId, "roleId", "", "", "")
	cmd.MarkFlagRequired("roleId")

	cmd.Flags().StringVarP(&AssignAppInstanceTargetToAppAdminRoleForGroupappName, "appName", "", "", "")
	cmd.MarkFlagRequired("appName")

	cmd.Flags().StringVarP(&AssignAppInstanceTargetToAppAdminRoleForGroupappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().BoolVarP(&AssignAppInstanceTargetToAppAdminRoleForGroupBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	AssignAppInstanceTargetToAppAdminRoleForGroupCmd := NewAssignAppInstanceTargetToAppAdminRoleForGroupCmd()
	RoleTargetCmd.AddCommand(AssignAppInstanceTargetToAppAdminRoleForGroupCmd)
}

var (
	UnassignAppInstanceTargetToAppAdminRoleForGroupgroupId string

	UnassignAppInstanceTargetToAppAdminRoleForGrouproleId string

	UnassignAppInstanceTargetToAppAdminRoleForGroupappName string

	UnassignAppInstanceTargetToAppAdminRoleForGroupappId string

	UnassignAppInstanceTargetToAppAdminRoleForGroupBackup bool
)

func NewUnassignAppInstanceTargetToAppAdminRoleForGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "unassignAppInstanceTargetToAppAdminRoleForGroup",
		Long: "Unassign an Application Instance Target from an Application Administrator Role",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RoleTargetAPI.UnassignAppInstanceTargetToAppAdminRoleForGroup(apiClient.GetConfig().Context, UnassignAppInstanceTargetToAppAdminRoleForGroupgroupId, UnassignAppInstanceTargetToAppAdminRoleForGrouproleId, UnassignAppInstanceTargetToAppAdminRoleForGroupappName, UnassignAppInstanceTargetToAppAdminRoleForGroupappId)

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
			if UnassignAppInstanceTargetToAppAdminRoleForGroupBackup {

				idParam := UnassignAppInstanceTargetToAppAdminRoleForGroupgroupId
				err := utils.BackupObject(d, "RoleTarget", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UnassignAppInstanceTargetToAppAdminRoleForGroupgroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	cmd.Flags().StringVarP(&UnassignAppInstanceTargetToAppAdminRoleForGrouproleId, "roleId", "", "", "")
	cmd.MarkFlagRequired("roleId")

	cmd.Flags().StringVarP(&UnassignAppInstanceTargetToAppAdminRoleForGroupappName, "appName", "", "", "")
	cmd.MarkFlagRequired("appName")

	cmd.Flags().StringVarP(&UnassignAppInstanceTargetToAppAdminRoleForGroupappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().BoolVarP(&UnassignAppInstanceTargetToAppAdminRoleForGroupBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UnassignAppInstanceTargetToAppAdminRoleForGroupCmd := NewUnassignAppInstanceTargetToAppAdminRoleForGroupCmd()
	RoleTargetCmd.AddCommand(UnassignAppInstanceTargetToAppAdminRoleForGroupCmd)
}

var (
	ListGroupTargetsForGroupRolegroupId string

	ListGroupTargetsForGroupRoleroleId string

	ListGroupTargetsForGroupRoleBackup bool
)

func NewListGroupTargetsForGroupRoleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listGroupTargetsForGroupRole",
		Long: "List all Group Targets for a Group Role",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RoleTargetAPI.ListGroupTargetsForGroupRole(apiClient.GetConfig().Context, ListGroupTargetsForGroupRolegroupId, ListGroupTargetsForGroupRoleroleId)

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
			if ListGroupTargetsForGroupRoleBackup {

				idParam := ListGroupTargetsForGroupRolegroupId
				err := utils.BackupObject(d, "RoleTarget", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListGroupTargetsForGroupRolegroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	cmd.Flags().StringVarP(&ListGroupTargetsForGroupRoleroleId, "roleId", "", "", "")
	cmd.MarkFlagRequired("roleId")

	cmd.Flags().BoolVarP(&ListGroupTargetsForGroupRoleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListGroupTargetsForGroupRoleCmd := NewListGroupTargetsForGroupRoleCmd()
	RoleTargetCmd.AddCommand(ListGroupTargetsForGroupRoleCmd)
}

var (
	AssignGroupTargetToGroupAdminRolegroupId string

	AssignGroupTargetToGroupAdminRoleroleId string

	AssignGroupTargetToGroupAdminRoletargetGroupId string

	AssignGroupTargetToGroupAdminRoleBackup bool
)

func NewAssignGroupTargetToGroupAdminRoleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "assignGroupTargetToGroupAdminRole",
		Long: "Assign a Group Target to a Group Role",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RoleTargetAPI.AssignGroupTargetToGroupAdminRole(apiClient.GetConfig().Context, AssignGroupTargetToGroupAdminRolegroupId, AssignGroupTargetToGroupAdminRoleroleId, AssignGroupTargetToGroupAdminRoletargetGroupId)

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
			if AssignGroupTargetToGroupAdminRoleBackup {

				idParam := AssignGroupTargetToGroupAdminRolegroupId
				err := utils.BackupObject(d, "RoleTarget", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&AssignGroupTargetToGroupAdminRolegroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	cmd.Flags().StringVarP(&AssignGroupTargetToGroupAdminRoleroleId, "roleId", "", "", "")
	cmd.MarkFlagRequired("roleId")

	cmd.Flags().StringVarP(&AssignGroupTargetToGroupAdminRoletargetGroupId, "targetGroupId", "", "", "")
	cmd.MarkFlagRequired("targetGroupId")

	cmd.Flags().BoolVarP(&AssignGroupTargetToGroupAdminRoleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	AssignGroupTargetToGroupAdminRoleCmd := NewAssignGroupTargetToGroupAdminRoleCmd()
	RoleTargetCmd.AddCommand(AssignGroupTargetToGroupAdminRoleCmd)
}

var (
	UnassignGroupTargetFromGroupAdminRolegroupId string

	UnassignGroupTargetFromGroupAdminRoleroleId string

	UnassignGroupTargetFromGroupAdminRoletargetGroupId string

	UnassignGroupTargetFromGroupAdminRoleBackup bool
)

func NewUnassignGroupTargetFromGroupAdminRoleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "unassignGroupTargetFromGroupAdminRole",
		Long: "Unassign a Group Target from a Group Role",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RoleTargetAPI.UnassignGroupTargetFromGroupAdminRole(apiClient.GetConfig().Context, UnassignGroupTargetFromGroupAdminRolegroupId, UnassignGroupTargetFromGroupAdminRoleroleId, UnassignGroupTargetFromGroupAdminRoletargetGroupId)

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
			if UnassignGroupTargetFromGroupAdminRoleBackup {

				idParam := UnassignGroupTargetFromGroupAdminRolegroupId
				err := utils.BackupObject(d, "RoleTarget", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UnassignGroupTargetFromGroupAdminRolegroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	cmd.Flags().StringVarP(&UnassignGroupTargetFromGroupAdminRoleroleId, "roleId", "", "", "")
	cmd.MarkFlagRequired("roleId")

	cmd.Flags().StringVarP(&UnassignGroupTargetFromGroupAdminRoletargetGroupId, "targetGroupId", "", "", "")
	cmd.MarkFlagRequired("targetGroupId")

	cmd.Flags().BoolVarP(&UnassignGroupTargetFromGroupAdminRoleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UnassignGroupTargetFromGroupAdminRoleCmd := NewUnassignGroupTargetFromGroupAdminRoleCmd()
	RoleTargetCmd.AddCommand(UnassignGroupTargetFromGroupAdminRoleCmd)
}

var (
	ListApplicationTargetsForApplicationAdministratorRoleForUseruserId string

	ListApplicationTargetsForApplicationAdministratorRoleForUserroleId string

	ListApplicationTargetsForApplicationAdministratorRoleForUserBackup bool
)

func NewListApplicationTargetsForApplicationAdministratorRoleForUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listApplicationTargetsForApplicationAdministratorRoleForUser",
		Long: "List all Application Targets for Application Administrator Role",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RoleTargetAPI.ListApplicationTargetsForApplicationAdministratorRoleForUser(apiClient.GetConfig().Context, ListApplicationTargetsForApplicationAdministratorRoleForUseruserId, ListApplicationTargetsForApplicationAdministratorRoleForUserroleId)

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
			if ListApplicationTargetsForApplicationAdministratorRoleForUserBackup {

				idParam := ListApplicationTargetsForApplicationAdministratorRoleForUseruserId
				err := utils.BackupObject(d, "RoleTarget", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListApplicationTargetsForApplicationAdministratorRoleForUseruserId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	cmd.Flags().StringVarP(&ListApplicationTargetsForApplicationAdministratorRoleForUserroleId, "roleId", "", "", "")
	cmd.MarkFlagRequired("roleId")

	cmd.Flags().BoolVarP(&ListApplicationTargetsForApplicationAdministratorRoleForUserBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListApplicationTargetsForApplicationAdministratorRoleForUserCmd := NewListApplicationTargetsForApplicationAdministratorRoleForUserCmd()
	RoleTargetCmd.AddCommand(ListApplicationTargetsForApplicationAdministratorRoleForUserCmd)
}

var (
	AssignAllAppsAsTargetToRoleForUseruserId string

	AssignAllAppsAsTargetToRoleForUserroleId string

	AssignAllAppsAsTargetToRoleForUserBackup bool
)

func NewAssignAllAppsAsTargetToRoleForUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "assignAllAppsAsTargetToRoleForUser",
		Long: "Assign all Apps as Target to Role",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RoleTargetAPI.AssignAllAppsAsTargetToRoleForUser(apiClient.GetConfig().Context, AssignAllAppsAsTargetToRoleForUseruserId, AssignAllAppsAsTargetToRoleForUserroleId)

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
			if AssignAllAppsAsTargetToRoleForUserBackup {

				idParam := AssignAllAppsAsTargetToRoleForUseruserId
				err := utils.BackupObject(d, "RoleTarget", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&AssignAllAppsAsTargetToRoleForUseruserId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	cmd.Flags().StringVarP(&AssignAllAppsAsTargetToRoleForUserroleId, "roleId", "", "", "")
	cmd.MarkFlagRequired("roleId")

	cmd.Flags().BoolVarP(&AssignAllAppsAsTargetToRoleForUserBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	AssignAllAppsAsTargetToRoleForUserCmd := NewAssignAllAppsAsTargetToRoleForUserCmd()
	RoleTargetCmd.AddCommand(AssignAllAppsAsTargetToRoleForUserCmd)
}

var (
	AssignAppTargetToAdminRoleForUseruserId string

	AssignAppTargetToAdminRoleForUserroleId string

	AssignAppTargetToAdminRoleForUserappName string

	AssignAppTargetToAdminRoleForUserBackup bool
)

func NewAssignAppTargetToAdminRoleForUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "assignAppTargetToAdminRoleForUser",
		Long: "Assign an Application Target to Administrator Role",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RoleTargetAPI.AssignAppTargetToAdminRoleForUser(apiClient.GetConfig().Context, AssignAppTargetToAdminRoleForUseruserId, AssignAppTargetToAdminRoleForUserroleId, AssignAppTargetToAdminRoleForUserappName)

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
			if AssignAppTargetToAdminRoleForUserBackup {

				idParam := AssignAppTargetToAdminRoleForUseruserId
				err := utils.BackupObject(d, "RoleTarget", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&AssignAppTargetToAdminRoleForUseruserId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	cmd.Flags().StringVarP(&AssignAppTargetToAdminRoleForUserroleId, "roleId", "", "", "")
	cmd.MarkFlagRequired("roleId")

	cmd.Flags().StringVarP(&AssignAppTargetToAdminRoleForUserappName, "appName", "", "", "")
	cmd.MarkFlagRequired("appName")

	cmd.Flags().BoolVarP(&AssignAppTargetToAdminRoleForUserBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	AssignAppTargetToAdminRoleForUserCmd := NewAssignAppTargetToAdminRoleForUserCmd()
	RoleTargetCmd.AddCommand(AssignAppTargetToAdminRoleForUserCmd)
}

var (
	UnassignAppTargetFromAppAdminRoleForUseruserId string

	UnassignAppTargetFromAppAdminRoleForUserroleId string

	UnassignAppTargetFromAppAdminRoleForUserappName string

	UnassignAppTargetFromAppAdminRoleForUserBackup bool
)

func NewUnassignAppTargetFromAppAdminRoleForUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "unassignAppTargetFromAppAdminRoleForUser",
		Long: "Unassign an Application Target from an Application Administrator Role",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RoleTargetAPI.UnassignAppTargetFromAppAdminRoleForUser(apiClient.GetConfig().Context, UnassignAppTargetFromAppAdminRoleForUseruserId, UnassignAppTargetFromAppAdminRoleForUserroleId, UnassignAppTargetFromAppAdminRoleForUserappName)

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
			if UnassignAppTargetFromAppAdminRoleForUserBackup {

				idParam := UnassignAppTargetFromAppAdminRoleForUseruserId
				err := utils.BackupObject(d, "RoleTarget", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UnassignAppTargetFromAppAdminRoleForUseruserId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	cmd.Flags().StringVarP(&UnassignAppTargetFromAppAdminRoleForUserroleId, "roleId", "", "", "")
	cmd.MarkFlagRequired("roleId")

	cmd.Flags().StringVarP(&UnassignAppTargetFromAppAdminRoleForUserappName, "appName", "", "", "")
	cmd.MarkFlagRequired("appName")

	cmd.Flags().BoolVarP(&UnassignAppTargetFromAppAdminRoleForUserBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UnassignAppTargetFromAppAdminRoleForUserCmd := NewUnassignAppTargetFromAppAdminRoleForUserCmd()
	RoleTargetCmd.AddCommand(UnassignAppTargetFromAppAdminRoleForUserCmd)
}

var (
	AssignAppInstanceTargetToAppAdminRoleForUseruserId string

	AssignAppInstanceTargetToAppAdminRoleForUserroleId string

	AssignAppInstanceTargetToAppAdminRoleForUserappName string

	AssignAppInstanceTargetToAppAdminRoleForUserappId string

	AssignAppInstanceTargetToAppAdminRoleForUserBackup bool
)

func NewAssignAppInstanceTargetToAppAdminRoleForUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "assignAppInstanceTargetToAppAdminRoleForUser",
		Long: "Assign an Application Instance Target to an Application Administrator Role",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RoleTargetAPI.AssignAppInstanceTargetToAppAdminRoleForUser(apiClient.GetConfig().Context, AssignAppInstanceTargetToAppAdminRoleForUseruserId, AssignAppInstanceTargetToAppAdminRoleForUserroleId, AssignAppInstanceTargetToAppAdminRoleForUserappName, AssignAppInstanceTargetToAppAdminRoleForUserappId)

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
			if AssignAppInstanceTargetToAppAdminRoleForUserBackup {

				idParam := AssignAppInstanceTargetToAppAdminRoleForUseruserId
				err := utils.BackupObject(d, "RoleTarget", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&AssignAppInstanceTargetToAppAdminRoleForUseruserId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	cmd.Flags().StringVarP(&AssignAppInstanceTargetToAppAdminRoleForUserroleId, "roleId", "", "", "")
	cmd.MarkFlagRequired("roleId")

	cmd.Flags().StringVarP(&AssignAppInstanceTargetToAppAdminRoleForUserappName, "appName", "", "", "")
	cmd.MarkFlagRequired("appName")

	cmd.Flags().StringVarP(&AssignAppInstanceTargetToAppAdminRoleForUserappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().BoolVarP(&AssignAppInstanceTargetToAppAdminRoleForUserBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	AssignAppInstanceTargetToAppAdminRoleForUserCmd := NewAssignAppInstanceTargetToAppAdminRoleForUserCmd()
	RoleTargetCmd.AddCommand(AssignAppInstanceTargetToAppAdminRoleForUserCmd)
}

var (
	UnassignAppInstanceTargetFromAdminRoleForUseruserId string

	UnassignAppInstanceTargetFromAdminRoleForUserroleId string

	UnassignAppInstanceTargetFromAdminRoleForUserappName string

	UnassignAppInstanceTargetFromAdminRoleForUserappId string

	UnassignAppInstanceTargetFromAdminRoleForUserBackup bool
)

func NewUnassignAppInstanceTargetFromAdminRoleForUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "unassignAppInstanceTargetFromAdminRoleForUser",
		Long: "Unassign an Application Instance Target from an Application Administrator Role",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RoleTargetAPI.UnassignAppInstanceTargetFromAdminRoleForUser(apiClient.GetConfig().Context, UnassignAppInstanceTargetFromAdminRoleForUseruserId, UnassignAppInstanceTargetFromAdminRoleForUserroleId, UnassignAppInstanceTargetFromAdminRoleForUserappName, UnassignAppInstanceTargetFromAdminRoleForUserappId)

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
			if UnassignAppInstanceTargetFromAdminRoleForUserBackup {

				idParam := UnassignAppInstanceTargetFromAdminRoleForUseruserId
				err := utils.BackupObject(d, "RoleTarget", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UnassignAppInstanceTargetFromAdminRoleForUseruserId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	cmd.Flags().StringVarP(&UnassignAppInstanceTargetFromAdminRoleForUserroleId, "roleId", "", "", "")
	cmd.MarkFlagRequired("roleId")

	cmd.Flags().StringVarP(&UnassignAppInstanceTargetFromAdminRoleForUserappName, "appName", "", "", "")
	cmd.MarkFlagRequired("appName")

	cmd.Flags().StringVarP(&UnassignAppInstanceTargetFromAdminRoleForUserappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().BoolVarP(&UnassignAppInstanceTargetFromAdminRoleForUserBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UnassignAppInstanceTargetFromAdminRoleForUserCmd := NewUnassignAppInstanceTargetFromAdminRoleForUserCmd()
	RoleTargetCmd.AddCommand(UnassignAppInstanceTargetFromAdminRoleForUserCmd)
}

var (
	ListGroupTargetsForRoleuserId string

	ListGroupTargetsForRoleroleId string

	ListGroupTargetsForRoleBackup bool
)

func NewListGroupTargetsForRoleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listGroupTargetsForRole",
		Long: "List all Group Targets for Role",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RoleTargetAPI.ListGroupTargetsForRole(apiClient.GetConfig().Context, ListGroupTargetsForRoleuserId, ListGroupTargetsForRoleroleId)

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
			if ListGroupTargetsForRoleBackup {

				idParam := ListGroupTargetsForRoleuserId
				err := utils.BackupObject(d, "RoleTarget", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListGroupTargetsForRoleuserId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	cmd.Flags().StringVarP(&ListGroupTargetsForRoleroleId, "roleId", "", "", "")
	cmd.MarkFlagRequired("roleId")

	cmd.Flags().BoolVarP(&ListGroupTargetsForRoleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListGroupTargetsForRoleCmd := NewListGroupTargetsForRoleCmd()
	RoleTargetCmd.AddCommand(ListGroupTargetsForRoleCmd)
}

var (
	AssignGroupTargetToUserRoleuserId string

	AssignGroupTargetToUserRoleroleId string

	AssignGroupTargetToUserRolegroupId string

	AssignGroupTargetToUserRoleBackup bool
)

func NewAssignGroupTargetToUserRoleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "assignGroupTargetToUserRole",
		Long: "Assign a Group Target to Role",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RoleTargetAPI.AssignGroupTargetToUserRole(apiClient.GetConfig().Context, AssignGroupTargetToUserRoleuserId, AssignGroupTargetToUserRoleroleId, AssignGroupTargetToUserRolegroupId)

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
			if AssignGroupTargetToUserRoleBackup {

				idParam := AssignGroupTargetToUserRoleuserId
				err := utils.BackupObject(d, "RoleTarget", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&AssignGroupTargetToUserRoleuserId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	cmd.Flags().StringVarP(&AssignGroupTargetToUserRoleroleId, "roleId", "", "", "")
	cmd.MarkFlagRequired("roleId")

	cmd.Flags().StringVarP(&AssignGroupTargetToUserRolegroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	cmd.Flags().BoolVarP(&AssignGroupTargetToUserRoleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	AssignGroupTargetToUserRoleCmd := NewAssignGroupTargetToUserRoleCmd()
	RoleTargetCmd.AddCommand(AssignGroupTargetToUserRoleCmd)
}

var (
	UnassignGroupTargetFromUserAdminRoleuserId string

	UnassignGroupTargetFromUserAdminRoleroleId string

	UnassignGroupTargetFromUserAdminRolegroupId string

	UnassignGroupTargetFromUserAdminRoleBackup bool
)

func NewUnassignGroupTargetFromUserAdminRoleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "unassignGroupTargetFromUserAdminRole",
		Long: "Unassign a Group Target from Role",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RoleTargetAPI.UnassignGroupTargetFromUserAdminRole(apiClient.GetConfig().Context, UnassignGroupTargetFromUserAdminRoleuserId, UnassignGroupTargetFromUserAdminRoleroleId, UnassignGroupTargetFromUserAdminRolegroupId)

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
			if UnassignGroupTargetFromUserAdminRoleBackup {

				idParam := UnassignGroupTargetFromUserAdminRoleuserId
				err := utils.BackupObject(d, "RoleTarget", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UnassignGroupTargetFromUserAdminRoleuserId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	cmd.Flags().StringVarP(&UnassignGroupTargetFromUserAdminRoleroleId, "roleId", "", "", "")
	cmd.MarkFlagRequired("roleId")

	cmd.Flags().StringVarP(&UnassignGroupTargetFromUserAdminRolegroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	cmd.Flags().BoolVarP(&UnassignGroupTargetFromUserAdminRoleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UnassignGroupTargetFromUserAdminRoleCmd := NewUnassignGroupTargetFromUserAdminRoleCmd()
	RoleTargetCmd.AddCommand(UnassignGroupTargetFromUserAdminRoleCmd)
}
