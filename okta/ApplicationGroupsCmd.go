package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var ApplicationGroupsCmd = &cobra.Command{
	Use:  "applicationGroups",
	Long: "Manage ApplicationGroupsAPI",
}

func init() {
	rootCmd.AddCommand(ApplicationGroupsCmd)
}

var (
	ListApplicationGroupAssignmentsappId string

	ListApplicationGroupAssignmentsBackup bool
)

func NewListApplicationGroupAssignmentsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listApplicationGroupAssignments",
		Long: "List all Assigned Groups",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationGroupsAPI.ListApplicationGroupAssignments(apiClient.GetConfig().Context, ListApplicationGroupAssignmentsappId)

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
			if ListApplicationGroupAssignmentsBackup {

				idParam := ListApplicationGroupAssignmentsappId
				err := utils.BackupObject(d, "ApplicationGroups", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListApplicationGroupAssignmentsappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().BoolVarP(&ListApplicationGroupAssignmentsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListApplicationGroupAssignmentsCmd := NewListApplicationGroupAssignmentsCmd()
	ApplicationGroupsCmd.AddCommand(ListApplicationGroupAssignmentsCmd)
}

var (
	GetApplicationGroupAssignmentappId string

	GetApplicationGroupAssignmentgroupId string

	GetApplicationGroupAssignmentBackup bool
)

func NewGetApplicationGroupAssignmentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getApplicationGroupAssignment",
		Long: "Retrieve an Assigned Group",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationGroupsAPI.GetApplicationGroupAssignment(apiClient.GetConfig().Context, GetApplicationGroupAssignmentappId, GetApplicationGroupAssignmentgroupId)

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
			if GetApplicationGroupAssignmentBackup {

				idParam := GetApplicationGroupAssignmentappId
				err := utils.BackupObject(d, "ApplicationGroups", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetApplicationGroupAssignmentappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().StringVarP(&GetApplicationGroupAssignmentgroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	cmd.Flags().BoolVarP(&GetApplicationGroupAssignmentBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetApplicationGroupAssignmentCmd := NewGetApplicationGroupAssignmentCmd()
	ApplicationGroupsCmd.AddCommand(GetApplicationGroupAssignmentCmd)
}

var (
	AssignGroupToApplicationappId string

	AssignGroupToApplicationgroupId string

	AssignGroupToApplicationdata string

	AssignGroupToApplicationBackup bool
)

func NewAssignGroupToApplicationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "assignGroupToApplication",
		Long: "Assign a Group",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationGroupsAPI.AssignGroupToApplication(apiClient.GetConfig().Context, AssignGroupToApplicationappId, AssignGroupToApplicationgroupId)

			if AssignGroupToApplicationdata != "" {
				req = req.Data(AssignGroupToApplicationdata)
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
			if AssignGroupToApplicationBackup {

				idParam := AssignGroupToApplicationappId
				err := utils.BackupObject(d, "ApplicationGroups", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&AssignGroupToApplicationappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().StringVarP(&AssignGroupToApplicationgroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	cmd.Flags().StringVarP(&AssignGroupToApplicationdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&AssignGroupToApplicationBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	AssignGroupToApplicationCmd := NewAssignGroupToApplicationCmd()
	ApplicationGroupsCmd.AddCommand(AssignGroupToApplicationCmd)
}

var (
	UnassignApplicationFromGroupappId string

	UnassignApplicationFromGroupgroupId string

	UnassignApplicationFromGroupBackup bool
)

func NewUnassignApplicationFromGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "unassignApplicationFromGroup",
		Long: "Unassign a Group",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationGroupsAPI.UnassignApplicationFromGroup(apiClient.GetConfig().Context, UnassignApplicationFromGroupappId, UnassignApplicationFromGroupgroupId)

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
			if UnassignApplicationFromGroupBackup {

				idParam := UnassignApplicationFromGroupappId
				err := utils.BackupObject(d, "ApplicationGroups", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UnassignApplicationFromGroupappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().StringVarP(&UnassignApplicationFromGroupgroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	cmd.Flags().BoolVarP(&UnassignApplicationFromGroupBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UnassignApplicationFromGroupCmd := NewUnassignApplicationFromGroupCmd()
	ApplicationGroupsCmd.AddCommand(UnassignApplicationFromGroupCmd)
}
