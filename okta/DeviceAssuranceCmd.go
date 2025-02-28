package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var DeviceAssuranceCmd = &cobra.Command{
	Use:  "deviceAssurance",
	Long: "Manage DeviceAssuranceAPI",
}

func init() {
	rootCmd.AddCommand(DeviceAssuranceCmd)
}

var (
	CreateDeviceAssurancePolicydata string

	CreateDeviceAssurancePolicyBackup bool
)

func NewCreateDeviceAssurancePolicyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "createPolicy",
		Long: "Create a Device Assurance Policy",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.DeviceAssuranceAPI.CreateDeviceAssurancePolicy(apiClient.GetConfig().Context)

			if CreateDeviceAssurancePolicydata != "" {
				req = req.Data(CreateDeviceAssurancePolicydata)
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
			if CreateDeviceAssurancePolicyBackup {

				err := utils.BackupObject(d, "DeviceAssurance", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateDeviceAssurancePolicydata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateDeviceAssurancePolicyBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateDeviceAssurancePolicyCmd := NewCreateDeviceAssurancePolicyCmd()
	DeviceAssuranceCmd.AddCommand(CreateDeviceAssurancePolicyCmd)
}

var ListDeviceAssurancePoliciesBackup bool

func NewListDeviceAssurancePoliciesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listPolicies",
		Long: "List all Device Assurance Policies",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.DeviceAssuranceAPI.ListDeviceAssurancePolicies(apiClient.GetConfig().Context)

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
			if ListDeviceAssurancePoliciesBackup {

				err := utils.BackupObject(d, "DeviceAssurance", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListDeviceAssurancePoliciesBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListDeviceAssurancePoliciesCmd := NewListDeviceAssurancePoliciesCmd()
	DeviceAssuranceCmd.AddCommand(ListDeviceAssurancePoliciesCmd)
}

var (
	GetDeviceAssurancePolicydeviceAssuranceId string

	GetDeviceAssurancePolicyBackup bool
)

func NewGetDeviceAssurancePolicyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getPolicy",
		Long: "Retrieve a Device Assurance Policy",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.DeviceAssuranceAPI.GetDeviceAssurancePolicy(apiClient.GetConfig().Context, GetDeviceAssurancePolicydeviceAssuranceId)

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
			if GetDeviceAssurancePolicyBackup {

				idParam := GetDeviceAssurancePolicydeviceAssuranceId
				err := utils.BackupObject(d, "DeviceAssurance", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetDeviceAssurancePolicydeviceAssuranceId, "deviceAssuranceId", "", "", "")
	cmd.MarkFlagRequired("deviceAssuranceId")

	cmd.Flags().BoolVarP(&GetDeviceAssurancePolicyBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetDeviceAssurancePolicyCmd := NewGetDeviceAssurancePolicyCmd()
	DeviceAssuranceCmd.AddCommand(GetDeviceAssurancePolicyCmd)
}

var (
	ReplaceDeviceAssurancePolicydeviceAssuranceId string

	ReplaceDeviceAssurancePolicydata string

	ReplaceDeviceAssurancePolicyBackup bool
)

func NewReplaceDeviceAssurancePolicyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replacePolicy",
		Long: "Replace a Device Assurance Policy",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.DeviceAssuranceAPI.ReplaceDeviceAssurancePolicy(apiClient.GetConfig().Context, ReplaceDeviceAssurancePolicydeviceAssuranceId)

			if ReplaceDeviceAssurancePolicydata != "" {
				req = req.Data(ReplaceDeviceAssurancePolicydata)
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
			if ReplaceDeviceAssurancePolicyBackup {

				idParam := ReplaceDeviceAssurancePolicydeviceAssuranceId
				err := utils.BackupObject(d, "DeviceAssurance", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceDeviceAssurancePolicydeviceAssuranceId, "deviceAssuranceId", "", "", "")
	cmd.MarkFlagRequired("deviceAssuranceId")

	cmd.Flags().StringVarP(&ReplaceDeviceAssurancePolicydata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceDeviceAssurancePolicyBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceDeviceAssurancePolicyCmd := NewReplaceDeviceAssurancePolicyCmd()
	DeviceAssuranceCmd.AddCommand(ReplaceDeviceAssurancePolicyCmd)
}

var (
	DeleteDeviceAssurancePolicydeviceAssuranceId string

	DeleteDeviceAssurancePolicyBackup bool
)

func NewDeleteDeviceAssurancePolicyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deletePolicy",
		Long: "Delete a Device Assurance Policy",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.DeviceAssuranceAPI.DeleteDeviceAssurancePolicy(apiClient.GetConfig().Context, DeleteDeviceAssurancePolicydeviceAssuranceId)

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
			if DeleteDeviceAssurancePolicyBackup {

				idParam := DeleteDeviceAssurancePolicydeviceAssuranceId
				err := utils.BackupObject(d, "DeviceAssurance", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteDeviceAssurancePolicydeviceAssuranceId, "deviceAssuranceId", "", "", "")
	cmd.MarkFlagRequired("deviceAssuranceId")

	cmd.Flags().BoolVarP(&DeleteDeviceAssurancePolicyBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteDeviceAssurancePolicyCmd := NewDeleteDeviceAssurancePolicyCmd()
	DeviceAssuranceCmd.AddCommand(DeleteDeviceAssurancePolicyCmd)
}
