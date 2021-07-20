package log

import (
	"fmt"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Module makes the injectable available for FX.
var Module = fx.Provide(New)

// New creates a new injectable.
func New() (*zap.SugaredLogger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		err = fmt.Errorf("impossible to init logger, %v", err)
		return nil, err
	}

	return logger.Sugar(), nil
}
