package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var EventHookCmd = &cobra.Command{
	Use:  "eventHook",
	Long: "Manage EventHookAPI",
}

func init() {
	rootCmd.AddCommand(EventHookCmd)
}

var (
	CreateEventHookdata string

	CreateEventHookBackup bool
)

func NewCreateEventHookCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create",
		Long: "Create an Event Hook",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.EventHookAPI.CreateEventHook(apiClient.GetConfig().Context)

			if CreateEventHookdata != "" {
				req = req.Data(CreateEventHookdata)
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
			if CreateEventHookBackup {

				err := utils.BackupObject(d, "EventHook", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateEventHookdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateEventHookBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateEventHookCmd := NewCreateEventHookCmd()
	EventHookCmd.AddCommand(CreateEventHookCmd)
}

var ListEventHooksBackup bool

func NewListEventHooksCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "lists",
		Long: "List all Event Hooks",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.EventHookAPI.ListEventHooks(apiClient.GetConfig().Context)

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
			if ListEventHooksBackup {

				err := utils.BackupObject(d, "EventHook", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListEventHooksBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListEventHooksCmd := NewListEventHooksCmd()
	EventHookCmd.AddCommand(ListEventHooksCmd)
}

var (
	GetEventHookeventHookId string

	GetEventHookBackup bool
)

func NewGetEventHookCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "get",
		Long: "Retrieve an Event Hook",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.EventHookAPI.GetEventHook(apiClient.GetConfig().Context, GetEventHookeventHookId)

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
			if GetEventHookBackup {

				idParam := GetEventHookeventHookId
				err := utils.BackupObject(d, "EventHook", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetEventHookeventHookId, "eventHookId", "", "", "")
	cmd.MarkFlagRequired("eventHookId")

	cmd.Flags().BoolVarP(&GetEventHookBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetEventHookCmd := NewGetEventHookCmd()
	EventHookCmd.AddCommand(GetEventHookCmd)
}

var (
	ReplaceEventHookeventHookId string

	ReplaceEventHookdata string

	ReplaceEventHookBackup bool
)

func NewReplaceEventHookCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replace",
		Long: "Replace an Event Hook",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.EventHookAPI.ReplaceEventHook(apiClient.GetConfig().Context, ReplaceEventHookeventHookId)

			if ReplaceEventHookdata != "" {
				req = req.Data(ReplaceEventHookdata)
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
			if ReplaceEventHookBackup {

				idParam := ReplaceEventHookeventHookId
				err := utils.BackupObject(d, "EventHook", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceEventHookeventHookId, "eventHookId", "", "", "")
	cmd.MarkFlagRequired("eventHookId")

	cmd.Flags().StringVarP(&ReplaceEventHookdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceEventHookBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceEventHookCmd := NewReplaceEventHookCmd()
	EventHookCmd.AddCommand(ReplaceEventHookCmd)
}

var (
	DeleteEventHookeventHookId string

	DeleteEventHookBackup bool
)

func NewDeleteEventHookCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete",
		Long: "Delete an Event Hook",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.EventHookAPI.DeleteEventHook(apiClient.GetConfig().Context, DeleteEventHookeventHookId)

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
			if DeleteEventHookBackup {

				idParam := DeleteEventHookeventHookId
				err := utils.BackupObject(d, "EventHook", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteEventHookeventHookId, "eventHookId", "", "", "")
	cmd.MarkFlagRequired("eventHookId")

	cmd.Flags().BoolVarP(&DeleteEventHookBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteEventHookCmd := NewDeleteEventHookCmd()
	EventHookCmd.AddCommand(DeleteEventHookCmd)
}

var (
	ActivateEventHookeventHookId string

	ActivateEventHookBackup bool
)

func NewActivateEventHookCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "activate",
		Long: "Activate an Event Hook",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.EventHookAPI.ActivateEventHook(apiClient.GetConfig().Context, ActivateEventHookeventHookId)

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
			if ActivateEventHookBackup {

				idParam := ActivateEventHookeventHookId
				err := utils.BackupObject(d, "EventHook", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ActivateEventHookeventHookId, "eventHookId", "", "", "")
	cmd.MarkFlagRequired("eventHookId")

	cmd.Flags().BoolVarP(&ActivateEventHookBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ActivateEventHookCmd := NewActivateEventHookCmd()
	EventHookCmd.AddCommand(ActivateEventHookCmd)
}

var (
	DeactivateEventHookeventHookId string

	DeactivateEventHookBackup bool
)

func NewDeactivateEventHookCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deactivate",
		Long: "Deactivate an Event Hook",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.EventHookAPI.DeactivateEventHook(apiClient.GetConfig().Context, DeactivateEventHookeventHookId)

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
			if DeactivateEventHookBackup {

				idParam := DeactivateEventHookeventHookId
				err := utils.BackupObject(d, "EventHook", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeactivateEventHookeventHookId, "eventHookId", "", "", "")
	cmd.MarkFlagRequired("eventHookId")

	cmd.Flags().BoolVarP(&DeactivateEventHookBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeactivateEventHookCmd := NewDeactivateEventHookCmd()
	EventHookCmd.AddCommand(DeactivateEventHookCmd)
}

var (
	VerifyEventHookeventHookId string

	VerifyEventHookBackup bool
)

func NewVerifyEventHookCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "verify",
		Long: "Verify an Event Hook",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.EventHookAPI.VerifyEventHook(apiClient.GetConfig().Context, VerifyEventHookeventHookId)

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
			if VerifyEventHookBackup {

				idParam := VerifyEventHookeventHookId
				err := utils.BackupObject(d, "EventHook", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&VerifyEventHookeventHookId, "eventHookId", "", "", "")
	cmd.MarkFlagRequired("eventHookId")

	cmd.Flags().BoolVarP(&VerifyEventHookBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	VerifyEventHookCmd := NewVerifyEventHookCmd()
	EventHookCmd.AddCommand(VerifyEventHookCmd)
}
