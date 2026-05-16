package config

import "github.com/joho/godotenv"

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	return nil
}

type GRPCConfig interface {
	GRPCAddress() string
}

type PGConfig interface {
	DSN() string
}
