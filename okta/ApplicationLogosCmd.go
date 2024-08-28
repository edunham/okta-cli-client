package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var ApplicationLogosCmd = &cobra.Command{
	Use:   "applicationLogos",
	Long:  "Manage ApplicationLogosAPI",
}

func NewApplicationLogosCmd() *cobra.Command {
    cmd := &cobra.Command{
		Use:   "applicationLogos",
		Long:  "Manage ApplicationLogosAPI",
	}
	return cmd
}

func init() {
    rootCmd.AddCommand(ApplicationLogosCmd)
}

var (
    
    
            UploadApplicationLogoappId string
        
            UploadApplicationLogodata string
        
    
)

func NewUploadApplicationLogoCmd() *cobra.Command {
    cmd := &cobra.Command{
	    Use:   "uploadApplicationLogo",
	  
        RunE: func(cmd *cobra.Command, args []string) error {
            
            
            
            req := apiClient.ApplicationLogosAPI.UploadApplicationLogo(apiClient.GetConfig().Context, UploadApplicationLogoappId)
            
            
            if UploadApplicationLogodata != "" {
                req = req.Data(UploadApplicationLogodata)
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

    
    
        cmd.Flags().StringVarP(&UploadApplicationLogoappId, "appId", "", "", "")
        cmd.MarkFlagRequired("appId")
        
        cmd.Flags().StringVarP(&UploadApplicationLogodata, "data", "", "", "")
        cmd.MarkFlagRequired("data")
        
    

	return cmd
}

func init() {
	UploadApplicationLogoCmd := NewUploadApplicationLogoCmd()
    ApplicationLogosCmd.AddCommand(UploadApplicationLogoCmd)
}