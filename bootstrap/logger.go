package bootstrap

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var _ log.Logger = (*ZapLogger)(nil)

// ZapLogger is a logger impl.
type ZapLogger struct {
	log  *zap.Logger
	Sync func() error
}

// NewLoggerProvider 配置zap日志,将zap日志库引入
func NewLoggerProvider(env string, logFile string, serviceInfo *ServiceInfo) log.Logger {
	var el zapcore.LevelEncoder
	if env == "dev" {
		// 开发环境下，level加色
		el = zapcore.CapitalColorLevelEncoder
	} else {
		el = zapcore.CapitalLevelEncoder
	}

	// 配置zap日志库的编码器
	encoder := zapcore.EncoderConfig{
		TimeKey:  "time",
		LevelKey: "level",
		NameKey:  "logger",
		//CallerKey:      "caller",
		//MessageKey:     "msg",
		//StacktraceKey:  "stack",
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    el,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var logger log.Logger
	logger = NewZapLogger(
		env, logFile, encoder,
		zap.NewAtomicLevelAt(zapcore.InfoLevel),
		zap.AddStacktrace(
			zap.NewAtomicLevelAt(zapcore.ErrorLevel)),
		zap.AddCaller(),
		zap.AddCallerSkip(2),
		zap.Development(),
	)

	// 添加全局字段
	logger = log.With(logger,
		"service.name", serviceInfo.Name,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
		"caller", log.DefaultCaller,
	)
	return logger
}

// NewZapLogger return a zap logger.
func NewZapLogger(env string, logFile string,
	encoder zapcore.EncoderConfig, level zap.AtomicLevel, opts ...zap.Option) *ZapLogger {

	var core zapcore.Core
	// 开发模式下打印到标准输出
	if env == "dev" {
		// 本地开发时，设置日志级别为debug
		level.SetLevel(zap.DebugLevel)
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoder),                      // 编码器配置
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), // 打印到控制台
			level, // 日志级别
		)
	} else {
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoder), // 编码器配置
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(getLogWriter(logFile))), // 打印到控制台和文件
			level, // 日志级别
		)
	}

	zapLogger := zap.New(core, opts...)
	return &ZapLogger{log: zapLogger, Sync: zapLogger.Sync}
}

// 日志自动切割，采用 lumberjack 实现的
func getLogWriter(logFile string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    100,   //日志的最大大小（M）
		MaxBackups: 5,     //日志的最大保存数量
		MaxAge:     30,    //日志文件存储最大天数
		Compress:   false, //是否执行压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

// Log 实现log接口
func (l *ZapLogger) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		l.log.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}
	// Zap.Field is used when keyvals pairs appear
	var data []zap.Field
	for i := 0; i < len(keyvals); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), fmt.Sprint(keyvals[i+1])))
	}
	switch level {
	case log.LevelDebug:
		l.log.Debug("", data...)
	case log.LevelInfo:
		l.log.Info("", data...)
	case log.LevelWarn:
		l.log.Warn("", data...)
	case log.LevelError:
		l.log.Error("", data...)
	}
	return nil
}
