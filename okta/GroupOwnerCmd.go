package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var GroupOwnerCmd = &cobra.Command{
	Use:  "groupOwner",
	Long: "Manage GroupOwnerAPI",
}

func init() {
	rootCmd.AddCommand(GroupOwnerCmd)
}

var (
	AssignGroupOwnergroupId string

	AssignGroupOwnerdata string

	AssignGroupOwnerBackup bool
)

func NewAssignGroupOwnerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "assign",
		Long: "Assign a Group Owner",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.GroupOwnerAPI.AssignGroupOwner(apiClient.GetConfig().Context, AssignGroupOwnergroupId)

			if AssignGroupOwnerdata != "" {
				req = req.Data(AssignGroupOwnerdata)
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
			if AssignGroupOwnerBackup {

				idParam := AssignGroupOwnergroupId
				err := utils.BackupObject(d, "GroupOwner", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&AssignGroupOwnergroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	cmd.Flags().StringVarP(&AssignGroupOwnerdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&AssignGroupOwnerBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	AssignGroupOwnerCmd := NewAssignGroupOwnerCmd()
	GroupOwnerCmd.AddCommand(AssignGroupOwnerCmd)
}

var (
	ListGroupOwnersgroupId string

	ListGroupOwnersBackup bool
)

func NewListGroupOwnersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "lists",
		Long: "List all Group Owners",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.GroupOwnerAPI.ListGroupOwners(apiClient.GetConfig().Context, ListGroupOwnersgroupId)

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
			if ListGroupOwnersBackup {

				idParam := ListGroupOwnersgroupId
				err := utils.BackupObject(d, "GroupOwner", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListGroupOwnersgroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	cmd.Flags().BoolVarP(&ListGroupOwnersBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListGroupOwnersCmd := NewListGroupOwnersCmd()
	GroupOwnerCmd.AddCommand(ListGroupOwnersCmd)
}

var (
	DeleteGroupOwnergroupId string

	DeleteGroupOwnerownerId string

	DeleteGroupOwnerBackup bool
)

func NewDeleteGroupOwnerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete",
		Long: "Delete a Group Owner",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.GroupOwnerAPI.DeleteGroupOwner(apiClient.GetConfig().Context, DeleteGroupOwnergroupId, DeleteGroupOwnerownerId)

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
			if DeleteGroupOwnerBackup {

				idParam := DeleteGroupOwnergroupId
				err := utils.BackupObject(d, "GroupOwner", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteGroupOwnergroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	cmd.Flags().StringVarP(&DeleteGroupOwnerownerId, "ownerId", "", "", "")
	cmd.MarkFlagRequired("ownerId")

	cmd.Flags().BoolVarP(&DeleteGroupOwnerBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteGroupOwnerCmd := NewDeleteGroupOwnerCmd()
	GroupOwnerCmd.AddCommand(DeleteGroupOwnerCmd)
}
