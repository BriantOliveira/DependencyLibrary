package library

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"bytes"
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

func NewClient(apiKey string) *Client {
	APIBaseURL, _ := url.Parse(baseURL)
	transport := &http.Transport{}
	client := &http.Client{Transport:transport}

	return &Client{
		apiKey: apiKey,
		client: client,
		transport: transport,
		UserAgent: userAgent,
		BaseURL: APIBaseURL,
	}
}

// NewRequest Creates a new request to API, creating a URL and serialize params
func (c *Client) NewRequest(method, urlString string, data interface{}) (*http.Request, error) {
	tempUrl, err := url.Parse(urlString)

	if err != nil {
		return nil, err
	}

	newUrl := c.BaseURL.ResolveReference(tempUrl)

	var payload io.ReadWriter
	if data != nil {
		payload = new(bytes.Buffer)

		err := json.NewEncoder(payload).Encode(data)
		if err != nil {
			return nil, err
		}

	}
	req, err := http.NewRequest(method, newUrl.String(), payload)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("api_key", c.apiKey)
	req.Header.Set("Accept", mediaType)
	req.Header.Set("User-Agent", c.UserAgent)

	if data != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return req, nil
}

// OverwriteAPIKey overwrites the api_key
func OverwriteAPIKey(url *url.URL) *url.URL {
	q := url.Query()
	q.Set("api_key", "REDACTED")
	url.RawQuery = q.Encode()
	return url
}

// ResError stores errors return by API call
type ResError struct {
	Response 	*http.Response
	Message 	string 	`json:"error"`
}

// Error handles the error
func (res *ResError) Error() string {
	return fmt.Sprintf(
		"%v %v: %d %q",
		res.Response.Request.Method,
		OverwriteAPIKey(res.Response.Request.URL),
		res.Response.StatusCode,
		res.Message,
		)
}

//Look Response check the response from API and see if there is any error
func LookResponse(res *http.Response) error {
	if stat := res.StatusCode; 200 <= stat && stat <= 299 {
		return nil
	}

	ResponseErr := &ResError{Response: res}

	data, err := ioutil.ReadAll(res.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, ResponseErr)
	}
	return ResponseErr
}


