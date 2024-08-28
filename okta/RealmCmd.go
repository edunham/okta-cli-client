package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var RealmCmd = &cobra.Command{
	Use:   "realm",
	Long:  "Manage RealmAPI",
}

func NewRealmCmd() *cobra.Command {
    cmd := &cobra.Command{
		Use:   "realm",
		Long:  "Manage RealmAPI",
	}
	return cmd
}

func init() {
    rootCmd.AddCommand(RealmCmd)
}

var (
    
    
            CreateRealmdata string
        
    
)

func NewCreateRealmCmd() *cobra.Command {
    cmd := &cobra.Command{
	    Use:   "create",
	  
        RunE: func(cmd *cobra.Command, args []string) error {
            
            
            
            req := apiClient.RealmAPI.CreateRealm(apiClient.GetConfig().Context)
            
            
            if CreateRealmdata != "" {
                req = req.Data(CreateRealmdata)
            }
            
            resp, err := req.Execute()
            if err != nil {
                return err
            }
		    d, err := io.ReadAll(resp.Body)
		    if err != nil {
			    return err
		    }
		    utils.PrettyPrintByte(d)
            cmd.Println(string(d))
		    return nil
        },
    }

    
    
        cmd.Flags().StringVarP(&CreateRealmdata, "data", "", "", "")
        cmd.MarkFlagRequired("data")
        
    

	return cmd
}

func init() {
	CreateRealmCmd := NewCreateRealmCmd()
    RealmCmd.AddCommand(CreateRealmCmd)
}

var (
    
    
    
)

func NewListRealmsCmd() *cobra.Command {
    cmd := &cobra.Command{
	    Use:   "lists",
	  
        RunE: func(cmd *cobra.Command, args []string) error {
            
            
            
            req := apiClient.RealmAPI.ListRealms(apiClient.GetConfig().Context)
            
            
            
            resp, err := req.Execute()
            if err != nil {
                return err
            }
		    d, err := io.ReadAll(resp.Body)
		    if err != nil {
			    return err
		    }
		    utils.PrettyPrintByte(d)
            cmd.Println(string(d))
		    return nil
        },
    }

    
    
    

	return cmd
}

func init() {
	ListRealmsCmd := NewListRealmsCmd()
    RealmCmd.AddCommand(ListRealmsCmd)
}

var (
    
    
            GetRealmrealmId string
        
    
)

func NewGetRealmCmd() *cobra.Command {
    cmd := &cobra.Command{
	    Use:   "get",
	  
        RunE: func(cmd *cobra.Command, args []string) error {
            
            
            
            req := apiClient.RealmAPI.GetRealm(apiClient.GetConfig().Context, GetRealmrealmId)
            
            
            
            resp, err := req.Execute()
            if err != nil {
                return err
            }
		    d, err := io.ReadAll(resp.Body)
		    if err != nil {
			    return err
		    }
		    utils.PrettyPrintByte(d)
            cmd.Println(string(d))
		    return nil
        },
    }

    
    
        cmd.Flags().StringVarP(&GetRealmrealmId, "realmId", "", "", "")
        cmd.MarkFlagRequired("realmId")
        
    

	return cmd
}

func init() {
	GetRealmCmd := NewGetRealmCmd()
    RealmCmd.AddCommand(GetRealmCmd)
}

var (
    
    
            ReplaceRealmrealmId string
        
            ReplaceRealmdata string
        
    
)

func NewReplaceRealmCmd() *cobra.Command {
    cmd := &cobra.Command{
	    Use:   "replace",
	  
        RunE: func(cmd *cobra.Command, args []string) error {
            
            
            
            req := apiClient.RealmAPI.ReplaceRealm(apiClient.GetConfig().Context, ReplaceRealmrealmId)
            
            
            if ReplaceRealmdata != "" {
                req = req.Data(ReplaceRealmdata)
            }
            
            resp, err := req.Execute()
            if err != nil {
                return err
            }
		    d, err := io.ReadAll(resp.Body)
		    if err != nil {
			    return err
		    }
		    utils.PrettyPrintByte(d)
            cmd.Println(string(d))
		    return nil
        },
    }

    
    
        cmd.Flags().StringVarP(&ReplaceRealmrealmId, "realmId", "", "", "")
        cmd.MarkFlagRequired("realmId")
        
        cmd.Flags().StringVarP(&ReplaceRealmdata, "data", "", "", "")
        cmd.MarkFlagRequired("data")
        
    

	return cmd
}

func init() {
	ReplaceRealmCmd := NewReplaceRealmCmd()
    RealmCmd.AddCommand(ReplaceRealmCmd)
}

var (
    
    
            DeleteRealmrealmId string
        
    
)

func NewDeleteRealmCmd() *cobra.Command {
    cmd := &cobra.Command{
	    Use:   "delete",
	  
        RunE: func(cmd *cobra.Command, args []string) error {
            
            
            
            req := apiClient.RealmAPI.DeleteRealm(apiClient.GetConfig().Context, DeleteRealmrealmId)
            
            
            
            resp, err := req.Execute()
            if err != nil {
                return err
            }
		    d, err := io.ReadAll(resp.Body)
		    if err != nil {
			    return err
		    }
		    utils.PrettyPrintByte(d)
            cmd.Println(string(d))
		    return nil
        },
    }

    
    
        cmd.Flags().StringVarP(&DeleteRealmrealmId, "realmId", "", "", "")
        cmd.MarkFlagRequired("realmId")
        
    

	return cmd
}

func init() {
	DeleteRealmCmd := NewDeleteRealmCmd()
    RealmCmd.AddCommand(DeleteRealmCmd)
}