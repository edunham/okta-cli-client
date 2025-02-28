package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var PolicyCmd = &cobra.Command{
	Use:  "policy",
	Long: "Manage PolicyAPI",
}

func init() {
	rootCmd.AddCommand(PolicyCmd)
}

var (
	CreatePolicydata string

	CreatePolicyBackup bool
)

func NewCreatePolicyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create",
		Long: "Create a Policy",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PolicyAPI.CreatePolicy(apiClient.GetConfig().Context)

			if CreatePolicydata != "" {
				req = req.Data(CreatePolicydata)
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
			if CreatePolicyBackup {

				err := utils.BackupObject(d, "Policy", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreatePolicydata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreatePolicyBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreatePolicyCmd := NewCreatePolicyCmd()
	PolicyCmd.AddCommand(CreatePolicyCmd)
}

var ListPoliciesBackup bool

func NewListPoliciesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listPolicies",
		Long: "List all Policies",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PolicyAPI.ListPolicies(apiClient.GetConfig().Context)

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
			if ListPoliciesBackup {

				err := utils.BackupObject(d, "Policy", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListPoliciesBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListPoliciesCmd := NewListPoliciesCmd()
	PolicyCmd.AddCommand(ListPoliciesCmd)
}

var (
	CreatePolicySimulationdata string

	CreatePolicySimulationBackup bool
)

func NewCreatePolicySimulationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "createSimulation",
		Long: "Create a Policy Simulation",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PolicyAPI.CreatePolicySimulation(apiClient.GetConfig().Context)

			if CreatePolicySimulationdata != "" {
				req = req.Data(CreatePolicySimulationdata)
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
			if CreatePolicySimulationBackup {

				err := utils.BackupObject(d, "Policy", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreatePolicySimulationdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreatePolicySimulationBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreatePolicySimulationCmd := NewCreatePolicySimulationCmd()
	PolicyCmd.AddCommand(CreatePolicySimulationCmd)
}

var (
	GetPolicypolicyId string

	GetPolicyBackup bool
)

func NewGetPolicyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "get",
		Long: "Retrieve a Policy",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PolicyAPI.GetPolicy(apiClient.GetConfig().Context, GetPolicypolicyId)

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
			if GetPolicyBackup {

				idParam := GetPolicypolicyId
				err := utils.BackupObject(d, "Policy", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetPolicypolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().BoolVarP(&GetPolicyBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetPolicyCmd := NewGetPolicyCmd()
	PolicyCmd.AddCommand(GetPolicyCmd)
}

var (
	ReplacePolicypolicyId string

	ReplacePolicydata string

	ReplacePolicyBackup bool
)

func NewReplacePolicyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replace",
		Long: "Replace a Policy",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PolicyAPI.ReplacePolicy(apiClient.GetConfig().Context, ReplacePolicypolicyId)

			if ReplacePolicydata != "" {
				req = req.Data(ReplacePolicydata)
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
			if ReplacePolicyBackup {

				idParam := ReplacePolicypolicyId
				err := utils.BackupObject(d, "Policy", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplacePolicypolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().StringVarP(&ReplacePolicydata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplacePolicyBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplacePolicyCmd := NewReplacePolicyCmd()
	PolicyCmd.AddCommand(ReplacePolicyCmd)
}

var (
	DeletePolicypolicyId string

	DeletePolicyBackup bool
)

func NewDeletePolicyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete",
		Long: "Delete a Policy",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PolicyAPI.DeletePolicy(apiClient.GetConfig().Context, DeletePolicypolicyId)

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
			if DeletePolicyBackup {

				idParam := DeletePolicypolicyId
				err := utils.BackupObject(d, "Policy", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeletePolicypolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().BoolVarP(&DeletePolicyBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeletePolicyCmd := NewDeletePolicyCmd()
	PolicyCmd.AddCommand(DeletePolicyCmd)
}

var (
	ListPolicyAppspolicyId string

	ListPolicyAppsBackup bool
)

func NewListPolicyAppsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listApps",
		Long: "List all Applications mapped to a Policy",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PolicyAPI.ListPolicyApps(apiClient.GetConfig().Context, ListPolicyAppspolicyId)

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
			if ListPolicyAppsBackup {

				idParam := ListPolicyAppspolicyId
				err := utils.BackupObject(d, "Policy", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListPolicyAppspolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().BoolVarP(&ListPolicyAppsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListPolicyAppsCmd := NewListPolicyAppsCmd()
	PolicyCmd.AddCommand(ListPolicyAppsCmd)
}

var (
	ClonePolicypolicyId string

	ClonePolicyBackup bool
)

func NewClonePolicyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "clone",
		Long: "Clone an existing Policy",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PolicyAPI.ClonePolicy(apiClient.GetConfig().Context, ClonePolicypolicyId)

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
			if ClonePolicyBackup {

				idParam := ClonePolicypolicyId
				err := utils.BackupObject(d, "Policy", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ClonePolicypolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().BoolVarP(&ClonePolicyBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ClonePolicyCmd := NewClonePolicyCmd()
	PolicyCmd.AddCommand(ClonePolicyCmd)
}

var (
	ActivatePolicypolicyId string

	ActivatePolicyBackup bool
)

func NewActivatePolicyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "activate",
		Long: "Activate a Policy",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PolicyAPI.ActivatePolicy(apiClient.GetConfig().Context, ActivatePolicypolicyId)

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
			if ActivatePolicyBackup {

				idParam := ActivatePolicypolicyId
				err := utils.BackupObject(d, "Policy", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ActivatePolicypolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().BoolVarP(&ActivatePolicyBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ActivatePolicyCmd := NewActivatePolicyCmd()
	PolicyCmd.AddCommand(ActivatePolicyCmd)
}

var (
	DeactivatePolicypolicyId string

	DeactivatePolicyBackup bool
)

func NewDeactivatePolicyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deactivate",
		Long: "Deactivate a Policy",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PolicyAPI.DeactivatePolicy(apiClient.GetConfig().Context, DeactivatePolicypolicyId)

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
			if DeactivatePolicyBackup {

				idParam := DeactivatePolicypolicyId
				err := utils.BackupObject(d, "Policy", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeactivatePolicypolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().BoolVarP(&DeactivatePolicyBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeactivatePolicyCmd := NewDeactivatePolicyCmd()
	PolicyCmd.AddCommand(DeactivatePolicyCmd)
}

var (
	MapResourceToPolicypolicyId string

	MapResourceToPolicydata string

	MapResourceToPolicyBackup bool
)

func NewMapResourceToPolicyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "mapResourceTo",
		Long: "Map a resource to a Policy",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PolicyAPI.MapResourceToPolicy(apiClient.GetConfig().Context, MapResourceToPolicypolicyId)

			if MapResourceToPolicydata != "" {
				req = req.Data(MapResourceToPolicydata)
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
			if MapResourceToPolicyBackup {

				idParam := MapResourceToPolicypolicyId
				err := utils.BackupObject(d, "Policy", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&MapResourceToPolicypolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().StringVarP(&MapResourceToPolicydata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&MapResourceToPolicyBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	MapResourceToPolicyCmd := NewMapResourceToPolicyCmd()
	PolicyCmd.AddCommand(MapResourceToPolicyCmd)
}

var (
	ListPolicyMappingspolicyId string

	ListPolicyMappingsBackup bool
)

func NewListPolicyMappingsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listMappings",
		Long: "List all resources mapped to a Policy",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PolicyAPI.ListPolicyMappings(apiClient.GetConfig().Context, ListPolicyMappingspolicyId)

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
			if ListPolicyMappingsBackup {

				idParam := ListPolicyMappingspolicyId
				err := utils.BackupObject(d, "Policy", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListPolicyMappingspolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().BoolVarP(&ListPolicyMappingsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListPolicyMappingsCmd := NewListPolicyMappingsCmd()
	PolicyCmd.AddCommand(ListPolicyMappingsCmd)
}

var (
	GetPolicyMappingpolicyId string

	GetPolicyMappingmappingId string

	GetPolicyMappingBackup bool
)

func NewGetPolicyMappingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getMapping",
		Long: "Retrieve a policy resource Mapping",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PolicyAPI.GetPolicyMapping(apiClient.GetConfig().Context, GetPolicyMappingpolicyId, GetPolicyMappingmappingId)

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
			if GetPolicyMappingBackup {

				idParam := GetPolicyMappingpolicyId
				err := utils.BackupObject(d, "Policy", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetPolicyMappingpolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().StringVarP(&GetPolicyMappingmappingId, "mappingId", "", "", "")
	cmd.MarkFlagRequired("mappingId")

	cmd.Flags().BoolVarP(&GetPolicyMappingBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetPolicyMappingCmd := NewGetPolicyMappingCmd()
	PolicyCmd.AddCommand(GetPolicyMappingCmd)
}

var (
	DeletePolicyResourceMappingpolicyId string

	DeletePolicyResourceMappingmappingId string

	DeletePolicyResourceMappingBackup bool
)

func NewDeletePolicyResourceMappingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deleteResourceMapping",
		Long: "Delete a policy resource Mapping",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PolicyAPI.DeletePolicyResourceMapping(apiClient.GetConfig().Context, DeletePolicyResourceMappingpolicyId, DeletePolicyResourceMappingmappingId)

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
			if DeletePolicyResourceMappingBackup {

				idParam := DeletePolicyResourceMappingpolicyId
				err := utils.BackupObject(d, "Policy", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeletePolicyResourceMappingpolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().StringVarP(&DeletePolicyResourceMappingmappingId, "mappingId", "", "", "")
	cmd.MarkFlagRequired("mappingId")

	cmd.Flags().BoolVarP(&DeletePolicyResourceMappingBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeletePolicyResourceMappingCmd := NewDeletePolicyResourceMappingCmd()
	PolicyCmd.AddCommand(DeletePolicyResourceMappingCmd)
}

var (
	CreatePolicyRulepolicyId string

	CreatePolicyRuledata string

	CreatePolicyRuleBackup bool
)

func NewCreatePolicyRuleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "createRule",
		Long: "Create a Policy Rule",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PolicyAPI.CreatePolicyRule(apiClient.GetConfig().Context, CreatePolicyRulepolicyId)

			if CreatePolicyRuledata != "" {
				req = req.Data(CreatePolicyRuledata)
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
			if CreatePolicyRuleBackup {

				idParam := CreatePolicyRulepolicyId
				err := utils.BackupObject(d, "Policy", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreatePolicyRulepolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().StringVarP(&CreatePolicyRuledata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreatePolicyRuleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreatePolicyRuleCmd := NewCreatePolicyRuleCmd()
	PolicyCmd.AddCommand(CreatePolicyRuleCmd)
}

var (
	ListPolicyRulespolicyId string

	ListPolicyRulesBackup bool
)

func NewListPolicyRulesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listRules",
		Long: "List all Policy Rules",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PolicyAPI.ListPolicyRules(apiClient.GetConfig().Context, ListPolicyRulespolicyId)

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
			if ListPolicyRulesBackup {

				idParam := ListPolicyRulespolicyId
				err := utils.BackupObject(d, "Policy", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListPolicyRulespolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().BoolVarP(&ListPolicyRulesBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListPolicyRulesCmd := NewListPolicyRulesCmd()
	PolicyCmd.AddCommand(ListPolicyRulesCmd)
}

var (
	GetPolicyRulepolicyId string

	GetPolicyRuleruleId string

	GetPolicyRuleBackup bool
)

func NewGetPolicyRuleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getRule",
		Long: "Retrieve a Policy Rule",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PolicyAPI.GetPolicyRule(apiClient.GetConfig().Context, GetPolicyRulepolicyId, GetPolicyRuleruleId)

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
			if GetPolicyRuleBackup {

				idParam := GetPolicyRulepolicyId
				err := utils.BackupObject(d, "Policy", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetPolicyRulepolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().StringVarP(&GetPolicyRuleruleId, "ruleId", "", "", "")
	cmd.MarkFlagRequired("ruleId")

	cmd.Flags().BoolVarP(&GetPolicyRuleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetPolicyRuleCmd := NewGetPolicyRuleCmd()
	PolicyCmd.AddCommand(GetPolicyRuleCmd)
}

var (
	ReplacePolicyRulepolicyId string

	ReplacePolicyRuleruleId string

	ReplacePolicyRuledata string

	ReplacePolicyRuleBackup bool
)

func NewReplacePolicyRuleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replaceRule",
		Long: "Replace a Policy Rule",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PolicyAPI.ReplacePolicyRule(apiClient.GetConfig().Context, ReplacePolicyRulepolicyId, ReplacePolicyRuleruleId)

			if ReplacePolicyRuledata != "" {
				req = req.Data(ReplacePolicyRuledata)
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
			if ReplacePolicyRuleBackup {

				idParam := ReplacePolicyRulepolicyId
				err := utils.BackupObject(d, "Policy", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplacePolicyRulepolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().StringVarP(&ReplacePolicyRuleruleId, "ruleId", "", "", "")
	cmd.MarkFlagRequired("ruleId")

	cmd.Flags().StringVarP(&ReplacePolicyRuledata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplacePolicyRuleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplacePolicyRuleCmd := NewReplacePolicyRuleCmd()
	PolicyCmd.AddCommand(ReplacePolicyRuleCmd)
}

var (
	DeletePolicyRulepolicyId string

	DeletePolicyRuleruleId string

	DeletePolicyRuleBackup bool
)

func NewDeletePolicyRuleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deleteRule",
		Long: "Delete a Policy Rule",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PolicyAPI.DeletePolicyRule(apiClient.GetConfig().Context, DeletePolicyRulepolicyId, DeletePolicyRuleruleId)

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
			if DeletePolicyRuleBackup {

				idParam := DeletePolicyRulepolicyId
				err := utils.BackupObject(d, "Policy", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeletePolicyRulepolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().StringVarP(&DeletePolicyRuleruleId, "ruleId", "", "", "")
	cmd.MarkFlagRequired("ruleId")

	cmd.Flags().BoolVarP(&DeletePolicyRuleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeletePolicyRuleCmd := NewDeletePolicyRuleCmd()
	PolicyCmd.AddCommand(DeletePolicyRuleCmd)
}

var (
	ActivatePolicyRulepolicyId string

	ActivatePolicyRuleruleId string

	ActivatePolicyRuleBackup bool
)

func NewActivatePolicyRuleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "activateRule",
		Long: "Activate a Policy Rule",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PolicyAPI.ActivatePolicyRule(apiClient.GetConfig().Context, ActivatePolicyRulepolicyId, ActivatePolicyRuleruleId)

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
			if ActivatePolicyRuleBackup {

				idParam := ActivatePolicyRulepolicyId
				err := utils.BackupObject(d, "Policy", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ActivatePolicyRulepolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().StringVarP(&ActivatePolicyRuleruleId, "ruleId", "", "", "")
	cmd.MarkFlagRequired("ruleId")

	cmd.Flags().BoolVarP(&ActivatePolicyRuleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ActivatePolicyRuleCmd := NewActivatePolicyRuleCmd()
	PolicyCmd.AddCommand(ActivatePolicyRuleCmd)
}

var (
	DeactivatePolicyRulepolicyId string

	DeactivatePolicyRuleruleId string

	DeactivatePolicyRuleBackup bool
)

func NewDeactivatePolicyRuleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deactivateRule",
		Long: "Deactivate a Policy Rule",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PolicyAPI.DeactivatePolicyRule(apiClient.GetConfig().Context, DeactivatePolicyRulepolicyId, DeactivatePolicyRuleruleId)

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
			if DeactivatePolicyRuleBackup {

				idParam := DeactivatePolicyRulepolicyId
				err := utils.BackupObject(d, "Policy", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeactivatePolicyRulepolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().StringVarP(&DeactivatePolicyRuleruleId, "ruleId", "", "", "")
	cmd.MarkFlagRequired("ruleId")

	cmd.Flags().BoolVarP(&DeactivatePolicyRuleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeactivatePolicyRuleCmd := NewDeactivatePolicyRuleCmd()
	PolicyCmd.AddCommand(DeactivatePolicyRuleCmd)
}
