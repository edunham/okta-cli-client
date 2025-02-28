package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var ApplicationFeaturesCmd = &cobra.Command{
	Use:  "applicationFeatures",
	Long: "Manage ApplicationFeaturesAPI",
}

func init() {
	rootCmd.AddCommand(ApplicationFeaturesCmd)
}

var (
	ListFeaturesForApplicationappId string

	ListFeaturesForApplicationBackup bool
)

func NewListFeaturesForApplicationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listFeaturesForApplication",
		Long: "List all Features",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationFeaturesAPI.ListFeaturesForApplication(apiClient.GetConfig().Context, ListFeaturesForApplicationappId)

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
			if ListFeaturesForApplicationBackup {

				idParam := ListFeaturesForApplicationappId
				err := utils.BackupObject(d, "ApplicationFeatures", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListFeaturesForApplicationappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().BoolVarP(&ListFeaturesForApplicationBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListFeaturesForApplicationCmd := NewListFeaturesForApplicationCmd()
	ApplicationFeaturesCmd.AddCommand(ListFeaturesForApplicationCmd)
}

var (
	GetFeatureForApplicationappId string

	GetFeatureForApplicationfeatureName string

	GetFeatureForApplicationBackup bool
)

func NewGetFeatureForApplicationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getFeatureForApplication",
		Long: "Retrieve a Feature",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationFeaturesAPI.GetFeatureForApplication(apiClient.GetConfig().Context, GetFeatureForApplicationappId, GetFeatureForApplicationfeatureName)

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
			if GetFeatureForApplicationBackup {

				idParam := GetFeatureForApplicationappId
				err := utils.BackupObject(d, "ApplicationFeatures", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetFeatureForApplicationappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().StringVarP(&GetFeatureForApplicationfeatureName, "featureName", "", "", "")
	cmd.MarkFlagRequired("featureName")

	cmd.Flags().BoolVarP(&GetFeatureForApplicationBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetFeatureForApplicationCmd := NewGetFeatureForApplicationCmd()
	ApplicationFeaturesCmd.AddCommand(GetFeatureForApplicationCmd)
}

var (
	UpdateFeatureForApplicationappId string

	UpdateFeatureForApplicationfeatureName string

	UpdateFeatureForApplicationdata string

	UpdateFeatureForApplicationBackup bool
)

func NewUpdateFeatureForApplicationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "updateFeatureForApplication",
		Long: "Update a Feature",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.ApplicationFeaturesAPI.UpdateFeatureForApplication(apiClient.GetConfig().Context, UpdateFeatureForApplicationappId, UpdateFeatureForApplicationfeatureName)

			if UpdateFeatureForApplicationdata != "" {
				req = req.Data(UpdateFeatureForApplicationdata)
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
			if UpdateFeatureForApplicationBackup {

				idParam := UpdateFeatureForApplicationappId
				err := utils.BackupObject(d, "ApplicationFeatures", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UpdateFeatureForApplicationappId, "appId", "", "", "")
	cmd.MarkFlagRequired("appId")

	cmd.Flags().StringVarP(&UpdateFeatureForApplicationfeatureName, "featureName", "", "", "")
	cmd.MarkFlagRequired("featureName")

	cmd.Flags().StringVarP(&UpdateFeatureForApplicationdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&UpdateFeatureForApplicationBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UpdateFeatureForApplicationCmd := NewUpdateFeatureForApplicationCmd()
	ApplicationFeaturesCmd.AddCommand(UpdateFeatureForApplicationCmd)
}
