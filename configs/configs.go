package configs

var Config Configuration

type ServerConfiguration struct {
	Host     string `mapstructure:"host"`
	HTTPPort string `mapstructure:"httpport"`
}

type DatabaseConfiguration struct {
	Host                  string `mapstructure:"host"`
	Port                  int    `mapstructure:"port"`
	User                  string `mapstructure:"user"`
	Password              string `mapstructure:"password"`
	DB                    string `mapstructure:"db"`
	ConnectionMaxLifetime int    `mapstructure:"connectionMaxLifetime"`
	MaxIdleConnections    int    `mapstructure:"MaxIdleConnections"`
	MaxOpenConnections    int    `mapstructure:"MaxOpenConnections"`
}

type Configuration struct {
	Server   ServerConfiguration   `mapstructure:"server"`
	Database DatabaseConfiguration `mapstructure:"database"`
}
