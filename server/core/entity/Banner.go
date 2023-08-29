package entity

type Banner struct {
	ImageURL string
	URL      string
}

func RegenBanner(
	ImageURL string,
	URL string,
) *Banner {
	return &Banner{
		ImageURL: ImageURL,
		URL:      URL,
	}
}
