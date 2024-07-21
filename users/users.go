package users

import (
	"fmt"

	rbxtils "github.com/notiku/goblox/utils"
)

// GetById gets detatailed user information by ID.
func (c *Users) GetById(id int64) (*UserByIdResponse, error) {
	params := map[string]string{
		"id": fmt.Sprint(id),
	}

	var result UserByIdResponse
	err := rbxtils.DoRequest("GET", c.UsersURL+"/users/"+fmt.Sprint(id), params, &result)
	if err != nil {
		return nil, err
	}

	if result.Name == "" {
		return nil, fmt.Errorf("user with ID %d not found", id)
	}

	return &result, nil
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
