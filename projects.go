package library

import "time"

// Project represents the table of project in libraries.io
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

	// Dependencies are only populated for Project Dependencies
	Dependencies []*ProjectDependency `json:"dependencies,omitempty"`

	// RepositoryURL is only populated for User Projects
	RepositoryURL *string `json:"repository_url,omitempty"`
}
