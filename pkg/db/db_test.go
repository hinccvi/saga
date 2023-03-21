package db

import (
	"context"
	"flag"
	"reflect"
	"testing"

	"github.com/hinccvi/saga/internal/config"
	"github.com/stretchr/testify/assert"
)

//nolint:gochecknoglobals // environment flag that only used in main
var flagMode = flag.String("env", "local", "environment")

func TestConnect(t *testing.T) {
	flag.Parse()

	cfg, err := config.Load("customer", *flagMode)
	assert.Nil(t, err)
	assert.False(t, reflect.DeepEqual(config.Config{}, cfg))

	db, err := ConnectPostgres(context.TODO(), cfg.Postgres.Dsn)
	assert.Nil(t, err)
	assert.NotNil(t, db)
}

func TestConnect_WhenConfigIsEmpty(t *testing.T) {
	db, err := ConnectPostgres(context.TODO(), "")
	assert.NotNil(t, err)
	assert.Nil(t, db)
}

func TestConnect_WhenInvalidDSN(t *testing.T) {
	db, err := ConnectPostgres(context.TODO(), "xxx")
	assert.NotNil(t, err)
	assert.Nil(t, db)
}
