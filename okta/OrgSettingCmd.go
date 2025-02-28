package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var OrgSettingCmd = &cobra.Command{
	Use:  "orgSetting",
	Long: "Manage OrgSettingAPI",
}

func init() {
	rootCmd.AddCommand(OrgSettingCmd)
}

var GetWellknownOrgMetadataBackup bool

func NewGetWellknownOrgMetadataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getWellknownOrgMetadata",
		Long: "Retrieve the Well-Known Org Metadata",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.OrgSettingAPI.GetWellknownOrgMetadata(apiClient.GetConfig().Context)

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
			if GetWellknownOrgMetadataBackup {

				err := utils.BackupObject(d, "OrgSetting", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&GetWellknownOrgMetadataBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetWellknownOrgMetadataCmd := NewGetWellknownOrgMetadataCmd()
	OrgSettingCmd.AddCommand(GetWellknownOrgMetadataCmd)
}

var (
	UpdateOrgSettingsdata string

	UpdateOrgSettingsBackup bool
)

func NewUpdateOrgSettingsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "updates",
		Long: "Update the Org Settings",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.OrgSettingAPI.UpdateOrgSettings(apiClient.GetConfig().Context)

			if UpdateOrgSettingsdata != "" {
				req = req.Data(UpdateOrgSettingsdata)
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
			if UpdateOrgSettingsBackup {

				err := utils.BackupObject(d, "OrgSetting", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UpdateOrgSettingsdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&UpdateOrgSettingsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UpdateOrgSettingsCmd := NewUpdateOrgSettingsCmd()
	OrgSettingCmd.AddCommand(UpdateOrgSettingsCmd)
}

var GetOrgSettingsBackup bool

func NewGetOrgSettingsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "gets",
		Long: "Retrieve the Org Settings",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.OrgSettingAPI.GetOrgSettings(apiClient.GetConfig().Context)

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
			if GetOrgSettingsBackup {

				err := utils.BackupObject(d, "OrgSetting", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&GetOrgSettingsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetOrgSettingsCmd := NewGetOrgSettingsCmd()
	OrgSettingCmd.AddCommand(GetOrgSettingsCmd)
}

var (
	ReplaceOrgSettingsdata string

	ReplaceOrgSettingsBackup bool
)

func NewReplaceOrgSettingsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replaces",
		Long: "Replace the Org Settings",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.OrgSettingAPI.ReplaceOrgSettings(apiClient.GetConfig().Context)

			if ReplaceOrgSettingsdata != "" {
				req = req.Data(ReplaceOrgSettingsdata)
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
			if ReplaceOrgSettingsBackup {

				err := utils.BackupObject(d, "OrgSetting", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceOrgSettingsdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceOrgSettingsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceOrgSettingsCmd := NewReplaceOrgSettingsCmd()
	OrgSettingCmd.AddCommand(ReplaceOrgSettingsCmd)
}

var GetOrgContactTypesBackup bool

func NewGetOrgContactTypesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getOrgContactTypes",
		Long: "Retrieve the Org Contact Types",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.OrgSettingAPI.GetOrgContactTypes(apiClient.GetConfig().Context)

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
			if GetOrgContactTypesBackup {

				err := utils.BackupObject(d, "OrgSetting", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&GetOrgContactTypesBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetOrgContactTypesCmd := NewGetOrgContactTypesCmd()
	OrgSettingCmd.AddCommand(GetOrgContactTypesCmd)
}

var (
	GetOrgContactUsercontactType string

	GetOrgContactUserBackup bool
)

func NewGetOrgContactUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getOrgContactUser",
		Long: "Retrieve the User of the Contact Type",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.OrgSettingAPI.GetOrgContactUser(apiClient.GetConfig().Context, GetOrgContactUsercontactType)

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
			if GetOrgContactUserBackup {

				idParam := GetOrgContactUsercontactType
				err := utils.BackupObject(d, "OrgSetting", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetOrgContactUsercontactType, "contactType", "", "", "")
	cmd.MarkFlagRequired("contactType")

	cmd.Flags().BoolVarP(&GetOrgContactUserBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetOrgContactUserCmd := NewGetOrgContactUserCmd()
	OrgSettingCmd.AddCommand(GetOrgContactUserCmd)
}

var (
	ReplaceOrgContactUsercontactType string

	ReplaceOrgContactUserdata string

	ReplaceOrgContactUserBackup bool
)

func NewReplaceOrgContactUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replaceOrgContactUser",
		Long: "Replace the User of the Contact Type",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.OrgSettingAPI.ReplaceOrgContactUser(apiClient.GetConfig().Context, ReplaceOrgContactUsercontactType)

			if ReplaceOrgContactUserdata != "" {
				req = req.Data(ReplaceOrgContactUserdata)
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
			if ReplaceOrgContactUserBackup {

				idParam := ReplaceOrgContactUsercontactType
				err := utils.BackupObject(d, "OrgSetting", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceOrgContactUsercontactType, "contactType", "", "", "")
	cmd.MarkFlagRequired("contactType")

	cmd.Flags().StringVarP(&ReplaceOrgContactUserdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceOrgContactUserBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceOrgContactUserCmd := NewReplaceOrgContactUserCmd()
	OrgSettingCmd.AddCommand(ReplaceOrgContactUserCmd)
}

var (
	BulkRemoveEmailAddressBouncesdata string

	BulkRemoveEmailAddressBouncesBackup bool
)

func NewBulkRemoveEmailAddressBouncesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "bulkRemoveEmailAddressBounces",
		Long: "Remove Emails from Email Provider Bounce List",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.OrgSettingAPI.BulkRemoveEmailAddressBounces(apiClient.GetConfig().Context)

			if BulkRemoveEmailAddressBouncesdata != "" {
				req = req.Data(BulkRemoveEmailAddressBouncesdata)
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
			if BulkRemoveEmailAddressBouncesBackup {

				err := utils.BackupObject(d, "OrgSetting", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&BulkRemoveEmailAddressBouncesdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&BulkRemoveEmailAddressBouncesBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	BulkRemoveEmailAddressBouncesCmd := NewBulkRemoveEmailAddressBouncesCmd()
	OrgSettingCmd.AddCommand(BulkRemoveEmailAddressBouncesCmd)
}

var (
	UploadOrgLogodata string

	UploadOrgLogoBackup bool
)

func NewUploadOrgLogoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "uploadOrgLogo",
		Long: "Upload the Org Logo",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.OrgSettingAPI.UploadOrgLogo(apiClient.GetConfig().Context)

			if UploadOrgLogodata != "" {
				req = req.Data(UploadOrgLogodata)
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
			if UploadOrgLogoBackup {

				err := utils.BackupObject(d, "OrgSetting", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UploadOrgLogodata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&UploadOrgLogoBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UploadOrgLogoCmd := NewUploadOrgLogoCmd()
	OrgSettingCmd.AddCommand(UploadOrgLogoCmd)
}

var UpdateThirdPartyAdminSettingBackup bool

func NewUpdateThirdPartyAdminSettingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "updateThirdPartyAdminSetting",
		Long: "Update the Org Third-Party Admin setting",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.OrgSettingAPI.UpdateThirdPartyAdminSetting(apiClient.GetConfig().Context)

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
			if UpdateThirdPartyAdminSettingBackup {

				err := utils.BackupObject(d, "OrgSetting", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&UpdateThirdPartyAdminSettingBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UpdateThirdPartyAdminSettingCmd := NewUpdateThirdPartyAdminSettingCmd()
	OrgSettingCmd.AddCommand(UpdateThirdPartyAdminSettingCmd)
}

var GetThirdPartyAdminSettingBackup bool

func NewGetThirdPartyAdminSettingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getThirdPartyAdminSetting",
		Long: "Retrieve the Org Third-Party Admin setting",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.OrgSettingAPI.GetThirdPartyAdminSetting(apiClient.GetConfig().Context)

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
			if GetThirdPartyAdminSettingBackup {

				err := utils.BackupObject(d, "OrgSetting", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&GetThirdPartyAdminSettingBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetThirdPartyAdminSettingCmd := NewGetThirdPartyAdminSettingCmd()
	OrgSettingCmd.AddCommand(GetThirdPartyAdminSettingCmd)
}

var GetOrgPreferencesBackup bool

func NewGetOrgPreferencesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getOrgPreferences",
		Long: "Retrieve the Org Preferences",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.OrgSettingAPI.GetOrgPreferences(apiClient.GetConfig().Context)

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
			if GetOrgPreferencesBackup {

				err := utils.BackupObject(d, "OrgSetting", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&GetOrgPreferencesBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetOrgPreferencesCmd := NewGetOrgPreferencesCmd()
	OrgSettingCmd.AddCommand(GetOrgPreferencesCmd)
}

var UpdateOrgHideOktaUIFooterBackup bool

func NewUpdateOrgHideOktaUIFooterCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "updateOrgHideOktaUIFooter",
		Long: "Update the Preference to Hide the Okta Dashboard Footer",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.OrgSettingAPI.UpdateOrgHideOktaUIFooter(apiClient.GetConfig().Context)

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
			if UpdateOrgHideOktaUIFooterBackup {

				err := utils.BackupObject(d, "OrgSetting", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&UpdateOrgHideOktaUIFooterBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UpdateOrgHideOktaUIFooterCmd := NewUpdateOrgHideOktaUIFooterCmd()
	OrgSettingCmd.AddCommand(UpdateOrgHideOktaUIFooterCmd)
}

var UpdateOrgShowOktaUIFooterBackup bool

func NewUpdateOrgShowOktaUIFooterCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "updateOrgShowOktaUIFooter",
		Long: "Update the Preference to Show the Okta Dashboard Footer",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.OrgSettingAPI.UpdateOrgShowOktaUIFooter(apiClient.GetConfig().Context)

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
			if UpdateOrgShowOktaUIFooterBackup {

				err := utils.BackupObject(d, "OrgSetting", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&UpdateOrgShowOktaUIFooterBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UpdateOrgShowOktaUIFooterCmd := NewUpdateOrgShowOktaUIFooterCmd()
	OrgSettingCmd.AddCommand(UpdateOrgShowOktaUIFooterCmd)
}

var GetOktaCommunicationSettingsBackup bool

func NewGetOktaCommunicationSettingsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getOktaCommunicationSettings",
		Long: "Retrieve the Okta Communication Settings",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.OrgSettingAPI.GetOktaCommunicationSettings(apiClient.GetConfig().Context)

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
			if GetOktaCommunicationSettingsBackup {

				err := utils.BackupObject(d, "OrgSetting", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&GetOktaCommunicationSettingsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetOktaCommunicationSettingsCmd := NewGetOktaCommunicationSettingsCmd()
	OrgSettingCmd.AddCommand(GetOktaCommunicationSettingsCmd)
}

var OptInUsersToOktaCommunicationEmailsBackup bool

func NewOptInUsersToOktaCommunicationEmailsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "optInUsersToOktaCommunicationEmails",
		Long: "Opt in all Users to Okta Communication emails",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.OrgSettingAPI.OptInUsersToOktaCommunicationEmails(apiClient.GetConfig().Context)

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
			if OptInUsersToOktaCommunicationEmailsBackup {

				err := utils.BackupObject(d, "OrgSetting", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&OptInUsersToOktaCommunicationEmailsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	OptInUsersToOktaCommunicationEmailsCmd := NewOptInUsersToOktaCommunicationEmailsCmd()
	OrgSettingCmd.AddCommand(OptInUsersToOktaCommunicationEmailsCmd)
}

var OptOutUsersFromOktaCommunicationEmailsBackup bool

func NewOptOutUsersFromOktaCommunicationEmailsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "optOutUsersFromOktaCommunicationEmails",
		Long: "Opt out all Users from Okta Communication emails",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.OrgSettingAPI.OptOutUsersFromOktaCommunicationEmails(apiClient.GetConfig().Context)

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
			if OptOutUsersFromOktaCommunicationEmailsBackup {

				err := utils.BackupObject(d, "OrgSetting", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&OptOutUsersFromOktaCommunicationEmailsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	OptOutUsersFromOktaCommunicationEmailsCmd := NewOptOutUsersFromOktaCommunicationEmailsCmd()
	OrgSettingCmd.AddCommand(OptOutUsersFromOktaCommunicationEmailsCmd)
}

var GetOrgOktaSupportSettingsBackup bool

func NewGetOrgOktaSupportSettingsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getOrgOktaSupportSettings",
		Long: "Retrieve the Okta Support Settings",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.OrgSettingAPI.GetOrgOktaSupportSettings(apiClient.GetConfig().Context)

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
			if GetOrgOktaSupportSettingsBackup {

				err := utils.BackupObject(d, "OrgSetting", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&GetOrgOktaSupportSettingsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetOrgOktaSupportSettingsCmd := NewGetOrgOktaSupportSettingsCmd()
	OrgSettingCmd.AddCommand(GetOrgOktaSupportSettingsCmd)
}

var ExtendOktaSupportBackup bool

func NewExtendOktaSupportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "extendOktaSupport",
		Long: "Extend Okta Support Access",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.OrgSettingAPI.ExtendOktaSupport(apiClient.GetConfig().Context)

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
			if ExtendOktaSupportBackup {

				err := utils.BackupObject(d, "OrgSetting", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ExtendOktaSupportBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ExtendOktaSupportCmd := NewExtendOktaSupportCmd()
	OrgSettingCmd.AddCommand(ExtendOktaSupportCmd)
}

var GrantOktaSupportBackup bool

func NewGrantOktaSupportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "grantOktaSupport",
		Long: "Grant Okta Support Access to your Org",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.OrgSettingAPI.GrantOktaSupport(apiClient.GetConfig().Context)

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
			if GrantOktaSupportBackup {

				err := utils.BackupObject(d, "OrgSetting", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&GrantOktaSupportBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GrantOktaSupportCmd := NewGrantOktaSupportCmd()
	OrgSettingCmd.AddCommand(GrantOktaSupportCmd)
}

var RevokeOktaSupportBackup bool

func NewRevokeOktaSupportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "revokeOktaSupport",
		Long: "Revoke Okta Support Access",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.OrgSettingAPI.RevokeOktaSupport(apiClient.GetConfig().Context)

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
			if RevokeOktaSupportBackup {

				err := utils.BackupObject(d, "OrgSetting", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&RevokeOktaSupportBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	RevokeOktaSupportCmd := NewRevokeOktaSupportCmd()
	OrgSettingCmd.AddCommand(RevokeOktaSupportCmd)
}

var GetClientPrivilegesSettingBackup bool

func NewGetClientPrivilegesSettingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getClientPrivilegesSetting",
		Long: "Retrieve the Org settings to assign the Super Admin role",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.OrgSettingAPI.GetClientPrivilegesSetting(apiClient.GetConfig().Context)

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
			if GetClientPrivilegesSettingBackup {

				err := utils.BackupObject(d, "OrgSetting", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&GetClientPrivilegesSettingBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetClientPrivilegesSettingCmd := NewGetClientPrivilegesSettingCmd()
	OrgSettingCmd.AddCommand(GetClientPrivilegesSettingCmd)
}

var (
	AssignClientPrivilegesSettingdata string

	AssignClientPrivilegesSettingBackup bool
)

func NewAssignClientPrivilegesSettingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "assignClientPrivilegesSetting",
		Long: "Assign the Super Admin role to a public client app",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.OrgSettingAPI.AssignClientPrivilegesSetting(apiClient.GetConfig().Context)

			if AssignClientPrivilegesSettingdata != "" {
				req = req.Data(AssignClientPrivilegesSettingdata)
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
			if AssignClientPrivilegesSettingBackup {

				err := utils.BackupObject(d, "OrgSetting", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&AssignClientPrivilegesSettingdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&AssignClientPrivilegesSettingBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	AssignClientPrivilegesSettingCmd := NewAssignClientPrivilegesSettingCmd()
	OrgSettingCmd.AddCommand(AssignClientPrivilegesSettingCmd)
}
