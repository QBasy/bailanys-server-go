package logger

type Config struct {
	Level string `json:"level" yaml:"level"`

	WebRTC struct {
		Level string `yaml:"level"`
	} `yaml:"webrtc"`

	Database struct {
		Level string `yaml:"level"`
	} `yaml:"database"`

	Outputs struct {
		Stdout bool `yaml:"stdout"`
		db     bool `yaml:"db"`
	} `yaml:"outputs"`
}

const (
	DefaultLevel = "info"
)

func (c *Config) ApplyDefaults() {
	if c.Level == "" {
		c.Level = DefaultLevel
	}

	if c.WebRTC.Level == "" {
		c.WebRTC.Level = c.Level
	}

	if c.Database.Level == "" {
		c.Database.Level = c.Level
	}

	if !c.Outputs.Stdout && !c.Outputs.db {
		c.Outputs.Stdout = true
	}
}
