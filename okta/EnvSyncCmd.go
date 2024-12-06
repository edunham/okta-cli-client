package okta

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

/***************************************************************/
/* COMMON functions                                            */
/***************************************************************/

// Define ANSI color codes
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

func backupSchemaInFile(filePath string, jsonData []byte) error {
	// Create directory structure if it doesn't exist
	dirPath := filepath.Dir(filePath)
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		return err // Return the error if directory creation fails
	}

	// Write the JSON data to the specified file
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return err // Return the error if file writing fails
	}

	return nil // Return nil if everything succeeds
}

func getHomePath() string {
	homePath := os.Getenv("USERPROFILE") // Default for Windows
	if homePath == "" {
		homePath = os.Getenv("HOME") // Used in Unix-like systems
	}

	return homePath
}

var EnvSyncCmd = &cobra.Command{
	Use:  "envsync",
	Long: "backup and restore okta environments",
}

func init() {
	rootCmd.AddCommand(EnvSyncCmd)
}

/***************************************************************/
/* END COMMON functions                                        */
/***************************************************************/

/***************************************************************/
/* DEMO commands                                               */
/***************************************************************/

// PULL
var PullSchemas []string

func NewPullCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pull",
		Short: "A command that pulls selected schemas into your local system",
		RunE: func(cmd *cobra.Command, args []string) error {

			// Iterate over the list of values and perform a switch
			for _, value := range PullSchemas {
				switch value {
				case "user":
					fmt.Println("[INFO] Pulling users")
					pullAllUsersCmd := NewPullAllUsersCmd()

					if err := pullAllUsersCmd.RunE(pullAllUsersCmd, []string{}); err != nil {
						return fmt.Errorf(Red+"[ERROR]"+Reset+"failed to pull users %w", err)
					} else {
						fmt.Println(Green + "[SUCCESS] Users were backed up successfully" + Reset)
					}
				case "group":
					fmt.Println("[INFO] Pulling groups")
					pullAllGroupsCmd := NewPullAllGroupsCmd()

					if err := pullAllGroupsCmd.RunE(pullAllGroupsCmd, []string{}); err != nil {
						return fmt.Errorf(Red+"failed to pull groups %w"+Reset, err)
					} else {
						fmt.Println(Green + "[SUCCESS] Groups were backed up successfully" + Reset)
					}
				default:
					fmt.Println(Red+"[ERROR]"+Reset+"Unsupported schema:", value)
				}
			}

			return nil
		},
	}

	// Define the --schemas flag to accept a comma-separated list of strings
	cmd.Flags().StringSliceVar(&PullSchemas, "schemas", []string{}, "Comma-separated list of schemas")
	cmd.MarkFlagRequired("schemas")

	return cmd
}

func init() {
	PullCmd := NewPullCmd()
	EnvSyncCmd.AddCommand(PullCmd)
}

// PUSH
var PushSchemas []string
var BackupFolderPath string

func NewPushCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "push",
		Short: "A command that pushes your Okta resources data from your local system to your Okta org.",
		RunE: func(cmd *cobra.Command, args []string) error {

			// Iterate over the list of values and perform a switch
			for _, value := range PushSchemas {
				switch value {
				case "user":
					fmt.Println("[INFO] Pushing users")
					pushAllUsersCmd := NewPushAllUsersCmd()

					// Set the dir flag for this file
					if err := pushAllUsersCmd.Flags().Set("dir", BackupFolderPath); err != nil {
						return fmt.Errorf(Red+"[ERROR]"+Reset+"failed to set dir for %s: %w", BackupFolderPath, err)
					}

					if err := pushAllUsersCmd.RunE(pushAllUsersCmd, []string{}); err != nil {
						return fmt.Errorf(Red+"[ERROR]"+Reset+"failed to push users %w", err)
					} else {
						fmt.Println(Green + "[SUCCESS] Users were restored successfully" + Reset)
					}
				case "group":
					fmt.Println("[INFO] Pushing groups")
					pushAllGroupsCmd := NewPushAllGroupsCmd()

					// Set the dir flag for this file
					if err := pushAllGroupsCmd.Flags().Set("dir", BackupFolderPath); err != nil {
						return fmt.Errorf(Red+"[ERROR]"+Reset+"failed to set dir for %s: %w", BackupFolderPath, err)
					}

					if err := pushAllGroupsCmd.RunE(pushAllGroupsCmd, []string{}); err != nil {
						return fmt.Errorf(Red+"[ERROR]"+Reset+"failed to push groups %w", err)
					} else {
						fmt.Println(Green + "[SUCCESS] Groups were restored successfully" + Reset)
					}
				default:
					fmt.Println(Red+"[ERROR]"+Reset+"Unsupported schema:", value)
				}
			}

			return nil
		},
	}

	// Define the --schemas flag to accept a comma-separated list of strings
	cmd.Flags().StringSliceVar(&PushSchemas, "schemas", []string{}, "Comma-separated list of schemas")
	cmd.Flags().StringVarP(&BackupFolderPath, "dir", "", "", "The path to the folder where your backup files are stored")
	cmd.MarkFlagRequired("schemas")
	cmd.MarkFlagRequired("dir")

	return cmd
}

func init() {
	PushCmd := NewPushCmd()
	EnvSyncCmd.AddCommand(PushCmd)
}

/***************************************************************/
/* END DEMO commands                                           */
/***************************************************************/

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

type User struct {
	Profile struct {
		Login string `json:"login"`
	} `json:"profile"`
}

func NewEnvSyncCleanCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "clean",
		RunE: func(cmd *cobra.Command, args []string) error {

			homePath := os.Getenv("USERPROFILE") // Default for Windows
			if homePath == "" {
				homePath = os.Getenv("HOME") // Used in Unix-like systems
			}

			oktaDomain := apiClient.GetConfig().Host
			dirPath := fmt.Sprintf("%s/.okta/%s/", homePath, oktaDomain)
			err := os.RemoveAll(dirPath)
			if err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}

func init() {
	EnvSyncCleanCmd := NewEnvSyncCleanCmd()
	EnvSyncCmd.AddCommand(EnvSyncCleanCmd)
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

var PushFromDir string

func NewPushAllUsersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "pushAllUsers",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Construct the full path to the users directory
			usersDir := filepath.Join(PushFromDir, "users")

			// Read all files in the users directory
			files, err := os.ReadDir(usersDir)
			if err != nil {
				return fmt.Errorf("failed to read users directory: %w", err)
			}

			// Get the single-user push command
			pushUserCmd := NewEnvSyncPushUserCmd()

			// Process each .json file
			for _, file := range files {
				if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
					fullPath := filepath.Join(usersDir, file.Name())

					// Set the UserdataPath flag for this file
					if err := pushUserCmd.Flags().Set("userdata", fullPath); err != nil {
						return fmt.Errorf("failed to set dir for %s: %w", file.Name(), err)
					}

					// Execute the push command for this user
					if err := pushUserCmd.RunE(pushUserCmd, []string{}); err != nil {
						return fmt.Errorf("failed to push user %s: %w", file.Name(), err)
					}
				}
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&PushFromDir, "dir", "", "", "")
	cmd.MarkFlagRequired("dir")

	return cmd
}

func init() {
	PushAllUsersCmd := NewPushAllUsersCmd()
	EnvSyncCmd.AddCommand(PushAllUsersCmd)
}

///////////////// group stuff

var SyncGroupdata string
var skipGroups = map[string]bool{
	"Okta Administrators.json": true,
	"Everyone.json": true,
}

func NewEnvSyncPushGroupCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use: "pushGroup",

		RunE: func(cmd *cobra.Command, args []string) error {
			baseFileName := filepath.Base(SyncGroupdata)
			if skipGroups[baseFileName] {
				fmt.Printf(Yellow+"[SKIP]"+Reset+"Skipping file: %s\n", baseFileName)
				return nil
			}

			// Read and parse the userdata file
			groupBackupData, err := os.ReadFile(SyncGroupdata)

			if err != nil {
				return fmt.Errorf(Red+"[ERROR]"+Reset+"error reading group data file: %w", err)
			}

			var group Group

			if err = json.Unmarshal(groupBackupData, &group); err != nil {
				return fmt.Errorf(Red+"[ERROR]"+Reset+"error parsing group data file: %w", err)
			}

			req := apiClient.GroupAPI.CreateGroup(apiClient.GetConfig().Context)
			createData := string(groupBackupData)
			if createData != "" {
				req = req.Data(createData)
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

			_, err = io.ReadAll(resp.Body)

			if err != nil {
				return err
			}

			fmt.Println(Green+"[SUCCESS]"+Reset+"Group restored successfully:", group.Profile.Name)

			//FIXME
			//utils.PrettyPrintByte(d)
			// cmd.Println(string(d))
			return nil
		},
	}

	cmd.Flags().StringVarP(&SyncGroupdata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	return cmd
}

func init() {
	EnvSyncPushGroupCmd := NewEnvSyncPushGroupCmd()
	EnvSyncCmd.AddCommand(EnvSyncPushGroupCmd)
}

var SyncGroupgroupId string

type Group struct {
	Type    string `json:"type"`
	Profile struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"profile"`
}

func NewEnvSyncPullGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "pullGroup",

		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.GroupAPI.GetGroup(apiClient.GetConfig().Context, SyncGroupgroupId)

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

			group := &Group{}

			err = json.Unmarshal(d, group)
			if err != nil {
				return err
			}

			oktaDomain := apiClient.GetConfig().Host // Adjust this line to get the Okta domain correctly
			filePath := fmt.Sprintf("%s/.okta/%s/groups/%s.json", getHomePath(), oktaDomain, group.Profile.Name)

			cmd.Println(filePath)

			groupJsonData, err := json.Marshal(group)

			if err != nil {
				return err
			}

			err = backupSchemaInFile(filePath, groupJsonData)

			if err != nil {
				return err
			} else {
				fmt.Println(Green+"[SUCCESS] Group was backed up successfully: %s"+Reset, group.Profile.Name)
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&SyncGroupgroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	return cmd
}

func init() {
	EnvSyncPullGroupCmd := NewEnvSyncPullGroupCmd()
	EnvSyncCmd.AddCommand(EnvSyncPullGroupCmd)
}

func NewPullAllGroupsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "pullAllGroups",

		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.GroupAPI.ListGroups(apiClient.GetConfig().Context)

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

			var groups []Group

			// Unmarshal the JSON byte slice into the slice of structs
			err = json.Unmarshal(d, &groups)
			if err != nil {
				log.Fatal(err)
				return err
			}

			oktaDomain := apiClient.GetConfig().Host // Adjust this line to get the Okta domain correctly

			// Backup groups
			for _, group := range groups {
				filePath := fmt.Sprintf("%s/.okta/%s/groups/%s.json", getHomePath(), oktaDomain, group.Profile.Name)
				groupJsonData, err := json.Marshal(group)

				if err != nil {
					log.Fatal(err)
				}

				err = backupSchemaInFile(filePath, groupJsonData)

				if err != nil {
					log.Fatal(err)
				} else {
					fmt.Println(Green+"[SUCCESS] "+Reset+"Group was backed up successfully:", group.Profile.Name)
				}
			}

			return nil
		},
	}

	return cmd
}

func init() {
	PullAllGroupsCmd := NewPullAllGroupsCmd()
	EnvSyncCmd.AddCommand(PullAllGroupsCmd)
}

var GroupsPushFromDir string

func NewPushAllGroupsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "pushAllGroups",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Construct the full path to the users directory
			groupsDir := filepath.Join(GroupsPushFromDir, "groups")

			// Read all files in the users directory
			files, err := os.ReadDir(groupsDir)
			if err != nil {
				return fmt.Errorf("failed to read groups directory: %w", err)
			}

			// Get the single-group push command
			pushGroupCmd := NewEnvSyncPushGroupCmd()

			// Process each .json file
			for _, file := range files {
				if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
					fullPath := filepath.Join(groupsDir, file.Name())

					// Set the UserdataPath flag for this file
					if err := pushGroupCmd.Flags().Set("data", fullPath); err != nil {
						return fmt.Errorf("failed to set data for %s: %w", file.Name(), err)
					}

					// Execute the push command for this user
					if err := pushGroupCmd.RunE(pushGroupCmd, []string{}); err != nil {
						return fmt.Errorf("failed to push group %s: %w", file.Name(), err)
					}
				}
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&GroupsPushFromDir, "dir", "", "", "")
	cmd.MarkFlagRequired("dir")

	return cmd
}

func init() {
	PushAllGroupsCmd := NewPushAllGroupsCmd()
	EnvSyncCmd.AddCommand(PushAllGroupsCmd)
}

///////////////////// group membership stuff

var ListGroupUsersSyncgroupId string

func NewPullGroupUsersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "pullGroupUsers",

		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.GroupAPI.ListGroupUsers(apiClient.GetConfig().Context, ListGroupUsersSyncgroupId)
			rxq := apiClient.GroupAPI.GetGroup(apiClient.GetConfig().Context, ListGroupUsersSyncgroupId)

			rxsp, err := rxq.Execute()
			if err != nil {
				if rxsp != nil && rxsp.Body != nil {
					b, err := io.ReadAll(rxsp.Body)
					if err == nil {
						utils.PrettyPrintByte(b)
					}
				}
				return err
			}
			b, err := io.ReadAll(rxsp.Body)
			if err != nil {
				return err
			}
			group := &Group{}

			err = json.Unmarshal(b, group)
			if err != nil {
				return err
			}

			homePath := os.Getenv("USERPROFILE") // Default for Windows
			if homePath == "" {
				homePath = os.Getenv("HOME") // Used in Unix-like systems
			}

			oktaDomain := apiClient.GetConfig().Host // Adjust this line to get the Okta domain correctly
			filePath := fmt.Sprintf("%s/.okta/%s/groups_users/%s.json", homePath, oktaDomain, group.Profile.Name)
			cmd.Println(filePath)
			dirPath := filepath.Dir(filePath)
			err = os.MkdirAll(dirPath, 0755)
			if err != nil {
				return err
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
			err = os.WriteFile(filePath, d, 0644)
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&ListGroupUsersSyncgroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	return cmd
}

func init() {
	PullGroupUsersCmd := NewPullGroupUsersCmd()
	EnvSyncCmd.AddCommand(PullGroupUsersCmd)
}

var (
	AssignUserToGroupSyncgroupId string

	AssignUserToSyncGroupuserId string
)

func NewSyncUserToGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "assignUserTo",

		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.GroupAPI.AssignUserToGroup(apiClient.GetConfig().Context, AssignUserToGroupSyncgroupId, AssignUserToSyncGroupuserId)

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
			// cmd.Println(string(d))
			return nil
		},
	}

	cmd.Flags().StringVarP(&AssignUserToGroupSyncgroupId, "groupId", "", "", "")
	cmd.MarkFlagRequired("groupId")

	cmd.Flags().StringVarP(&AssignUserToSyncGroupuserId, "userId", "", "", "")
	cmd.MarkFlagRequired("userId")

	return cmd
}

func init() {
	SyncUserToGroupCmd := NewSyncUserToGroupCmd()
	GroupCmd.AddCommand(SyncUserToGroupCmd)
}

///////////////////// policy stubs

var SyncPolicydata string

func NewSyncPolicyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "pushPolicy",

		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PolicyAPI.CreatePolicy(apiClient.GetConfig().Context)

			if SyncPolicydata != "" {
				req = req.Data(SyncPolicydata)
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
			utils.PrettyPrintByte(d)
			// cmd.Println(string(d))
			return nil
		},
	}

	cmd.Flags().StringVarP(&SyncPolicydata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	return cmd
}

func init() {
	SyncPolicyCmd := NewSyncPolicyCmd()
	EnvSyncCmd.AddCommand(SyncPolicyCmd)
}

func NewGetAllPoliciesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "pullAllPolicies",

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
			utils.PrettyPrintByte(d)
			// cmd.Println(string(d))
			return nil
		},
	}

	return cmd
}

func init() {
	GetAllPoliciesCmd := NewGetAllPoliciesCmd()
	EnvSyncCmd.AddCommand(GetAllPoliciesCmd)
}

var GetSyncpolicyId string

func NewGetSyncPolicyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "pullPolicy",

		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PolicyAPI.GetPolicy(apiClient.GetConfig().Context, GetSyncpolicyId)

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
			// cmd.Println(string(d))
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetSyncpolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	return cmd
}

func init() {
	GetSyncPolicyCmd := NewGetSyncPolicyCmd()
	EnvSyncCmd.AddCommand(GetSyncPolicyCmd)
}

///////////////////// policy mapping stubs

var (
	SyncResourceToPolicypolicyId string

	SyncResourceToPolicydata string
)

func NewSyncResourceToPolicyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "pushResourcePolicy",

		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PolicyAPI.MapResourceToPolicy(apiClient.GetConfig().Context, SyncResourceToPolicypolicyId)

			if SyncResourceToPolicydata != "" {
				req = req.Data(SyncResourceToPolicydata)
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
			utils.PrettyPrintByte(d)
			// cmd.Println(string(d))
			return nil
		},
	}

	cmd.Flags().StringVarP(&SyncResourceToPolicypolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().StringVarP(&SyncResourceToPolicydata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	return cmd
}

func init() {
	SyncResourceToPolicyCmd := NewSyncResourceToPolicyCmd()
	EnvSyncCmd.AddCommand(SyncResourceToPolicyCmd)
}

var SyncPolicyMappingspolicyId string

func NewSyncPolicyMappingsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "pullAllMappings",

		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PolicyAPI.ListPolicyMappings(apiClient.GetConfig().Context, SyncPolicyMappingspolicyId)

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
			// cmd.Println(string(d))
			return nil
		},
	}

	cmd.Flags().StringVarP(&SyncPolicyMappingspolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	return cmd
}

func init() {
	SyncPolicyMappingsCmd := NewSyncPolicyMappingsCmd()
	EnvSyncCmd.AddCommand(SyncPolicyMappingsCmd)
}

var (
	GetSyncPolicyMappingpolicyId string

	GetSyncPolicyMappingmappingId string
)

func NewGetSyncPolicyMappingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "pullMapping",

		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PolicyAPI.GetPolicyMapping(apiClient.GetConfig().Context, GetSyncPolicyMappingpolicyId, GetSyncPolicyMappingmappingId)

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
			// cmd.Println(string(d))
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetSyncPolicyMappingpolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().StringVarP(&GetSyncPolicyMappingmappingId, "mappingId", "", "", "")
	cmd.MarkFlagRequired("mappingId")

	return cmd
}

func init() {
	GetSyncPolicyMappingCmd := NewGetSyncPolicyMappingCmd()
	EnvSyncCmd.AddCommand(GetSyncPolicyMappingCmd)
}

///////////////////// policy rule stubs

var (
	SyncPolicyRulepolicyId string

	SyncPolicyRuledata string
)

func NewSyncPolicyRuleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "pushRule",

		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PolicyAPI.CreatePolicyRule(apiClient.GetConfig().Context, SyncPolicyRulepolicyId)

			if SyncPolicyRuledata != "" {
				req = req.Data(SyncPolicyRuledata)
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
			utils.PrettyPrintByte(d)
			// cmd.Println(string(d))
			return nil
		},
	}

	cmd.Flags().StringVarP(&SyncPolicyRulepolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().StringVarP(&SyncPolicyRuledata, "data", "", "", "")
	cmd.MarkFlagRequired("data")

	return cmd
}

func init() {
	SyncPolicyRuleCmd := NewSyncPolicyRuleCmd()
	EnvSyncCmd.AddCommand(SyncPolicyRuleCmd)
}

var GetAllPolicyRulespolicyId string

func NewGetAllPolicyRulesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "pullAllRules",

		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PolicyAPI.ListPolicyRules(apiClient.GetConfig().Context, GetAllPolicyRulespolicyId)

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
			// cmd.Println(string(d))
			return nil
		},
	}

	cmd.Flags().StringVarP(&GetAllPolicyRulespolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	return cmd
}

func init() {
	GetAllPolicyRulesCmd := NewGetAllPolicyRulesCmd()
	EnvSyncCmd.AddCommand(GetAllPolicyRulesCmd)
}

var (
	PullPolicyRulepolicyId string

	PullPolicyRuleruleId string
)

func NewPullPolicyRuleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "pullRule",

		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.PolicyAPI.GetPolicyRule(apiClient.GetConfig().Context, PullPolicyRulepolicyId, PullPolicyRuleruleId)

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
			// cmd.Println(string(d))
			return nil
		},
	}

	cmd.Flags().StringVarP(&PullPolicyRulepolicyId, "policyId", "", "", "")
	cmd.MarkFlagRequired("policyId")

	cmd.Flags().StringVarP(&PullPolicyRuleruleId, "ruleId", "", "", "")
	cmd.MarkFlagRequired("ruleId")

	return cmd
}

func init() {
	PullPolicyRuleCmd := NewPullPolicyRuleCmd()
	EnvSyncCmd.AddCommand(PullPolicyRuleCmd)
}
