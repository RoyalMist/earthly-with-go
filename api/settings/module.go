package settings

import (
	"go.uber.org/fx"
	"wiatt/api/settings/config"
	"wiatt/api/settings/log"
	"wiatt/api/settings/server"
)

// Module makes the collection of injectables available for FX.
var Module = fx.Options(
	config.Module,
	log.Module,
	server.Module,
)
