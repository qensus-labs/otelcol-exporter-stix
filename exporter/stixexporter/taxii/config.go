package taxii

type Config struct {

	// TAXII 2.1 API Root.
	//
	// Example:
	// https://taxii.example.com/taxii2/root
	APIRoot string `mapstructure:"api_root"`

	// TAXII Collection ID.
	CollectionID string `mapstructure:"collection_id"`

	// Optional basic authentication.
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`

	// Optional bearer token.
	APIKey string `mapstructure:"api_key"`
}
