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

package session

import (
	"context"
	"github.com/argoproj/argo-cd/util/session"
	"github.com/argoproj/argo-cd/util/settings"
	"github.com/devtron-labs/template-cron-job/client/argocdServer"
	"github.com/devtron-labs/template-cron-job/pkg/dex"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	sessionManager *session.SessionManager
)

func SessionManager(settings *settings.SettingsManager, cfg *dex.Config) *session.SessionManager {
	//cfg, err := cfg
	//if err != nil {
	//	log.Fatal(err)
	//}
	return nil
}

func CDSettingsManager(settings *settings.SettingsManager) (*settings.ArgoCDSettings, error) {

	return nil, nil
}

func SettingsManager(cfg *argocdServer.Config) (*settings.SettingsManager, error) {

	return settings.NewSettingsManager(context.Background(), nil, ""), nil
}

func GetK8sclient() (k8sClient *kubernetes.Clientset, k8sConfig clientcmd.ClientConfig) {

	return nil, nil
}

//workaround for wire single value return
func NewK8sClient() *kubernetes.Clientset {
	return nil
}
