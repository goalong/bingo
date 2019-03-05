package conf

import (
	"github.com/jinzhu/configor"
	"github.com/joho/godotenv"
)

var Config = struct {
	APP struct {
		Debug    bool   `default:"true"`
		Host     string `default:"0.0.0.0"`
		Port     string `default:"8000"`
		PageSize int    `default:"10"`
		BaseURL  string `default:"https://api.example.com/"`
		JwtSecret string `default:"087047$086092"`
	}

	DB struct {
		Host     string `default:"127.0.0.1"`
		Port     string `default:"3306"`
		User     string `default:"user1"`
		Password string `default:"123456"`
		Name     string `default:"flask_demo"`
	}

	Redis struct {
		Host     string `default:"127.0.0.1"`
		Port     string `default:"6379"`
		Password string
		DB       int `default:"0"`
	}
}{}

func init() {
	godotenv.Load()
	configor.Load(&Config)
}