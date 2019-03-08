package library

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	ID 				*int		`json:"id, omitempty"`
	UUID 			*int		`json:"uuid, omitempty"`
	Login			*string		`json:"login, omitempty"`
	UserType		*string		`json:"user_type, omitempty"`
	CreatedAt		*time.Time	`json:"created_at, omitempty"`
	UpdatedAt		*time.Time	`json:"updated_at, omitempty"`
	Name 			*string 	`json:"name, omitempty"`
	Company			*string		`json:"company, omitempty"`
	Blog			*string		`json:"blog, omitempty"`
	Location 		*string 	`json:"location, omitempty"`
	Hidden 			*bool		`json:"hidden, omitempty"`
	LastSyncedAt 	*time.Time	`json:"last_synced, omitempty"`
	Email			*string		`json:"email, omitempty"`
	Bio 			*string 	`json:"bio, omitempty"`
	Followers		*int		`json:"followers, omitempty"`
	Following		*int		`json:"following, omitempty"`
	HostType		*string		`json:"host_type, omitempty"`
	Github			*int		`json:"hithub_id, ommitempty"`
}

type Repository struct {
	ContributionsCount       *int       `json:"contributions_count,omitempty"`
	CreatedAt                *time.Time `json:"created_at,omitempty"`
	DefaultBranch            *string    `json:"default_branch,omitempty"`
	Description              *string    `json:"description,omitempty"`
	Fork                     *bool      `json:"fork,omitempty"`
	ForkPolicy               *string    `json:"fork_policy,omitempty"`
	ForksCount               *int       `json:"forks_count,omitempty"`
	FullName                 *string    `json:"full_name,omitempty"`
	GithubContributionsCount *int       `json:"github_contributions_count,omitempty"`
	GithubID                 *string    `json:"github_id,omitempty"`
	HasAudit                 *string    `json:"has_audit,omitempty"`
	HasChangelog             *string    `json:"has_changelog,omitempty"`
	HasCoc                   *string    `json:"has_coc,omitempty"`
	HasContributing          *string    `json:"has_contributing,omitempty"`
	HasIssues                *bool      `json:"has_issues,omitempty"`
	HasLicense               *string    `json:"has_license,omitempty"`
	HasPages                 *bool      `json:"has_pages,omitempty"`
	HasReadme                *string    `json:"has_readme,omitempty"`
	HasThreatModel           *string    `json:"has_threat_model,omitempty"`
	HasWiki                  *bool      `json:"has_wiki,omitempty"`
	Homepage                 *string    `json:"homepage,omitempty"`
	HostDomain               *string    `json:"host_domain,omitempty"`
	HostType                 *string    `json:"host_type,omitempty"`
	Keywords                 []*string  `json:"keywords,omitempty"`
	Language                 *string    `json:"language,omitempty"`
	LastSyncedAt             *time.Time `json:"last_synced_at,omitempty"`
	License                  *string    `json:"license,omitempty"`
	LogoURL                  *string    `json:"logo_url,omitempty"`
	MirrorURL                *string    `json:"mirror_url,omitempty"`
	Name                     *string    `json:"name,omitempty"`
	OpenIssuesCount          *int       `json:"open_issues_count,omitempty"`
	Private                  *bool      `json:"private,omitempty"`
	PullRequestsEnabled      *bool      `json:"pull_requests_enabled,omitempty"`
	PushedAt                 *time.Time `json:"pushed_at,omitempty"`
	Rank                     *int       `json:"rank,omitempty"`
	Scm                      *string    `json:"scm,omitempty"`
	Size                     *int       `json:"size,omitempty"`
	SourceName               *string    `json:"source_name,omitempty"`
	StargazersCount          *int       `json:"stargazers_count,omitempty"`
	Status                   *string    `json:"status,omitempty"`
	SubscribersCount         *int       `json:"subscribers_count,omitempty"`
	UUID                     *string    `json:"uuid,omitempty"`
	UpdatedAt                *time.Time `json:"updated_at,omitempty"`
}

// Login gets information for a given user or organization.
func (c *Client) UserInfo(ctx context.Context, login string) (*User, *http.Response, error) {
	urlString := fmt.Sprintf("github/%v", login)

	req, err := c.NewRequest("GET", urlString, nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(User)

	res, err := c.makeCall(ctx, req, user)
	if err != nil {
		return nil, res, err
	}
	return user, res, nil
}

// UserPackages gets a list of packages referencing the given user's repositories.
func (c *Client) UserPackages(ctx context.Context, login string) ([]*Project, *http.Response, error) {
	urlString := fmt.Sprintf("github/%v/projects", login)

	req, err := c.NewRequest("GET", urlString, nil)

	if err != nil {
		return nil, nil, err
	}

	var projects []*Project

	res, err := c.makeCall(ctx, req, &projects)
	if err != nil {
		return nil, res, err
	}
	return projects, res, nil
}

// UserRepositories gets repositories owned by a user.
func (c *Client) UserRepositories(ctx context.Context, login string) ([]*Repository, *http.Response, error) {
	urlString := fmt.Sprintf("github/%v/repositories", login)

	req, err := c.NewRequest("GET", urlString, nil)

	if err != nil {
		return nil, nil, err
	}
	var repo []*Repository

	res, err := c.makeCall(ctx, req, &repo)
	if err != nil {
		return nil, res, err
	}
	return repo, res, nil
}
