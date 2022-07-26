package cron

import (
	"github.com/devtron-labs/template-cron-job/pkg/app"
	"github.com/devtron-labs/template-cron-job/pkg/appStore/deployment/service"
	"github.com/devtron-labs/template-cron-job/pkg/pipeline"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type HelmApplicationStatusUpdateHandler interface {
	HelmApplicationStatusUpdate()
}

type HelmApplicationStatusUpdateHandlerImpl struct {
	logger              *zap.SugaredLogger
	cron                *cron.Cron
	appService          app.AppService
	workflowDagExecutor pipeline.WorkflowDagExecutor
	installedAppService service.InstalledAppService
	CdHandler           pipeline.CdHandler
}

const HelmAppStatusUpdateCronExpr string = "*/2 * * * *"

func NewHelmApplicationStatusUpdateHandlerImpl(logger *zap.SugaredLogger, appService app.AppService,
	workflowDagExecutor pipeline.WorkflowDagExecutor, installedAppService service.InstalledAppService,
	CdHandler pipeline.CdHandler) *HelmApplicationStatusUpdateHandlerImpl {
	impl := &HelmApplicationStatusUpdateHandlerImpl{
		logger:              logger,
		appService:          appService,
		workflowDagExecutor: workflowDagExecutor,
		installedAppService: installedAppService,
		CdHandler:           CdHandler,
	}

	return impl
}

func (impl *HelmApplicationStatusUpdateHandlerImpl) HelmApplicationStatusUpdate() {
	err := impl.CdHandler.CheckHelmAppStatusPeriodicallyAndUpdateInDb()
	if err != nil {
		impl.logger.Errorw("error helm app status update - cron job", "err", err)
		return
	}
	return
}
