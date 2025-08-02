package interfaces

type Config struct {
	IntervalPulling   int `yaml:"interval_pulling"`
	HTTPTimeout       int `yaml:"http_timeout"`
	ConnectionTimeout int `yaml:"connection_timeout"`
	ReadTimeout       int `yaml:"read_timeout"`
	WriteTimeout      int `yaml:"write_timeout"`
}
