package api

import (
	"go.uber.org/fx"
	"wiatt/api/settings"
)

// Module makes the collection of injectables available for FX.
var Module = fx.Options(
	settings.Module,
)
