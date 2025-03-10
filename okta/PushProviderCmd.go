package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var PushProviderCmd = &cobra.Command{
	Use:  "pushProvider",
	Long: "Manage PushProviderAPI",
}

func init() {
	rootCmd.AddCommand(PushProviderCmd)
}

var (
	CreatePushProviderdata string

	CreatePushProviderBackup bool
)

func NewCreatePushProviderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create",
		Long: "Create a Push Provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PushProviderAPI.CreatePushProvider(apiClient.GetConfig().Context)

			if CreatePushProviderdata != "" {
				req = req.Data(CreatePushProviderdata)
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
			if CreatePushProviderBackup {

				err := utils.BackupObject(d, "PushProvider", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreatePushProviderdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreatePushProviderBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreatePushProviderCmd := NewCreatePushProviderCmd()
	PushProviderCmd.AddCommand(CreatePushProviderCmd)
}

var ListPushProvidersBackup bool

func NewListPushProvidersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "lists",
		Long: "List all Push Providers",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PushProviderAPI.ListPushProviders(apiClient.GetConfig().Context)

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
			if ListPushProvidersBackup {

				err := utils.BackupObject(d, "PushProvider", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListPushProvidersBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListPushProvidersCmd := NewListPushProvidersCmd()
	PushProviderCmd.AddCommand(ListPushProvidersCmd)
}

var (
	GetPushProviderpushProviderId string

	GetPushProviderBackup bool
)

func NewGetPushProviderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "get",
		Long: "Retrieve a Push Provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PushProviderAPI.GetPushProvider(apiClient.GetConfig().Context, GetPushProviderpushProviderId)

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
			if GetPushProviderBackup {

				idParam := GetPushProviderpushProviderId
				err := utils.BackupObject(d, "PushProvider", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetPushProviderpushProviderId, "pushProviderId", "", "", "")
	cmd.MarkFlagRequired("pushProviderId")

	cmd.Flags().BoolVarP(&GetPushProviderBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetPushProviderCmd := NewGetPushProviderCmd()
	PushProviderCmd.AddCommand(GetPushProviderCmd)
}

var (
	ReplacePushProviderpushProviderId string

	ReplacePushProviderdata string

	ReplacePushProviderBackup bool
)

func NewReplacePushProviderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replace",
		Long: "Replace a Push Provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PushProviderAPI.ReplacePushProvider(apiClient.GetConfig().Context, ReplacePushProviderpushProviderId)

			if ReplacePushProviderdata != "" {
				req = req.Data(ReplacePushProviderdata)
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
			if ReplacePushProviderBackup {

				idParam := ReplacePushProviderpushProviderId
				err := utils.BackupObject(d, "PushProvider", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplacePushProviderpushProviderId, "pushProviderId", "", "", "")
	cmd.MarkFlagRequired("pushProviderId")

	cmd.Flags().StringVarP(&ReplacePushProviderdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplacePushProviderBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplacePushProviderCmd := NewReplacePushProviderCmd()
	PushProviderCmd.AddCommand(ReplacePushProviderCmd)
}

var (
	DeletePushProviderpushProviderId string

	DeletePushProviderBackup bool
)

func NewDeletePushProviderCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete",
		Long: "Delete a Push Provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PushProviderAPI.DeletePushProvider(apiClient.GetConfig().Context, DeletePushProviderpushProviderId)

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
			if DeletePushProviderBackup {

				idParam := DeletePushProviderpushProviderId
				err := utils.BackupObject(d, "PushProvider", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeletePushProviderpushProviderId, "pushProviderId", "", "", "")
	cmd.MarkFlagRequired("pushProviderId")

	cmd.Flags().BoolVarP(&DeletePushProviderBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeletePushProviderCmd := NewDeletePushProviderCmd()
	PushProviderCmd.AddCommand(DeletePushProviderCmd)
}
