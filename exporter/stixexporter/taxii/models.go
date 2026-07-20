package taxii

type Discovery struct {
	Title string `json:"title"`

	APIRoots []string `json:"api_roots"`
}

type Collection struct {
	ID string `json:"id"`

	Title string `json:"title"`
}

type Collections struct {
	Collections []Collection `json:"collections"`
}
