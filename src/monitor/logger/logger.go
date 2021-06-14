package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Logger struct {
	Logger *zap.Logger
}

// InitLogger 初始化日志 /**
func (logger Logger) InitLogger() Logger {
	//writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
	logger.Logger = zap.New(core, zap.AddCaller())
	return logger
}

/**
 * @Description: 日志设置
 * @return zapcore.Encoder
 */
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

/**
 * @Description: 日志写入文件
 * @return zapcore.WriteSyncer
 */
func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./test.log")
	return zapcore.AddSync(file)
}
