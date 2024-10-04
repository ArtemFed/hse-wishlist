package xlogger

import (
	"github.com/ArtemFed/hse-wishlist/pkg/xapp"
	"github.com/TheZeroSlave/zapsentry"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

const (
	serviceTag     = "SYSTEM_SERVICE"
	environmentTag = "SYSTEM_ENV"
)

// TODO support using sentry

func Init(cfg *Config, appCfg *xapp.Config) (*zap.Logger, error) {
	var levelZap, levelSentry zapcore.Level
	err := levelZap.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		log.Printf("Zap logs level with value=%s not initiolized", levelZap)
		return nil, err
	}

	var cfgZap zap.Config
	if cfg.Env == xapp.DevelopmentEnv {
		cfgZap = zap.NewDevelopmentConfig()
	} else {
		cfgZap = zap.NewProductionConfig()
	}
	cfgZap.Level = zap.NewAtomicLevelAt(levelZap)
	cfgZap.OutputPaths = cfg.Outputs
	cfgZap.ErrorOutputPaths = cfg.ErrorOutputs

	logger, err := cfgZap.Build()
	if err != nil {
		log.Println("Zap logger build failed")
		return nil, err
	}

	err = levelSentry.UnmarshalText([]byte(cfg.SentryLevel))
	if err != nil {
		log.Printf("Sentry logs level with value=%s not initialized", levelZap)
		return nil, err
	}
	cfgSentry := zapsentry.Configuration{
		Level: levelSentry,
		Tags: map[string]string{
			"environment": cfg.Env,
			"app":         appCfg.Name,
		},
	}
	core, err := zapsentry.NewCore(
		cfgSentry,
		zapsentry.NewSentryClientFromDSN(cfg.SentryDSN),
	)
	if err != nil {
		log.Println("Zapsentry NewCore not initialized")
		return nil, err
	}

	logger = zapsentry.AttachCoreToLogger(core, logger)

	return logger, nil
}
