package config

import "GoChallenge/internal/infrastructure/entrypoints/router"

func Init() {
	router.Serve()
}
