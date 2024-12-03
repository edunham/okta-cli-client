package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

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

func main() {
	// Read the userdata.json file
	userData, err := os.ReadFile("userdata.json")
	if err != nil {
		fmt.Println("Error reading userdata.json:", err)
		return
	}

	// Unmarshal the JSON data
	var data UserData
	err = json.Unmarshal(userData, &data)
	if err != nil {
		fmt.Println("Error parsing userdata.json:", err)
		return
	}

	// Prepare the CLI command
	cliCommand := fmt.Sprintf(
		"okta-cli-client user create --data '{\"credentials\":{\"password\":{\"value\":\"Hell4W0rld\"}},\"profile\":{\"email\":\"%s\",\"firstName\":\"%s\",\"lastName\":\"%s\",\"login\":\"%s\"}}'",
		data.Profile.Email, data.Profile.FirstName, data.Profile.LastName, data.Profile.Login,
	)

	// Execute the CLI command
	cmd := exec.Command("bash", "-c", cliCommand)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing CLI command:", err)
		fmt.Println("Output:", string(output))
		return
	}

	fmt.Println("User created successfully.")
}
