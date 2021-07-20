package log

import (
	"fmt"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Module permits create a FX module.
var Module = fx.Provide(New)

// New create a new instance.
func New() (*zap.SugaredLogger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		err = fmt.Errorf("impossible to init logger, %v", err)
		return nil, err
	}

	return logger.Sugar(), nil
}
