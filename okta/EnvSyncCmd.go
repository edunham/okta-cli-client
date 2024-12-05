package okta

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

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
					if err := pushUserCmd.Flags().Set("UserdataPath", fullPath); err != nil {
						return fmt.Errorf("failed to set UserdataPath for %s: %w", file.Name(), err)
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

func NewEnvSyncPushGroupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create",

		RunE: func(cmd *cobra.Command, args []string) error {
			req := apiClient.GroupAPI.CreateGroup(apiClient.GetConfig().Context)

			if SyncGroupdata != "" {
				req = req.Data(SyncGroupdata)
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
			//FIXME
			utils.PrettyPrintByte(d)
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
			//utils.PrettyPrintByte(d)
			// cmd.Println(string(d))
			group := &Group{}

			err = json.Unmarshal(d, group)
			if err != nil {
				return err
			}

			homePath := os.Getenv("USERPROFILE") // Default for Windows
			if homePath == "" {
				homePath = os.Getenv("HOME") // Used in Unix-like systems
			}

			oktaDomain := apiClient.GetConfig().Host // Adjust this line to get the Okta domain correctly
			filePath := fmt.Sprintf("%s/.okta/%s/groups/%s.json", homePath, oktaDomain, group.Profile.Name)
			cmd.Println(filePath)
			dirPath := filepath.Dir(filePath)
			err = os.MkdirAll(dirPath, 0755)
			if err != nil {
				return err
			}

			groupJsonData, err := json.Marshal(group)

			if err != nil {
				return err
			}

			err = os.WriteFile(filePath, groupJsonData, 0644)
			if err != nil {
				return err
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
			//FIXME
			utils.PrettyPrintByte(d)
			// cmd.Println(string(d))
			return nil
		},
	}

	return cmd
}

// func ProcessGroups(v []byte) error {
// 	var groups []Group
// 	err := json.Unmarshal(v, &groups)
// 	if err != nil {
// 		return err
// 	}

// 	for _, group := range groups {

// 	}
// 	return nil
// }

// func backupSchemaInFile(){

// }

func init() {
	PullAllGroupsCmd := NewPullAllGroupsCmd()
	EnvSyncCmd.AddCommand(PullAllGroupsCmd)
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
		Use: "create",

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
		Use: "listPolicies",

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
		Use: "get",

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
		Use: "mapResourceTo",

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
		Use: "listMappings",

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
		Use: "getMapping",

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
		Use: "createRule",

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
		Use: "listRules",

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
		Use: "getRule",

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
