package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var ProfileMappingCmd = &cobra.Command{
	Use:  "profileMapping",
	Long: "Manage ProfileMappingAPI",
}

func init() {
	rootCmd.AddCommand(ProfileMappingCmd)
}

var ListProfileMappingsBackup bool

func NewListProfileMappingsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "lists",
		Long: "List all Profile Mappings",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ProfileMappingAPI.ListProfileMappings(apiClient.GetConfig().Context)

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
			if ListProfileMappingsBackup {

				err := utils.BackupObject(d, "ProfileMapping", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListProfileMappingsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListProfileMappingsCmd := NewListProfileMappingsCmd()
	ProfileMappingCmd.AddCommand(ListProfileMappingsCmd)
}

var (
	UpdateProfileMappingmappingId string

	UpdateProfileMappingdata string

	UpdateProfileMappingBackup bool
)

func NewUpdateProfileMappingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "update",
		Long: "Update a Profile Mapping",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ProfileMappingAPI.UpdateProfileMapping(apiClient.GetConfig().Context, UpdateProfileMappingmappingId)

			if UpdateProfileMappingdata != "" {
				req = req.Data(UpdateProfileMappingdata)
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
			if UpdateProfileMappingBackup {

				idParam := UpdateProfileMappingmappingId
				err := utils.BackupObject(d, "ProfileMapping", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UpdateProfileMappingmappingId, "mappingId", "", "", "")
	cmd.MarkFlagRequired("mappingId")

	cmd.Flags().StringVarP(&UpdateProfileMappingdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&UpdateProfileMappingBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UpdateProfileMappingCmd := NewUpdateProfileMappingCmd()
	ProfileMappingCmd.AddCommand(UpdateProfileMappingCmd)
}

var (
	GetProfileMappingmappingId string

	GetProfileMappingBackup bool
)

func NewGetProfileMappingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "get",
		Long: "Retrieve a Profile Mapping",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ProfileMappingAPI.GetProfileMapping(apiClient.GetConfig().Context, GetProfileMappingmappingId)

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
			if GetProfileMappingBackup {

				idParam := GetProfileMappingmappingId
				err := utils.BackupObject(d, "ProfileMapping", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetProfileMappingmappingId, "mappingId", "", "", "")
	cmd.MarkFlagRequired("mappingId")

	cmd.Flags().BoolVarP(&GetProfileMappingBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetProfileMappingCmd := NewGetProfileMappingCmd()
	ProfileMappingCmd.AddCommand(GetProfileMappingCmd)
}
