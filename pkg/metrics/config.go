package metrics

type Config struct {
	Enable  bool   `mapstructure:"enable"`
	Address string `mapstructure:"address"`
}
