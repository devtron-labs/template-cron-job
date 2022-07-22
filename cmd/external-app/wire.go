//go:build wireinject
// +build wireinject

package main

import (
	"github.com/devtron-labs/template-cron-job/pkg/sql"
	"github.com/google/wire"
)

func InitializeApp() (*App, error) {
	wire.Build(
		sql.PgSqlWireSet,
	)
	return &App{}, nil
}
