package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"io/ioutil"
	"path/filepath"
)

func Migration(db *pgxpool.Pool) error {
	path := filepath.Join("internal", "repository", "postgres", "migrations")
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		c, err := ioutil.ReadFile(path + "/" + file.Name())
		if err != nil {
			return err
		}

		_, err = db.Exec(context.Background(), string(c))
		if err != nil {
			return err
		}
	}
	return nil
}
