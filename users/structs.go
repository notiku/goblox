package users

type Users struct {
	APIKey   string
	UsersURL string
}

type UserByIdResponse struct {
	Description            string `json:"description"`
	Created                string `json:"created"`
	IsBanned               bool   `json:"isBanned"`
	ExternalAppDisplayName string `json:"externalAppDisplayName"`
	HasVerifiedBadge       bool   `json:"hasVerifiedBadge"`
	ID                     int    `json:"id"`
	Name                   string `json:"name"`
	DisplayName            string `json:"displayName"`
}

type UserAuthenticatedResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
}
