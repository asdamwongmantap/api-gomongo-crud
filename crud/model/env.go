package model

import "github.com/asdamwongmantap/api-gomongo-crud/lib/db"

type (
	EnvConfig struct {
		AppName string
		Env     string

		Mongo db.MongoConfig
	}
)
