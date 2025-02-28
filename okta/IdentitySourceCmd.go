package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var IdentitySourceCmd = &cobra.Command{
	Use:  "identitySource",
	Long: "Manage IdentitySourceAPI",
}

func init() {
	rootCmd.AddCommand(IdentitySourceCmd)
}

var (
	CreateIdentitySourceSessionidentitySourceId string

	CreateIdentitySourceSessionBackup bool
)

func NewCreateIdentitySourceSessionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "createSession",
		Long: "Create an Identity Source Session",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentitySourceAPI.CreateIdentitySourceSession(apiClient.GetConfig().Context, CreateIdentitySourceSessionidentitySourceId)

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
			if CreateIdentitySourceSessionBackup {

				idParam := CreateIdentitySourceSessionidentitySourceId
				err := utils.BackupObject(d, "IdentitySource", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateIdentitySourceSessionidentitySourceId, "identitySourceId", "", "", "")
	cmd.MarkFlagRequired("identitySourceId")

	cmd.Flags().BoolVarP(&CreateIdentitySourceSessionBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateIdentitySourceSessionCmd := NewCreateIdentitySourceSessionCmd()
	IdentitySourceCmd.AddCommand(CreateIdentitySourceSessionCmd)
}

var (
	ListIdentitySourceSessionsidentitySourceId string

	ListIdentitySourceSessionsBackup bool
)

func NewListIdentitySourceSessionsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listSessions",
		Long: "List all Identity Source Sessions",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentitySourceAPI.ListIdentitySourceSessions(apiClient.GetConfig().Context, ListIdentitySourceSessionsidentitySourceId)

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
			if ListIdentitySourceSessionsBackup {

				idParam := ListIdentitySourceSessionsidentitySourceId
				err := utils.BackupObject(d, "IdentitySource", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListIdentitySourceSessionsidentitySourceId, "identitySourceId", "", "", "")
	cmd.MarkFlagRequired("identitySourceId")

	cmd.Flags().BoolVarP(&ListIdentitySourceSessionsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListIdentitySourceSessionsCmd := NewListIdentitySourceSessionsCmd()
	IdentitySourceCmd.AddCommand(ListIdentitySourceSessionsCmd)
}

var (
	GetIdentitySourceSessionidentitySourceId string

	GetIdentitySourceSessionsessionId string

	GetIdentitySourceSessionBackup bool
)

func NewGetIdentitySourceSessionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getSession",
		Long: "Retrieve an Identity Source Session",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentitySourceAPI.GetIdentitySourceSession(apiClient.GetConfig().Context, GetIdentitySourceSessionidentitySourceId, GetIdentitySourceSessionsessionId)

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
			if GetIdentitySourceSessionBackup {

				idParam := GetIdentitySourceSessionidentitySourceId
				err := utils.BackupObject(d, "IdentitySource", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetIdentitySourceSessionidentitySourceId, "identitySourceId", "", "", "")
	cmd.MarkFlagRequired("identitySourceId")

	cmd.Flags().StringVarP(&GetIdentitySourceSessionsessionId, "sessionId", "", "", "")
	cmd.MarkFlagRequired("sessionId")

	cmd.Flags().BoolVarP(&GetIdentitySourceSessionBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetIdentitySourceSessionCmd := NewGetIdentitySourceSessionCmd()
	IdentitySourceCmd.AddCommand(GetIdentitySourceSessionCmd)
}

var (
	DeleteIdentitySourceSessionidentitySourceId string

	DeleteIdentitySourceSessionsessionId string

	DeleteIdentitySourceSessionBackup bool
)

func NewDeleteIdentitySourceSessionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deleteSession",
		Long: "Delete an Identity Source Session",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentitySourceAPI.DeleteIdentitySourceSession(apiClient.GetConfig().Context, DeleteIdentitySourceSessionidentitySourceId, DeleteIdentitySourceSessionsessionId)

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
			if DeleteIdentitySourceSessionBackup {

				idParam := DeleteIdentitySourceSessionidentitySourceId
				err := utils.BackupObject(d, "IdentitySource", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteIdentitySourceSessionidentitySourceId, "identitySourceId", "", "", "")
	cmd.MarkFlagRequired("identitySourceId")

	cmd.Flags().StringVarP(&DeleteIdentitySourceSessionsessionId, "sessionId", "", "", "")
	cmd.MarkFlagRequired("sessionId")

	cmd.Flags().BoolVarP(&DeleteIdentitySourceSessionBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteIdentitySourceSessionCmd := NewDeleteIdentitySourceSessionCmd()
	IdentitySourceCmd.AddCommand(DeleteIdentitySourceSessionCmd)
}

var (
	UploadIdentitySourceDataForDeleteidentitySourceId string

	UploadIdentitySourceDataForDeletesessionId string

	UploadIdentitySourceDataForDeletedata string

	UploadIdentitySourceDataForDeleteBackup bool
)

func NewUploadIdentitySourceDataForDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "uploadDataForDelete",
		Long: "Upload the data to be deleted in Okta",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentitySourceAPI.UploadIdentitySourceDataForDelete(apiClient.GetConfig().Context, UploadIdentitySourceDataForDeleteidentitySourceId, UploadIdentitySourceDataForDeletesessionId)

			if UploadIdentitySourceDataForDeletedata != "" {
				req = req.Data(UploadIdentitySourceDataForDeletedata)
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
			if UploadIdentitySourceDataForDeleteBackup {

				idParam := UploadIdentitySourceDataForDeleteidentitySourceId
				err := utils.BackupObject(d, "IdentitySource", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UploadIdentitySourceDataForDeleteidentitySourceId, "identitySourceId", "", "", "")
	cmd.MarkFlagRequired("identitySourceId")

	cmd.Flags().StringVarP(&UploadIdentitySourceDataForDeletesessionId, "sessionId", "", "", "")
	cmd.MarkFlagRequired("sessionId")

	cmd.Flags().StringVarP(&UploadIdentitySourceDataForDeletedata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&UploadIdentitySourceDataForDeleteBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UploadIdentitySourceDataForDeleteCmd := NewUploadIdentitySourceDataForDeleteCmd()
	IdentitySourceCmd.AddCommand(UploadIdentitySourceDataForDeleteCmd)
}

var (
	UploadIdentitySourceDataForUpsertidentitySourceId string

	UploadIdentitySourceDataForUpsertsessionId string

	UploadIdentitySourceDataForUpsertdata string

	UploadIdentitySourceDataForUpsertBackup bool
)

func NewUploadIdentitySourceDataForUpsertCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "uploadDataForUpsert",
		Long: "Upload the data to be upserted in Okta",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentitySourceAPI.UploadIdentitySourceDataForUpsert(apiClient.GetConfig().Context, UploadIdentitySourceDataForUpsertidentitySourceId, UploadIdentitySourceDataForUpsertsessionId)

			if UploadIdentitySourceDataForUpsertdata != "" {
				req = req.Data(UploadIdentitySourceDataForUpsertdata)
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
			if UploadIdentitySourceDataForUpsertBackup {

				idParam := UploadIdentitySourceDataForUpsertidentitySourceId
				err := utils.BackupObject(d, "IdentitySource", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UploadIdentitySourceDataForUpsertidentitySourceId, "identitySourceId", "", "", "")
	cmd.MarkFlagRequired("identitySourceId")

	cmd.Flags().StringVarP(&UploadIdentitySourceDataForUpsertsessionId, "sessionId", "", "", "")
	cmd.MarkFlagRequired("sessionId")

	cmd.Flags().StringVarP(&UploadIdentitySourceDataForUpsertdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&UploadIdentitySourceDataForUpsertBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UploadIdentitySourceDataForUpsertCmd := NewUploadIdentitySourceDataForUpsertCmd()
	IdentitySourceCmd.AddCommand(UploadIdentitySourceDataForUpsertCmd)
}

var (
	StartImportFromIdentitySourceidentitySourceId string

	StartImportFromIdentitySourcesessionId string

	StartImportFromIdentitySourceBackup bool
)

func NewStartImportFromIdentitySourceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "startImportFrom",
		Long: "Start the import from the Identity Source",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.IdentitySourceAPI.StartImportFromIdentitySource(apiClient.GetConfig().Context, StartImportFromIdentitySourceidentitySourceId, StartImportFromIdentitySourcesessionId)

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
			if StartImportFromIdentitySourceBackup {

				idParam := StartImportFromIdentitySourceidentitySourceId
				err := utils.BackupObject(d, "IdentitySource", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&StartImportFromIdentitySourceidentitySourceId, "identitySourceId", "", "", "")
	cmd.MarkFlagRequired("identitySourceId")

	cmd.Flags().StringVarP(&StartImportFromIdentitySourcesessionId, "sessionId", "", "", "")
	cmd.MarkFlagRequired("sessionId")

	cmd.Flags().BoolVarP(&StartImportFromIdentitySourceBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	StartImportFromIdentitySourceCmd := NewStartImportFromIdentitySourceCmd()
	IdentitySourceCmd.AddCommand(StartImportFromIdentitySourceCmd)
}
