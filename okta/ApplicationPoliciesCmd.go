package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var ApplicationPoliciesCmd = &cobra.Command{
	Use:  "applicationPolicies",
	Long: "Manage ApplicationPoliciesAPI",
}

func init() {
	rootCmd.AddCommand(ApplicationPoliciesCmd)
}

var (
	AssignApplicationPolicyappId string

	AssignApplicationPolicypolicyId string

	AssignApplicationPolicyBackup bool
)

func NewAssignApplicationPolicyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "assignApplicationPolicy",
		Long: "Assign an application to a Policy",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationPoliciesAPI.AssignApplicationPolicy(apiClient.GetConfig().Context, AssignApplicationPolicyappId, AssignApplicationPolicypolicyId)

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
			if AssignApplicationPolicyBackup {

				idParam := AssignApplicationPolicyappId
				err := utils.BackupObject(d, "ApplicationPolicies", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&AssignApplicationPolicyappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().StringVarP(&AssignApplicationPolicypolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().BoolVarP(&AssignApplicationPolicyBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	AssignApplicationPolicyCmd := NewAssignApplicationPolicyCmd()
	ApplicationPoliciesCmd.AddCommand(AssignApplicationPolicyCmd)
}
