# GoBlox

[Roblox](https://roblox.com/) API wrapper for [Go](https://go.dev/).

## NOTE: I don't recommend using this at the moment as it's not even close to a first beta release.

# Installation

```
go get github.com/notiku/goblox
```

# (Very) Basic Example

```go
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
}
```

More examples can be found [here](https://github.com/notiku/goblox/tree/main/examples).