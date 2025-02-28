package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var BehaviorCmd = &cobra.Command{
	Use:  "behavior",
	Long: "Manage BehaviorAPI",
}

func init() {
	rootCmd.AddCommand(BehaviorCmd)
}

var (
	CreateBehaviorDetectionRuledata string

	CreateBehaviorDetectionRuleBackup bool
)

func NewCreateBehaviorDetectionRuleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "createDetectionRule",
		Long: "Create a Behavior Detection Rule",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.BehaviorAPI.CreateBehaviorDetectionRule(apiClient.GetConfig().Context)

			if CreateBehaviorDetectionRuledata != "" {
				req = req.Data(CreateBehaviorDetectionRuledata)
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
			if CreateBehaviorDetectionRuleBackup {

				err := utils.BackupObject(d, "Behavior", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateBehaviorDetectionRuledata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateBehaviorDetectionRuleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateBehaviorDetectionRuleCmd := NewCreateBehaviorDetectionRuleCmd()
	BehaviorCmd.AddCommand(CreateBehaviorDetectionRuleCmd)
}

var ListBehaviorDetectionRulesBackup bool

func NewListBehaviorDetectionRulesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listDetectionRules",
		Long: "List all Behavior Detection Rules",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.BehaviorAPI.ListBehaviorDetectionRules(apiClient.GetConfig().Context)

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
			if ListBehaviorDetectionRulesBackup {

				err := utils.BackupObject(d, "Behavior", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListBehaviorDetectionRulesBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListBehaviorDetectionRulesCmd := NewListBehaviorDetectionRulesCmd()
	BehaviorCmd.AddCommand(ListBehaviorDetectionRulesCmd)
}

var (
	GetBehaviorDetectionRulebehaviorId string

	GetBehaviorDetectionRuleBackup bool
)

func NewGetBehaviorDetectionRuleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getDetectionRule",
		Long: "Retrieve a Behavior Detection Rule",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.BehaviorAPI.GetBehaviorDetectionRule(apiClient.GetConfig().Context, GetBehaviorDetectionRulebehaviorId)

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
			if GetBehaviorDetectionRuleBackup {

				idParam := GetBehaviorDetectionRulebehaviorId
				err := utils.BackupObject(d, "Behavior", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetBehaviorDetectionRulebehaviorId, "behaviorId", "", "", "")
	cmd.MarkFlagRequired("behaviorId")

	cmd.Flags().BoolVarP(&GetBehaviorDetectionRuleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetBehaviorDetectionRuleCmd := NewGetBehaviorDetectionRuleCmd()
	BehaviorCmd.AddCommand(GetBehaviorDetectionRuleCmd)
}

var (
	ReplaceBehaviorDetectionRulebehaviorId string

	ReplaceBehaviorDetectionRuledata string

	ReplaceBehaviorDetectionRuleBackup bool
)

func NewReplaceBehaviorDetectionRuleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replaceDetectionRule",
		Long: "Replace a Behavior Detection Rule",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.BehaviorAPI.ReplaceBehaviorDetectionRule(apiClient.GetConfig().Context, ReplaceBehaviorDetectionRulebehaviorId)

			if ReplaceBehaviorDetectionRuledata != "" {
				req = req.Data(ReplaceBehaviorDetectionRuledata)
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
			if ReplaceBehaviorDetectionRuleBackup {

				idParam := ReplaceBehaviorDetectionRulebehaviorId
				err := utils.BackupObject(d, "Behavior", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceBehaviorDetectionRulebehaviorId, "behaviorId", "", "", "")
	cmd.MarkFlagRequired("behaviorId")

	cmd.Flags().StringVarP(&ReplaceBehaviorDetectionRuledata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceBehaviorDetectionRuleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceBehaviorDetectionRuleCmd := NewReplaceBehaviorDetectionRuleCmd()
	BehaviorCmd.AddCommand(ReplaceBehaviorDetectionRuleCmd)
}

var (
	DeleteBehaviorDetectionRulebehaviorId string

	DeleteBehaviorDetectionRuleBackup bool
)

func NewDeleteBehaviorDetectionRuleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deleteDetectionRule",
		Long: "Delete a Behavior Detection Rule",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.BehaviorAPI.DeleteBehaviorDetectionRule(apiClient.GetConfig().Context, DeleteBehaviorDetectionRulebehaviorId)

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
			if DeleteBehaviorDetectionRuleBackup {

				idParam := DeleteBehaviorDetectionRulebehaviorId
				err := utils.BackupObject(d, "Behavior", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteBehaviorDetectionRulebehaviorId, "behaviorId", "", "", "")
	cmd.MarkFlagRequired("behaviorId")

	cmd.Flags().BoolVarP(&DeleteBehaviorDetectionRuleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteBehaviorDetectionRuleCmd := NewDeleteBehaviorDetectionRuleCmd()
	BehaviorCmd.AddCommand(DeleteBehaviorDetectionRuleCmd)
}

var (
	ActivateBehaviorDetectionRulebehaviorId string

	ActivateBehaviorDetectionRuleBackup bool
)

func NewActivateBehaviorDetectionRuleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "activateDetectionRule",
		Long: "Activate a Behavior Detection Rule",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.BehaviorAPI.ActivateBehaviorDetectionRule(apiClient.GetConfig().Context, ActivateBehaviorDetectionRulebehaviorId)

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
			if ActivateBehaviorDetectionRuleBackup {

				idParam := ActivateBehaviorDetectionRulebehaviorId
				err := utils.BackupObject(d, "Behavior", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ActivateBehaviorDetectionRulebehaviorId, "behaviorId", "", "", "")
	cmd.MarkFlagRequired("behaviorId")

	cmd.Flags().BoolVarP(&ActivateBehaviorDetectionRuleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ActivateBehaviorDetectionRuleCmd := NewActivateBehaviorDetectionRuleCmd()
	BehaviorCmd.AddCommand(ActivateBehaviorDetectionRuleCmd)
}

var (
	DeactivateBehaviorDetectionRulebehaviorId string

	DeactivateBehaviorDetectionRuleBackup bool
)

func NewDeactivateBehaviorDetectionRuleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deactivateDetectionRule",
		Long: "Deactivate a Behavior Detection Rule",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.BehaviorAPI.DeactivateBehaviorDetectionRule(apiClient.GetConfig().Context, DeactivateBehaviorDetectionRulebehaviorId)

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
			if DeactivateBehaviorDetectionRuleBackup {

				idParam := DeactivateBehaviorDetectionRulebehaviorId
				err := utils.BackupObject(d, "Behavior", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeactivateBehaviorDetectionRulebehaviorId, "behaviorId", "", "", "")
	cmd.MarkFlagRequired("behaviorId")

	cmd.Flags().BoolVarP(&DeactivateBehaviorDetectionRuleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeactivateBehaviorDetectionRuleCmd := NewDeactivateBehaviorDetectionRuleCmd()
	BehaviorCmd.AddCommand(DeactivateBehaviorDetectionRuleCmd)
}
