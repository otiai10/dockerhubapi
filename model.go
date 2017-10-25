package dockerhubapi

import (
	"path"
	"time"
)

// Resource ...
type Resource interface {
	Path(...string) string
}

// Repository ...
type Repository struct {
	User            string      `json:"user"`
	Name            string      `json:"name"`
	Namespace       string      `json:"namespace"`
	RepositoryType  string      `json:"repository_type"`
	Status          int         `json:"status"`
	Description     string      `json:"description"`
	IsPrivate       bool        `json:"is_private"`
	IsAutomated     bool        `json:"is_automated"`
	CanEdit         bool        `json:"can_edit"`
	StarCount       int         `json:"star_count"`
	PullCount       int         `json:"pull_count"`
	LastUpdated     time.Time   `json:"last_updated"`
	BuildOnCloud    interface{} `json:"build_on_cloud"`
	HasStarred      bool        `json:"has_starred"`
	FullDescription string      `json:"full_description"`
	Affiliation     interface{} `json:"affiliation"`
	Permissions     struct {
		Read  bool `json:"read"`
		Write bool `json:"write"`
		Admin bool `json:"admin"`
	} `json:"permissions"`
}

// Path ...
func (r *Repository) Path(queries ...string) string {
	return path.Join("/repositories", r.User, r.Name)
}

// UserRepositories ...
type UserRepositories struct {
	User     string       `json:"-"`
	Count    int          `json:"count"`
	Next     interface{}  `json:"next"`
	Previous interface{}  `json:"previous"`
	Results  []Repository `json:"results"`
}

// Path ...
func (u *UserRepositories) Path(queries ...string) string {
	return path.Join("/repositories", u.User)
}

// User ...
type User struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	FullName    string    `json:"full_name"`
	Location    string    `json:"location"`
	Company     string    `json:"company"`
	ProfileURL  string    `json:"profile_url"`
	DateJoined  time.Time `json:"date_joined"`
	GravatarURL string    `json:"gravatar_url"`
	Type        string    `json:"type"`
}

// Path ...
func (u *User) Path(queries ...string) string {
	return path.Join("/users", u.Username)
}
