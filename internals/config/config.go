package config

type Config struct {
	Port int
	Env  string
	DB   struct {
		DSN string
	}
	Jwt struct {
		Secret string
	}
}
