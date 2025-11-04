package configs

import (
	"errors"
	"os"
)

type AppConfig struct {
	ServerPort string
}

func SetUpEnv() (cfg AppConfig, err error) {
	httpPort := os.Getenv("PORT")

	if len(httpPort) < 1 {
		return AppConfig{}, errors.New("no enviroment variables found")
	}

	return AppConfig{ServerPort: httpPort}, nil

}
