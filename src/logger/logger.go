package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"io/ioutil"
	"time"
)

const (
	TimeFormat     = "2006-01-02 15:04:05"
	TimeFormatDate = "2006-01-02"
)

var (
	logger *zap.Logger
	level  string
	Sugar  *zap.SugaredLogger
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
	// 自定义级别展示
	debugLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DebugLevel && lvl >= logLevel
	})
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel && lvl >= logLevel
	})
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel && lvl >= logLevel
	})

	writeDebugSyncer := getLogWriter("DEBUG")
	writeInfoSyncer := getLogWriter("INFO")
	writeWarnSyncer := getLogWriter("WARN")

	encoder := getEncoder()

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(writeDebugSyncer), debugLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(writeInfoSyncer), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(writeWarnSyncer), warnLevel),
		//zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), logLevel),
	)

	logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.DebugLevel))
	Sugar = logger.Sugar()
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
		CallerKey:      "file",
		EncodeCaller:   zapcore.ShortCallerEncoder,
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
		Filename:   `./logs/` + filename + "." + time.Now().Format(TimeFormatDate) + ".log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}
}

func FormatTitle(title string) string {
	format := "\t| " + title + ": "
	return format
}

func WriteFile(message string, filename string, file []byte) {
	Sugar.Info(message)
	err := ioutil.WriteFile("./logs/"+filename+".txt", file, 0666)
	if err != nil {
		Sugar.Error("写入文件失败", false, err)
	} //写入文件(字节数组)
}
