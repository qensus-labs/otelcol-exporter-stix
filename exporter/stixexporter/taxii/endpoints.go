package taxii

import (
	"fmt"
	"net/url"
)

// DiscoveryEndpoint returns the Discovery URL, which TAXII 2.1
// serves at /taxii2/ on the host of the API Root.
func (c Config) DiscoveryEndpoint() string {

	apiRoot, err := url.Parse(c.APIRoot)

	if err != nil {
		return ""
	}

	return apiRoot.ResolveReference(
		&url.URL{Path: "/taxii2/"},
	).String()
}

func (c Config) CollectionsEndpoint() string {

	return fmt.Sprintf(
		"%s/collections",
		c.APIRoot,
	)
}

func (c Config) ObjectsEndpoint() string {

	return fmt.Sprintf(
		"%s/collections/%s/objects/",
		c.APIRoot,
		c.CollectionID,
	)
}
