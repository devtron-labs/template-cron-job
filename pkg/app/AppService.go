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

package app

import (
	"encoding/json"
	"fmt"
	"github.com/caarlos0/env"
	"github.com/devtron-labs/template-cron-job/internal/sql/repository/app"
	chartRepoRepository "github.com/devtron-labs/template-cron-job/pkg/chartRepo/repository"
	"os"
	"strings"
	"time"

	"github.com/devtron-labs/template-cron-job/internal/sql/repository"
	. "github.com/devtron-labs/template-cron-job/internal/util"
	"github.com/go-pg/pg"
	"go.uber.org/zap"
)

type AppServiceImpl struct {
	logger           *zap.SugaredLogger
	gitFactory       *GitFactory
	appRepository    app.AppRepository
	chartRepository  chartRepoRepository.ChartRepository
	gitOpsRepository repository.GitOpsConfigRepository
	PatchConfig      *PatchConfig
}

type AppService interface {
	RefChartPatchForGitOpsApps()
}

func NewAppService(
	logger *zap.SugaredLogger,
	appRepository app.AppRepository,
	chartRepository chartRepoRepository.ChartRepository,
	gitFactory *GitFactory, gitOpsRepository repository.GitOpsConfigRepository, PatchConfig *PatchConfig) *AppServiceImpl {
	appServiceImpl := &AppServiceImpl{
		logger:           logger,
		appRepository:    appRepository,
		chartRepository:  chartRepository,
		gitFactory:       gitFactory,
		gitOpsRepository: gitOpsRepository,
		PatchConfig:      PatchConfig,
	}
	return appServiceImpl
}

type PatchConfig struct {
	PatchBatchLimit int `env:"PATCH_BATCH_LIMIT" envDefault:"10"` //no of apps
	PatchInterval   int `env:"PATCH_INTERVAL" envDefault:"120"`   //seconds
	FirstAppId      int `env:"FIRST_APP_ID" envDefault:"0"`       //seconds
	LastAppId       int `env:"LAST_APP_ID" envDefault:"0"`        //seconds
}

func GetPatchConfig() (*PatchConfig, error) {
	cfg := &PatchConfig{}
	err := env.Parse(cfg)
	return cfg, err
}

func (impl AppServiceImpl) writeTemplateEditResponseToFile(processedAppResult map[int]map[int]bool) {
	impl.logger.Infow("DONE - data fix for apps completes, check result in /tmp/template-edit-result.txt file")
	jsonString, err := json.Marshal(processedAppResult)
	if err != nil {
		impl.logger.Errorw(" error on writing result file", "err", err)
		return
	}
	f, err := os.Create("/tmp/template-edit-result.txt")
	if err != nil {
		impl.logger.Errorw(" error on writing result file", "err", err)
		return
	}
	defer f.Close()
	_, err2 := f.WriteString(string(jsonString))
	if err2 != nil {
		impl.logger.Errorw(" error on writing result file", "err", err)
		return
	}
}

func (impl AppServiceImpl) GetChartRepoName(gitRepoUrl string) string {
	gitRepoUrl = gitRepoUrl[strings.LastIndex(gitRepoUrl, "/")+1:]
	chartRepoName := strings.ReplaceAll(gitRepoUrl, ".git", "")
	return chartRepoName
}

func (impl AppServiceImpl) RefChartPatchForGitOpsApps() {
	processedAppResult := make(map[int]map[int]bool)
	defer impl.writeTemplateEditResponseToFile(processedAppResult)
	impl.logger.Info("data fix for app, cron started")
	var apps []*app.App
	var err error
	if impl.PatchConfig.FirstAppId > 0 && impl.PatchConfig.LastAppId > 0 {
		apps, err = impl.appRepository.FindAllByAppIdRange(impl.PatchConfig.FirstAppId, impl.PatchConfig.LastAppId)
	} else {
		apps, err = impl.appRepository.FindAll()
	}
	if err != nil {
		impl.logger.Errorw("data fix for app, error fetching apps", "err", err)
		return
	}
	content := "{{- range .Values.rawYaml }}\n---\n{{ toYaml . }}\n  {{- end -}}"
	impl.logger.Infow("data fix for app", "content", content, "app size", len(apps))
	batchCounter := 1
	batchLimit := impl.PatchConfig.PatchBatchLimit
	for _, app := range apps {
		impl.updateGenericInTemplateForApp(app.Id, content, processedAppResult)
		if batchCounter == batchLimit {
			time.Sleep(time.Duration(impl.PatchConfig.PatchInterval) * time.Second)
			batchCounter = 0
		}
		batchCounter = batchCounter + 1
	}
	impl.logger.Infow("data fix for apps completes", "processedAppResult", processedAppResult)
}

func (impl AppServiceImpl) updateGenericInTemplateForApp(appId int, content string, processedAppResult map[int]map[int]bool) {
	charts, err := impl.chartRepository.FindActiveChartsByAppId(appId)
	if err != nil {
		impl.logger.Errorw("data fix for app, error fetching charts for app", "err", err, "appId", appId)
		return
	}
	impl.logger.Infow("data fix for app", "appId", appId, "charts size", len(charts))
	count := 0
	if _, ok := processedAppResult[appId]; !ok {
		processedAppResult[appId] = make(map[int]bool)
	}

	for _, chart := range charts {
		appResult := processedAppResult[appId]
		if _, ok := appResult[chart.Id]; !ok {
			appResult[chart.Id] = false
			processedAppResult[appId] = appResult
		}
		if len(chart.GitRepoUrl) == 0 {
			impl.logger.Warnw("data fix for app, skipped", "appId", appId, "chartId", chart.Id, "gitRepoUrl", chart.GitRepoUrl, "chartLocation", chart.ChartLocation)
			continue
		}
		chartRepoName := impl.GetChartRepoName(chart.GitRepoUrl)
		chartGitAttr := &ChartConfig{
			FileName:       "generic.yaml",
			FileContent:    content,
			ChartName:      chart.ChartName,
			ChartLocation:  fmt.Sprintf("%s/templates", chart.ChartLocation),
			ChartRepoName:  chartRepoName,
			ReleaseMessage: "generic content fix for reference template",
			UserName:       "devtron",
			UserEmailId:    "devtron@devtron.ai",
		}
		gitOpsConfigBitbucket, err := impl.gitOpsRepository.GetGitOpsConfigByProvider(BITBUCKET_PROVIDER)
		if err != nil {
			if err == pg.ErrNoRows {
				gitOpsConfigBitbucket.BitBucketWorkspaceId = ""
			} else {
				impl.logger.Errorw("data fix for app, failed", "err", err, "appId", appId, "chartId", chart.Id, "gitRepoUrl", chart.GitRepoUrl, "chartLocation", chart.ChartLocation)
				continue
			}
		}
		_, err = impl.gitFactory.Client.CommitValues(chartGitAttr, gitOpsConfigBitbucket.BitBucketWorkspaceId)
		if err != nil {
			impl.logger.Errorw("data fix for app, failed", "err", err, "appId", appId, "chartId", chart.Id, "gitRepoUrl", chart.GitRepoUrl, "chartLocation", chart.ChartLocation)
			continue
		}
		count = count + 1
		appResult[chart.Id] = true
		processedAppResult[appId] = appResult
	}
	impl.logger.Infow("data fix for app completed", "appId", appId, "charts size", len(charts), "succeeded count", count)
}
