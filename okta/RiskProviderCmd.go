package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var RiskProviderCmd = &cobra.Command{
	Use:  "riskProvider",
	Long: "Manage RiskProviderAPI",
}

func init() {
	rootCmd.AddCommand(RiskProviderCmd)
}

var (
	CreateRiskProviderdata string

	CreateRiskProviderBackup bool
)

func NewCreateRiskProviderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create",
		Long: "Create a Risk Provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RiskProviderAPI.CreateRiskProvider(apiClient.GetConfig().Context)

			if CreateRiskProviderdata != "" {
				req = req.Data(CreateRiskProviderdata)
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
			if CreateRiskProviderBackup {

				err := utils.BackupObject(d, "RiskProvider", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateRiskProviderdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateRiskProviderBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateRiskProviderCmd := NewCreateRiskProviderCmd()
	RiskProviderCmd.AddCommand(CreateRiskProviderCmd)
}

var ListRiskProvidersBackup bool

func NewListRiskProvidersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "lists",
		Long: "List all Risk Providers",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RiskProviderAPI.ListRiskProviders(apiClient.GetConfig().Context)

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
			if ListRiskProvidersBackup {

				err := utils.BackupObject(d, "RiskProvider", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListRiskProvidersBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListRiskProvidersCmd := NewListRiskProvidersCmd()
	RiskProviderCmd.AddCommand(ListRiskProvidersCmd)
}

var (
	GetRiskProviderriskProviderId string

	GetRiskProviderBackup bool
)

func NewGetRiskProviderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "get",
		Long: "Retrieve a Risk Provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RiskProviderAPI.GetRiskProvider(apiClient.GetConfig().Context, GetRiskProviderriskProviderId)

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
			if GetRiskProviderBackup {

				idParam := GetRiskProviderriskProviderId
				err := utils.BackupObject(d, "RiskProvider", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetRiskProviderriskProviderId, "riskProviderId", "", "", "")
	cmd.MarkFlagRequired("riskProviderId")

	cmd.Flags().BoolVarP(&GetRiskProviderBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetRiskProviderCmd := NewGetRiskProviderCmd()
	RiskProviderCmd.AddCommand(GetRiskProviderCmd)
}

var (
	ReplaceRiskProviderriskProviderId string

	ReplaceRiskProviderdata string

	ReplaceRiskProviderBackup bool
)

func NewReplaceRiskProviderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replace",
		Long: "Replace a Risk Provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RiskProviderAPI.ReplaceRiskProvider(apiClient.GetConfig().Context, ReplaceRiskProviderriskProviderId)

			if ReplaceRiskProviderdata != "" {
				req = req.Data(ReplaceRiskProviderdata)
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
			if ReplaceRiskProviderBackup {

				idParam := ReplaceRiskProviderriskProviderId
				err := utils.BackupObject(d, "RiskProvider", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceRiskProviderriskProviderId, "riskProviderId", "", "", "")
	cmd.MarkFlagRequired("riskProviderId")

	cmd.Flags().StringVarP(&ReplaceRiskProviderdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceRiskProviderBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceRiskProviderCmd := NewReplaceRiskProviderCmd()
	RiskProviderCmd.AddCommand(ReplaceRiskProviderCmd)
}

var (
	DeleteRiskProviderriskProviderId string

	DeleteRiskProviderBackup bool
)

func NewDeleteRiskProviderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete",
		Long: "Delete a Risk Provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.RiskProviderAPI.DeleteRiskProvider(apiClient.GetConfig().Context, DeleteRiskProviderriskProviderId)

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
			if DeleteRiskProviderBackup {

				idParam := DeleteRiskProviderriskProviderId
				err := utils.BackupObject(d, "RiskProvider", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteRiskProviderriskProviderId, "riskProviderId", "", "", "")
	cmd.MarkFlagRequired("riskProviderId")

	cmd.Flags().BoolVarP(&DeleteRiskProviderBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteRiskProviderCmd := NewDeleteRiskProviderCmd()
	RiskProviderCmd.AddCommand(DeleteRiskProviderCmd)
}
