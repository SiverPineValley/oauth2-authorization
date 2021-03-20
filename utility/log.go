package utility

import (
	"fmt"
	"net/http"
	"oauth2-authorization/config"
	"os"
	"strings"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	LogFormat           = "%s %s [%s,%s] --- %s"
	LogFormatNotRequest = "%s %s --- %s"

	DEBUG = "debug"
	INFO  = "info"
	WARN  = "warn"
	ERROR = "error"
)

var logger *zap.Logger
var host string

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./log/oauth2-authorization-" + time.Now().Format("2006-10-02") + ".log",
		MaxSize:    1024,
		MaxBackups: 30,
		MaxAge:     30,
		LocalTime:  true,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoder(func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02T15:04:05:00.000"))
	})
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func InitLogger() {
	writerSyncer := getLogWriter()
	encoder := getEncoder()
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel),
		zapcore.NewCore(encoder, consoleDebugging, zap.DebugLevel),
		zapcore.NewCore(encoder, consoleErrors, zap.ErrorLevel),
	)
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(2))
	defer logger.Sync()

	host = config.GetHost()

	Log("debug", nil, nil, "Logger Init Complete!!")
}

func Log(level string, w http.ResponseWriter, r *http.Request, body ...string) {
	hostname, _ := os.Hostname()
	logBody := ""
	msg := ""

	for _, each := range body {
		logBody = logBody + " " + each
	}

	if r != nil {
		uri := r.RequestURI
		method := r.Method
		msg = fmt.Sprintf(LogFormat, hostname, config.GetProcessName(), method, host+uri, logBody)
	} else {
		msg = fmt.Sprintf(LogFormatNotRequest, hostname, config.GetProcessName(), logBody)
	}

	switch strings.ToLower(level) {
	case INFO:
		info(msg)
	case DEBUG:
		debug(msg)
	case WARN:
		warn(msg)
	case ERROR:
		error(msg)
	default:
		info(msg)
	}
}

func info(msg string) {
	logger.Info(msg)
}

func debug(msg string) {
	logger.Debug(msg)
}

func warn(msg string) {
	logger.Warn(msg)
}

func error(msg string) {
	logger.Error(msg)
}
