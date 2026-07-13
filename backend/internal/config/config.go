package config

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Host         string `env:"RK_HOST, default=0.0.0.0"`
	BindAddress  string `env:"RK_BIND_ADDRESS, default=0.0.0.0"`
	Port         int    `env:"RK_PORT, default=8080"`
	EnabledHTTPS bool   `env:"RK_ENABLE_HTTPS, default=false"`
}

type DatabaseConfig struct {
	Host     string `env:"RK_DB_HOST, default=localhost"`
	Port     int    `env:"RK_DB_PORT, default=5432"`
	Name     string `env:"RK_DB_NAME, default=postgres"`
	User     string `env:"RK_DB_USER, default=postgres"`
	Password string `env:"RK_DB_PASSWORD, default=postgres"`
}
