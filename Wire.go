//go:build wireinject
// +build wireinject

/*
 * Copyright (c) 2020 Devtron Labs
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"github.com/devtron-labs/template-cron-job/api/router"
	"github.com/devtron-labs/template-cron-job/internal/sql/repository"
	"github.com/devtron-labs/template-cron-job/internal/sql/repository/app"
	"github.com/devtron-labs/template-cron-job/internal/util"
	chartRepoRepository "github.com/devtron-labs/template-cron-job/pkg/chartRepo/repository"
	"github.com/devtron-labs/template-cron-job/pkg/sql"
	"github.com/google/wire"
)

func InitializeApp() (*App, error) {

	wire.Build(
		util.NewSugardLogger,
		sql.PgSqlWireSet,
		router.NewMuxRouter,
		app.NewAppRepositoryImpl,
		wire.Bind(new(app.AppRepository), new(*app.AppRepositoryImpl)),
		chartRepoRepository.NewChartRepoRepositoryImpl,
		wire.Bind(new(chartRepoRepository.ChartRepoRepository), new(*chartRepoRepository.ChartRepoRepositoryImpl)),
		repository.NewGitOpsConfigRepositoryImpl,
		wire.Bind(new(repository.GitOpsConfigRepository), new(*repository.GitOpsConfigRepositoryImpl)),

		util.NewGitFactory,
		util.NewGitCliUtil,

	)
	return &App{}, nil
}
