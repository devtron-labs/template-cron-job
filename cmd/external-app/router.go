package main

import (
	"encoding/json"
	"github.com/devtron-labs/template-cron-job/api/apiToken"
	appStoreDeployment "github.com/devtron-labs/template-cron-job/api/appStore/deployment"
	appStoreDiscover "github.com/devtron-labs/template-cron-job/api/appStore/discover"
	appStoreValues "github.com/devtron-labs/template-cron-job/api/appStore/values"
	"github.com/devtron-labs/template-cron-job/api/chartRepo"
	"github.com/devtron-labs/template-cron-job/api/cluster"
	"github.com/devtron-labs/template-cron-job/api/dashboardEvent"
	"github.com/devtron-labs/template-cron-job/api/externalLink"
	client "github.com/devtron-labs/template-cron-job/api/helm-app"
	"github.com/devtron-labs/template-cron-job/api/module"
	"github.com/devtron-labs/template-cron-job/api/restHandler/common"
	"github.com/devtron-labs/template-cron-job/api/server"
	"github.com/devtron-labs/template-cron-job/api/sso"
	"github.com/devtron-labs/template-cron-job/api/team"
	"github.com/devtron-labs/template-cron-job/api/user"
	"github.com/devtron-labs/template-cron-job/client/dashboard"
	"github.com/devtron-labs/template-cron-job/util/k8s"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type MuxRouter struct {
	Router                   *mux.Router
	logger                   *zap.SugaredLogger
	ssoLoginRouter           sso.SsoLoginRouter
	teamRouter               team.TeamRouter
	UserAuthRouter           user.UserAuthRouter
	userRouter               user.UserRouter
	clusterRouter            cluster.ClusterRouter
	dashboardRouter          dashboard.DashboardRouter
	helmAppRouter            client.HelmAppRouter
	environmentRouter        cluster.EnvironmentRouter
	k8sApplicationRouter     k8s.K8sApplicationRouter
	chartRepositoryRouter    chartRepo.ChartRepositoryRouter
	appStoreDiscoverRouter   appStoreDiscover.AppStoreDiscoverRouter
	appStoreValuesRouter     appStoreValues.AppStoreValuesRouter
	appStoreDeploymentRouter appStoreDeployment.AppStoreDeploymentRouter
	dashboardTelemetryRouter dashboardEvent.DashboardTelemetryRouter
	commonDeploymentRouter   appStoreDeployment.CommonDeploymentRouter
	externalLinksRouter      externalLink.ExternalLinkRouter
	moduleRouter             module.ModuleRouter
	serverRouter             server.ServerRouter
	apiTokenRouter           apiToken.ApiTokenRouter
}

func NewMuxRouter(
	logger *zap.SugaredLogger,
	ssoLoginRouter sso.SsoLoginRouter,
	teamRouter team.TeamRouter,
	UserAuthRouter user.UserAuthRouter,
	userRouter user.UserRouter,
	clusterRouter cluster.ClusterRouter,
	dashboardRouter dashboard.DashboardRouter,
	helmAppRouter client.HelmAppRouter,
	environmentRouter cluster.EnvironmentRouter,
	k8sApplicationRouter k8s.K8sApplicationRouter,
	chartRepositoryRouter chartRepo.ChartRepositoryRouter,
	appStoreDiscoverRouter appStoreDiscover.AppStoreDiscoverRouter,
	appStoreValuesRouter appStoreValues.AppStoreValuesRouter,
	appStoreDeploymentRouter appStoreDeployment.AppStoreDeploymentRouter,
	dashboardTelemetryRouter dashboardEvent.DashboardTelemetryRouter,
	commonDeploymentRouter appStoreDeployment.CommonDeploymentRouter,
	externalLinkRouter externalLink.ExternalLinkRouter,
	moduleRouter module.ModuleRouter,
	serverRouter server.ServerRouter, apiTokenRouter apiToken.ApiTokenRouter,
) *MuxRouter {
	r := &MuxRouter{
		Router:                   mux.NewRouter(),
		logger:                   logger,
		ssoLoginRouter:           ssoLoginRouter,
		teamRouter:               teamRouter,
		UserAuthRouter:           UserAuthRouter,
		userRouter:               userRouter,
		clusterRouter:            clusterRouter,
		dashboardRouter:          dashboardRouter,
		helmAppRouter:            helmAppRouter,
		environmentRouter:        environmentRouter,
		k8sApplicationRouter:     k8sApplicationRouter,
		chartRepositoryRouter:    chartRepositoryRouter,
		appStoreDiscoverRouter:   appStoreDiscoverRouter,
		appStoreValuesRouter:     appStoreValuesRouter,
		appStoreDeploymentRouter: appStoreDeploymentRouter,
		dashboardTelemetryRouter: dashboardTelemetryRouter,
		commonDeploymentRouter:   commonDeploymentRouter,
		externalLinksRouter:      externalLinkRouter,
		moduleRouter:             moduleRouter,
		serverRouter:             serverRouter,
		apiTokenRouter:           apiTokenRouter,
	}
	return r
}
func (r *MuxRouter) Init() {
	r.Router.StrictSlash(true)
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
