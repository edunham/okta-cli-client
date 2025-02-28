package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var EmailDomainCmd = &cobra.Command{
	Use:  "emailDomain",
	Long: "Manage EmailDomainAPI",
}

func init() {
	rootCmd.AddCommand(EmailDomainCmd)
}

var (
	CreateEmailDomaindata string

	CreateEmailDomainBackup bool
)

func NewCreateEmailDomainCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create",
		Long: "Create an Email Domain",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.EmailDomainAPI.CreateEmailDomain(apiClient.GetConfig().Context)

			if CreateEmailDomaindata != "" {
				req = req.Data(CreateEmailDomaindata)
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
			if CreateEmailDomainBackup {

				err := utils.BackupObject(d, "EmailDomain", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateEmailDomaindata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateEmailDomainBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateEmailDomainCmd := NewCreateEmailDomainCmd()
	EmailDomainCmd.AddCommand(CreateEmailDomainCmd)
}

var ListEmailDomainsBackup bool

func NewListEmailDomainsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "lists",
		Long: "List all Email Domains",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.EmailDomainAPI.ListEmailDomains(apiClient.GetConfig().Context)

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
			if ListEmailDomainsBackup {

				err := utils.BackupObject(d, "EmailDomain", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListEmailDomainsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListEmailDomainsCmd := NewListEmailDomainsCmd()
	EmailDomainCmd.AddCommand(ListEmailDomainsCmd)
}

var (
	GetEmailDomainemailDomainId string

	GetEmailDomainBackup bool
)

func NewGetEmailDomainCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "get",
		Long: "Retrieve an Email Domain",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.EmailDomainAPI.GetEmailDomain(apiClient.GetConfig().Context, GetEmailDomainemailDomainId)

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
			if GetEmailDomainBackup {

				idParam := GetEmailDomainemailDomainId
				err := utils.BackupObject(d, "EmailDomain", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetEmailDomainemailDomainId, "emailDomainId", "", "", "")
	cmd.MarkFlagRequired("emailDomainId")

	cmd.Flags().BoolVarP(&GetEmailDomainBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetEmailDomainCmd := NewGetEmailDomainCmd()
	EmailDomainCmd.AddCommand(GetEmailDomainCmd)
}

var (
	ReplaceEmailDomainemailDomainId string

	ReplaceEmailDomaindata string

	ReplaceEmailDomainBackup bool
)

func NewReplaceEmailDomainCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replace",
		Long: "Replace an Email Domain",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.EmailDomainAPI.ReplaceEmailDomain(apiClient.GetConfig().Context, ReplaceEmailDomainemailDomainId)

			if ReplaceEmailDomaindata != "" {
				req = req.Data(ReplaceEmailDomaindata)
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
			if ReplaceEmailDomainBackup {

				idParam := ReplaceEmailDomainemailDomainId
				err := utils.BackupObject(d, "EmailDomain", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceEmailDomainemailDomainId, "emailDomainId", "", "", "")
	cmd.MarkFlagRequired("emailDomainId")

	cmd.Flags().StringVarP(&ReplaceEmailDomaindata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceEmailDomainBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceEmailDomainCmd := NewReplaceEmailDomainCmd()
	EmailDomainCmd.AddCommand(ReplaceEmailDomainCmd)
}

var (
	DeleteEmailDomainemailDomainId string

	DeleteEmailDomainBackup bool
)

func NewDeleteEmailDomainCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete",
		Long: "Delete an Email Domain",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.EmailDomainAPI.DeleteEmailDomain(apiClient.GetConfig().Context, DeleteEmailDomainemailDomainId)

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
			if DeleteEmailDomainBackup {

				idParam := DeleteEmailDomainemailDomainId
				err := utils.BackupObject(d, "EmailDomain", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteEmailDomainemailDomainId, "emailDomainId", "", "", "")
	cmd.MarkFlagRequired("emailDomainId")

	cmd.Flags().BoolVarP(&DeleteEmailDomainBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteEmailDomainCmd := NewDeleteEmailDomainCmd()
	EmailDomainCmd.AddCommand(DeleteEmailDomainCmd)
}

var (
	VerifyEmailDomainemailDomainId string

	VerifyEmailDomainBackup bool
)

func NewVerifyEmailDomainCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "verify",
		Long: "Verify an Email Domain",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.EmailDomainAPI.VerifyEmailDomain(apiClient.GetConfig().Context, VerifyEmailDomainemailDomainId)

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
			if VerifyEmailDomainBackup {

				idParam := VerifyEmailDomainemailDomainId
				err := utils.BackupObject(d, "EmailDomain", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&VerifyEmailDomainemailDomainId, "emailDomainId", "", "", "")
	cmd.MarkFlagRequired("emailDomainId")

	cmd.Flags().BoolVarP(&VerifyEmailDomainBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	VerifyEmailDomainCmd := NewVerifyEmailDomainCmd()
	EmailDomainCmd.AddCommand(VerifyEmailDomainCmd)
}
