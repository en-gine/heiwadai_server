package env

import "os"

type (
	EnvKey string
	Mode   string
)

var Env map[EnvKey]string

const (
	ModeDev  Mode = "dev"
	ModeProd Mode = "prod"
)

const (
	EnvMode EnvKey = "ENV_MODE"

	EnvServerPort EnvKey = "SERVER_PORT"

	EnvPsqlDbname EnvKey = "PSQL_DBNAME"
	EnvPsqlUser   EnvKey = "PSQL_USER"
	EnvPsqlPass   EnvKey = "PSQL_PASS"
	EnvPsqlHost   EnvKey = "PSQL_HOST"
	EnvPsqlPort   EnvKey = "PSQL_PORT"

	EnvSupabaseUrl EnvKey = "SUPABASE_URL"
	EnvSupabaseKey EnvKey = "SUPABASE_KEY"

	EnvRabbitmqDefaultUser EnvKey = "RABBITMQ_DEFAULT_USER"
	EnvRabbitmqDefaultPass EnvKey = "RABBITMQ_DEFAULT_PASS"

	EnvRedisHost EnvKey = "REDIS_HOST"
	EnvRedisPort EnvKey = "REDIS_PORT"
	EnvRedisPass EnvKey = "REDIS_PASS"

	EnvMailHost EnvKey = "MAIL_HOST"
	EnvMailPort EnvKey = "MAIL_PORT"
	EnvMailFrom EnvKey = "MAIL_FROM"
	EnvMailPass EnvKey = "MAIL_PASS"

	EnvTlbookingAvailApiUrl   EnvKey = "TLBOOKING_AVAIL_API_URL"
	EnvTlbookingBookingApiUrl EnvKey = "TLBOOKING_BOOKING_API_URL"
	EnvTlbookingUsername      EnvKey = "TLBOOKING_USERNAME"
	EnvTlbookingPassword      EnvKey = "TLBOOKING_PASSWORD"

	EnvTestUserMail  EnvKey = "TEST_USER_MAIL"
	EnvTestUserPass  EnvKey = "TEST_USER_PASS"
	EnvTestAdminMail EnvKey = "TEST_ADMIN_MAIL"

	EnvFirebaseType                    EnvKey = "FIREBASE_type"
	EnvFirebaseProjectId               EnvKey = "FIREBASE_project_id"
	EnvFirebasePrivateKeyId            EnvKey = "FIREBASE_private_key_id"
	EnvFirebasePrivateKey              EnvKey = "FIREBASE_private_key"
	EnvFirebaseClientEmail             EnvKey = "FIREBASE_client_email"
	EnvFirebaseClientId                EnvKey = "FIREBASE_client_id"
	EnvFirebaseAuthUri                 EnvKey = "FIREBASE_auth_uri"
	EnvFirebaseTokenUri                EnvKey = "FIREBASE_token_uri"
	EnvFirebaseAuthProviderX509CertUrl EnvKey = "FIREBASE_auth_provider_x509_cert_url"
	EnvFirebaseClientX509CertUrl       EnvKey = "FIREBASE_client_x509_cert_url"
	EnvFirebaseUniverseDomain          EnvKey = "FIREBASE_universe_domain"
)

func InitEnv() {
	Env = map[EnvKey]string{
		EnvMode: os.Getenv(string(EnvMode)),

		EnvServerPort: os.Getenv(string(EnvServerPort)),

		EnvPsqlDbname: os.Getenv(string(EnvPsqlDbname)),
		EnvPsqlUser:   os.Getenv(string(EnvPsqlUser)),
		EnvPsqlPass:   os.Getenv(string(EnvPsqlPass)),
		EnvPsqlHost:   os.Getenv(string(EnvPsqlHost)),
		EnvPsqlPort:   os.Getenv(string(EnvPsqlPort)),

		EnvSupabaseUrl: os.Getenv(string(EnvSupabaseUrl)),
		EnvSupabaseKey: os.Getenv(string(EnvSupabaseKey)),

		EnvRabbitmqDefaultUser: os.Getenv(string(EnvRabbitmqDefaultUser)),
		EnvRabbitmqDefaultPass: os.Getenv(string(EnvRabbitmqDefaultPass)),

		EnvRedisHost: os.Getenv(string(EnvRedisHost)),
		EnvRedisPort: os.Getenv(string(EnvRedisPort)),
		EnvRedisPass: os.Getenv(string(EnvRedisPass)),

		EnvMailHost: os.Getenv(string(EnvMailHost)),
		EnvMailPort: os.Getenv(string(EnvMailPort)),
		EnvMailFrom: os.Getenv(string(EnvMailFrom)),
		EnvMailPass: os.Getenv(string(EnvMailPass)),

		EnvTlbookingAvailApiUrl:   os.Getenv(string(EnvTlbookingAvailApiUrl)),
		EnvTlbookingBookingApiUrl: os.Getenv(string(EnvTlbookingBookingApiUrl)),
		EnvTlbookingUsername:      os.Getenv(string(EnvTlbookingUsername)),
		EnvTlbookingPassword:      os.Getenv(string(EnvTlbookingPassword)),

		EnvTestUserMail:  os.Getenv(string(EnvTestUserMail)),
		EnvTestUserPass:  os.Getenv(string(EnvTestUserPass)),
		EnvTestAdminMail: os.Getenv(string(EnvTestAdminMail)),

		EnvFirebaseType:                    os.Getenv(string(EnvFirebaseType)),
		EnvFirebaseProjectId:               os.Getenv(string(EnvFirebaseProjectId)),
		EnvFirebasePrivateKeyId:            os.Getenv(string(EnvFirebasePrivateKeyId)),
		EnvFirebasePrivateKey:              os.Getenv(string(EnvFirebasePrivateKey)),
		EnvFirebaseClientEmail:             os.Getenv(string(EnvFirebaseClientEmail)),
		EnvFirebaseClientId:                os.Getenv(string(EnvFirebaseClientId)),
		EnvFirebaseAuthUri:                 os.Getenv(string(EnvFirebaseAuthUri)),
		EnvFirebaseTokenUri:                os.Getenv(string(EnvFirebaseTokenUri)),
		EnvFirebaseAuthProviderX509CertUrl: os.Getenv(string(EnvFirebaseAuthProviderX509CertUrl)),
		EnvFirebaseClientX509CertUrl:       os.Getenv(string(EnvFirebaseClientX509CertUrl)),
		EnvFirebaseUniverseDomain:          os.Getenv(string(EnvFirebaseUniverseDomain)),
	}
	if Env[EnvMode] != string(ModeDev) && Env[EnvMode] != string(ModeProd) {
		panic("Env " + string(EnvMode) + " is invalid. EnvMode must be dev or prod.")
	}

	// envをループして値が入ってなければpanic
	for k, v := range Env {
		if v == "" {
			panic("Env " + string(k) + " is empty.")
		}
	}
}

func GetEnv(key EnvKey) string {
	return Env[key]
}
