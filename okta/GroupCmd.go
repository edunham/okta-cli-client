package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var GroupCmd = &cobra.Command{
	Use:  "group",
	Long: "Manage GroupAPI",
}

func init() {
	rootCmd.AddCommand(GroupCmd)
}

var (
	CreateGroupdata string

	CreateGroupBackup bool
)

func NewCreateGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create",
		Long: "Create a Group",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.GroupAPI.CreateGroup(apiClient.GetConfig().Context)

			if CreateGroupdata != "" {
				req = req.Data(CreateGroupdata)
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
			if CreateGroupBackup {

				err := utils.BackupObject(d, "Group", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateGroupdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateGroupBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateGroupCmd := NewCreateGroupCmd()
	GroupCmd.AddCommand(CreateGroupCmd)
}

var ListGroupsBackup bool

func NewListGroupsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "lists",
		Long: "List all Groups",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.GroupAPI.ListGroups(apiClient.GetConfig().Context)

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
			if ListGroupsBackup {

				err := utils.BackupObject(d, "Group", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListGroupsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListGroupsCmd := NewListGroupsCmd()
	GroupCmd.AddCommand(ListGroupsCmd)
}

var (
	CreateGroupRuledata string

	CreateGroupRuleBackup bool
)

func NewCreateGroupRuleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "createRule",
		Long: "Create a Group Rule",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.GroupAPI.CreateGroupRule(apiClient.GetConfig().Context)

			if CreateGroupRuledata != "" {
				req = req.Data(CreateGroupRuledata)
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
			if CreateGroupRuleBackup {

				err := utils.BackupObject(d, "Group", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateGroupRuledata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateGroupRuleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateGroupRuleCmd := NewCreateGroupRuleCmd()
	GroupCmd.AddCommand(CreateGroupRuleCmd)
}

var ListGroupRulesBackup bool

func NewListGroupRulesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listRules",
		Long: "List all Group Rules",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.GroupAPI.ListGroupRules(apiClient.GetConfig().Context)

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
			if ListGroupRulesBackup {

				err := utils.BackupObject(d, "Group", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListGroupRulesBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListGroupRulesCmd := NewListGroupRulesCmd()
	GroupCmd.AddCommand(ListGroupRulesCmd)
}

var (
	GetGroupRulegroupRuleId string

	GetGroupRuleBackup bool
)

func NewGetGroupRuleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getRule",
		Long: "Retrieve a Group Rule",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.GroupAPI.GetGroupRule(apiClient.GetConfig().Context, GetGroupRulegroupRuleId)

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
			if GetGroupRuleBackup {

				idParam := GetGroupRulegroupRuleId
				err := utils.BackupObject(d, "Group", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetGroupRulegroupRuleId, "groupRuleId", "", "", "")
	cmd.MarkFlagRequired("groupRuleId")

	cmd.Flags().BoolVarP(&GetGroupRuleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetGroupRuleCmd := NewGetGroupRuleCmd()
	GroupCmd.AddCommand(GetGroupRuleCmd)
}

var (
	ReplaceGroupRulegroupRuleId string

	ReplaceGroupRuledata string

	ReplaceGroupRuleBackup bool
)

func NewReplaceGroupRuleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replaceRule",
		Long: "Replace a Group Rule",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.GroupAPI.ReplaceGroupRule(apiClient.GetConfig().Context, ReplaceGroupRulegroupRuleId)

			if ReplaceGroupRuledata != "" {
				req = req.Data(ReplaceGroupRuledata)
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
			if ReplaceGroupRuleBackup {

				idParam := ReplaceGroupRulegroupRuleId
				err := utils.BackupObject(d, "Group", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceGroupRulegroupRuleId, "groupRuleId", "", "", "")
	cmd.MarkFlagRequired("groupRuleId")

	cmd.Flags().StringVarP(&ReplaceGroupRuledata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceGroupRuleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceGroupRuleCmd := NewReplaceGroupRuleCmd()
	GroupCmd.AddCommand(ReplaceGroupRuleCmd)
}

var (
	DeleteGroupRulegroupRuleId string

	DeleteGroupRuleBackup bool
)

func NewDeleteGroupRuleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deleteRule",
		Long: "Delete a group Rule",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.GroupAPI.DeleteGroupRule(apiClient.GetConfig().Context, DeleteGroupRulegroupRuleId)

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
			if DeleteGroupRuleBackup {

				idParam := DeleteGroupRulegroupRuleId
				err := utils.BackupObject(d, "Group", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteGroupRulegroupRuleId, "groupRuleId", "", "", "")
	cmd.MarkFlagRequired("groupRuleId")

	cmd.Flags().BoolVarP(&DeleteGroupRuleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteGroupRuleCmd := NewDeleteGroupRuleCmd()
	GroupCmd.AddCommand(DeleteGroupRuleCmd)
}

var (
	ActivateGroupRulegroupRuleId string

	ActivateGroupRuleBackup bool
)

func NewActivateGroupRuleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "activateRule",
		Long: "Activate a Group Rule",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.GroupAPI.ActivateGroupRule(apiClient.GetConfig().Context, ActivateGroupRulegroupRuleId)

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
			if ActivateGroupRuleBackup {

				idParam := ActivateGroupRulegroupRuleId
				err := utils.BackupObject(d, "Group", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ActivateGroupRulegroupRuleId, "groupRuleId", "", "", "")
	cmd.MarkFlagRequired("groupRuleId")

	cmd.Flags().BoolVarP(&ActivateGroupRuleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ActivateGroupRuleCmd := NewActivateGroupRuleCmd()
	GroupCmd.AddCommand(ActivateGroupRuleCmd)
}

var (
	DeactivateGroupRulegroupRuleId string

	DeactivateGroupRuleBackup bool
)

func NewDeactivateGroupRuleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deactivateRule",
		Long: "Deactivate a Group Rule",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.GroupAPI.DeactivateGroupRule(apiClient.GetConfig().Context, DeactivateGroupRulegroupRuleId)

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
			if DeactivateGroupRuleBackup {

				idParam := DeactivateGroupRulegroupRuleId
				err := utils.BackupObject(d, "Group", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeactivateGroupRulegroupRuleId, "groupRuleId", "", "", "")
	cmd.MarkFlagRequired("groupRuleId")

	cmd.Flags().BoolVarP(&DeactivateGroupRuleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeactivateGroupRuleCmd := NewDeactivateGroupRuleCmd()
	GroupCmd.AddCommand(DeactivateGroupRuleCmd)
}

var (
	GetGroupgroupId string

	GetGroupBackup bool
)

func NewGetGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "get",
		Long: "Retrieve a Group",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.GroupAPI.GetGroup(apiClient.GetConfig().Context, GetGroupgroupId)

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
			if GetGroupBackup {

				idParam := GetGroupgroupId
				err := utils.BackupObject(d, "Group", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetGroupgroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	cmd.Flags().BoolVarP(&GetGroupBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetGroupCmd := NewGetGroupCmd()
	GroupCmd.AddCommand(GetGroupCmd)
}

var (
	ReplaceGroupgroupId string

	ReplaceGroupdata string

	ReplaceGroupBackup bool
)

func NewReplaceGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replace",
		Long: "Replace a Group",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.GroupAPI.ReplaceGroup(apiClient.GetConfig().Context, ReplaceGroupgroupId)

			if ReplaceGroupdata != "" {
				req = req.Data(ReplaceGroupdata)
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
			if ReplaceGroupBackup {

				idParam := ReplaceGroupgroupId
				err := utils.BackupObject(d, "Group", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceGroupgroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	cmd.Flags().StringVarP(&ReplaceGroupdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceGroupBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceGroupCmd := NewReplaceGroupCmd()
	GroupCmd.AddCommand(ReplaceGroupCmd)
}

var (
	DeleteGroupgroupId string

	DeleteGroupBackup bool
)

func NewDeleteGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete",
		Long: "Delete a Group",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.GroupAPI.DeleteGroup(apiClient.GetConfig().Context, DeleteGroupgroupId)

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
			if DeleteGroupBackup {

				idParam := DeleteGroupgroupId
				err := utils.BackupObject(d, "Group", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteGroupgroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	cmd.Flags().BoolVarP(&DeleteGroupBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteGroupCmd := NewDeleteGroupCmd()
	GroupCmd.AddCommand(DeleteGroupCmd)
}

var (
	ListAssignedApplicationsForGroupgroupId string

	ListAssignedApplicationsForGroupBackup bool
)

func NewListAssignedApplicationsForGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listAssignedApplicationsFor",
		Long: "List all Assigned Applications",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.GroupAPI.ListAssignedApplicationsForGroup(apiClient.GetConfig().Context, ListAssignedApplicationsForGroupgroupId)

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
			if ListAssignedApplicationsForGroupBackup {

				idParam := ListAssignedApplicationsForGroupgroupId
				err := utils.BackupObject(d, "Group", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListAssignedApplicationsForGroupgroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	cmd.Flags().BoolVarP(&ListAssignedApplicationsForGroupBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListAssignedApplicationsForGroupCmd := NewListAssignedApplicationsForGroupCmd()
	GroupCmd.AddCommand(ListAssignedApplicationsForGroupCmd)
}

var (
	ListGroupUsersgroupId string

	ListGroupUsersBackup bool
)

func NewListGroupUsersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listUsers",
		Long: "List all Member Users",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.GroupAPI.ListGroupUsers(apiClient.GetConfig().Context, ListGroupUsersgroupId)

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
			if ListGroupUsersBackup {

				idParam := ListGroupUsersgroupId
				err := utils.BackupObject(d, "Group", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListGroupUsersgroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	cmd.Flags().BoolVarP(&ListGroupUsersBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListGroupUsersCmd := NewListGroupUsersCmd()
	GroupCmd.AddCommand(ListGroupUsersCmd)
}

var (
	AssignUserToGroupgroupId string

	AssignUserToGroupuserId string

	AssignUserToGroupBackup bool
)

func NewAssignUserToGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "assignUserTo",
		Long: "Assign a User",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.GroupAPI.AssignUserToGroup(apiClient.GetConfig().Context, AssignUserToGroupgroupId, AssignUserToGroupuserId)

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
			if AssignUserToGroupBackup {

				idParam := AssignUserToGroupgroupId
				err := utils.BackupObject(d, "Group", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&AssignUserToGroupgroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	cmd.Flags().StringVarP(&AssignUserToGroupuserId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	cmd.Flags().BoolVarP(&AssignUserToGroupBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	AssignUserToGroupCmd := NewAssignUserToGroupCmd()
	GroupCmd.AddCommand(AssignUserToGroupCmd)
}

var (
	UnassignUserFromGroupgroupId string

	UnassignUserFromGroupuserId string

	UnassignUserFromGroupBackup bool
)

func NewUnassignUserFromGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "unassignUserFrom",
		Long: "Unassign a User",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.GroupAPI.UnassignUserFromGroup(apiClient.GetConfig().Context, UnassignUserFromGroupgroupId, UnassignUserFromGroupuserId)

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
			if UnassignUserFromGroupBackup {

				idParam := UnassignUserFromGroupgroupId
				err := utils.BackupObject(d, "Group", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UnassignUserFromGroupgroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	cmd.Flags().StringVarP(&UnassignUserFromGroupuserId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	cmd.Flags().BoolVarP(&UnassignUserFromGroupBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UnassignUserFromGroupCmd := NewUnassignUserFromGroupCmd()
	GroupCmd.AddCommand(UnassignUserFromGroupCmd)
}
