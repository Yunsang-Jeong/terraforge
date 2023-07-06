package logger

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Feedback(msg string)
	Debug(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
}

type simpleLogger struct {
	debugLogger *zap.SugaredLogger
	errorLogger *zap.SugaredLogger
}

func NewSimpleLogger(enableDebug bool) *simpleLogger {
	logger := simpleLogger{}

	logger.errorLogger = newLogger(zapcore.ErrorLevel)

	if enableDebug {
		logger.debugLogger = newLogger(zapcore.DebugLevel)
	}

	return &logger
}

func newLogger(level zapcore.Level) *zap.SugaredLogger {
	config := zapcore.EncoderConfig{
		MessageKey:       "msg",
		LevelKey:         "level",
		TimeKey:          "ts",
		NameKey:          "logger",
		CallerKey:        "caller",
		FunctionKey:      zapcore.OmitKey,
		SkipLineEnding:   false,
		StacktraceKey:    "stacktrace",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      zapcore.LowercaseLevelEncoder,
		EncodeTime:       zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05(-0700)"),
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		EncodeCaller:     zapcore.ShortCallerEncoder,
		ConsoleSeparator: "\t",
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(config),
		zapcore.Lock(os.Stderr),
		zap.LevelEnablerFunc(func(l zapcore.Level) bool {
			return l == level
		}),
	)

	logger := zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)

	return logger.Sugar()
}

func (l *simpleLogger) Feedback(msg string) {
	fmt.Printf("%s\n", color.HiGreenString(msg))
}

func (l *simpleLogger) Debug(msg string, keysAndValues ...interface{}) {
	l.debugLogger.Debugw(color.HiYellowString(msg), keysAndValues...)
}

func (l *simpleLogger) Error(msg string, keysAndValues ...interface{}) {
	l.errorLogger.Errorw(color.HiMagentaString(msg), keysAndValues...)
}
