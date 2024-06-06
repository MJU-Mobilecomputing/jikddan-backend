package config

type Config struct {
	App App `mapstructure:"app"`
	DB  DB  `mapstructure:"db"`
}

type App struct {
	Port int `mapstructure:"port"`
}

type DB struct {
	Name     string `mapstructure:"name"`
	Username string `mapstructure:"username"`
	Port     string `mapstructure:"port"`
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
}
