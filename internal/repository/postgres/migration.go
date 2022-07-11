package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"io/ioutil"
	"path/filepath"
	"time"
)

func Migration(db *pgxpool.Pool) error {
	path := filepath.Join("internal", "repository", "postgres", "migrations")
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	timeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, file := range files {
		c, err := ioutil.ReadFile(path + "/" + file.Name())
		if err != nil {
			return err
		}

		_, err = db.Exec(timeout, string(c))
		if err != nil {
			return err
		}
	}
	return nil
}
