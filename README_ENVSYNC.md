# Envsync

Synchronize one Okta environment with another

```
$ okta-cli-client envsync pulluser --userId abc123

# user is created in ~/.okta/orgid/user@host.json

$ okta-cli-client envsync pushuser --userdata ~/.okta/dev-000000.okta.com/users/user@host.json
```

```
$ okta-cli-client envsync -h
backup and restore okta environments

Usage:
  okta-cli-client envsync [command]

Available Commands:
  clean              
  pullAllGroups      
  pullAllMappings    
  pullAllPolicies    
  pullAllRules       
  pullAllUsers       
  pullGroup          
  pullGroupUsers     
  pullMapping        
  pullPolicy         
  pullRule           
  pullUser           
  pushAllUsers       
  pushGroup          
  pushPolicy         
  pushResourcePolicy 
  pushRule           
  pushUser           

Flags:
  -h, --help   help for envsync

Use "okta-cli-client envsync [command] --help" for more information about a command.
```
