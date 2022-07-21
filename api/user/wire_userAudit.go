package user

import (
	"github.com/devtron-labs/template-cron-job/pkg/user"
	"github.com/devtron-labs/template-cron-job/pkg/user/repository"
	"github.com/google/wire"
)

var UserAuditWireSet = wire.NewSet(
	repository.NewUserAuditRepositoryImpl,
	wire.Bind(new(repository.UserAuditRepository), new(*repository.UserAuditRepositoryImpl)),
	user.NewUserAuditServiceImpl,
	wire.Bind(new(user.UserAuditService), new(*user.UserAuditServiceImpl)),
)
