package taxii

import (
	"fmt"
	"net/url"
	"strings"
)

func (c Config) DiscoveryEndpoint() string {
	u, err := url.Parse(c.APIRoot)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return ""
	}

	// TAXII 2.1 discovery is at the server root,
	// not beneath the configured API Root.
	u.Path = "/taxii2/"
	u.RawPath = ""
	u.RawQuery = ""
	u.Fragment = ""

	return u.String()
}

func (c Config) CollectionsEndpoint() string {

	return fmt.Sprintf(
		"%s/collections/",
		strings.TrimRight(c.APIRoot, "/"),
	)
}

func (c Config) ObjectsEndpoint() string {

	return fmt.Sprintf(
		"%s/collections/%s/objects/",
		strings.TrimRight(c.APIRoot, "/"),
		c.CollectionID,
	)
}
