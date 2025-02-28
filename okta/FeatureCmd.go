package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var FeatureCmd = &cobra.Command{
	Use:  "feature",
	Long: "Manage FeatureAPI",
}

func init() {
	rootCmd.AddCommand(FeatureCmd)
}

var ListFeaturesBackup bool

func NewListFeaturesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "lists",
		Long: "List all Features",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.FeatureAPI.ListFeatures(apiClient.GetConfig().Context)

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
			if ListFeaturesBackup {

				err := utils.BackupObject(d, "Feature", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListFeaturesBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListFeaturesCmd := NewListFeaturesCmd()
	FeatureCmd.AddCommand(ListFeaturesCmd)
}

var (
	GetFeaturefeatureId string

	GetFeatureBackup bool
)

func NewGetFeatureCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "get",
		Long: "Retrieve a Feature",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.FeatureAPI.GetFeature(apiClient.GetConfig().Context, GetFeaturefeatureId)

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
			if GetFeatureBackup {

				idParam := GetFeaturefeatureId
				err := utils.BackupObject(d, "Feature", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetFeaturefeatureId, "featureId", "", "", "")
	cmd.MarkFlagRequired("featureId")

	cmd.Flags().BoolVarP(&GetFeatureBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetFeatureCmd := NewGetFeatureCmd()
	FeatureCmd.AddCommand(GetFeatureCmd)
}

var (
	ListFeatureDependenciesfeatureId string

	ListFeatureDependenciesBackup bool
)

func NewListFeatureDependenciesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listDependencies",
		Long: "List all dependencies",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.FeatureAPI.ListFeatureDependencies(apiClient.GetConfig().Context, ListFeatureDependenciesfeatureId)

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
			if ListFeatureDependenciesBackup {

				idParam := ListFeatureDependenciesfeatureId
				err := utils.BackupObject(d, "Feature", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListFeatureDependenciesfeatureId, "featureId", "", "", "")
	cmd.MarkFlagRequired("featureId")

	cmd.Flags().BoolVarP(&ListFeatureDependenciesBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListFeatureDependenciesCmd := NewListFeatureDependenciesCmd()
	FeatureCmd.AddCommand(ListFeatureDependenciesCmd)
}

var (
	ListFeatureDependentsfeatureId string

	ListFeatureDependentsBackup bool
)

func NewListFeatureDependentsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listDependents",
		Long: "List all dependents",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.FeatureAPI.ListFeatureDependents(apiClient.GetConfig().Context, ListFeatureDependentsfeatureId)

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
			if ListFeatureDependentsBackup {

				idParam := ListFeatureDependentsfeatureId
				err := utils.BackupObject(d, "Feature", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListFeatureDependentsfeatureId, "featureId", "", "", "")
	cmd.MarkFlagRequired("featureId")

	cmd.Flags().BoolVarP(&ListFeatureDependentsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListFeatureDependentsCmd := NewListFeatureDependentsCmd()
	FeatureCmd.AddCommand(ListFeatureDependentsCmd)
}

var (
	UpdateFeatureLifecyclefeatureId string

	UpdateFeatureLifecyclelifecycle string

	UpdateFeatureLifecycleBackup bool
)

func NewUpdateFeatureLifecycleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "updateLifecycle",
		Long: "Update a Feature lifecycle",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.FeatureAPI.UpdateFeatureLifecycle(apiClient.GetConfig().Context, UpdateFeatureLifecyclefeatureId, UpdateFeatureLifecyclelifecycle)

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
			if UpdateFeatureLifecycleBackup {

				idParam := UpdateFeatureLifecyclefeatureId
				err := utils.BackupObject(d, "Feature", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UpdateFeatureLifecyclefeatureId, "featureId", "", "", "")
	cmd.MarkFlagRequired("featureId")

	cmd.Flags().StringVarP(&UpdateFeatureLifecyclelifecycle, "lifecycle", "", "", "")
	cmd.MarkFlagRequired("lifecycle")

	cmd.Flags().BoolVarP(&UpdateFeatureLifecycleBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UpdateFeatureLifecycleCmd := NewUpdateFeatureLifecycleCmd()
	FeatureCmd.AddCommand(UpdateFeatureLifecycleCmd)
}
