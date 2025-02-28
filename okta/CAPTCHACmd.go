package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var CAPTCHACmd = &cobra.Command{
	Use:  "cAPTCHA",
	Long: "Manage CAPTCHAAPI",
}

func init() {
	rootCmd.AddCommand(CAPTCHACmd)
}

var (
	CreateCaptchaInstancedata string

	CreateCaptchaInstanceBackup bool
)

func NewCreateCaptchaInstanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "createCaptchaInstance",
		Long: "Create a CAPTCHA instance",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.CAPTCHAAPI.CreateCaptchaInstance(apiClient.GetConfig().Context)

			if CreateCaptchaInstancedata != "" {
				req = req.Data(CreateCaptchaInstancedata)
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
			if CreateCaptchaInstanceBackup {

				err := utils.BackupObject(d, "CAPTCHA", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateCaptchaInstancedata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateCaptchaInstanceBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateCaptchaInstanceCmd := NewCreateCaptchaInstanceCmd()
	CAPTCHACmd.AddCommand(CreateCaptchaInstanceCmd)
}

var ListCaptchaInstancesBackup bool

func NewListCaptchaInstancesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listCaptchaInstances",
		Long: "List all CAPTCHA Instances",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.CAPTCHAAPI.ListCaptchaInstances(apiClient.GetConfig().Context)

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
			if ListCaptchaInstancesBackup {

				err := utils.BackupObject(d, "CAPTCHA", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListCaptchaInstancesBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListCaptchaInstancesCmd := NewListCaptchaInstancesCmd()
	CAPTCHACmd.AddCommand(ListCaptchaInstancesCmd)
}

var (
	UpdateCaptchaInstancecaptchaId string

	UpdateCaptchaInstancedata string

	UpdateCaptchaInstanceBackup bool
)

func NewUpdateCaptchaInstanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "updateCaptchaInstance",
		Long: "Update a CAPTCHA Instance",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.CAPTCHAAPI.UpdateCaptchaInstance(apiClient.GetConfig().Context, UpdateCaptchaInstancecaptchaId)

			if UpdateCaptchaInstancedata != "" {
				req = req.Data(UpdateCaptchaInstancedata)
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
			if UpdateCaptchaInstanceBackup {

				idParam := UpdateCaptchaInstancecaptchaId
				err := utils.BackupObject(d, "CAPTCHA", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UpdateCaptchaInstancecaptchaId, "captchaId", "", "", "")
	cmd.MarkFlagRequired("captchaId")

	cmd.Flags().StringVarP(&UpdateCaptchaInstancedata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&UpdateCaptchaInstanceBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UpdateCaptchaInstanceCmd := NewUpdateCaptchaInstanceCmd()
	CAPTCHACmd.AddCommand(UpdateCaptchaInstanceCmd)
}

var (
	GetCaptchaInstancecaptchaId string

	GetCaptchaInstanceBackup bool
)

func NewGetCaptchaInstanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getCaptchaInstance",
		Long: "Retrieve a CAPTCHA Instance",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.CAPTCHAAPI.GetCaptchaInstance(apiClient.GetConfig().Context, GetCaptchaInstancecaptchaId)

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
			if GetCaptchaInstanceBackup {

				idParam := GetCaptchaInstancecaptchaId
				err := utils.BackupObject(d, "CAPTCHA", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetCaptchaInstancecaptchaId, "captchaId", "", "", "")
	cmd.MarkFlagRequired("captchaId")

	cmd.Flags().BoolVarP(&GetCaptchaInstanceBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetCaptchaInstanceCmd := NewGetCaptchaInstanceCmd()
	CAPTCHACmd.AddCommand(GetCaptchaInstanceCmd)
}

var (
	ReplaceCaptchaInstancecaptchaId string

	ReplaceCaptchaInstancedata string

	ReplaceCaptchaInstanceBackup bool
)

func NewReplaceCaptchaInstanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replaceCaptchaInstance",
		Long: "Replace a CAPTCHA Instance",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.CAPTCHAAPI.ReplaceCaptchaInstance(apiClient.GetConfig().Context, ReplaceCaptchaInstancecaptchaId)

			if ReplaceCaptchaInstancedata != "" {
				req = req.Data(ReplaceCaptchaInstancedata)
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
			if ReplaceCaptchaInstanceBackup {

				idParam := ReplaceCaptchaInstancecaptchaId
				err := utils.BackupObject(d, "CAPTCHA", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceCaptchaInstancecaptchaId, "captchaId", "", "", "")
	cmd.MarkFlagRequired("captchaId")

	cmd.Flags().StringVarP(&ReplaceCaptchaInstancedata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceCaptchaInstanceBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceCaptchaInstanceCmd := NewReplaceCaptchaInstanceCmd()
	CAPTCHACmd.AddCommand(ReplaceCaptchaInstanceCmd)
}

var (
	DeleteCaptchaInstancecaptchaId string

	DeleteCaptchaInstanceBackup bool
)

func NewDeleteCaptchaInstanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deleteCaptchaInstance",
		Long: "Delete a CAPTCHA Instance",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.CAPTCHAAPI.DeleteCaptchaInstance(apiClient.GetConfig().Context, DeleteCaptchaInstancecaptchaId)

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
			if DeleteCaptchaInstanceBackup {

				idParam := DeleteCaptchaInstancecaptchaId
				err := utils.BackupObject(d, "CAPTCHA", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteCaptchaInstancecaptchaId, "captchaId", "", "", "")
	cmd.MarkFlagRequired("captchaId")

	cmd.Flags().BoolVarP(&DeleteCaptchaInstanceBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteCaptchaInstanceCmd := NewDeleteCaptchaInstanceCmd()
	CAPTCHACmd.AddCommand(DeleteCaptchaInstanceCmd)
}

var GetOrgCaptchaSettingsBackup bool

func NewGetOrgCaptchaSettingsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getOrgCaptchaSettings",
		Long: "Retrieve the Org-wide CAPTCHA Settings",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.CAPTCHAAPI.GetOrgCaptchaSettings(apiClient.GetConfig().Context)

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
			if GetOrgCaptchaSettingsBackup {

				err := utils.BackupObject(d, "CAPTCHA", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&GetOrgCaptchaSettingsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetOrgCaptchaSettingsCmd := NewGetOrgCaptchaSettingsCmd()
	CAPTCHACmd.AddCommand(GetOrgCaptchaSettingsCmd)
}

var (
	ReplacesOrgCaptchaSettingsdata string

	ReplacesOrgCaptchaSettingsBackup bool
)

func NewReplacesOrgCaptchaSettingsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replacesOrgCaptchaSettings",
		Long: "Replace the Org-wide CAPTCHA Settings",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.CAPTCHAAPI.ReplacesOrgCaptchaSettings(apiClient.GetConfig().Context)

			if ReplacesOrgCaptchaSettingsdata != "" {
				req = req.Data(ReplacesOrgCaptchaSettingsdata)
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
			if ReplacesOrgCaptchaSettingsBackup {

				err := utils.BackupObject(d, "CAPTCHA", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplacesOrgCaptchaSettingsdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplacesOrgCaptchaSettingsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplacesOrgCaptchaSettingsCmd := NewReplacesOrgCaptchaSettingsCmd()
	CAPTCHACmd.AddCommand(ReplacesOrgCaptchaSettingsCmd)
}

var DeleteOrgCaptchaSettingsBackup bool

func NewDeleteOrgCaptchaSettingsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deleteOrgCaptchaSettings",
		Long: "Delete the Org-wide CAPTCHA Settings",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.CAPTCHAAPI.DeleteOrgCaptchaSettings(apiClient.GetConfig().Context)

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
			if DeleteOrgCaptchaSettingsBackup {

				err := utils.BackupObject(d, "CAPTCHA", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&DeleteOrgCaptchaSettingsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteOrgCaptchaSettingsCmd := NewDeleteOrgCaptchaSettingsCmd()
	CAPTCHACmd.AddCommand(DeleteOrgCaptchaSettingsCmd)
}
