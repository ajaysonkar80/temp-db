package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Connect() error {
	connConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}

	Pool, err = pgxpool.Connect(context.Background(), connConfig)
	if err != nil {
		return err
	}

	return nil
}

func Close() {
	if Pool != nil {
		Pool.Close()
	}
}

//CREATE TABLE C10M1MCQ (QNo SERIAL PRIMARY KEY, question VARCHAR(500) NOT NULL, option1 VARCHAR(250) NOT NULL, option2 VARCHAR(250) NOT NULL, option3 VARCHAR(250) NOT NULL, option4 VARCHAR(250) NOT NULL, Answer INT NOT NULL, Explanation VARCHAR(1000));
