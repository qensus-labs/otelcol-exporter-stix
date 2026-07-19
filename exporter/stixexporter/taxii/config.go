package taxii

type Config struct {

	// TAXII 2.1 collection endpoint.
	URL string

	// Optional basic authentication.
	Username string
	Password string
}
