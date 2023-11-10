package env

import (
	"os"
	"sync"
)

type (
	EnvKey    string
	Mode      string
	EnvConfig map[EnvKey]string
)

var (
	Env  *EnvConfig
	once sync.Once
)

const (
	ModeDev   Mode = "dev"
	ModeProd  Mode = "prod"
	ModeDebug Mode = "debug"
)

const (
	EnvMode EnvKey = "ENV_MODE"

	ServerPort EnvKey = "PORT"

	PsqlDbname EnvKey = "PSQL_DBNAME"
	PsqlUser   EnvKey = "PSQL_USER"
	PsqlPass   EnvKey = "PSQL_PASS"
	PsqlHost   EnvKey = "PSQL_HOST"
	PsqlPort   EnvKey = "PSQL_PORT"

	SupabaseUrl EnvKey = "SUPABASE_URL"
	SupabaseKey EnvKey = "SUPABASE_KEY"

	RabbitmqDefaultUser EnvKey = "RABBITMQ_DEFAULT_USER"
	RabbitmqDefaultPass EnvKey = "RABBITMQ_DEFAULT_PASS"

	// RedisHost EnvKey = "REDISHOST"
	// RedisPort EnvKey = "REDISPORT"

	MailHost EnvKey = "MAIL_HOST"
	MailPort EnvKey = "MAIL_PORT"
	MailFrom EnvKey = "MAIL_FROM"
	MailPass EnvKey = "MAIL_PASS"

	TlbookingIsTest        EnvKey = "TLBOOKING_IS_TEST"
	TlbookingAvailApiUrl   EnvKey = "TLBOOKING_AVAIL_API_URL"
	TlbookingBookingApiUrl EnvKey = "TLBOOKING_BOOKING_API_URL"
	TlbookingCancelApiUrl  EnvKey = "TLBOOKING_CANCEL_API_URL"
	TlbookingUsername      EnvKey = "TLBOOKING_USERNAME"
	TlbookingPassword      EnvKey = "TLBOOKING_PASSWORD"

	TestUserMail  EnvKey = "TEST_USER_MAIL"
	TestUserPass  EnvKey = "TEST_USER_PASS"
	TestAdminMail EnvKey = "TEST_ADMIN_MAIL"

	FirebaseType                    EnvKey = "FIREBASE_type"
	FirebaseProjectId               EnvKey = "FIREBASE_project_id"
	FirebasePrivateKeyId            EnvKey = "FIREBASE_private_key_id"
	FirebasePrivateKey              EnvKey = "FIREBASE_private_key"
	FirebaseClientEmail             EnvKey = "FIREBASE_client_email"
	FirebaseClientId                EnvKey = "FIREBASE_client_id"
	FirebaseAuthUri                 EnvKey = "FIREBASE_auth_uri"
	FirebaseTokenUri                EnvKey = "FIREBASE_token_uri"
	FirebaseAuthProviderX509CertUrl EnvKey = "FIREBASE_auth_provider_x509_cert_url"
	FirebaseClientX509CertUrl       EnvKey = "FIREBASE_client_x509_cert_url"
	FirebaseUniverseDomain          EnvKey = "FIREBASE_universe_domain"
)

func InitEnv() {
	once.Do(func() {
		Env = &EnvConfig{
			EnvMode: os.Getenv(string(EnvMode)),

			ServerPort: os.Getenv(string(ServerPort)),

			PsqlDbname: os.Getenv(string(PsqlDbname)),
			PsqlUser:   os.Getenv(string(PsqlUser)),
			PsqlPass:   os.Getenv(string(PsqlPass)),
			PsqlHost:   os.Getenv(string(PsqlHost)),
			PsqlPort:   os.Getenv(string(PsqlPort)),

			SupabaseUrl: os.Getenv(string(SupabaseUrl)),
			SupabaseKey: os.Getenv(string(SupabaseKey)),

			RabbitmqDefaultUser: os.Getenv(string(RabbitmqDefaultUser)),
			RabbitmqDefaultPass: os.Getenv(string(RabbitmqDefaultPass)),

			// RedisHost: os.Getenv(string(RedisHost)),
			// RedisPort: os.Getenv(string(RedisPort)),

			MailHost: os.Getenv(string(MailHost)),
			MailPort: os.Getenv(string(MailPort)),
			MailFrom: os.Getenv(string(MailFrom)),
			MailPass: os.Getenv(string(MailPass)),

			TlbookingIsTest:        os.Getenv(string(TlbookingIsTest)),
			TlbookingAvailApiUrl:   os.Getenv(string(TlbookingAvailApiUrl)),
			TlbookingBookingApiUrl: os.Getenv(string(TlbookingBookingApiUrl)),
			TlbookingCancelApiUrl:  os.Getenv(string(TlbookingCancelApiUrl)),
			TlbookingUsername:      os.Getenv(string(TlbookingUsername)),
			TlbookingPassword:      os.Getenv(string(TlbookingPassword)),

			TestUserMail:  os.Getenv(string(TestUserMail)),
			TestUserPass:  os.Getenv(string(TestUserPass)),
			TestAdminMail: os.Getenv(string(TestAdminMail)),

			FirebaseType:                    os.Getenv(string(FirebaseType)),
			FirebaseProjectId:               os.Getenv(string(FirebaseProjectId)),
			FirebasePrivateKeyId:            os.Getenv(string(FirebasePrivateKeyId)),
			FirebasePrivateKey:              os.Getenv(string(FirebasePrivateKey)),
			FirebaseClientEmail:             os.Getenv(string(FirebaseClientEmail)),
			FirebaseClientId:                os.Getenv(string(FirebaseClientId)),
			FirebaseAuthUri:                 os.Getenv(string(FirebaseAuthUri)),
			FirebaseTokenUri:                os.Getenv(string(FirebaseTokenUri)),
			FirebaseAuthProviderX509CertUrl: os.Getenv(string(FirebaseAuthProviderX509CertUrl)),
			FirebaseClientX509CertUrl:       os.Getenv(string(FirebaseClientX509CertUrl)),
			FirebaseUniverseDomain:          os.Getenv(string(FirebaseUniverseDomain)),
		}
	})
	env := *Env
	if env[EnvMode] != string(ModeDev) && env[EnvMode] != string(ModeProd) && env[EnvMode] != string(ModeDebug) {
		panic("Env " + string(EnvMode) + " is invalid. EnvMode must be dev or prod or debug.")
	}

	// envをループして値が入ってなければpanic
	for k, v := range env {
		if v == "" {
			panic("Env " + string(k) + " is empty.")
		}
	}
}

func GetEnv(key EnvKey) string {
	InitEnv()
	env := *Env
	return env[key]
}
