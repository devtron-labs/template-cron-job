package delete

import (
	"fmt"
	"github.com/devtron-labs/template-cron-job/internal/sql/repository/pipelineConfig"
	"github.com/devtron-labs/template-cron-job/pkg/pipeline"
	"github.com/go-pg/pg"
	"go.uber.org/zap"
)

type DeleteServiceFullMode interface {
	DeleteGitProvider(deleteRequest *pipeline.GitRegistry) error
	DeleteDockerRegistryConfig(deleteRequest *pipeline.DockerArtifactStoreBean) error
}

type DeleteServiceFullModeImpl struct {
	logger                *zap.SugaredLogger
	gitMaterialRepository pipelineConfig.MaterialRepository
	gitRegistryConfig     pipeline.GitRegistryConfig
	ciTemplateRepository  pipelineConfig.CiTemplateRepository
	dockerRegistryConfig  pipeline.DockerRegistryConfig
}

func NewDeleteServiceFullModeImpl(logger *zap.SugaredLogger,
	gitMaterialRepository pipelineConfig.MaterialRepository,
	gitRegistryConfig pipeline.GitRegistryConfig,
	ciTemplateRepository pipelineConfig.CiTemplateRepository,
	dockerRegistryConfig pipeline.DockerRegistryConfig,
) *DeleteServiceFullModeImpl {
	return &DeleteServiceFullModeImpl{
		logger:                logger,
		gitMaterialRepository: gitMaterialRepository,
		gitRegistryConfig:     gitRegistryConfig,
		ciTemplateRepository:  ciTemplateRepository,
		dockerRegistryConfig:  dockerRegistryConfig,
	}
}
func (impl DeleteServiceFullModeImpl) DeleteGitProvider(deleteRequest *pipeline.GitRegistry) error {
	//finding if this git account is used in any git material, if yes then will not delete
	materials, err := impl.gitMaterialRepository.FindByGitProviderId(deleteRequest.Id)
	if err != nil && err != pg.ErrNoRows {
		impl.logger.Errorw("err in deleting git provider", "gitProvider", deleteRequest.Name, "err", err)
		return err
	}
	if len(materials) > 0 {
		impl.logger.Errorw("err in deleting git provider, found git materials using provider", "gitProvider", deleteRequest.Name)
		return fmt.Errorf(" Please delete all related git materials before deleting this git account")
	}
	err = impl.gitRegistryConfig.Delete(deleteRequest)
	if err != nil {
		impl.logger.Errorw("error in deleting git account", "err", err, "deleteRequest", deleteRequest)
		return err
	}
	return nil
}

func (impl DeleteServiceFullModeImpl) DeleteDockerRegistryConfig(deleteRequest *pipeline.DockerArtifactStoreBean) error {
	//finding if docker reg is used in any app, if yes then will not delete
	ciTemplates, err := impl.ciTemplateRepository.FindByDockerRegistryId(deleteRequest.Id)
	if err != nil && err != pg.ErrNoRows {
		impl.logger.Errorw("err in deleting docker registry", "dockerRegistry", deleteRequest.Id, "err", err)
		return err
	}
	if len(ciTemplates) > 0 {
		impl.logger.Errorw("err in deleting docker registry, found docker build config using registry", "dockerRegistry", deleteRequest.Id, "err", err)
		return fmt.Errorf(" Please update all related docker config before deleting this registry")
	}
	err = impl.dockerRegistryConfig.DeleteReg(deleteRequest)
	if err != nil {
		impl.logger.Errorw("error in deleting docker registry", "err", err, "deleteRequest", deleteRequest)
		return err
	}
	return nil
}
