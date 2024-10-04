package log

import (
	"github.com/ArtemFed/hse-wishlist/pkg/xapp"
	"github.com/ArtemFed/hse-wishlist/pkg/xlogger"
	"go.uber.org/zap"
)

var Logger *zap.Logger

func Init(cfgLogger *xlogger.Config, cfgApp *xapp.Config) error {
	var err error
	Logger, err = xlogger.Init(cfgLogger, cfgApp)
	if err != nil {
		return err
	}

	return err
}
