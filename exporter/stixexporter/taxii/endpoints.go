package taxii

import "fmt"

func (c Config) DiscoveryEndpoint() string {

	return fmt.Sprintf(
		"%s/discovery",
		c.APIRoot,
	)
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
