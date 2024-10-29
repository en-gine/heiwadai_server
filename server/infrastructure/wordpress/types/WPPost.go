package types

type WPPost struct {
	ID            int      `json:"id"`
	Date          string   `json:"date"`
	DateGMT       string   `json:"date_gmt"`
	GUID          GUID     `json:"guid"`
	Modified      string   `json:"modified"`
	ModifiedGMT   string   `json:"modified_gmt"`
	Slug          string   `json:"slug"`
	Status        string   `json:"status"`
	Type          string   `json:"type"`
	Link          string   `json:"link"`
	Title         Title    `json:"title"`
	Content       Content  `json:"content"`
	Excerpt       Content  `json:"excerpt"`
	Author        int      `json:"author"`
	FeaturedMedia int      `json:"featured_media"`
	CommentStatus string   `json:"comment_status"`
	PingStatus    string   `json:"ping_status"`
	Sticky        bool     `json:"sticky"`
	Template      string   `json:"template"`
	Format        string   `json:"format"`
	Meta          Meta     `json:"meta"` // Changed from []string to Meta struct
	Categories    []int    `json:"categories"`
	Tags          []int    `json:"tags"`
	ACF           []any    `json:"acf"`             // Added ACF field
	YoastHead     string   `json:"yoast_head"`      // Added YoastHead field
	YoastHeadJSON any      `json:"yoast_head_json"` // Added YoastHeadJSON field
	Links         Links    `json:"_links"`
	Embedded      Embedded `json:"_embedded"`
}

// New struct for Meta
type Meta struct {
	Footnotes string `json:"footnotes"`
}

type GUID struct {
	Rendered string `json:"rendered"`
}

type Title struct {
	Rendered string `json:"rendered"`
}

type Content struct {
	Rendered  string `json:"rendered"`
	Protected bool   `json:"protected"`
}

type Links struct {
	Self               []HrefLink           `json:"self"`
	Collection         []HrefLink           `json:"collection"`
	About              []HrefLink           `json:"about"`
	Author             []AuthorLink         `json:"author"`              // Changed to AuthorLink
	Replies            []AuthorLink         `json:"replies"`             // Changed to AuthorLink
	VersionHistory     []VersionLink        `json:"version-history"`     // Changed to VersionLink
	PredecessorVersion []PredecessorVersion `json:"predecessor-version"` // Changed to PredecessorVersion
	WPAttachment       []HrefLink           `json:"wp:attachment"`
	WPTerm             []WPTermLink         `json:"wp:term"`
	Curies             []Curies             `json:"curies"`
}

type HrefLink struct {
	Href string `json:"href"`
}

// New struct for Author Link
type AuthorLink struct {
	Embeddable bool   `json:"embeddable"`
	Href       string `json:"href"`
}

// New struct for Version Link
type VersionLink struct {
	Count int    `json:"count"`
	Href  string `json:"href"`
}

// New struct for Predecessor Version
type PredecessorVersion struct {
	ID   int    `json:"id"`
	Href string `json:"href"`
}

type WPTermLink struct {
	Taxonomy   string `json:"taxonomy"`
	Embeddable bool   `json:"embeddable"`
	Href       string `json:"href"`
}

type Curies struct {
	Name      string `json:"name"`
	Href      string `json:"href"`
	Templated bool   `json:"templated"`
}

type Embedded struct {
	Author []Author `json:"author"`
	WPTerm [][]Term `json:"wp:term"`
}

type Author struct {
	ID            int               `json:"id"`
	Name          string            `json:"name"`
	URL           string            `json:"url"`
	Description   string            `json:"description"`
	Link          string            `json:"link"`
	Slug          string            `json:"slug"`
	AvatarUrls    map[string]string `json:"avatar_urls"`
	YoastHead     string            `json:"yoast_head"`      // Added YoastHead field
	YoastHeadJSON any               `json:"yoast_head_json"` // Added YoastHeadJSON field
	ACF           []any             `json:"acf"`             // Added ACF field
	Links         Links             `json:"_links"`
}

type Term struct {
	ID            int    `json:"id"`
	Link          string `json:"link"`
	Name          string `json:"name"`
	Slug          string `json:"slug"`
	Taxonomy      string `json:"taxonomy"`
	YoastHead     string `json:"yoast_head"`      // Added YoastHead field
	YoastHeadJSON any    `json:"yoast_head_json"` // Added YoastHeadJSON field
	ACF           []any  `json:"acf"`             // Added ACF field
	Links         Links  `json:"_links"`
}
