package backend


import (
	m "com/privatesquare/jenkins-cli/model"
	u "com/privatesquare/jenkins-cli/utils"
	"fmt"
	"encoding/json"
	"log"
)

func GetFolders(baseURL string, user m.AuthUser, verbose bool) []string{
	url := fmt.Sprintf("%s/api/json?pretty", baseURL)
	req := u.CreateBaseRequest("GET", url, nil, user, verbose)

	respBody, _ := u.HTTPRequest(user, req, verbose)

	var (
		folders m.Folders
		folderList []string
	)

	json.Unmarshal(respBody, &folders)

	for _, folder := range folders.Jobs {
		folderList = append(folderList, folder.Name)
	}
	return folderList
}

func folderExists(baseURL, folderName string, user m.AuthUser, verbose bool) bool{
	if folderName == "" {
		log.Fatal("folderName is a required parameter for checking if a folder exists")
	}
	var folderExists bool
	folderList := GetFolders(baseURL, user, verbose)

	for _, folder := range folderList{
		if folder == folderName{
			folderExists = true
		}else{
			folderExists = false
		}
	}
	return folderExists
}

func CreateFolder(baseURL string, folderName string, user m.AuthUser, verbose bool){
	if folderName == "" {
		log.Fatal("folderName is a required parameter for creating a folder")
	}
	url := fmt.Sprintf("%s/createItem", baseURL)
	req := u.CreateBaseRequest("POST", url, nil, user, verbose)

	if folderExists(baseURL, folderName, user, verbose){
		log.Printf("'%s' folder already exists", folderName)
	}else{
		query := req.URL.Query()
		query.Add("name", folderName)
		query.Add("mode", "com.cloudbees.hudson.plugins.folder.Folder")
		req.URL.RawQuery = query.Encode()

		_, status := u.HTTPRequest(user, req, verbose)

		if status == "200 OK"{
			log.Printf("'%s' folder is created", folderName)
		}
	}
}

func DeleteFolder(baseURL string, folderName string, user m.AuthUser, verbose bool){
	if folderName == "" {
		log.Fatal("folderName is a required parameter for deleting a folder")
	}
	url := fmt.Sprintf("%s/job/%s/doDelete", baseURL, folderName)
	req := u.CreateBaseRequest("POST", url, nil, user, verbose)

	query := req.URL.Query()
	query.Add("Submit", "Yes")
	req.URL.RawQuery = query.Encode()

	_, status := u.HTTPRequest(user, req, verbose)

	if status == "302 Found"{
		log.Printf("A job is created with the name %s", folderName)
	}
}