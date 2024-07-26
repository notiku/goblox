package users

import (
	"fmt"

	rbxutils "github.com/notiku/goblox/utils"
)

// GetById gets detatailed user information by ID.
func (c *Users) GetById(id int64) (*UserByIdResponse, error) {
	params := map[string]string{
		"id": fmt.Sprint(id),
	}

	var result UserByIdResponse
	err := rbxutils.DoRequest("GET", c.UsersURL+"/users/"+fmt.Sprint(id), params, &result)
	if err != nil {
		return nil, err
	}

	if result.Name == "" {
		return nil, fmt.Errorf("user with ID %d not found", id)
	}

	return &result, nil
}

// GetByUsername gets detailed user information by username.
func (c *Users) GetByUsernames(username string, excludeBannedUsers bool) (*UserByUsernameResponse, error) {
	usernames := []string{username}

	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	params := map[string]interface{}{
		"usernames":          usernames,
		"excludeBannedUsers": excludeBannedUsers,
	}

	var result UserByUsernameResponseData
	if err := rbxutils.DoRequest2("POST", c.UsersURL+"/usernames/users", headers, params, &result); err != nil {
		fmt.Println("Error making request:", err)
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("user with username %s not found", username)
	}

	return &result.Data[0], nil
}

// GetAuthenticatedUser gets the minimal authenticated user.
// func (c *Users) GetAuthenticatedUser() (*UserAuthenticatedResponse, error) {
// 	var result UserAuthenticatedResponse
// 	err := rbxtils.DoRequest("GET", c.UsersURL+"/users/authenticated", nil, &result)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if result.Name == "" {
// 		return nil, fmt.Errorf("authenticated user not found")
// 	}

// 	return &result, nil
// }
