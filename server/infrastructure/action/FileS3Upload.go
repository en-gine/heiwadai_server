package action

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"mime"
	"strings"

	"server/core/infra/action"
	"server/infrastructure/env"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var _ action.IFileAction = &FileS3Client{}

type FileS3Client struct {
	client *s3.S3
}

var (
	S3Bucket    = env.GetEnv(env.S3Bucket)
	S3AccessKey = env.GetEnv(env.S3AccessKey)
	S3Secretkey = env.GetEnv(env.S3SecretKey)
	S3Region    = env.GetEnv(env.S3Region)
)

func NewS3FileS3Client() *FileS3Client {
	newSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// S3クライアントを作成します
	client := s3.New(newSession, &aws.Config{
		Region:      aws.String(S3Region),
		Credentials: credentials.NewStaticCredentials(S3AccessKey, S3Secretkey, ""),
	})
	return &FileS3Client{
		client: client,
	}
}

func (fc *FileS3Client) Upload(base64Image *string, filename string) (*string, error) {
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
	bucket := aws.String(S3Bucket)
	key := aws.String(postFileName)

	_, err = fc.client.PutObject(&s3.PutObjectInput{
		Bucket:      bucket,
		Key:         key,
		Body:        reader,
		ContentType: aws.String(mimetype),
	})
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://%s.amazonaws.com/%s/%s", S3Region, S3Bucket, postFileName)

	// 画像のURLを返す
	return &url, nil
}

func (fc *FileS3Client) getFileExtensionFromBase64(base64Image *string) (string, error) {
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
