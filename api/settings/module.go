package settings

import (
	"go.uber.org/fx"
	"wiatt/api/settings/config"
	"wiatt/api/settings/log"
	"wiatt/api/settings/server"
)

var Module = fx.Options(
	config.Module,
	log.Module,
	server.Module,
)
