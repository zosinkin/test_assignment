package core_logger

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type loggerContextKey struct{}

var (
	key = loggerContextKey{}
)

//это не создание логгера с нуля, а только добавление доп. функционала
type Logger struct {
	*zap.Logger
	file *os.File
}


//функция, которая кладет логгер в контексе
func ToContext(ctx context.Context, log *Logger) context.Context {
	return context.WithValue(
		ctx,
		key,
		log,
	)
}


//функция, которая получает логгер из контекста
func FromContext(ctx context.Context) *Logger {
	log, ok := ctx.Value(key).(*Logger)
	if !ok {
		panic("No logger in context")
	}
	return log
}


//создание логгера
func NewLogger(config Config) (*Logger, error) {
	zapLvl := zap.NewAtomicLevel()
	if err := zapLvl.UnmarshalText([]byte(config.Level)); err != nil {
		return nil, fmt.Errorf("Unmarshal log level: %w", err)
	}

	if err := os.MkdirAll(config.Folder, 0755); err != nil {
		return nil, fmt.Errorf("mkdir log Folder: %w", err)
	}

	timestamp := time.Now().UTC().Format("2006-01-02T15-04-05.00000")

	logFilePath := filepath.Join(
		config.Folder,
		fmt.Sprintf("%s.log", timestamp),	
	)

	logFile, err := os.OpenFile(logFilePath, os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("open log file: %w", err)
	}

	zapConfig := zap.NewDevelopmentEncoderConfig()

	zapConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2005-01-02T15:04:05.00000")

	zapEncoder := zapcore.NewConsoleEncoder(zapConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(zapEncoder, zapcore.AddSync(os.Stdout), zapLvl),
		zapcore.NewCore(zapEncoder, zapcore.AddSync(logFile), zapLvl),
	)

	
	zapLogger := zap.New(core, zap.AddCaller())


	return &Logger{
		Logger: zapLogger,
		file:   logFile,
	}, nil
}


func (l *Logger) With(field ...zap.Field) *Logger {
	return &Logger{
		Logger: l.Logger.With(field...),
		file: 	l.file,
	}
}



func (l *Logger) Close() {
	if err := l.file.Close(); err != nil {
		fmt.Println("Failed to close application logger:", err)
	}
}




