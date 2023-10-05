package firebase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"server/infrastructure/logger"

	firebase "firebase.google.com/go"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

type FirebaseConfig struct {
	Type                    string `json:"type"`
	ProjectId               string `json:"project_id"`
	PrivateKeyId            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientId                string `json:"client_id"`
	AuthUri                 string `json:"auth_uri"`
	TokenUri                string `json:"token_uri"`
	AuthProviderX509CertUrl string `json:"auth_provider_x509_cert_url"`
	ClientX509CertUrl       string `json:"client_x509_cert_url"`
	UniverseDomain          string `json:"universe_domain"`
}

func InitFirebase() (*firebase.App, error) {
	config := FirebaseConfig{
		Type:                    os.Getenv("TYPE"),
		ProjectId:               os.Getenv("PROJECT_ID"),
		PrivateKeyId:            os.Getenv("PRIVATE_KEY_ID"),
		PrivateKey:              os.Getenv("PRIVATE_KEY"),
		ClientEmail:             os.Getenv("CLIENT_EMAIL"),
		ClientId:                os.Getenv("CLIENT_ID"),
		AuthUri:                 os.Getenv("AUTH_URI"),
		TokenUri:                os.Getenv("TOKEN_URI"),
		AuthProviderX509CertUrl: os.Getenv("AUTH_PROVIDER_X509_CERT_URL"),
		ClientX509CertUrl:       os.Getenv("CLIENT_X509_CERT_URL"),
		UniverseDomain:          os.Getenv("UNIVERSE_DOMAIN"),
	}

	jsonData, err := json.Marshal(config)
	if err != nil {
		log.Fatalf("Error marshaling to JSON: %v", err)
	}

	credentials, err := google.CredentialsFromJSON(context.Background(), jsonData)
	if err != nil {
		logger.Errorf("error credentials from json: %v\n", err)
	}
	opt := option.WithCredentials(credentials)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return app, nil
}
