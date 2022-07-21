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
	"context"
	"github.com/devtron-labs/template-cron-job/api/router"
	"github.com/devtron-labs/template-cron-job/pkg/app"
	"github.com/go-pg/pg"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type App struct {
	MuxRouter  *router.MuxRouter
	Logger     *zap.SugaredLogger
	server     *http.Server
	db         *pg.DB
	serveTls   bool
	appService app.AppService
}

func NewApp(router *router.MuxRouter,
	Logger *zap.SugaredLogger,
	db *pg.DB,
	appService app.AppService,
) *App {
	app := &App{
		MuxRouter:  router,
		Logger:     Logger,
		db:         db,
		serveTls:   false,
		appService: appService,
	}
	return app
}

func (app *App) Start() {
	port := 8089 //TODO: extract from environment variable
	app.Logger.Debugw("starting server")
	app.Logger.Infow("starting server on ", "port", port)
	app.MuxRouter.Init()
	app.appService.TemplateFixForAllApps()
}

func (app *App) Stop() {
	app.Logger.Infow("orchestrator shutdown initiating")
	timeoutContext, _ := context.WithTimeout(context.Background(), 5*time.Second)
	app.Logger.Infow("closing router")
	err := app.server.Shutdown(timeoutContext)
	if err != nil {
		app.Logger.Errorw("error in mux router shutdown", "err", err)
	}
	app.Logger.Infow("closing db connection")
	err = app.db.Close()
	if err != nil {
		app.Logger.Errorw("error in closing db connection", "err", err)
	}

	app.Logger.Infow("housekeeping done. exiting now")
}
