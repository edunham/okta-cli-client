package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var ResourceSetCmd = &cobra.Command{
	Use:  "resourceSet",
	Long: "Manage ResourceSetAPI",
}

func init() {
	rootCmd.AddCommand(ResourceSetCmd)
}

var (
	CreateResourceSetdata string

	CreateResourceSetBackup bool
)

func NewCreateResourceSetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create",
		Long: "Create a Resource Set",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ResourceSetAPI.CreateResourceSet(apiClient.GetConfig().Context)

			if CreateResourceSetdata != "" {
				req = req.Data(CreateResourceSetdata)
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
			if CreateResourceSetBackup {

				err := utils.BackupObject(d, "ResourceSet", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateResourceSetdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateResourceSetBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateResourceSetCmd := NewCreateResourceSetCmd()
	ResourceSetCmd.AddCommand(CreateResourceSetCmd)
}

var ListResourceSetsBackup bool

func NewListResourceSetsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "lists",
		Long: "List all Resource Sets",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ResourceSetAPI.ListResourceSets(apiClient.GetConfig().Context)

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
			if ListResourceSetsBackup {

				err := utils.BackupObject(d, "ResourceSet", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListResourceSetsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListResourceSetsCmd := NewListResourceSetsCmd()
	ResourceSetCmd.AddCommand(ListResourceSetsCmd)
}

var (
	GetResourceSetresourceSetId string

	GetResourceSetBackup bool
)

func NewGetResourceSetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "get",
		Long: "Retrieve a Resource Set",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ResourceSetAPI.GetResourceSet(apiClient.GetConfig().Context, GetResourceSetresourceSetId)

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
			if GetResourceSetBackup {

				idParam := GetResourceSetresourceSetId
				err := utils.BackupObject(d, "ResourceSet", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetResourceSetresourceSetId, "resourceSetId", "", "", "")
	cmd.MarkFlagRequired("resourceSetId")

	cmd.Flags().BoolVarP(&GetResourceSetBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetResourceSetCmd := NewGetResourceSetCmd()
	ResourceSetCmd.AddCommand(GetResourceSetCmd)
}

var (
	ReplaceResourceSetresourceSetId string

	ReplaceResourceSetdata string

	ReplaceResourceSetBackup bool
)

func NewReplaceResourceSetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replace",
		Long: "Replace a Resource Set",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ResourceSetAPI.ReplaceResourceSet(apiClient.GetConfig().Context, ReplaceResourceSetresourceSetId)

			if ReplaceResourceSetdata != "" {
				req = req.Data(ReplaceResourceSetdata)
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
			if ReplaceResourceSetBackup {

				idParam := ReplaceResourceSetresourceSetId
				err := utils.BackupObject(d, "ResourceSet", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceResourceSetresourceSetId, "resourceSetId", "", "", "")
	cmd.MarkFlagRequired("resourceSetId")

	cmd.Flags().StringVarP(&ReplaceResourceSetdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceResourceSetBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceResourceSetCmd := NewReplaceResourceSetCmd()
	ResourceSetCmd.AddCommand(ReplaceResourceSetCmd)
}

var (
	DeleteResourceSetresourceSetId string

	DeleteResourceSetBackup bool
)

func NewDeleteResourceSetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete",
		Long: "Delete a Resource Set",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ResourceSetAPI.DeleteResourceSet(apiClient.GetConfig().Context, DeleteResourceSetresourceSetId)

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
			if DeleteResourceSetBackup {

				idParam := DeleteResourceSetresourceSetId
				err := utils.BackupObject(d, "ResourceSet", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteResourceSetresourceSetId, "resourceSetId", "", "", "")
	cmd.MarkFlagRequired("resourceSetId")

	cmd.Flags().BoolVarP(&DeleteResourceSetBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteResourceSetCmd := NewDeleteResourceSetCmd()
	ResourceSetCmd.AddCommand(DeleteResourceSetCmd)
}

var (
	CreateResourceSetBindingresourceSetId string

	CreateResourceSetBindingdata string

	CreateResourceSetBindingBackup bool
)

func NewCreateResourceSetBindingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "createBinding",
		Long: "Create a Resource Set Binding",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ResourceSetAPI.CreateResourceSetBinding(apiClient.GetConfig().Context, CreateResourceSetBindingresourceSetId)

			if CreateResourceSetBindingdata != "" {
				req = req.Data(CreateResourceSetBindingdata)
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
			if CreateResourceSetBindingBackup {

				idParam := CreateResourceSetBindingresourceSetId
				err := utils.BackupObject(d, "ResourceSet", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateResourceSetBindingresourceSetId, "resourceSetId", "", "", "")
	cmd.MarkFlagRequired("resourceSetId")

	cmd.Flags().StringVarP(&CreateResourceSetBindingdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateResourceSetBindingBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateResourceSetBindingCmd := NewCreateResourceSetBindingCmd()
	ResourceSetCmd.AddCommand(CreateResourceSetBindingCmd)
}

var (
	ListBindingsresourceSetId string

	ListBindingsBackup bool
)

func NewListBindingsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listBindings",
		Long: "List all Bindings",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ResourceSetAPI.ListBindings(apiClient.GetConfig().Context, ListBindingsresourceSetId)

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
			if ListBindingsBackup {

				idParam := ListBindingsresourceSetId
				err := utils.BackupObject(d, "ResourceSet", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListBindingsresourceSetId, "resourceSetId", "", "", "")
	cmd.MarkFlagRequired("resourceSetId")

	cmd.Flags().BoolVarP(&ListBindingsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListBindingsCmd := NewListBindingsCmd()
	ResourceSetCmd.AddCommand(ListBindingsCmd)
}

var (
	GetBindingresourceSetId string

	GetBindingroleIdOrLabel string

	GetBindingBackup bool
)

func NewGetBindingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getBinding",
		Long: "Retrieve a Binding",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ResourceSetAPI.GetBinding(apiClient.GetConfig().Context, GetBindingresourceSetId, GetBindingroleIdOrLabel)

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
			if GetBindingBackup {

				idParam := GetBindingresourceSetId
				err := utils.BackupObject(d, "ResourceSet", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetBindingresourceSetId, "resourceSetId", "", "", "")
	cmd.MarkFlagRequired("resourceSetId")

	cmd.Flags().StringVarP(&GetBindingroleIdOrLabel, "roleIdOrLabel", "", "", "")
	cmd.MarkFlagRequired("roleIdOrLabel")

	cmd.Flags().BoolVarP(&GetBindingBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetBindingCmd := NewGetBindingCmd()
	ResourceSetCmd.AddCommand(GetBindingCmd)
}

var (
	DeleteBindingresourceSetId string

	DeleteBindingroleIdOrLabel string

	DeleteBindingBackup bool
)

func NewDeleteBindingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deleteBinding",
		Long: "Delete a Binding",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ResourceSetAPI.DeleteBinding(apiClient.GetConfig().Context, DeleteBindingresourceSetId, DeleteBindingroleIdOrLabel)

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
			if DeleteBindingBackup {

				idParam := DeleteBindingresourceSetId
				err := utils.BackupObject(d, "ResourceSet", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteBindingresourceSetId, "resourceSetId", "", "", "")
	cmd.MarkFlagRequired("resourceSetId")

	cmd.Flags().StringVarP(&DeleteBindingroleIdOrLabel, "roleIdOrLabel", "", "", "")
	cmd.MarkFlagRequired("roleIdOrLabel")

	cmd.Flags().BoolVarP(&DeleteBindingBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteBindingCmd := NewDeleteBindingCmd()
	ResourceSetCmd.AddCommand(DeleteBindingCmd)
}

var (
	ListMembersOfBindingresourceSetId string

	ListMembersOfBindingroleIdOrLabel string

	ListMembersOfBindingBackup bool
)

func NewListMembersOfBindingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listMembersOfBinding",
		Long: "List all Members of a binding",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ResourceSetAPI.ListMembersOfBinding(apiClient.GetConfig().Context, ListMembersOfBindingresourceSetId, ListMembersOfBindingroleIdOrLabel)

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
			if ListMembersOfBindingBackup {

				idParam := ListMembersOfBindingresourceSetId
				err := utils.BackupObject(d, "ResourceSet", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListMembersOfBindingresourceSetId, "resourceSetId", "", "", "")
	cmd.MarkFlagRequired("resourceSetId")

	cmd.Flags().StringVarP(&ListMembersOfBindingroleIdOrLabel, "roleIdOrLabel", "", "", "")
	cmd.MarkFlagRequired("roleIdOrLabel")

	cmd.Flags().BoolVarP(&ListMembersOfBindingBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListMembersOfBindingCmd := NewListMembersOfBindingCmd()
	ResourceSetCmd.AddCommand(ListMembersOfBindingCmd)
}

var (
	AddMembersToBindingresourceSetId string

	AddMembersToBindingroleIdOrLabel string

	AddMembersToBindingdata string

	AddMembersToBindingBackup bool
)

func NewAddMembersToBindingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "addMembersToBinding",
		Long: "Add more Members to a binding",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ResourceSetAPI.AddMembersToBinding(apiClient.GetConfig().Context, AddMembersToBindingresourceSetId, AddMembersToBindingroleIdOrLabel)

			if AddMembersToBindingdata != "" {
				req = req.Data(AddMembersToBindingdata)
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
			if AddMembersToBindingBackup {

				idParam := AddMembersToBindingresourceSetId
				err := utils.BackupObject(d, "ResourceSet", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&AddMembersToBindingresourceSetId, "resourceSetId", "", "", "")
	cmd.MarkFlagRequired("resourceSetId")

	cmd.Flags().StringVarP(&AddMembersToBindingroleIdOrLabel, "roleIdOrLabel", "", "", "")
	cmd.MarkFlagRequired("roleIdOrLabel")

	cmd.Flags().StringVarP(&AddMembersToBindingdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&AddMembersToBindingBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	AddMembersToBindingCmd := NewAddMembersToBindingCmd()
	ResourceSetCmd.AddCommand(AddMembersToBindingCmd)
}

var (
	GetMemberOfBindingresourceSetId string

	GetMemberOfBindingroleIdOrLabel string

	GetMemberOfBindingmemberId string

	GetMemberOfBindingBackup bool
)

func NewGetMemberOfBindingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getMemberOfBinding",
		Long: "Retrieve a Member of a binding",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ResourceSetAPI.GetMemberOfBinding(apiClient.GetConfig().Context, GetMemberOfBindingresourceSetId, GetMemberOfBindingroleIdOrLabel, GetMemberOfBindingmemberId)

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
			if GetMemberOfBindingBackup {

				idParam := GetMemberOfBindingresourceSetId
				err := utils.BackupObject(d, "ResourceSet", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetMemberOfBindingresourceSetId, "resourceSetId", "", "", "")
	cmd.MarkFlagRequired("resourceSetId")

	cmd.Flags().StringVarP(&GetMemberOfBindingroleIdOrLabel, "roleIdOrLabel", "", "", "")
	cmd.MarkFlagRequired("roleIdOrLabel")

	cmd.Flags().StringVarP(&GetMemberOfBindingmemberId, "memberId", "", "", "")
	cmd.MarkFlagRequired("memberId")

	cmd.Flags().BoolVarP(&GetMemberOfBindingBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetMemberOfBindingCmd := NewGetMemberOfBindingCmd()
	ResourceSetCmd.AddCommand(GetMemberOfBindingCmd)
}

var (
	UnassignMemberFromBindingresourceSetId string

	UnassignMemberFromBindingroleIdOrLabel string

	UnassignMemberFromBindingmemberId string

	UnassignMemberFromBindingBackup bool
)

func NewUnassignMemberFromBindingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "unassignMemberFromBinding",
		Long: "Unassign a Member from a binding",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ResourceSetAPI.UnassignMemberFromBinding(apiClient.GetConfig().Context, UnassignMemberFromBindingresourceSetId, UnassignMemberFromBindingroleIdOrLabel, UnassignMemberFromBindingmemberId)

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
			if UnassignMemberFromBindingBackup {

				idParam := UnassignMemberFromBindingresourceSetId
				err := utils.BackupObject(d, "ResourceSet", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UnassignMemberFromBindingresourceSetId, "resourceSetId", "", "", "")
	cmd.MarkFlagRequired("resourceSetId")

	cmd.Flags().StringVarP(&UnassignMemberFromBindingroleIdOrLabel, "roleIdOrLabel", "", "", "")
	cmd.MarkFlagRequired("roleIdOrLabel")

	cmd.Flags().StringVarP(&UnassignMemberFromBindingmemberId, "memberId", "", "", "")
	cmd.MarkFlagRequired("memberId")

	cmd.Flags().BoolVarP(&UnassignMemberFromBindingBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UnassignMemberFromBindingCmd := NewUnassignMemberFromBindingCmd()
	ResourceSetCmd.AddCommand(UnassignMemberFromBindingCmd)
}

var (
	ListResourceSetResourcesresourceSetId string

	ListResourceSetResourcesBackup bool
)

func NewListResourceSetResourcesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listResources",
		Long: "List all Resources of a Resource Set",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ResourceSetAPI.ListResourceSetResources(apiClient.GetConfig().Context, ListResourceSetResourcesresourceSetId)

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
			if ListResourceSetResourcesBackup {

				idParam := ListResourceSetResourcesresourceSetId
				err := utils.BackupObject(d, "ResourceSet", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListResourceSetResourcesresourceSetId, "resourceSetId", "", "", "")
	cmd.MarkFlagRequired("resourceSetId")

	cmd.Flags().BoolVarP(&ListResourceSetResourcesBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListResourceSetResourcesCmd := NewListResourceSetResourcesCmd()
	ResourceSetCmd.AddCommand(ListResourceSetResourcesCmd)
}

var (
	AddResourceSetResourceresourceSetId string

	AddResourceSetResourcedata string

	AddResourceSetResourceBackup bool
)

func NewAddResourceSetResourceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "addResource",
		Long: "Add more Resource to a Resource Set",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ResourceSetAPI.AddResourceSetResource(apiClient.GetConfig().Context, AddResourceSetResourceresourceSetId)

			if AddResourceSetResourcedata != "" {
				req = req.Data(AddResourceSetResourcedata)
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
			if AddResourceSetResourceBackup {

				idParam := AddResourceSetResourceresourceSetId
				err := utils.BackupObject(d, "ResourceSet", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&AddResourceSetResourceresourceSetId, "resourceSetId", "", "", "")
	cmd.MarkFlagRequired("resourceSetId")

	cmd.Flags().StringVarP(&AddResourceSetResourcedata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&AddResourceSetResourceBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	AddResourceSetResourceCmd := NewAddResourceSetResourceCmd()
	ResourceSetCmd.AddCommand(AddResourceSetResourceCmd)
}

var (
	DeleteResourceSetResourceresourceSetId string

	DeleteResourceSetResourceresourceId string

	DeleteResourceSetResourceBackup bool
)

func NewDeleteResourceSetResourceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deleteResource",
		Long: "Delete a Resource from a Resource Set",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ResourceSetAPI.DeleteResourceSetResource(apiClient.GetConfig().Context, DeleteResourceSetResourceresourceSetId, DeleteResourceSetResourceresourceId)

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
			if DeleteResourceSetResourceBackup {

				idParam := DeleteResourceSetResourceresourceSetId
				err := utils.BackupObject(d, "ResourceSet", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteResourceSetResourceresourceSetId, "resourceSetId", "", "", "")
	cmd.MarkFlagRequired("resourceSetId")

	cmd.Flags().StringVarP(&DeleteResourceSetResourceresourceId, "resourceId", "", "", "")
	cmd.MarkFlagRequired("resourceId")

	cmd.Flags().BoolVarP(&DeleteResourceSetResourceBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteResourceSetResourceCmd := NewDeleteResourceSetResourceCmd()
	ResourceSetCmd.AddCommand(DeleteResourceSetResourceCmd)
}
