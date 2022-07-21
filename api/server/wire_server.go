package server

import (
	"github.com/devtron-labs/template-cron-job/pkg/server"
	serverEnvConfig "github.com/devtron-labs/template-cron-job/pkg/server/config"
	serverDataStore "github.com/devtron-labs/template-cron-job/pkg/server/store"
	"github.com/google/wire"
)

var ServerWireSet = wire.NewSet(
	server.NewServerActionAuditLogRepositoryImpl,
	wire.Bind(new(server.ServerActionAuditLogRepository), new(*server.ServerActionAuditLogRepositoryImpl)),
	serverEnvConfig.ParseServerEnvConfig,
	serverDataStore.InitServerDataStore,
	server.NewServerServiceImpl,
	wire.Bind(new(server.ServerService), new(*server.ServerServiceImpl)),
	server.NewServerCacheServiceImpl,
	wire.Bind(new(server.ServerCacheService), new(*server.ServerCacheServiceImpl)),
	NewServerRestHandlerImpl,
	wire.Bind(new(ServerRestHandler), new(*ServerRestHandlerImpl)),
	NewServerRouterImpl,
	wire.Bind(new(ServerRouter), new(*ServerRouterImpl)),
)
