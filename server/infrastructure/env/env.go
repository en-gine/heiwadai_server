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

	// AdminClientIps EnvKey = "ADMIN_CLIENT_IPs"
	// ClonClientIps  EnvKey = "CRON_CLIENT_IPs"

	ServerPort EnvKey = "PORT"

	PsqlDbname EnvKey = "PSQL_DBNAME"
	PsqlUser   EnvKey = "PSQL_USER"
	PsqlPass   EnvKey = "PSQL_PASS"
	PsqlHost   EnvKey = "PSQL_HOST"
	PsqlPort   EnvKey = "PSQL_PORT"

	SupabaseUrl EnvKey = "SUPABASE_URL"
	SupabaseKey EnvKey = "SUPABASE_KEY"

	SupabaseProjectID EnvKey = "SUPABASE_PROJECT_ID"
	SupabaseBucket    EnvKey = "SUPABASE_BUCKET"

	AdminAuthRedirectURL EnvKey = "ADMIN_AUTH_REDIRECT_URL"
	UserAuthRedirectURL  EnvKey = "USER_AUTH_REDIRECT_URL"

	EncryptKey EnvKey = "ENCRYPT_KEY"
	RedisHost  EnvKey = "REDISHOST"
	RedisPort  EnvKey = "REDISPORT"
	RedisUser  EnvKey = "REDISUSER"
	RedisPass  EnvKey = "REDISPASS"

	MailHost EnvKey = "MAIL_HOST"
	MailPort EnvKey = "MAIL_PORT"
	MailFrom EnvKey = "MAIL_FROM"
	MailPass EnvKey = "MAIL_PASS"
	MailUser EnvKey = "MAIL_USER"

	TlbookingIsTest        EnvKey = "TLBOOKING_IS_TEST"
	TlbookingAvailApiUrl   EnvKey = "TLBOOKING_AVAIL_API_URL"
	TlbookingBookingApiUrl EnvKey = "TLBOOKING_BOOKING_API_URL"
	TlbookingCancelApiUrl  EnvKey = "TLBOOKING_CANCEL_API_URL"
	TlbookingUsername      EnvKey = "TLBOOKING_USERNAME"
	TlbookingPassword      EnvKey = "TLBOOKING_PASSWORD"

	TestUserMail  EnvKey = "TEST_USER_MAIL"
	TestUserPass  EnvKey = "TEST_USER_PASS"
	TestAdminMail EnvKey = "TEST_ADMIN_MAIL"

	CronAccessKey    EnvKey = "CRON_ACCESS_KEY"
	CronAccessSecret EnvKey = "CRON_ACCESS_SECRET"
	S3Region         EnvKey = "S3_REGION"
	S3Bucket         EnvKey = "S3_BUCKET"
	S3AccessKey      EnvKey = "S3_ACCESS_KEY"
	S3SecretKey      EnvKey = "S3_SECRET_KEY"
)

func InitEnv() {
	once.Do(func() {
		Env = &EnvConfig{
			EnvMode: os.Getenv(string(EnvMode)),

			ServerPort: os.Getenv(string(ServerPort)),

			// AdminClientIps: os.Getenv(string(AdminClientIps)),
			// ClonClientIps:  os.Getenv(string(ClonClientIps)),

			PsqlDbname: os.Getenv(string(PsqlDbname)),
			PsqlUser:   os.Getenv(string(PsqlUser)),
			PsqlPass:   os.Getenv(string(PsqlPass)),
			PsqlHost:   os.Getenv(string(PsqlHost)),
			PsqlPort:   os.Getenv(string(PsqlPort)),

			SupabaseUrl: os.Getenv(string(SupabaseUrl)),
			SupabaseKey: os.Getenv(string(SupabaseKey)),

			EncryptKey: os.Getenv(string(EncryptKey)),

			SupabaseProjectID: os.Getenv(string(SupabaseProjectID)),
			SupabaseBucket:    os.Getenv(string(SupabaseBucket)),

			AdminAuthRedirectURL: os.Getenv(string(AdminAuthRedirectURL)),
			UserAuthRedirectURL:  os.Getenv(string(UserAuthRedirectURL)),

			RedisHost: os.Getenv(string(RedisHost)),
			RedisPort: os.Getenv(string(RedisPort)),
			RedisUser: os.Getenv(string(RedisUser)),
			RedisPass: os.Getenv(string(RedisPass)),

			MailHost: os.Getenv(string(MailHost)),
			MailPort: os.Getenv(string(MailPort)),
			MailFrom: os.Getenv(string(MailFrom)),
			MailPass: os.Getenv(string(MailPass)),
			MailUser: os.Getenv(string(MailUser)),

			TlbookingIsTest:        os.Getenv(string(TlbookingIsTest)),
			TlbookingAvailApiUrl:   os.Getenv(string(TlbookingAvailApiUrl)),
			TlbookingBookingApiUrl: os.Getenv(string(TlbookingBookingApiUrl)),
			TlbookingCancelApiUrl:  os.Getenv(string(TlbookingCancelApiUrl)),
			TlbookingUsername:      os.Getenv(string(TlbookingUsername)),
			TlbookingPassword:      os.Getenv(string(TlbookingPassword)),

			TestUserMail:  os.Getenv(string(TestUserMail)),
			TestUserPass:  os.Getenv(string(TestUserPass)),
			TestAdminMail: os.Getenv(string(TestAdminMail)),

			CronAccessKey:    os.Getenv(string(CronAccessKey)),
			CronAccessSecret: os.Getenv(string(CronAccessSecret)),

			S3Region:    os.Getenv(string(S3Region)),
			S3Bucket:    os.Getenv(string(S3Bucket)),
			S3AccessKey: os.Getenv(string(S3AccessKey)),
			S3SecretKey: os.Getenv(string(S3SecretKey)),
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
