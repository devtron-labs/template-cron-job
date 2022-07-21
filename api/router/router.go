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

package router

import (
	"encoding/json"
	"github.com/devtron-labs/template-cron-job/api/restHandler/common"
	"github.com/devtron-labs/template-cron-job/internal/sql/repository"
	"github.com/devtron-labs/template-cron-job/pkg/app"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type MuxRouter struct {
	logger                 *zap.SugaredLogger
	Router                 *mux.Router
	gitOpsConfigRepository repository.GitOpsConfigRepository
	appService             app.AppService
}

func NewMuxRouter(logger *zap.SugaredLogger, gitOpsConfigRepository repository.GitOpsConfigRepository,
	appService app.AppService) *MuxRouter {
	r := &MuxRouter{
		Router:                 mux.NewRouter(),
		logger:                 logger,
		gitOpsConfigRepository: gitOpsConfigRepository,
		appService:             appService,
	}
	return r
}

func (r MuxRouter) Init() {

	r.Router.Path("/health").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(200)
		response := common.Response{}
		response.Code = 200
		response.Result = "OK"
		b, err := json.Marshal(response)
		if err != nil {
			b = []byte("OK")
			r.logger.Errorw("Unexpected error in apiError", "err", err)
		}
		_, _ = writer.Write(b)
	})

}
