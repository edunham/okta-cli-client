package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var AgentPoolsCmd = &cobra.Command{
	Use:  "agentPools",
	Long: "Manage AgentPoolsAPI",
}

func init() {
	rootCmd.AddCommand(AgentPoolsCmd)
}

var ListAgentPoolsBackup bool

func NewListAgentPoolsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "list",
		Long: "List all Agent Pools",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AgentPoolsAPI.ListAgentPools(apiClient.GetConfig().Context)

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
			if ListAgentPoolsBackup {

				err := utils.BackupObject(d, "AgentPools", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListAgentPoolsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListAgentPoolsCmd := NewListAgentPoolsCmd()
	AgentPoolsCmd.AddCommand(ListAgentPoolsCmd)
}

var (
	CreateAgentPoolsUpdatepoolId string

	CreateAgentPoolsUpdatedata string

	CreateAgentPoolsUpdateBackup bool
)

func NewCreateAgentPoolsUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "createUpdate",
		Long: "Create an Agent Pool update",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AgentPoolsAPI.CreateAgentPoolsUpdate(apiClient.GetConfig().Context, CreateAgentPoolsUpdatepoolId)

			if CreateAgentPoolsUpdatedata != "" {
				req = req.Data(CreateAgentPoolsUpdatedata)
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
			if CreateAgentPoolsUpdateBackup {

				idParam := CreateAgentPoolsUpdatepoolId
				err := utils.BackupObject(d, "AgentPools", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateAgentPoolsUpdatepoolId, "poolId", "", "", "")
	cmd.MarkFlagRequired("poolId")

	cmd.Flags().StringVarP(&CreateAgentPoolsUpdatedata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateAgentPoolsUpdateBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateAgentPoolsUpdateCmd := NewCreateAgentPoolsUpdateCmd()
	AgentPoolsCmd.AddCommand(CreateAgentPoolsUpdateCmd)
}

var (
	ListAgentPoolsUpdatespoolId string

	ListAgentPoolsUpdatesBackup bool
)

func NewListAgentPoolsUpdatesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listUpdates",
		Long: "List all Agent Pool updates",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AgentPoolsAPI.ListAgentPoolsUpdates(apiClient.GetConfig().Context, ListAgentPoolsUpdatespoolId)

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
			if ListAgentPoolsUpdatesBackup {

				idParam := ListAgentPoolsUpdatespoolId
				err := utils.BackupObject(d, "AgentPools", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListAgentPoolsUpdatespoolId, "poolId", "", "", "")
	cmd.MarkFlagRequired("poolId")

	cmd.Flags().BoolVarP(&ListAgentPoolsUpdatesBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListAgentPoolsUpdatesCmd := NewListAgentPoolsUpdatesCmd()
	AgentPoolsCmd.AddCommand(ListAgentPoolsUpdatesCmd)
}

var (
	UpdateAgentPoolsUpdateSettingspoolId string

	UpdateAgentPoolsUpdateSettingsdata string

	UpdateAgentPoolsUpdateSettingsBackup bool
)

func NewUpdateAgentPoolsUpdateSettingsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "updateUpdateSettings",
		Long: "Update an Agent Pool update settings",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AgentPoolsAPI.UpdateAgentPoolsUpdateSettings(apiClient.GetConfig().Context, UpdateAgentPoolsUpdateSettingspoolId)

			if UpdateAgentPoolsUpdateSettingsdata != "" {
				req = req.Data(UpdateAgentPoolsUpdateSettingsdata)
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
			if UpdateAgentPoolsUpdateSettingsBackup {

				idParam := UpdateAgentPoolsUpdateSettingspoolId
				err := utils.BackupObject(d, "AgentPools", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UpdateAgentPoolsUpdateSettingspoolId, "poolId", "", "", "")
	cmd.MarkFlagRequired("poolId")

	cmd.Flags().StringVarP(&UpdateAgentPoolsUpdateSettingsdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&UpdateAgentPoolsUpdateSettingsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UpdateAgentPoolsUpdateSettingsCmd := NewUpdateAgentPoolsUpdateSettingsCmd()
	AgentPoolsCmd.AddCommand(UpdateAgentPoolsUpdateSettingsCmd)
}

var (
	GetAgentPoolsUpdateSettingspoolId string

	GetAgentPoolsUpdateSettingsBackup bool
)

func NewGetAgentPoolsUpdateSettingsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getUpdateSettings",
		Long: "Retrieve an Agent Pool update's settings",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AgentPoolsAPI.GetAgentPoolsUpdateSettings(apiClient.GetConfig().Context, GetAgentPoolsUpdateSettingspoolId)

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
			if GetAgentPoolsUpdateSettingsBackup {

				idParam := GetAgentPoolsUpdateSettingspoolId
				err := utils.BackupObject(d, "AgentPools", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetAgentPoolsUpdateSettingspoolId, "poolId", "", "", "")
	cmd.MarkFlagRequired("poolId")

	cmd.Flags().BoolVarP(&GetAgentPoolsUpdateSettingsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetAgentPoolsUpdateSettingsCmd := NewGetAgentPoolsUpdateSettingsCmd()
	AgentPoolsCmd.AddCommand(GetAgentPoolsUpdateSettingsCmd)
}

var (
	UpdateAgentPoolsUpdatepoolId string

	UpdateAgentPoolsUpdateupdateId string

	UpdateAgentPoolsUpdatedata string

	UpdateAgentPoolsUpdateBackup bool
)

func NewUpdateAgentPoolsUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "updateUpdate",
		Long: "Update an Agent Pool update by id",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AgentPoolsAPI.UpdateAgentPoolsUpdate(apiClient.GetConfig().Context, UpdateAgentPoolsUpdatepoolId, UpdateAgentPoolsUpdateupdateId)

			if UpdateAgentPoolsUpdatedata != "" {
				req = req.Data(UpdateAgentPoolsUpdatedata)
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
			if UpdateAgentPoolsUpdateBackup {

				idParam := UpdateAgentPoolsUpdatepoolId
				err := utils.BackupObject(d, "AgentPools", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UpdateAgentPoolsUpdatepoolId, "poolId", "", "", "")
	cmd.MarkFlagRequired("poolId")

	cmd.Flags().StringVarP(&UpdateAgentPoolsUpdateupdateId, "updateId", "", "", "")
	cmd.MarkFlagRequired("updateId")

	cmd.Flags().StringVarP(&UpdateAgentPoolsUpdatedata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&UpdateAgentPoolsUpdateBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UpdateAgentPoolsUpdateCmd := NewUpdateAgentPoolsUpdateCmd()
	AgentPoolsCmd.AddCommand(UpdateAgentPoolsUpdateCmd)
}

var (
	GetAgentPoolsUpdateInstancepoolId string

	GetAgentPoolsUpdateInstanceupdateId string

	GetAgentPoolsUpdateInstanceBackup bool
)

func NewGetAgentPoolsUpdateInstanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getUpdateInstance",
		Long: "Retrieve an Agent Pool update by id",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AgentPoolsAPI.GetAgentPoolsUpdateInstance(apiClient.GetConfig().Context, GetAgentPoolsUpdateInstancepoolId, GetAgentPoolsUpdateInstanceupdateId)

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
			if GetAgentPoolsUpdateInstanceBackup {

				idParam := GetAgentPoolsUpdateInstancepoolId
				err := utils.BackupObject(d, "AgentPools", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetAgentPoolsUpdateInstancepoolId, "poolId", "", "", "")
	cmd.MarkFlagRequired("poolId")

	cmd.Flags().StringVarP(&GetAgentPoolsUpdateInstanceupdateId, "updateId", "", "", "")
	cmd.MarkFlagRequired("updateId")

	cmd.Flags().BoolVarP(&GetAgentPoolsUpdateInstanceBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetAgentPoolsUpdateInstanceCmd := NewGetAgentPoolsUpdateInstanceCmd()
	AgentPoolsCmd.AddCommand(GetAgentPoolsUpdateInstanceCmd)
}

var (
	DeleteAgentPoolsUpdatepoolId string

	DeleteAgentPoolsUpdateupdateId string

	DeleteAgentPoolsUpdateBackup bool
)

func NewDeleteAgentPoolsUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deleteUpdate",
		Long: "Delete an Agent Pool update",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AgentPoolsAPI.DeleteAgentPoolsUpdate(apiClient.GetConfig().Context, DeleteAgentPoolsUpdatepoolId, DeleteAgentPoolsUpdateupdateId)

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
			if DeleteAgentPoolsUpdateBackup {

				idParam := DeleteAgentPoolsUpdatepoolId
				err := utils.BackupObject(d, "AgentPools", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteAgentPoolsUpdatepoolId, "poolId", "", "", "")
	cmd.MarkFlagRequired("poolId")

	cmd.Flags().StringVarP(&DeleteAgentPoolsUpdateupdateId, "updateId", "", "", "")
	cmd.MarkFlagRequired("updateId")

	cmd.Flags().BoolVarP(&DeleteAgentPoolsUpdateBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteAgentPoolsUpdateCmd := NewDeleteAgentPoolsUpdateCmd()
	AgentPoolsCmd.AddCommand(DeleteAgentPoolsUpdateCmd)
}

var (
	ActivateAgentPoolsUpdatepoolId string

	ActivateAgentPoolsUpdateupdateId string

	ActivateAgentPoolsUpdateBackup bool
)

func NewActivateAgentPoolsUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "activateUpdate",
		Long: "Activate an Agent Pool update",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AgentPoolsAPI.ActivateAgentPoolsUpdate(apiClient.GetConfig().Context, ActivateAgentPoolsUpdatepoolId, ActivateAgentPoolsUpdateupdateId)

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
			if ActivateAgentPoolsUpdateBackup {

				idParam := ActivateAgentPoolsUpdatepoolId
				err := utils.BackupObject(d, "AgentPools", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ActivateAgentPoolsUpdatepoolId, "poolId", "", "", "")
	cmd.MarkFlagRequired("poolId")

	cmd.Flags().StringVarP(&ActivateAgentPoolsUpdateupdateId, "updateId", "", "", "")
	cmd.MarkFlagRequired("updateId")

	cmd.Flags().BoolVarP(&ActivateAgentPoolsUpdateBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ActivateAgentPoolsUpdateCmd := NewActivateAgentPoolsUpdateCmd()
	AgentPoolsCmd.AddCommand(ActivateAgentPoolsUpdateCmd)
}

var (
	DeactivateAgentPoolsUpdatepoolId string

	DeactivateAgentPoolsUpdateupdateId string

	DeactivateAgentPoolsUpdateBackup bool
)

func NewDeactivateAgentPoolsUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deactivateUpdate",
		Long: "Deactivate an Agent Pool update",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AgentPoolsAPI.DeactivateAgentPoolsUpdate(apiClient.GetConfig().Context, DeactivateAgentPoolsUpdatepoolId, DeactivateAgentPoolsUpdateupdateId)

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
			if DeactivateAgentPoolsUpdateBackup {

				idParam := DeactivateAgentPoolsUpdatepoolId
				err := utils.BackupObject(d, "AgentPools", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeactivateAgentPoolsUpdatepoolId, "poolId", "", "", "")
	cmd.MarkFlagRequired("poolId")

	cmd.Flags().StringVarP(&DeactivateAgentPoolsUpdateupdateId, "updateId", "", "", "")
	cmd.MarkFlagRequired("updateId")

	cmd.Flags().BoolVarP(&DeactivateAgentPoolsUpdateBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeactivateAgentPoolsUpdateCmd := NewDeactivateAgentPoolsUpdateCmd()
	AgentPoolsCmd.AddCommand(DeactivateAgentPoolsUpdateCmd)
}

var (
	PauseAgentPoolsUpdatepoolId string

	PauseAgentPoolsUpdateupdateId string

	PauseAgentPoolsUpdateBackup bool
)

func NewPauseAgentPoolsUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "pauseUpdate",
		Long: "Pause an Agent Pool update",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AgentPoolsAPI.PauseAgentPoolsUpdate(apiClient.GetConfig().Context, PauseAgentPoolsUpdatepoolId, PauseAgentPoolsUpdateupdateId)

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
			if PauseAgentPoolsUpdateBackup {

				idParam := PauseAgentPoolsUpdatepoolId
				err := utils.BackupObject(d, "AgentPools", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&PauseAgentPoolsUpdatepoolId, "poolId", "", "", "")
	cmd.MarkFlagRequired("poolId")

	cmd.Flags().StringVarP(&PauseAgentPoolsUpdateupdateId, "updateId", "", "", "")
	cmd.MarkFlagRequired("updateId")

	cmd.Flags().BoolVarP(&PauseAgentPoolsUpdateBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	PauseAgentPoolsUpdateCmd := NewPauseAgentPoolsUpdateCmd()
	AgentPoolsCmd.AddCommand(PauseAgentPoolsUpdateCmd)
}

var (
	ResumeAgentPoolsUpdatepoolId string

	ResumeAgentPoolsUpdateupdateId string

	ResumeAgentPoolsUpdateBackup bool
)

func NewResumeAgentPoolsUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "resumeUpdate",
		Long: "Resume an Agent Pool update",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AgentPoolsAPI.ResumeAgentPoolsUpdate(apiClient.GetConfig().Context, ResumeAgentPoolsUpdatepoolId, ResumeAgentPoolsUpdateupdateId)

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
			if ResumeAgentPoolsUpdateBackup {

				idParam := ResumeAgentPoolsUpdatepoolId
				err := utils.BackupObject(d, "AgentPools", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ResumeAgentPoolsUpdatepoolId, "poolId", "", "", "")
	cmd.MarkFlagRequired("poolId")

	cmd.Flags().StringVarP(&ResumeAgentPoolsUpdateupdateId, "updateId", "", "", "")
	cmd.MarkFlagRequired("updateId")

	cmd.Flags().BoolVarP(&ResumeAgentPoolsUpdateBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ResumeAgentPoolsUpdateCmd := NewResumeAgentPoolsUpdateCmd()
	AgentPoolsCmd.AddCommand(ResumeAgentPoolsUpdateCmd)
}

var (
	RetryAgentPoolsUpdatepoolId string

	RetryAgentPoolsUpdateupdateId string

	RetryAgentPoolsUpdateBackup bool
)

func NewRetryAgentPoolsUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "retryUpdate",
		Long: "Retry an Agent Pool update",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AgentPoolsAPI.RetryAgentPoolsUpdate(apiClient.GetConfig().Context, RetryAgentPoolsUpdatepoolId, RetryAgentPoolsUpdateupdateId)

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
			if RetryAgentPoolsUpdateBackup {

				idParam := RetryAgentPoolsUpdatepoolId
				err := utils.BackupObject(d, "AgentPools", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&RetryAgentPoolsUpdatepoolId, "poolId", "", "", "")
	cmd.MarkFlagRequired("poolId")

	cmd.Flags().StringVarP(&RetryAgentPoolsUpdateupdateId, "updateId", "", "", "")
	cmd.MarkFlagRequired("updateId")

	cmd.Flags().BoolVarP(&RetryAgentPoolsUpdateBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	RetryAgentPoolsUpdateCmd := NewRetryAgentPoolsUpdateCmd()
	AgentPoolsCmd.AddCommand(RetryAgentPoolsUpdateCmd)
}

var (
	StopAgentPoolsUpdatepoolId string

	StopAgentPoolsUpdateupdateId string

	StopAgentPoolsUpdateBackup bool
)

func NewStopAgentPoolsUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "stopUpdate",
		Long: "Stop an Agent Pool update",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.AgentPoolsAPI.StopAgentPoolsUpdate(apiClient.GetConfig().Context, StopAgentPoolsUpdatepoolId, StopAgentPoolsUpdateupdateId)

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
			if StopAgentPoolsUpdateBackup {

				idParam := StopAgentPoolsUpdatepoolId
				err := utils.BackupObject(d, "AgentPools", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&StopAgentPoolsUpdatepoolId, "poolId", "", "", "")
	cmd.MarkFlagRequired("poolId")

	cmd.Flags().StringVarP(&StopAgentPoolsUpdateupdateId, "updateId", "", "", "")
	cmd.MarkFlagRequired("updateId")

	cmd.Flags().BoolVarP(&StopAgentPoolsUpdateBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	StopAgentPoolsUpdateCmd := NewStopAgentPoolsUpdateCmd()
	AgentPoolsCmd.AddCommand(StopAgentPoolsUpdateCmd)
}
