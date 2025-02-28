package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var SessionCmd = &cobra.Command{
	Use:  "session",
	Long: "Manage SessionAPI",
}

func init() {
	rootCmd.AddCommand(SessionCmd)
}

var (
	CreateSessiondata string

	CreateSessionBackup bool
)

func NewCreateSessionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "create",
		Long: "Create a Session with session token",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.SessionAPI.CreateSession(apiClient.GetConfig().Context)

			if CreateSessiondata != "" {
				req = req.Data(CreateSessiondata)
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
			if CreateSessionBackup {

				err := utils.BackupObject(d, "Session", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&CreateSessiondata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&CreateSessionBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CreateSessionCmd := NewCreateSessionCmd()
	SessionCmd.AddCommand(CreateSessionCmd)
}

var GetCurrentSessionBackup bool

func NewGetCurrentSessionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getCurrent",
		Long: "Retrieve the current Session",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.SessionAPI.GetCurrentSession(apiClient.GetConfig().Context)

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
			if GetCurrentSessionBackup {

				err := utils.BackupObject(d, "Session", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&GetCurrentSessionBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetCurrentSessionCmd := NewGetCurrentSessionCmd()
	SessionCmd.AddCommand(GetCurrentSessionCmd)
}

var CloseCurrentSessionBackup bool

func NewCloseCurrentSessionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "closeCurrent",
		Long: "Close the current Session",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.SessionAPI.CloseCurrentSession(apiClient.GetConfig().Context)

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
			if CloseCurrentSessionBackup {

				err := utils.BackupObject(d, "Session", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&CloseCurrentSessionBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	CloseCurrentSessionCmd := NewCloseCurrentSessionCmd()
	SessionCmd.AddCommand(CloseCurrentSessionCmd)
}

var RefreshCurrentSessionBackup bool

func NewRefreshCurrentSessionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "refreshCurrent",
		Long: "Refresh the current Session",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.SessionAPI.RefreshCurrentSession(apiClient.GetConfig().Context)

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
			if RefreshCurrentSessionBackup {

				err := utils.BackupObject(d, "Session", "hasNoIdParam")
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&RefreshCurrentSessionBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	RefreshCurrentSessionCmd := NewRefreshCurrentSessionCmd()
	SessionCmd.AddCommand(RefreshCurrentSessionCmd)
}

var (
	GetSessionsessionId string

	GetSessionBackup bool
)

func NewGetSessionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "get",
		Long: "Retrieve a Session",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.SessionAPI.GetSession(apiClient.GetConfig().Context, GetSessionsessionId)

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
			if GetSessionBackup {

				idParam := GetSessionsessionId
				err := utils.BackupObject(d, "Session", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetSessionsessionId, "sessionId", "", "", "")
	cmd.MarkFlagRequired("sessionId")

	cmd.Flags().BoolVarP(&GetSessionBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetSessionCmd := NewGetSessionCmd()
	SessionCmd.AddCommand(GetSessionCmd)
}

var (
	RevokeSessionsessionId string

	RevokeSessionBackup bool
)

func NewRevokeSessionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "revoke",
		Long: "Revoke a Session",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.SessionAPI.RevokeSession(apiClient.GetConfig().Context, RevokeSessionsessionId)

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
			if RevokeSessionBackup {

				idParam := RevokeSessionsessionId
				err := utils.BackupObject(d, "Session", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&RevokeSessionsessionId, "sessionId", "", "", "")
	cmd.MarkFlagRequired("sessionId")

	cmd.Flags().BoolVarP(&RevokeSessionBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	RevokeSessionCmd := NewRevokeSessionCmd()
	SessionCmd.AddCommand(RevokeSessionCmd)
}

var (
	RefreshSessionsessionId string

	RefreshSessionBackup bool
)

func NewRefreshSessionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "refresh",
		Long: "Refresh a Session",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.SessionAPI.RefreshSession(apiClient.GetConfig().Context, RefreshSessionsessionId)

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
			if RefreshSessionBackup {

				idParam := RefreshSessionsessionId
				err := utils.BackupObject(d, "Session", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&RefreshSessionsessionId, "sessionId", "", "", "")
	cmd.MarkFlagRequired("sessionId")

	cmd.Flags().BoolVarP(&RefreshSessionBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	RefreshSessionCmd := NewRefreshSessionCmd()
	SessionCmd.AddCommand(RefreshSessionCmd)
}
