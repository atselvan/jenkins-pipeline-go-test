package main

import (
	b "com/privatesquare/jenkins-cli/backend"
	m "com/privatesquare/jenkins-cli/model"
	u "com/privatesquare/jenkins-cli/utils"
	"flag"
	"log"
	"os"
)

func main() {

	//options
	listFolders := flag.Bool("listFolders", false, "Get a list of all the folder on the root level of Jenkins")
	createFolder := flag.Bool("createFolder", false, "Create a folder in the root level of Jenkins. Required: folderName")
	deleteFolder := flag.Bool("deleteFolder", false, "Deletes a folder on the root level of Jenkins. Required: folderName")
	createRole := flag.Bool("createRole", false, "Create a RBAC role on the root level of Jenkins. Required: roleName, rolePermissions")
	deleteRole := flag.Bool("deleteRole", false, "Delete a RBAC role from jenkins.")
	createGroup := flag.Bool("createGroup", false, "Create a group on the root level of Jenkins. Required: groupName, roleName, memberId")
	deleteGroup := flag.Bool("deleteGroup", false, "Delete a group from the root level of Jenkins")

	//parameters
	jenkinsURL := flag.String("jenkinsURL", "http://cicd.privatesquare.in:8080", "Jenkins server URL")
	username := flag.String("username", "allan.selvan", "Username for authentication")
	password := flag.String("password", "welkom", "Password for authentication")
	folderName := flag.String("folderName", "", "Name of the Jenkins folder")
	roleName := flag.String("roleName", "", "Role name.")
	rolePermissions := flag.String("rolePermissions", "", "Comma separated values of role permissions")
	groupName := flag.String("groupName", "", "Group Name.")
	memberId := flag.String("memberID", "", "Member ID to grant access to a group")
	verbose := flag.Bool("verbose", false, "For debug logs set this flag")

	flag.Parse()

	if *jenkinsURL == "" || *username == "" || *password == "" {
		log.Fatal("jenkinsURL, username and password are required parameters for running the CLI")
	}

	user := m.AuthUser{Username: *username, Password: *password}

	if *listFolders == true {
		folderList := b.GetFolders(*jenkinsURL, user, *verbose)
		u.PrintFolderList(folderList)
	} else if *createFolder == true {
		b.CreateFolder(*jenkinsURL, *folderName, user, *verbose)
	} else if *deleteFolder == true {
		b.DeleteFolder(*jenkinsURL, *folderName, user, *verbose)
	} else if *createRole == true {
		b.CreateRole(*jenkinsURL, *roleName, *rolePermissions, user, *verbose)
	} else if *deleteRole == true {
		b.DeleteRole(*jenkinsURL, *roleName, user, *verbose)
	} else if *createGroup == true {
		b.CreateGroup(*jenkinsURL, *groupName, *roleName, *memberId, user, *verbose)
	} else if *deleteGroup == true {
		b.DeleteGroup(*jenkinsURL, *groupName, user, *verbose)
	} else {
		flag.Usage()
		log.Printf("Select a valid action flag")
		os.Exit(1)
	}
}
