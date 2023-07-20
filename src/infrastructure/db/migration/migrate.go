package migration

import (
	"url-shortener/src/infrastructure/db"
	"url-shortener/src/links"
)

func MigrateDB() error {

	dbProvider := db.PostgresDBProvider

	err := dbProvider.DB.AutoMigrate(&links.Link{})
	if err != nil {
		return err
	}

	return nil
}
