package main

import (
	"fmt"

	"github.com/notiku/goblox"
)

func main() {
	// Create a new Roblox API client.
	// Note that only some endpoints
	// require an valid API key.
	client := goblox.New("your-api-key")

	var userID int64 = 748671568

	user, err := client.Users.GetById(userID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf(
		"Description: %s\nCreated: %s\nIs Banned: %t\nExternal App Display Name: %s\nHas Verified Badge: %t\nID: %d\nName: %s\nDisplay Name: %s\n",
		user.Description,
		user.Created,
		user.IsBanned,
		user.ExternalAppDisplayName,
		user.HasVerifiedBadge,
		user.ID,
		user.Name,
		user.DisplayName,
	)

	var username string = "ikucoder"

	user2, err := client.Users.GetByUsernames(username, true)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf(
		"Requested Username: %s\nHas Verified Badge: %t\nID: %d\nName: %s\nDisplay Name: %s\n",
		user2.RequestedUsername,
		user2.HasVerifiedBadge,
		user2.ID,
		user2.Name,
		user2.DisplayName,
	)
}
