package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var UserFactorCmd = &cobra.Command{
	Use:  "userFactor",
	Long: "Manage UserFactorAPI",
}

func init() {
	rootCmd.AddCommand(UserFactorCmd)
}

var (
	EnrollFactoruserId string

	EnrollFactordata string

	EnrollFactorBackup bool
)

func NewEnrollFactorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "enrollFactor",
		Long: "Enroll a Factor",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.UserFactorAPI.EnrollFactor(apiClient.GetConfig().Context, EnrollFactoruserId)

			if EnrollFactordata != "" {
				req = req.Data(EnrollFactordata)
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
			if EnrollFactorBackup {

				idParam := EnrollFactoruserId
				err := utils.BackupObject(d, "UserFactor", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&EnrollFactoruserId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	cmd.Flags().StringVarP(&EnrollFactordata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&EnrollFactorBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	EnrollFactorCmd := NewEnrollFactorCmd()
	UserFactorCmd.AddCommand(EnrollFactorCmd)
}

var (
	ListFactorsuserId string

	ListFactorsBackup bool
)

func NewListFactorsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listFactors",
		Long: "List all enrolled Factors",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.UserFactorAPI.ListFactors(apiClient.GetConfig().Context, ListFactorsuserId)

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
			if ListFactorsBackup {

				idParam := ListFactorsuserId
				err := utils.BackupObject(d, "UserFactor", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListFactorsuserId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	cmd.Flags().BoolVarP(&ListFactorsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListFactorsCmd := NewListFactorsCmd()
	UserFactorCmd.AddCommand(ListFactorsCmd)
}

var (
	ListSupportedFactorsuserId string

	ListSupportedFactorsBackup bool
)

func NewListSupportedFactorsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listSupportedFactors",
		Long: "List all supported Factors",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.UserFactorAPI.ListSupportedFactors(apiClient.GetConfig().Context, ListSupportedFactorsuserId)

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
			if ListSupportedFactorsBackup {

				idParam := ListSupportedFactorsuserId
				err := utils.BackupObject(d, "UserFactor", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListSupportedFactorsuserId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	cmd.Flags().BoolVarP(&ListSupportedFactorsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListSupportedFactorsCmd := NewListSupportedFactorsCmd()
	UserFactorCmd.AddCommand(ListSupportedFactorsCmd)
}

var (
	ListSupportedSecurityQuestionsuserId string

	ListSupportedSecurityQuestionsBackup bool
)

func NewListSupportedSecurityQuestionsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "listSupportedSecurityQuestions",
		Long: "List all supported Security Questions",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.UserFactorAPI.ListSupportedSecurityQuestions(apiClient.GetConfig().Context, ListSupportedSecurityQuestionsuserId)

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
			if ListSupportedSecurityQuestionsBackup {

				idParam := ListSupportedSecurityQuestionsuserId
				err := utils.BackupObject(d, "UserFactor", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListSupportedSecurityQuestionsuserId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	cmd.Flags().BoolVarP(&ListSupportedSecurityQuestionsBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ListSupportedSecurityQuestionsCmd := NewListSupportedSecurityQuestionsCmd()
	UserFactorCmd.AddCommand(ListSupportedSecurityQuestionsCmd)
}

var (
	GetFactoruserId string

	GetFactorfactorId string

	GetFactorBackup bool
)

func NewGetFactorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getFactor",
		Long: "Retrieve a Factor",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.UserFactorAPI.GetFactor(apiClient.GetConfig().Context, GetFactoruserId, GetFactorfactorId)

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
			if GetFactorBackup {

				idParam := GetFactoruserId
				err := utils.BackupObject(d, "UserFactor", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetFactoruserId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	cmd.Flags().StringVarP(&GetFactorfactorId, "factorId", "", "", "")
	cmd.MarkFlagRequired("factorId")

	cmd.Flags().BoolVarP(&GetFactorBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetFactorCmd := NewGetFactorCmd()
	UserFactorCmd.AddCommand(GetFactorCmd)
}

var (
	UnenrollFactoruserId string

	UnenrollFactorfactorId string

	UnenrollFactorBackup bool
)

func NewUnenrollFactorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "unenrollFactor",
		Long: "Unenroll a Factor",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.UserFactorAPI.UnenrollFactor(apiClient.GetConfig().Context, UnenrollFactoruserId, UnenrollFactorfactorId)

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
			if UnenrollFactorBackup {

				idParam := UnenrollFactoruserId
				err := utils.BackupObject(d, "UserFactor", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&UnenrollFactoruserId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	cmd.Flags().StringVarP(&UnenrollFactorfactorId, "factorId", "", "", "")
	cmd.MarkFlagRequired("factorId")

	cmd.Flags().BoolVarP(&UnenrollFactorBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	UnenrollFactorCmd := NewUnenrollFactorCmd()
	UserFactorCmd.AddCommand(UnenrollFactorCmd)
}

var (
	ActivateFactoruserId string

	ActivateFactorfactorId string

	ActivateFactordata string

	ActivateFactorBackup bool
)

func NewActivateFactorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "activateFactor",
		Long: "Activate a Factor",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.UserFactorAPI.ActivateFactor(apiClient.GetConfig().Context, ActivateFactoruserId, ActivateFactorfactorId)

			if ActivateFactordata != "" {
				req = req.Data(ActivateFactordata)
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
			if ActivateFactorBackup {

				idParam := ActivateFactoruserId
				err := utils.BackupObject(d, "UserFactor", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ActivateFactoruserId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	cmd.Flags().StringVarP(&ActivateFactorfactorId, "factorId", "", "", "")
	cmd.MarkFlagRequired("factorId")

	cmd.Flags().StringVarP(&ActivateFactordata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ActivateFactorBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ActivateFactorCmd := NewActivateFactorCmd()
	UserFactorCmd.AddCommand(ActivateFactorCmd)
}

var (
	ResendEnrollFactoruserId string

	ResendEnrollFactorfactorId string

	ResendEnrollFactordata string

	ResendEnrollFactorBackup bool
)

func NewResendEnrollFactorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "resendEnrollFactor",
		Long: "Resend a Factor enrollment",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.UserFactorAPI.ResendEnrollFactor(apiClient.GetConfig().Context, ResendEnrollFactoruserId, ResendEnrollFactorfactorId)

			if ResendEnrollFactordata != "" {
				req = req.Data(ResendEnrollFactordata)
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
			if ResendEnrollFactorBackup {

				idParam := ResendEnrollFactoruserId
				err := utils.BackupObject(d, "UserFactor", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&ResendEnrollFactoruserId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	cmd.Flags().StringVarP(&ResendEnrollFactorfactorId, "factorId", "", "", "")
	cmd.MarkFlagRequired("factorId")

	cmd.Flags().StringVarP(&ResendEnrollFactordata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&ResendEnrollFactorBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	ResendEnrollFactorCmd := NewResendEnrollFactorCmd()
	UserFactorCmd.AddCommand(ResendEnrollFactorCmd)
}

var (
	GetFactorTransactionStatususerId string

	GetFactorTransactionStatusfactorId string

	GetFactorTransactionStatustransactionId string

	GetFactorTransactionStatusBackup bool
)

func NewGetFactorTransactionStatusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "getFactorTransactionStatus",
		Long: "Retrieve a Factor transaction status",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.UserFactorAPI.GetFactorTransactionStatus(apiClient.GetConfig().Context, GetFactorTransactionStatususerId, GetFactorTransactionStatusfactorId, GetFactorTransactionStatustransactionId)

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
			if GetFactorTransactionStatusBackup {

				idParam := GetFactorTransactionStatususerId
				err := utils.BackupObject(d, "UserFactor", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetFactorTransactionStatususerId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	cmd.Flags().StringVarP(&GetFactorTransactionStatusfactorId, "factorId", "", "", "")
	cmd.MarkFlagRequired("factorId")

	cmd.Flags().StringVarP(&GetFactorTransactionStatustransactionId, "transactionId", "", "", "")
	cmd.MarkFlagRequired("transactionId")

	cmd.Flags().BoolVarP(&GetFactorTransactionStatusBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	GetFactorTransactionStatusCmd := NewGetFactorTransactionStatusCmd()
	UserFactorCmd.AddCommand(GetFactorTransactionStatusCmd)
}

var (
	VerifyFactoruserId string

	VerifyFactorfactorId string

	VerifyFactordata string

	VerifyFactorBackup bool
)

func NewVerifyFactorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "verifyFactor",
		Long: "Verify a Factor",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.UserFactorAPI.VerifyFactor(apiClient.GetConfig().Context, VerifyFactoruserId, VerifyFactorfactorId)

			if VerifyFactordata != "" {
				req = req.Data(VerifyFactordata)
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
			if VerifyFactorBackup {

				idParam := VerifyFactoruserId
				err := utils.BackupObject(d, "UserFactor", idParam)
				if err != nil {
					return err
				}
			}
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	cmd.Flags().StringVarP(&VerifyFactoruserId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	cmd.Flags().StringVarP(&VerifyFactorfactorId, "factorId", "", "", "")
	cmd.MarkFlagRequired("factorId")

	cmd.Flags().StringVarP(&VerifyFactordata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	cmd.Flags().BoolVarP(&VerifyFactorBackup, "backup", "b", false, "Backup the object to a file")

	return cmd
}

func init() {
	VerifyFactorCmd := NewVerifyFactorCmd()
	UserFactorCmd.AddCommand(VerifyFactorCmd)
}
