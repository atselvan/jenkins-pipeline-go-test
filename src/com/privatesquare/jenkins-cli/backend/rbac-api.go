package backend

import (
	m "com/privatesquare/jenkins-cli/model"
	u "com/privatesquare/jenkins-cli/utils"
	"fmt"
	"log"
	"encoding/json"
)

func CreateRole(baseURL, roleName, permissions string, user m.AuthUser, verbose bool){
	if roleName == "" || permissions == ""  {
		log.Fatal("roleName and rolePermissions are required parameters for creating a role")
	}
	url := fmt.Sprintf("%s/roles/createRole/api/json", baseURL)
	req := u.CreateBaseRequest("POST", url, nil, user, verbose)

	query := req.URL.Query()
	query.Add("name", roleName)
	req.URL.RawQuery = query.Encode()

	_, status := u.HTTPRequest(user, req, verbose)

	CheckAuth(user, status)

	if status == "200 OK"{
		log.Printf("'%s' role is created", roleName)
	} else if status == "400 Bad Request" {
		log.Printf("A role with the name '%s' already exists", roleName)
	}
	GrantPermissionsToRole(baseURL, roleName, permissions, user, verbose)
}

func DeleteRole(baseURL, roleName string, user m.AuthUser, verbose bool){
	if roleName == "" {
		log.Fatal("roleName is a required parameter for deleting a role")
	}
	RevokeAllPermissionsFromRole(baseURL, roleName, user, verbose)
	url := fmt.Sprintf("%s/roles/deleteRole/api/json", baseURL)
	req := u.CreateBaseRequest("POST", url, nil, user, verbose)

	query := req.URL.Query()
	query.Add("name", roleName)
	req.URL.RawQuery = query.Encode()

	_, status := u.HTTPRequest(user, req, verbose)

	CheckAuth(user, status)

	if status == "200 OK"{
		log.Printf("'%s' role is deleted", roleName)
	} else if status == "400 Bad Request" {
		log.Printf("A role with the name '%s' does not exists", roleName)
	}
}

// TODO: Handle 500 Internal server error : when role permissions provided are not available
func GrantPermissionsToRole(baseURL, roleName, permissions string, user m.AuthUser, verbose bool){
	if roleName == "" || permissions == ""  {
		log.Fatal("roleName and rolePermissions are required parameters for granting permission to a role")
	}
	url := fmt.Sprintf("%s/roles/%s/grantPermissions/api/json", baseURL, roleName)
	req := u.CreateBaseRequest("POST", url, nil, user, verbose)

	query := req.URL.Query()
	query.Add("permissions", permissions)
	req.URL.RawQuery = query.Encode()

	_, status := u.HTTPRequest(user, req, verbose)

	CheckAuth(user, status)

	if status == "200 OK"{
		log.Printf("Permissions granted")
	} else if status == "400 Bad Request" {
		log.Printf("Permissions not granted")
	}
}

func RevokePermissionsFromRole(baseURL, roleName, permissions string, user m.AuthUser, verbose bool){
	if roleName == "" || permissions == ""  {
		log.Fatal("roleName and rolePermissions are required parameters for revoking permissions from a role")
	}
	url := fmt.Sprintf("%s/roles/%s/revokePermissions/api/json", baseURL, roleName)
	req := u.CreateBaseRequest("POST", url, nil, user, verbose)

	query := req.URL.Query()
	query.Add("permissions", permissions)
	req.URL.RawQuery = query.Encode()

	_, status := u.HTTPRequest(user, req, verbose)

	CheckAuth(user, status)

	if status == "200 OK"{
		log.Printf("Permissions revoked")
	} else if status == "400 Bad Request" {
		log.Printf("Permissions not revoked")
	}
}

func RevokeAllPermissionsFromRole(baseURL, roleName string, user m.AuthUser, verbose bool){
	if roleName == "" {
		log.Fatal("roleName is a required parameter for revokig permissions from a role")
	}
	url := fmt.Sprintf("%s/roles/%s/api/json", baseURL, roleName)
	req := u.CreateBaseRequest("GET", url, nil, user, verbose)

	respBody, status := u.HTTPRequest(user, req, verbose)

	CheckAuth(user, status)

	var jsonOutput m.RolePermissions
	json.Unmarshal(respBody, &jsonOutput)

	for _, permission := range jsonOutput.GrantedPermissions {
		RevokePermissionsFromRole(baseURL, roleName, permission, user, verbose)
	}
}

func CreateGroup(baseURL, groupName, roleName, memberId string, user m.AuthUser, verbose bool){
	if groupName == "" || roleName == "" || memberId == "" {
		log.Fatal("groupName is a required parameter for creating a group")
	}
	url := fmt.Sprintf("%s/groups/createGroup/api/json", baseURL)
	req := u.CreateBaseRequest("POST", url, nil, user, verbose)

	query := req.URL.Query()
	query.Add("name", groupName)
	req.URL.RawQuery = query.Encode()

	_, status := u.HTTPRequest(user, req, verbose)

	CheckAuth(user, status)

	if status == "200 OK"{
		log.Printf("'%s' group is created", groupName)
	} else if status == "400 Bad Request" {
		log.Printf("A group with the name '%s' already exists", groupName)
	}
	AddRoleToGroup(baseURL, groupName, roleName, user, verbose)
	AddMemberToGroup(baseURL, groupName, memberId, user, verbose)
}

func DeleteGroup(baseURL, groupName string, user m.AuthUser, verbose bool){
	if groupName == "" {
		log.Fatal("groupName is a required parameter for deleting a group")
	}
	url := fmt.Sprintf("%s/groups/deleteGroup/api/json", baseURL)
	req := u.CreateBaseRequest("POST", url, nil, user, verbose)

	query := req.URL.Query()
	query.Add("name", groupName)
	req.URL.RawQuery = query.Encode()

	_, status := u.HTTPRequest(user, req, verbose)

	CheckAuth(user, status)

	if status == "200 OK"{
		log.Printf("'%s' group is deleted", groupName)
	} else if status == "400 Bad Request" {
		log.Printf("A group with the name '%s' does not exists", groupName)
	}
}

func AddRoleToGroup(baseURL, groupName, roleName string, user m.AuthUser, verbose bool){
	if groupName == "" || roleName == "" {
		log.Fatal("groupName and roleName are required parameters for adding a role to a group")
	}
	url := fmt.Sprintf("%s/groups/%s/grantRole/api/json", baseURL, groupName)
	req := u.CreateBaseRequest("POST", url, nil, user, verbose)

	query := req.URL.Query()
	query.Add("role", roleName)
	query.Add("offset", "0")
	query.Add("inherited", "true")
	req.URL.RawQuery = query.Encode()

	_, status := u.HTTPRequest(user, req, verbose)

	CheckAuth(user, status)

	if status == "200 OK"{
		log.Printf("'%s' role is added to the group %s", roleName, groupName)
	} else if status == "400 Bad Request" {
		log.Printf("A group with the name '%s' already exists", groupName)
	}
}

func AddMemberToGroup(baseURL, groupName, memberId string, user m.AuthUser, verbose bool){
	if groupName == "" || memberId == "" {
		log.Fatal("groupName and memberId are required parameters for adding a member to a group")
	}
	url := fmt.Sprintf("%s/groups/%s/addMember/api/json", baseURL, groupName)
	req := u.CreateBaseRequest("POST", url, nil, user, verbose)

	query := req.URL.Query()
	query.Add("name", memberId)
	req.URL.RawQuery = query.Encode()

	_, status := u.HTTPRequest(user, req, verbose)

	CheckAuth(user, status)

	if status == "200 OK"{
		log.Printf("'%s' member is added to the group %s", memberId, groupName)
	} else if status == "400 Bad Request" {
		log.Printf("A group with the name '%s' does not exists", groupName)
	}
}