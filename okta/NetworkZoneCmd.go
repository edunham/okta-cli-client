package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var NetworkZoneCmd = &cobra.Command{
	Use:  "networkZone",
	Long: "Manage NetworkZoneAPI",
}

func init() {
	rootCmd.AddCommand(NetworkZoneCmd)
}

var (
	CreateNetworkZonedata string

	CreateNetworkZoneBackup bool
)

func NewCreateNetworkZoneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create",
		Long: "Create a Network Zone",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.NetworkZoneAPI.CreateNetworkZone(apiClient.GetConfig().Context)

			if CreateNetworkZonedata != "" {
				req = req.Data(CreateNetworkZonedata)
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
			if CreateNetworkZoneBackup {

				err := utils.BackupObject(d, "NetworkZone", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateNetworkZonedata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateNetworkZoneBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateNetworkZoneCmd := NewCreateNetworkZoneCmd()
	NetworkZoneCmd.AddCommand(CreateNetworkZoneCmd)
}

var ListNetworkZonesBackup bool

func NewListNetworkZonesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "lists",
		Long: "List all Network Zones",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.NetworkZoneAPI.ListNetworkZones(apiClient.GetConfig().Context)

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
			if ListNetworkZonesBackup {

				err := utils.BackupObject(d, "NetworkZone", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&ListNetworkZonesBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListNetworkZonesCmd := NewListNetworkZonesCmd()
	NetworkZoneCmd.AddCommand(ListNetworkZonesCmd)
}

var (
	GetNetworkZonezoneId string

	GetNetworkZoneBackup bool
)

func NewGetNetworkZoneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "get",
		Long: "Retrieve a Network Zone",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.NetworkZoneAPI.GetNetworkZone(apiClient.GetConfig().Context, GetNetworkZonezoneId)

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
			if GetNetworkZoneBackup {

				idParam := GetNetworkZonezoneId
				err := utils.BackupObject(d, "NetworkZone", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetNetworkZonezoneId, "zoneId", "", "", "")
	cmd.MarkFlagRequired("zoneId")

	cmd.Flags().BoolVarP(&GetNetworkZoneBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetNetworkZoneCmd := NewGetNetworkZoneCmd()
	NetworkZoneCmd.AddCommand(GetNetworkZoneCmd)
}

var (
	ReplaceNetworkZonezoneId string

	ReplaceNetworkZonedata string

	ReplaceNetworkZoneBackup bool
)

func NewReplaceNetworkZoneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "replace",
		Long: "Replace a Network Zone",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.NetworkZoneAPI.ReplaceNetworkZone(apiClient.GetConfig().Context, ReplaceNetworkZonezoneId)

			if ReplaceNetworkZonedata != "" {
				req = req.Data(ReplaceNetworkZonedata)
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
			if ReplaceNetworkZoneBackup {

				idParam := ReplaceNetworkZonezoneId
				err := utils.BackupObject(d, "NetworkZone", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ReplaceNetworkZonezoneId, "zoneId", "", "", "")
	cmd.MarkFlagRequired("zoneId")

	cmd.Flags().StringVarP(&ReplaceNetworkZonedata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ReplaceNetworkZoneBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ReplaceNetworkZoneCmd := NewReplaceNetworkZoneCmd()
	NetworkZoneCmd.AddCommand(ReplaceNetworkZoneCmd)
}

var (
	DeleteNetworkZonezoneId string

	DeleteNetworkZoneBackup bool
)

func NewDeleteNetworkZoneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "delete",
		Long: "Delete a Network Zone",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.NetworkZoneAPI.DeleteNetworkZone(apiClient.GetConfig().Context, DeleteNetworkZonezoneId)

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
			if DeleteNetworkZoneBackup {

				idParam := DeleteNetworkZonezoneId
				err := utils.BackupObject(d, "NetworkZone", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeleteNetworkZonezoneId, "zoneId", "", "", "")
	cmd.MarkFlagRequired("zoneId")

	cmd.Flags().BoolVarP(&DeleteNetworkZoneBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeleteNetworkZoneCmd := NewDeleteNetworkZoneCmd()
	NetworkZoneCmd.AddCommand(DeleteNetworkZoneCmd)
}

var (
	ActivateNetworkZonezoneId string

	ActivateNetworkZoneBackup bool
)

func NewActivateNetworkZoneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "activate",
		Long: "Activate a Network Zone",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.NetworkZoneAPI.ActivateNetworkZone(apiClient.GetConfig().Context, ActivateNetworkZonezoneId)

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
			if ActivateNetworkZoneBackup {

				idParam := ActivateNetworkZonezoneId
				err := utils.BackupObject(d, "NetworkZone", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ActivateNetworkZonezoneId, "zoneId", "", "", "")
	cmd.MarkFlagRequired("zoneId")

	cmd.Flags().BoolVarP(&ActivateNetworkZoneBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ActivateNetworkZoneCmd := NewActivateNetworkZoneCmd()
	NetworkZoneCmd.AddCommand(ActivateNetworkZoneCmd)
}

var (
	DeactivateNetworkZonezoneId string

	DeactivateNetworkZoneBackup bool
)

func NewDeactivateNetworkZoneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "deactivate",
		Long: "Deactivate a Network Zone",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.NetworkZoneAPI.DeactivateNetworkZone(apiClient.GetConfig().Context, DeactivateNetworkZonezoneId)

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
			if DeactivateNetworkZoneBackup {

				idParam := DeactivateNetworkZonezoneId
				err := utils.BackupObject(d, "NetworkZone", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&DeactivateNetworkZonezoneId, "zoneId", "", "", "")
	cmd.MarkFlagRequired("zoneId")

	cmd.Flags().BoolVarP(&DeactivateNetworkZoneBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	DeactivateNetworkZoneCmd := NewDeactivateNetworkZoneCmd()
	NetworkZoneCmd.AddCommand(DeactivateNetworkZoneCmd)
}
