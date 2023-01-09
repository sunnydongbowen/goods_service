package logger

// @program:     goods_servies
// @file:        logger.go
// @author:      bowen
// @create:      2023-01-09 21:04
// @description:

import (
	"go.uber.org/zap"
	"goods_service/config"
)

var lg *zap.Logger

func Init(cfg *config.LogConfig, mode string)
