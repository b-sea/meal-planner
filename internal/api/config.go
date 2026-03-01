package api

// Config is the main API configuration.
type Config struct {
	HTTP struct {
		Port         int   `config:"port"`
		ReadTimeout  int64 `config:"readTimeout"`
		WriteTimeout int64 `config:"writeTimeout"`
	} `config:"http"`

	Logger struct {
		Level string `config:"level"`
	} `config:"logger"`

	MariaDB struct {
		Host string `config:"host"`
		User string `config:"user"`
		Pwd  string `config:"pwd"`
	} `config:"mariadb"`
}
