package logger

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"log"
	"time"
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
	//core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(writeInfoSyncer), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(writeWarnSyncer), warnLevel),
		//zapcore.NewCore(encoder,
		//	zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), logLevel),
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
		EncodeLevel: zapcore.CapitalLevelEncoder, //将级别转换成大写
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	}
	encoderConfig := zapcore.NewConsoleEncoder(config)
	return encoderConfig
}

/**
 * @Description: 日志写入文件
 * @return zapcore.WriteSyncer
 */
func getLogWriter(filename string) io.Writer {
	// demo.log是指向最新日志的链接
	hook, err := rotatelogs.New(
		`./logs/`+filename+".%Y%m%d%H.log",
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*30),    // 保存30天
		rotatelogs.WithRotationTime(time.Hour*24), //切割频率 24小时
	)
	if err != nil {
		log.Println("日志启动异常")
		panic(err)
	}
	return hook
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
