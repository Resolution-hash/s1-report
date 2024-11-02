package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type UserConfig struct {
	Login    string
	Password string
}

type BOTCongig struct {
	TelegramAPIToken string
}

func LoadUserConfig(isSolarWinds bool) (*UserConfig, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	envFilePath := filepath.Join(cwd, ".env")

	err = godotenv.Load(envFilePath)
	if err != nil {
		return nil, err
	}
	if isSolarWinds {
		return &UserConfig{Login: os.Getenv("SOLARWINDS_LOGIN"), Password: os.Getenv("SOLARWINDS_PASSWORD")}, nil
	}

	return &UserConfig{Login: os.Getenv("LOGIN"), Password: os.Getenv("PASSWORD")}, nil
}

func LoadBOTConfig() (*BOTCongig, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	envFilePath := filepath.Join(cwd, ".env")

	err = godotenv.Load(envFilePath)
	if err != nil {
		return nil, err
	}

	return &BOTCongig{
		TelegramAPIToken: os.Getenv("TELEGRAM_TOKEN"),
	}, nil
}
