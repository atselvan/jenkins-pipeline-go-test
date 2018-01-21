package backend

import (
	m "com/privatesquare/jenkins-cli/model"
	u "com/privatesquare/jenkins-cli/utils"
	"fmt"
	"encoding/json"
	"log"
	"os"
)

func RequestCrumbToken(baseURL string, user m.AuthUser, verbose bool) string {
	url := fmt.Sprintf("%s/crumbIssuer/api/json", baseURL)
	req := u.CreateBaseRequest("GET", url, nil, user, verbose)

	respBody, _ := u.HTTPRequest(user, req, verbose)

	var crumb m.Crumb
	json.Unmarshal(respBody, &crumb)

	return crumb.Crumb
}

func CheckAuth(user m.AuthUser, status string){
	if status == fmt.Sprintf("401 Invalid password/token for user: %s", user.Username){
		log.Printf("Invalid password/token for user: %s", user.Username)
		os.Exit(1)
	} else if status == "403 No valid crumb was included in the request"{
		log.Printf("No valid crumb was included in the request")
		os.Exit(1)
	}
}
