package api

import (
	"go.uber.org/fx"
	"wiatt/api/settings"
)

var Module = fx.Options(
	settings.Module,
)
