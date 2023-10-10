package firebase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"server/infrastructure/env"
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
		Type:                    env.GetEnv(env.FirebaseType),
		ProjectId:               env.GetEnv(env.FirebaseProjectId),
		PrivateKeyId:            env.GetEnv(env.FirebasePrivateKeyId),
		PrivateKey:              env.GetEnv(env.FirebasePrivateKey),
		ClientEmail:             env.GetEnv(env.FirebaseClientEmail),
		ClientId:                env.GetEnv(env.FirebaseClientId),
		AuthUri:                 env.GetEnv(env.FirebaseAuthUri),
		TokenUri:                env.GetEnv(env.FirebaseTokenUri),
		AuthProviderX509CertUrl: env.GetEnv(env.FirebaseAuthProviderX509CertUrl),
		ClientX509CertUrl:       env.GetEnv(env.FirebaseClientX509CertUrl),
		UniverseDomain:          env.GetEnv(env.FirebaseUniverseDomain),
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
