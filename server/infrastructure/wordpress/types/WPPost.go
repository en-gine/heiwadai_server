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
	Meta          []string `json:"meta"`
	Categories    []int    `json:"categories"`
	Tags          []int    `json:"tags"`
	Links         Links    `json:"_links"`
	Embedded      Embedded `json:"_embedded"`
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
	Self               []HrefLink   `json:"self"`
	Collection         []HrefLink   `json:"collection"`
	About              []HrefLink   `json:"about"`
	Author             []HrefLink   `json:"author"`
	Replies            []HrefLink   `json:"replies"`
	VersionHistory     []HrefLink   `json:"version-history"`
	PredecessorVersion []HrefLink   `json:"predecessor-version"`
	WPAttachment       []HrefLink   `json:"wp:attachment"`
	WPTerm             []WPTermLink `json:"wp:term"`
	Curies             []Curies     `json:"curies"`
}

type HrefLink struct {
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
	ID          int               `json:"id"`
	Name        string            `json:"name"`
	URL         string            `json:"url"`
	Description string            `json:"description"`
	Link        string            `json:"link"`
	Slug        string            `json:"slug"`
	AvatarUrls  map[string]string `json:"avatar_urls"`
	Links       Links             `json:"_links"`
}

type Term struct {
	ID       int    `json:"id"`
	Link     string `json:"link"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	Taxonomy string `json:"taxonomy"`
	Links    Links  `json:"_links"`
}
