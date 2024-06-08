package config

type Config struct {
	App    App    `mapstructure:"app"`
	DB     DB     `mapstructure:"db"`
	AWS    AWS    `mapstructure:"aws"`
	OpenAI OpenAI `mapstructure:"openai"`
}

type App struct {
	Port int `mapstructure:"port"`
}

type AWS struct {
	Bucket     string `mapstructure:"bucket"`
	Key        string `mapstructure:"key"`
	PrivateKey string `mapstructure:"private_key"`
}

type OpenAI struct {
	Key string `mapstructure:"key"`
}
type DB struct {
	Name     string `mapstructure:"name"`
	Username string `mapstructure:"username"`
	Port     string `mapstructure:"port"`
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
}
