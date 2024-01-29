package action

type IFileAction interface {
	PostFile(base64Image *string, fileName string) (*string, error)
}
