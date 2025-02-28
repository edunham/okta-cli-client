package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var PrincipalRateLimitCmd = &cobra.Command{
	Use:  "principalRateLimit",
	Long: "Manage PrincipalRateLimitAPI",
}

func init() {
	rootCmd.AddCommand(PrincipalRateLimitCmd)
}

var (
	CreatePrincipalRateLimitEntitydata string

	CreatePrincipalRateLimitEntityBackup bool
)

func NewCreatePrincipalRateLimitEntityCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "createEntity",
		Long: "Create a Principal Rate Limit",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PrincipalRateLimitAPI.CreatePrincipalRateLimitEntity(apiClient.GetConfig().Context)

			if CreatePrincipalRateLimitEntitydata != "" {
				req = req.Data(CreatePrincipalRateLimitEntitydata)
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
			if CreatePrincipalRateLimitEntityBackup {

				err := utils.BackupObject(d, "PrincipalRateLimit", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreatePrincipalRateLimitEntitydata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreatePrincipalRateLimitEntityBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreatePrincipalRateLimitEntityCmd := NewCreatePrincipalRateLimitEntityCmd()
	PrincipalRateLimitCmd.AddCommand(CreatePrincipalRateLimitEntityCmd)
}

var ListPrincipalRateLimitEntitiesBackup bool

func NewListPrincipalRateLimitEntitiesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listEntities",
		Long: "List all Principal Rate Limits",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PrincipalRateLimitAPI.ListPrincipalRateLimitEntities(apiClient.GetConfig().Context)

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
			if ListPrincipalRateLimitEntitiesBackup {

				err := utils.BackupObject(d, "PrincipalRateLimit", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListPrincipalRateLimitEntitiesBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListPrincipalRateLimitEntitiesCmd := NewListPrincipalRateLimitEntitiesCmd()
	PrincipalRateLimitCmd.AddCommand(ListPrincipalRateLimitEntitiesCmd)
}

var (
	GetPrincipalRateLimitEntityprincipalRateLimitId string

	GetPrincipalRateLimitEntityBackup bool
)

func NewGetPrincipalRateLimitEntityCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getEntity",
		Long: "Retrieve a Principal Rate Limit",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PrincipalRateLimitAPI.GetPrincipalRateLimitEntity(apiClient.GetConfig().Context, GetPrincipalRateLimitEntityprincipalRateLimitId)

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
			if GetPrincipalRateLimitEntityBackup {

				idParam := GetPrincipalRateLimitEntityprincipalRateLimitId
				err := utils.BackupObject(d, "PrincipalRateLimit", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetPrincipalRateLimitEntityprincipalRateLimitId, "principalRateLimitId", "", "", "")
	cmd.MarkFlagRequired("principalRateLimitId")

	cmd.Flags().BoolVarP(&GetPrincipalRateLimitEntityBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetPrincipalRateLimitEntityCmd := NewGetPrincipalRateLimitEntityCmd()
	PrincipalRateLimitCmd.AddCommand(GetPrincipalRateLimitEntityCmd)
}

var (
	ReplacePrincipalRateLimitEntityprincipalRateLimitId string

	ReplacePrincipalRateLimitEntitydata string

	ReplacePrincipalRateLimitEntityBackup bool
)

func NewReplacePrincipalRateLimitEntityCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replaceEntity",
		Long: "Replace a Principal Rate Limit",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PrincipalRateLimitAPI.ReplacePrincipalRateLimitEntity(apiClient.GetConfig().Context, ReplacePrincipalRateLimitEntityprincipalRateLimitId)

			if ReplacePrincipalRateLimitEntitydata != "" {
				req = req.Data(ReplacePrincipalRateLimitEntitydata)
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
			if ReplacePrincipalRateLimitEntityBackup {

				idParam := ReplacePrincipalRateLimitEntityprincipalRateLimitId
				err := utils.BackupObject(d, "PrincipalRateLimit", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplacePrincipalRateLimitEntityprincipalRateLimitId, "principalRateLimitId", "", "", "")
	cmd.MarkFlagRequired("principalRateLimitId")

	cmd.Flags().StringVarP(&ReplacePrincipalRateLimitEntitydata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplacePrincipalRateLimitEntityBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplacePrincipalRateLimitEntityCmd := NewReplacePrincipalRateLimitEntityCmd()
	PrincipalRateLimitCmd.AddCommand(ReplacePrincipalRateLimitEntityCmd)
}
