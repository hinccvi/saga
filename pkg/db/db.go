package db

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	// postgres driver required by database/sql.
	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	maxOpenConns   int           = 25
	maxIdleConns   int           = 25
	maxLifetime    time.Duration = 5 * time.Minute
	contextTimeout time.Duration = 5 * time.Second
	mongoDBName    string        = "saga"
)

func ConnectPostgres(ctx context.Context, dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(maxLifetime)

	ctx, cancel := context.WithTimeout(ctx, contextTimeout)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectMongo(ctx context.Context, dsn string) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		return &mongo.Client{}, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return &mongo.Client{}, err
	}

	return client, nil
}
