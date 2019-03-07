package library

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Project represents the table of project in libraries.io BigQuery
type Project struct {
	Description              *string    `json:"description,omitempty"`
	Forks                    *int       `json:"forks,omitempty"`
	Homepage                 *string    `json:"homepage,omitempty"`
	Keywords                 []*string  `json:"keywords,omitempty"`
	Language                 *string    `json:"language,omitempty"`
	LatestReleaseNumber      *string    `json:"latest_release_number,omitempty"`
	LatestReleasePublishedAt *time.Time `json:"latest_release_published_at,omitempty"`
	LatestStableRelease      *Releases   `json:"latest_stable_release,omitempty"`
	Name                     *string    `json:"name,omitempty"`
	NormalizedLicenses       []*string  `json:"normalized_licenses,omitempty"`
	PackageManagerURL        *string    `json:"package_manager_url,omitempty"`
	Platform                 *string    `json:"platform,omitempty"`
	Rank                     *int       `json:"rank,omitempty"`
	Stars                    *int       `json:"stars,omitempty"`
	Status                   *string    `json:"status,omitempty"`
	Versions                 []*Releases`json:"versions,omitempty"`

	// Dependencies are only for Project Dependencies
	Dependencies []*ProjectDependencies `json:"dependencies,omitempty"`

	// RepositoryURL is only for User Projects
	RepositoryURL *string `json:"repository_url,omitempty"`
}

// Releases are a for the releases of the project
type Releases struct {
	Number 			*string 	`json:"number,omitempty"`
	PublishedAt		*time.Time	`json:"published_at, omitempty"`
}

// Project Dependencies shows the dependencies of a given project
type ProjectDependencies struct {
	Deprecated 		*bool 		`json:"deprecated, omitempty"`
	Latest			*string		`json:"latest, omitempty"`
	LatestStable	*string 	`json:"latest_stable, omitempty"`
	Name			*string		`json:"outdated, omitempty"`
	Outdated		*bool		`json:"outdated, omitempty"`
	Platform		*string		`json:"platform, omitempty"`
	ProjectName		*string		`json:"platform, omitempty"`
	Requirements	*string		`json:"requirements, omitempty"`
}


// Get information about a package and its versions.
func (c * Client) Project(ctx context.Context, platform, name string) (*Project, *http.Response, error) {
	urlString := fmt.Sprintf("%v/%v", platform, name)

	req, err := c.NewRequest("GET", urlString, nil)

	if err != nil {
		return nil, nil, err
	}

	projects := new(Project)
	res, err := c.Do(ctx, req, projects)
	if err != nil {
		return nil, res, err
	}
	return projects, res, nil
}

// Get a list of dependencies for a version of a project, pass latest to get dependency
// info for the latest available version
func (c *Client) ProjectDependency(ctx context.Context, platform, name, version string) (*Project, *http.Response, error) {
	urlString = fmt.Sprintf("v%/%v/%v/dependencies", platform, name, version)

	req, err := c.NewRequest("GET", urlString, nil)
	if err != nil {
		return nil, nil, err
	}
	project := new(Project)

	res, err := c.Do(ctx, req, project)
	if err != nil {
		return nil, response, err
	}
	return project, res, nil
}