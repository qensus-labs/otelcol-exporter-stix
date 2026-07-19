package taxii

type Config struct {

	// TAXII 2.1 API Root.
	//
	// Example:
	// https://taxii.example.com/taxii2/root
	APIRoot string

	// TAXII Collection ID.
	CollectionID string

	// Optional basic authentication.
	Username string
	Password string

	// Optional bearer token.
	APIKey string
}
