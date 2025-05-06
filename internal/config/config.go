package config

type Project struct {
	Name    string `yaml:"name" validate:"required"`
	Domain  string `env:"DOMAIN" validate:"required"`
	Version string `yaml:"version" validate:"required"`
}

type Logger struct {
	Level  string `yaml:"level" validate:"required,oneof=debug info warn error disabled"`
	Format string `yaml:"format" validate:"required,oneof=text json"`
}

type Postgres struct {
	Host     string `env:"POSTGRES_HOST"     validate:"required"`
	Port     int32  `env:"POSTGRES_PORT"     validate:"required"`
	Db       string `env:"POSTGRES_DB"       validate:"required"`
	User     string `env:"POSTGRES_USER"     validate:"required"`
	Password string `env:"POSTGRES_PASSWORD" validate:"required"`
}
type Config struct {
	AppMode  string  `env:"APP_MODE"`
	Project  Project `yaml:"project"`
	Logger   Logger  `yaml:"logger"`
	Postgres Postgres
}
