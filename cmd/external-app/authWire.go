package main

import (
	"github.com/devtron-labs/authenticator/client"
	"github.com/devtron-labs/authenticator/middleware"
	"github.com/devtron-labs/template-cron-job/api/apiToken"
	"github.com/google/wire"
)

// AuthWireSet:	 set of components used to initialise authentication with dex
var AuthWireSet = wire.NewSet(
	client.GetRuntimeConfig,
	client.NewK8sClient,
	client.BuildDexConfig,
	client.GetSettings,
	apiToken.ApiTokenSecretWireSet,
	middleware.NewSessionManager,
	middleware.NewUserLogin,
)
