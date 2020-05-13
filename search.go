package trakt

type SearchField string

func (s SearchField) String() string { return string(s) }

const (
	// Shared search fields for media elements (movie|episode|show)
	SearchFieldTitle       SearchField = "title"
	SearchFieldOverview    SearchField = "overview"
	SearchFieldPeople      SearchField = "people"
	SearchFieldTranslation SearchField = "translations"
	SearchFieldAlias       SearchField = "aliases"

	// search fields specifically for movies.
	SearchFieldTagline SearchField = "tagline"

	// Shared search fields for both person and list.
	SearchFieldName SearchField = "name"

	// search fields specifically for person
	SearchFieldBiography SearchField = "biography"

	// search fields specifically for list.
	SearchFieldDescription SearchField = "description"
)

type SearchQueryParams struct {
	BasicListParams

	Type     Type          `json:"-" url:"-"`
	Query    string        `json:"-" url:"query"`
	Fields   []SearchField `json:"-" url:"fields,comma,omitempty"`
	Extended ExtendedType  `json:"-" url:"extended,omitempty"`
}

type IDLookupParams struct {
	BasicListParams

	Type     Type         `json:"-" url:"-"`
	Extended ExtendedType `json:"-" url:"extended,omitempty"`
}

// SearchResult represents a result from performing a search either by
// test search or ID lookup.
type SearchResult struct {
	GenericElement

	Person *Person `json:"person"`
}

// SearchResultIterator an instance of an iterator which allows us to
// return the current pointer as a concrete SearchResult struct.
type SearchResultIterator struct{ Iterator }

func (s *SearchResultIterator) Result() *SearchResult { return s.Current().(*SearchResult) }
