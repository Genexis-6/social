package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Genexis-6/social/internal/env"
)

func DbPool(SetMaxIdleConns int, SetMaxOpenConns int, SetConnMaxIdleTime time.Duration) (*sql.DB, error) {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		env.GetEnvString("DB_USER", ""),
		env.GetEnvString("DB_PASSWORD", "Genexis123*"),
		env.GetEnvString("DB_HOST", "localhost"),
		env.GetEnvString("DB_PORT", "5432"),
		env.GetEnvString("DB_NAME", ""),
	)

	db, err := sql.Open("postgres", url)

	ctx,cancle := context.WithTimeout(context.Background(), 5 * time.Second)

	
	defer cancle()

	db.SetMaxIdleConns(SetMaxIdleConns)
	db.SetConnMaxIdleTime(SetConnMaxIdleTime)
	db.SetMaxOpenConns(SetMaxOpenConns)

		if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return db, nil

}
