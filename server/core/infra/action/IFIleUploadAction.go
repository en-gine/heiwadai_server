package action

type IFileAction interface {
	// Upload base64 image to cloud storage and return the URL
	Upload(base64Image *string, fileName string) (*string, error)
}
