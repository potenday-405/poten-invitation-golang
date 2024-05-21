package util

import "github.com/joho/godotenv"

func EnvInitializer() error {
	if err := godotenv.Load("./.env"); err != nil {
		return err
	}
	return nil
}