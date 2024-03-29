package language

// Tag defined by RFC5646: https://tools.ietf.org/html/rfc5646
type Tag struct {
	// MUST contain at least one each
	Type        string
	Description []string
	Added       string

	// MUST contain one of them.
	Tag    string
	Subtag string

	// MAY also contain the following fields
	Deprecated     string
	PreferredValue string `json:"Preferred-Value"`
	Prefix         []string
	SuppressScript string `json:"Suppress-Script"`
	Macrolanguage  string
	Scope          string
	Comments       string
}
