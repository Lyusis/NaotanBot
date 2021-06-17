package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"time"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
	TimeFormatDate = "2006-01-02"
)

var (
	logger *zap.Logger
	level  string
)

// InitLogger 初始化日志 /**
func init() {
	// 设置级别
	logLevel := zap.DebugLevel
	switch level {
	case "debug":
		logLevel = zap.DebugLevel
	case "info":
		logLevel = zap.InfoLevel
	case "warn":
		logLevel = zap.WarnLevel
	case "error":
		logLevel = zap.ErrorLevel
	case "panic":
		logLevel = zap.PanicLevel
	case "fatal":
		logLevel = zap.FatalLevel
	default:
		logLevel = zap.InfoLevel
	}
	// 实现两个判断日志等级的interface  可以自定义级别展示
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel && lvl >= logLevel
	})

	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel && lvl >= logLevel
	})

	writeInfoSyncer := getLogWriter("INFO")
	writeWarnSyncer := getLogWriter("WARN")

	encoder := getEncoder()

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(writeInfoSyncer), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(writeWarnSyncer), warnLevel),
		//zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), logLevel),
	)

	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
}

/**
 * @Description: 设置基本日志格式
 * @return zapcore.Encoder
 */
func getEncoder() zapcore.Encoder {
	config := zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder, 
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(TimeFormat))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}
	encoderConfig := zapcore.NewConsoleEncoder(config)
	return encoderConfig
}

/**
 * @Description: 日志写入文件
 * @return zapcore.WriteSyncer
 */
func getLogWriter(filename string) io.Writer {
	return &lumberjack.Logger{
		Filename: `./logs/`+filename+"." + time.Now().Format(TimeFormatDate) +".log",
		MaxSize: 10,
		MaxBackups:5,
		MaxAge: 30,
		Compress: true, 
	  }
}

// Debug logs.Debug(...)
func Debug(format string, v ...interface{}) {
	logger.Sugar().Debugf(format, v...)
}

func Info(format string, v ...interface{}) {
	logger.Sugar().Infof(format, v...)
}

func Warn(format string, v ...interface{}) {
	logger.Sugar().Warnf(format, v...)
}

func Error(format string, v ...interface{}) {
	logger.Sugar().Errorf(format, v...)
}

func Panic(format string, v ...interface{}) {
	logger.Sugar().Panicf(format, v...)
}
