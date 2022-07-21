package serverEnvConfig

import (
	"fmt"
	"github.com/caarlos0/env"
)

type ServerEnvConfig struct {
	DevtronInstallationType              string `env:"DEVTRON_INSTALLATION_TYPE"`
	InstallerCrdObjectGroupName          string `env:"INSTALLER_CRD_OBJECT_GROUP_NAME" envDefault:"installer.template-cron-job.ai"`
	InstallerCrdObjectVersion            string `env:"INSTALLER_CRD_OBJECT_VERSION" envDefault:"v1alpha1"`
	InstallerCrdObjectResource           string `env:"INSTALLER_CRD_OBJECT_RESOURCE" envDefault:"installers"`
	InstallerCrdNamespace                string `env:"INSTALLER_CRD_NAMESPACE" envDefault:"devtroncd"`
	DevtronHelmRepoName                  string `env:"DEVTRON_HELM_REPO_NAME" envDefault:"template-cron-job"`
	DevtronHelmRepoUrl                   string `env:"DEVTRON_HELM_REPO_URL" envDefault:"https://helm.template-cron-job.ai"`
	DevtronHelmReleaseName               string `env:"DEVTRON_HELM_RELEASE_NAME" envDefault:"template-cron-job"`
	DevtronHelmReleaseNamespace          string `env:"DEVTRON_HELM_RELEASE_NAMESPACE" envDefault:"devtroncd"`
	DevtronHelmReleaseChartName          string `env:"DEVTRON_HELM_RELEASE_CHART_NAME" envDefault:"template-cron-job-operator"`
	DevtronVersionIdentifierInHelmValues string `env:"DEVTRON_VERSION_IDENTIFIER_IN_HELM_VALUES" envDefault:"installer.release"`
	DevtronModulesIdentifierInHelmValues string `env:"DEVTRON_MODULES_IDENTIFIER_IN_HELM_VALUES" envDefault:"installer.modules"`
	DevtronBomUrl                        string `env:"DEVTRON_BOM_URL" envDefault:"https://raw.githubusercontent.com/template-cron-job-labs/template-cron-job/%s/charts/template-cron-job/template-cron-job-bom.yaml"`
}

func ParseServerEnvConfig() (*ServerEnvConfig, error) {
	cfg := &ServerEnvConfig{}
	err := env.Parse(cfg)
	if err != nil {
		fmt.Println("failed to parse server env config: " + err.Error())
		return nil, err
	}
	return cfg, nil
}
