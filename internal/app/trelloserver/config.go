package trelloserver

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`
	Email struct {
		Address  string `yaml:"address"`
		Password string `yaml:"password"`
	} `yaml:"email"`
}

func NewConfig() *Config {
	return &Config{
		Server: struct {
			Port string `yaml:"port"`
		}{
			Port: "8080",
		},
		Database: struct {
			User     string `yaml:"user"`
			Password string `yaml:"password"`
		}{
			User:     "user",
			Password: "qwerty",
		},
		Email: struct {
			Address  string `yaml:"address"`
			Password string `yaml:"password"`
		}{
			Address:  "user",
			Password: "qwerty",
		},
	}
}
