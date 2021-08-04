package config

const (
	DEV = "DEV"

	CURRENT_PHASE = DEV
)

type DbConfig struct {
	Driver   string
	User     string
	Password string
	DbName   string
}

var DB_CONFIGS map[string]DbConfig = map[string]DbConfig{
	DEV: {
		Driver:   "mysql",
		User:     "root",
		Password: "password",
		DbName:   "data",
	},
}
