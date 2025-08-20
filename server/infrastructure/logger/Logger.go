package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"server/infrastructure/env"

	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

func (l LogLevel) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return "Unknown"
	}
}

var LogConfig = &lumberjack.Logger{
	Filename:   "/server/log/app.log", // ファイル名
	MaxSize:    500,                   // ローテーションするファイルサイズ(megabytes)
	MaxBackups: 3,                     // 保持する古いログの最大ファイル数
	MaxAge:     365,                   // 古いログを保持する日数
	LocalTime:  true,                  // バックアップファイルの時刻フォーマットをサーバローカル時間指定
	Compress:   true,                  // ローテーションされたファイルのgzip圧縮
}

func Log(level LogLevel, message string) {
	_, file, line, _ := runtime.Caller(2)
	file = strings.Replace(file, os.Getenv("GOPATH")+"/src/", "", 1)
	now := time.Now()
	// log.SetOutput(LogConfig)
	// log.Printf("%s [%s] %s:%d %s\n", now.Format("2006-01-02 15:04:05"), level.String(), file, line, message)

	// cloud loggingのために標準出力に変更
	fmt.Printf("%s [%s] %s:%d %s\n", now.Format("2006-01-02 15:04:05"), level.String(), file, line, message)
}

func Trace() string {
	buf := make([]byte, 1024)
	n := runtime.Stack(buf, false) // trueを指定すると、すべてのゴルーチンのスタックトレースを取得します
	return string(buf[:n])
}

func Error(message string) {
	Log(LevelError, message+"\nStacktrace:\n"+Trace())
}

func Errorf(format string, a ...interface{}) {
	Log(LevelError, fmt.Sprintf(format, a...)+"\nStacktrace:\n"+Trace())
}

func Info(message string) {
	Log(LevelInfo, message)
}

func Infof(format string, a ...interface{}) {
	Log(LevelInfo, fmt.Sprintf(format, a...))
}

func Debugf(format string, a ...interface{}) {
	if env.GetEnv(env.EnvMode) == "dev" {
		Log(LevelDebug, fmt.Sprintf(format, a...))
	}
}

func Debug(message string) {
	if env.GetEnv(env.EnvMode) == "dev" {
		Log(LevelDebug, message)
	}
}

func DebugPrint(object interface{}) {
	if env.GetEnv(env.EnvMode) == "dev" {
		fmt.Print("============================\n")
		fmt.Printf("Type: %T\n", object)
		fmt.Printf("Value: %+v\n", object)
		
		// JSON形式での出力も試みる
		if jsonBytes, err := json.MarshalIndent(object, "", "  "); err == nil {
			fmt.Printf("JSON:\n%s\n", string(jsonBytes))
		}
		fmt.Print("============================\n")
	}
}

func Fatal(message string) {
	Log(LevelFatal, message)
}

func Fatalf(format string, a ...interface{}) {
	Log(LevelFatal, fmt.Sprintf(format, a...))
}

func Warn(message string) {
	Log(LevelWarn, message)
}

func Warnf(format string, a ...interface{}) {
	Log(LevelWarn, fmt.Sprintf(format, a...))
}

// ログディレクトリを作成
func ensureLogDir() error {
	logDir := "./log"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		return os.MkdirAll(logDir, 0755)
	}
	return nil
}

// 単純にテキストをログファイルに出力（コンソールと同じ内容）
func LogToFile(content string) error {
	if err := ensureLogDir(); err != nil {
		return fmt.Errorf("failed to create log directory: %w", err)
	}

	filename := fmt.Sprintf("%d.log", time.Now().UnixMilli())
	logPath := fmt.Sprintf("./log/%s", filename)
	
	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}
	defer f.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	if _, err := f.WriteString(fmt.Sprintf("[%s]\n%s\n", timestamp, content)); err != nil {
		return fmt.Errorf("failed to write to log file: %w", err)
	}

	if env.GetEnv(env.EnvMode) == "dev" {
		fmt.Printf("Log written to: %s\n", logPath)
	}

	return nil
}

// フォーマット付きでログファイルに出力
func LogToFilef(format string, a ...interface{}) error {
	return LogToFile(fmt.Sprintf(format, a...))
}

// 複数行のテキストをログファイルに出力
func LogLinesToFile(lines []string) error {
	return LogToFile(strings.Join(lines, "\n"))
}

// Println相当（改行付き）でログファイルに出力
func Println(a ...interface{}) error {
	return LogToFile(fmt.Sprintln(a...))
}

// Printf相当でログファイルに出力  
func Printf(format string, a ...interface{}) error {
	return LogToFile(fmt.Sprintf(format, a...))
}
