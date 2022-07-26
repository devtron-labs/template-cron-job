package k8s

import (
	application2 "github.com/devtron-labs/template-cron-job/client/k8s/application"
	"github.com/devtron-labs/template-cron-job/client/k8s/informer"
	"github.com/devtron-labs/template-cron-job/pkg/terminal"
	"github.com/google/wire"
)

var K8sApplicationWireSet = wire.NewSet(
	NewK8sApplicationRouterImpl,
	wire.Bind(new(K8sApplicationRouter), new(*K8sApplicationRouterImpl)),
	NewK8sApplicationRestHandlerImpl,
	wire.Bind(new(K8sApplicationRestHandler), new(*K8sApplicationRestHandlerImpl)),
	NewK8sApplicationServiceImpl,
	wire.Bind(new(K8sApplicationService), new(*K8sApplicationServiceImpl)),
	application2.NewK8sClientServiceImpl,
	wire.Bind(new(application2.K8sClientService), new(*application2.K8sClientServiceImpl)),
	terminal.NewTerminalSessionHandlerImpl,
	wire.Bind(new(terminal.TerminalSessionHandler), new(*terminal.TerminalSessionHandlerImpl)),

	informer.NewGlobalMapClusterNamespace,
	informer.NewK8sInformerFactoryImpl,
	wire.Bind(new(informer.K8sInformerFactory), new(*informer.K8sInformerFactoryImpl)),
)
