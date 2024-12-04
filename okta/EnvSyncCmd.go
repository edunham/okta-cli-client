package okta

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var EnvSyncCmd = &cobra.Command{
	Use:  "envsync",
	Long: "backup and restore okta environments",
}

var UserdataPath string

type UserProfile struct {
	Email       string `json:"email"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Login       string `json:"login"`
	MobilePhone string `json:"mobilePhone"`
	SecondEmail string `json:"secondEmail"`
}

type UserData struct {
	Profile UserProfile `json:"profile"`
}

func init() {
	rootCmd.AddCommand(EnvSyncCmd)
}

type User struct {
	Profile struct {
		Login string `json:"login"`
	} `json:"profile"`
}

func NewEnvSyncPullUserCmd() *cobra.Command {
	var GetUseruserId string

	cmd := &cobra.Command{
		Use: "pullUser",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.UserAPI.GetUser(apiClient.GetConfig().Context, GetUseruserId)
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

			user := &User{}

			err = json.Unmarshal(d, user)
			if err != nil {
				return err
			}

			homePath := os.Getenv("USERPROFILE") // Default for Windows
			if homePath == "" {
				homePath = os.Getenv("HOME") // Used in Unix-like systems
			}

			oktaDomain := apiClient.GetConfig().Host // Adjust this line to get the Okta domain correctly
			filePath := fmt.Sprintf("%s/.okta/%s/users/%s.json", homePath, oktaDomain, user.Profile.Login)
			cmd.Println(filePath)
			dirPath := filepath.Dir(filePath)
			err = os.MkdirAll(dirPath, 0755)
			if err != nil {
				return err
			}
			//END OF PART THAT NEEDS TO MOVE OUT
			err = os.WriteFile(filePath, d, 0644)
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&GetUseruserId, "userId", "", "", "ID of the user to fetch")
	cmd.MarkFlagRequired("userId")

	return cmd
}

func init() {
	EnvSyncPullUserCmd := NewEnvSyncPullUserCmd()
	EnvSyncCmd.AddCommand(EnvSyncPullUserCmd)
}

func NewEnvSyncPushUserCmd() *cobra.Command {
	createCmd := &cobra.Command{
		Use: "pushUser",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Read and parse the userdata file
			userData, err := os.ReadFile(UserdataPath)
			if err != nil {
				return fmt.Errorf("error reading user data file: %w", err)
			}

			var data UserData
			if err = json.Unmarshal(userData, &data); err != nil {
				return fmt.Errorf("error parsing user data file: %w", err)
			}

			// Create the API request data
			createData := fmt.Sprintf(
				`{"credentials":{"password":{"value":"Hell4W0rld"}},"profile":{"email":"%s","firstName":"%s","lastName":"%s","login":"%s"}}`,
				data.Profile.Email, data.Profile.FirstName, data.Profile.LastName, data.Profile.Login,
			)

			// Make the API request
			req := apiClient.UserAPI.CreateUser(apiClient.GetConfig().Context)
			req = req.Data(createData)

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
			utils.PrettyPrintByte(d)
			return nil
		},
	}

	createCmd.Flags().StringVarP(&UserdataPath, "userdata", "u", "", "Path to the userdata file")
	createCmd.MarkFlagRequired("userdata")

	return createCmd
}

func init() {
	EnvSyncPushUserCmd := NewEnvSyncPushUserCmd()
	EnvSyncCmd.AddCommand(EnvSyncPushUserCmd)
}

type SyncUser struct {
    ID string `json:"id"`
}

func ProcessUserIDs(v []byte) error {
    var users []SyncUser
    err := json.Unmarshal(v, &users)
    if err != nil {
        return err
    }

    cmd := NewEnvSyncPullUserCmd()
    for _, user := range users {
        if err := cmd.Flags().Set("userId", user.ID); err != nil {
            return fmt.Errorf("failed to set userId flag for %s: %w", user.ID, err)
        }
        
        if err := cmd.RunE(cmd, []string{}); err != nil {
            return fmt.Errorf("failed to process user %s: %w", user.ID, err)
        }
    }
    return nil
}
func NewPullAllUsersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "pullAllUsers",

		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.UserAPI.ListUsers(apiClient.GetConfig().Context)

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
			ProcessUserIDs(d)
			return nil
		},
	}

	return cmd
}

func init() {
	PullAllUsersCmd := NewPullAllUsersCmd()
	EnvSyncCmd.AddCommand(PullAllUsersCmd)
}
