package xlogger

type Config struct {
	Level        string   `mapstructure:"level" env:"LEVEL"`
	Env          string   `mapstructure:"env" env:"ENV"`
	Encoding     string   `mapstructure:"encoding" env:"ENCODING"`
	Outputs      []string `mapstructure:"outputs" env:"OUTPUT_PATHS"`
	ErrorOutputs []string `mapstructure:"error_outputs" env:"ERROR_OUTPUT_PATHS"`
	SentryLevel  string   `mapstructure:"sentry_level"`
	SentryDSN    string   `mapstructure:"sentry_dsn"`
}
