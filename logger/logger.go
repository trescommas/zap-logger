package logger

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(file_path string, max_size int32, max_backups int32, max_age int32, compress bool) *zap.Logger {
	writerSyncer := getLogWriter(file_path, max_size, max_backups, max_age, compress)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)

	console_encoder := getEncoder()
	core_console := zapcore.NewCore(console_encoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel)

	cores := zapcore.NewTee(core, core_console)
	logger := zap.New(cores, zap.AddCaller())

	return logger
}

func getLogWriter(file_path string, max_size int32, max_backups int32, max_age int32, compress bool) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file_path,
		MaxSize:    max_size,
		MaxBackups: max_backups,
		MaxAge:     max_age,
		Compress:   compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(config)
}
