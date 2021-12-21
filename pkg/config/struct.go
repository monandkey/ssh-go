package config

// ConfigAction is an interface that defines methods for manipulating the configuration.
type ConfigAction interface {
	SetParams(string, string, string, string, string)
	Load() []configParams
	Write() error
}

// configParams is a structure that handles the config
type configParams struct {
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Keypath  string `yaml:"keypath"`
	Jump     []jump `yaml:"jump"`
}

// configDetail is a structure that manages parameters for configuring
type jump struct {
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	Keypath  string `yaml:"keypath"`
}

// params is a structure that manages parameters for configuring
type params struct {
	name     string
	user     string
	host     string
	password string
	keypath  string
}

// baseParams is a structure that inherits from params
type baseParams struct {
	params
}

// - name: server with jump
//   user: appuser
//   host: 192.168.8.35
//   port: 22
//   password: 123456
//   keypath: ""
//   jump:
//   - user: appuser
//     host: 192.168.8.36
//     port: 2222
