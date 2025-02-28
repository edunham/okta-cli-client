package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var TemplateCmd = &cobra.Command{
	Use:  "template",
	Long: "Manage TemplateAPI",
}

func init() {
	rootCmd.AddCommand(TemplateCmd)
}

var (
	CreateSmsTemplatedata string

	CreateSmsTemplateBackup bool
)

func NewCreateSmsTemplateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "createSms",
		Long: "Create an SMS Template",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.TemplateAPI.CreateSmsTemplate(apiClient.GetConfig().Context)

			if CreateSmsTemplatedata != "" {
				req = req.Data(CreateSmsTemplatedata)
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
			if CreateSmsTemplateBackup {

				err := utils.BackupObject(d, "Template", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateSmsTemplatedata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateSmsTemplateBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateSmsTemplateCmd := NewCreateSmsTemplateCmd()
	TemplateCmd.AddCommand(CreateSmsTemplateCmd)
}

var ListSmsTemplatesBackup bool

func NewListSmsTemplatesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listSmss",
		Long: "List all SMS Templates",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.TemplateAPI.ListSmsTemplates(apiClient.GetConfig().Context)

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
			if ListSmsTemplatesBackup {

				err := utils.BackupObject(d, "Template", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListSmsTemplatesBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListSmsTemplatesCmd := NewListSmsTemplatesCmd()
	TemplateCmd.AddCommand(ListSmsTemplatesCmd)
}

var (
	UpdateSmsTemplatetemplateId string

	UpdateSmsTemplatedata string

	UpdateSmsTemplateBackup bool
)

func NewUpdateSmsTemplateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "updateSms",
		Long: "Update an SMS Template",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.TemplateAPI.UpdateSmsTemplate(apiClient.GetConfig().Context, UpdateSmsTemplatetemplateId)

			if UpdateSmsTemplatedata != "" {
				req = req.Data(UpdateSmsTemplatedata)
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
			if UpdateSmsTemplateBackup {

				idParam := UpdateSmsTemplatetemplateId
				err := utils.BackupObject(d, "Template", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UpdateSmsTemplatetemplateId, "templateId", "", "", "")
	cmd.MarkFlagRequired("templateId")

	cmd.Flags().StringVarP(&UpdateSmsTemplatedata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&UpdateSmsTemplateBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UpdateSmsTemplateCmd := NewUpdateSmsTemplateCmd()
	TemplateCmd.AddCommand(UpdateSmsTemplateCmd)
}

var (
	GetSmsTemplatetemplateId string

	GetSmsTemplateBackup bool
)

func NewGetSmsTemplateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getSms",
		Long: "Retrieve an SMS Template",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.TemplateAPI.GetSmsTemplate(apiClient.GetConfig().Context, GetSmsTemplatetemplateId)

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
			if GetSmsTemplateBackup {

				idParam := GetSmsTemplatetemplateId
				err := utils.BackupObject(d, "Template", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetSmsTemplatetemplateId, "templateId", "", "", "")
	cmd.MarkFlagRequired("templateId")

	cmd.Flags().BoolVarP(&GetSmsTemplateBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetSmsTemplateCmd := NewGetSmsTemplateCmd()
	TemplateCmd.AddCommand(GetSmsTemplateCmd)
}

var (
	ReplaceSmsTemplatetemplateId string

	ReplaceSmsTemplatedata string

	ReplaceSmsTemplateBackup bool
)

func NewReplaceSmsTemplateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replaceSms",
		Long: "Replace an SMS Template",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.TemplateAPI.ReplaceSmsTemplate(apiClient.GetConfig().Context, ReplaceSmsTemplatetemplateId)

			if ReplaceSmsTemplatedata != "" {
				req = req.Data(ReplaceSmsTemplatedata)
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
			if ReplaceSmsTemplateBackup {

				idParam := ReplaceSmsTemplatetemplateId
				err := utils.BackupObject(d, "Template", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceSmsTemplatetemplateId, "templateId", "", "", "")
	cmd.MarkFlagRequired("templateId")

	cmd.Flags().StringVarP(&ReplaceSmsTemplatedata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceSmsTemplateBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceSmsTemplateCmd := NewReplaceSmsTemplateCmd()
	TemplateCmd.AddCommand(ReplaceSmsTemplateCmd)
}

var (
	DeleteSmsTemplatetemplateId string

	DeleteSmsTemplateBackup bool
)

func NewDeleteSmsTemplateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deleteSms",
		Long: "Delete an SMS Template",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.TemplateAPI.DeleteSmsTemplate(apiClient.GetConfig().Context, DeleteSmsTemplatetemplateId)

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
			if DeleteSmsTemplateBackup {

				idParam := DeleteSmsTemplatetemplateId
				err := utils.BackupObject(d, "Template", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteSmsTemplatetemplateId, "templateId", "", "", "")
	cmd.MarkFlagRequired("templateId")

	cmd.Flags().BoolVarP(&DeleteSmsTemplateBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteSmsTemplateCmd := NewDeleteSmsTemplateCmd()
	TemplateCmd.AddCommand(DeleteSmsTemplateCmd)
}
