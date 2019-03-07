package library

import (
	"net/http"
	"net/url"
)

const (
	libraryVersion = "1"
	baseURL        = "https://libraries.io/api/"
	userAgent      = "go-librariesio/" + libraryVersion
	contentType    = "application/json"
	mediaType      = "application/json"
)


// Making connection to libraries.io API
type Client struct {
	apiKey 		string
	transport 	*http.Transport
	client		*http.Client
	UserAgent	string
	BaseURL		*url.URL
}