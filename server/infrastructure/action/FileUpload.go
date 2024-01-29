package action

import (
	"bytes"
	"encoding/base64"
	"errors"
	"log"
	"mime"
	"strings"

	"server/core/infra/action"
	"server/infrastructure/env"

	storage_go "github.com/supabase-community/storage-go"
)

var _ action.IFileAction = &FileClient{}

type FileClient struct {
	client *storage_go.Client
}

var (
	projectID = env.GetEnv(env.SupabaseProjectID)
	bucket    = env.GetEnv(env.SupabaseBucket)
	authkey   = env.GetEnv(env.SupabaseKey)
)

func NewFileClient() *FileClient {
	client := storage_go.NewClient("https://"+projectID+".supabase.co/storage/v1", authkey, nil)

	return &FileClient{
		client: client,
	}
}

func (fc *FileClient) Upload(base64Image *string, filename string) (*string, error) {
	if base64Image == nil {
		return nil, errors.New("base64エンコードされたImageデータが見つかりません。")
	}
	// base64エンコードされた画像をデコード
	base64String := *base64Image
	b64data := base64String[strings.IndexByte(base64String, ',')+1:]
	data, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		log.Fatalf("Base64デコードに失敗しました: %v", err)
	}

	ext, err := fc.getFileExtensionFromBase64(base64Image)
	if err != nil {
		return nil, err
	}

	// bytes.Readerを使用してio.Readerを作成
	reader := bytes.NewReader(data)
	postFileName := filename + ext
	mimetype := mime.TypeByExtension(ext)
	upsert := true
	_, err = fc.client.UploadFile(bucket, postFileName, reader, storage_go.FileOptions{
		ContentType: &mimetype,
		Upsert:      &upsert,
	})
	if err != nil {
		return nil, err
	}
	result := fc.client.GetPublicUrl(bucket, postFileName)
	// 画像のURLを返す
	return &result.SignedURL, nil
}

func (fc *FileClient) getFileExtensionFromBase64(base64Image *string) (string, error) {
	// MIMEタイプを抽出するために';'と','で文字列を分割
	data := *base64Image
	parts := strings.Split(data, ";")
	if len(parts) < 2 {
		return "", errors.New("無効な形式のデータです")
	}

	mimePart := strings.Split(parts[0], ":")
	if len(mimePart) < 2 {
		return "", errors.New("MIMEタイプが見つかりません")
	}

	// MIMEタイプに基づいて拡張子を決定
	switch mimePart[1] {
	case "image/png":
		return ".png", nil
	case "image/jpeg":
		return ".jpg", nil
	case "image/gif":
		return ".gif", nil
	case "image/svg+xml":
		return ".svg", nil
	case "image/webp":
		return ".webp", nil
	default:
		return "", errors.New("未知のMIMEタイプ: " + mimePart[1])
	}
}
